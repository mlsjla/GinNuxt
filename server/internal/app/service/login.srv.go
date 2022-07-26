package service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dchest/captcha"
	"github.com/google/wire"

	"github.com/mlsjla/gin-nuxt/server/internal/app/dao"
	"github.com/mlsjla/gin-nuxt/server/internal/app/schema"
	"github.com/mlsjla/gin-nuxt/server/pkg/auth"
	"github.com/mlsjla/gin-nuxt/server/pkg/errors"
	"github.com/mlsjla/gin-nuxt/server/pkg/util/hash"
)

var LoginSet = wire.NewSet(wire.Struct(new(LoginSrv), "*"))

type LoginSrv struct {
	Auth         auth.Auther
	UserRepo     *dao.UserRepo
	UserRoleRepo *dao.UserRoleRepo
	RoleRepo     *dao.RoleRepo
	MenuRepo     *dao.MenuRepo
}

func (a *LoginSrv) GetCaptcha(ctx context.Context, length int) (*schema.LoginCaptcha, error) {
	captchaID := captcha.NewLen(length)
	item := &schema.LoginCaptcha{
		CaptchaID: captchaID,
	}
	return item, nil
}

func (a *LoginSrv) ResCaptcha(ctx context.Context, w http.ResponseWriter, captchaID string, width, height int) error {
	err := captcha.WriteImage(w, captchaID, width, height)
	if err != nil {
		if err == captcha.ErrNotFound {
			return errors.ErrNotFound
		}
		return errors.WithStack(err)
	}

	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	w.Header().Set("Content-Type", "image/png")
	return nil
}

func (a *LoginSrv) Verify(ctx context.Context, userName, password string) (*schema.User, error) {
	root := schema.GetRootUser()
	v := hash.BcryptCompare(root.Password, password)
	fmt.Println(v, root)
	if userName == root.Username && v == nil && root.Username != "" {
		return root, nil
	}

	result, err := a.UserRepo.Query(ctx, schema.UserQueryParam{
		Username: userName,
	})
	if err != nil {
		return nil, err
	} else if len(result.Data) == 0 {
		return nil, errors.New400Response("not found username")
	}

	item := result.Data[0]
	if item.Password != hash.SHA1String(password) {
		return nil, errors.New400Response("password incorrect")
	} else if item.Status != 1 {
		return nil, errors.ErrUserDisable
	}

	return item, nil
}

func (a *LoginSrv) GenerateToken(ctx context.Context, userID string) (*schema.LoginTokenInfo, error) {
	tokenInfo, err := a.Auth.GenerateToken(ctx, userID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	item := &schema.LoginTokenInfo{
		AccessToken: tokenInfo.GetAccessToken(),
		TokenType:   tokenInfo.GetTokenType(),
		ExpiresAt:   tokenInfo.GetExpiresAt(),
	}
	return item, nil
}

func (a *LoginSrv) DestroyToken(ctx context.Context, tokenString string) error {
	err := a.Auth.DestroyToken(ctx, tokenString)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (a *LoginSrv) checkAndGetUser(ctx context.Context, userID uint64) (*schema.User, error) {
	user, err := a.UserRepo.Get(ctx, userID)
	if err != nil {
		return nil, err
	} else if user == nil {
		return nil, errors.ErrNotFound
	} else if !schema.CheckIsRootUser(ctx, userID) && user.Status != 1 {
		return nil, errors.ErrUserDisable
	}
	return user, nil
}

func (a *LoginSrv) GetLoginInfo(ctx context.Context, userID uint64) (*schema.UserLoginInfo, error) {
	user, err := a.checkAndGetUser(ctx, userID)
	if err != nil {
		if isRoot := schema.CheckIsRootUser(ctx, userID); isRoot {
			root := schema.GetRootUser()
			loginInfo := &schema.UserLoginInfo{
				Username: root.Username,
				Realname: root.Realname,
				UserID:   root.ID,
			}
			return loginInfo, nil
		}
		return nil, err
	}

	info := &schema.UserLoginInfo{
		UserID:        user.ID,
		Username:      user.Username,
		Realname:      user.Realname,
		Nickname:      user.Nickname,
		ThreadCount:   user.ThreadCount,
		FollowCount:   user.FollowCount,
		FansCount:     user.FansCount,
		LickdCount:    user.LickdCount,
		QuestionCount: user.QuestionCount,
		Avatar:        user.Avatar,
		Introduce:     user.Introduce,
	}

	userRoleResult, err := a.UserRoleRepo.Query(ctx, schema.UserRoleQueryParam{
		UserID: userID,
	})
	if err != nil {
		return nil, err
	}

	if roleIDs := userRoleResult.Data.ToRoleIDs(); len(roleIDs) > 0 {
		roleResult, err := a.RoleRepo.Query(ctx, schema.RoleQueryParam{
			IDs:    roleIDs,
			Status: 1,
		})
		if err != nil {
			return nil, err
		}
		info.Roles = roleResult.Data
	}

	return info, nil
}

func (a *LoginSrv) UpdateUserInfo(ctx context.Context, userID uint64, params schema.UpdateInfoParam) error {
	if userID < 1 {
		return errors.New400Response("请先登录")
	}
	if !schema.CheckIsRootUser(ctx, userID) {
		info, err := a.checkAndGetUser(ctx, userID)
		if err != nil || info.ID != userID {
			return errors.New400Response("当前用户不存在")
		}
	}

	var user schema.User
	if params.Nickname != "" {
		user.Nickname = params.Nickname
	}
	if params.Avatar != "" {
		user.Avatar = params.Avatar
	}
	user.Introduce = params.Introduce
	user.Realname = params.Realname

	// 如果是管理员，则检测而用户表是否有信息，如有，则更新，无没有，则创建
	if schema.CheckIsRootUser(ctx, userID) {
		_, err := a.checkAndGetUser(ctx, userID)
		root := schema.GetRootUser()
		if err == errors.ErrNotFound {
			// 则创建用户
			user.ID = userID
			user.Username = root.Username
			user.Status = 1
			user.Introduce = params.Introduce
			user.Nickname = params.Nickname
			user.Avatar = params.Avatar
			user.Realname = params.Realname
			return a.UserRepo.Create(ctx, user)
		} else if err != nil {
			return err
		}
	}

	return a.UserRepo.Update(ctx, userID, user)
}

func (a *LoginSrv) UpdatePassword(ctx context.Context, userID uint64, params schema.UpdatePasswordParam) error {
	if schema.CheckIsRootUser(ctx, userID) {
		return errors.New400Response("root用户不允许更新密码")
	}

	user, err := a.checkAndGetUser(ctx, userID)
	if err != nil {
		return err
	} else if hash.SHA1String(params.OldPassword) != user.Password {
		return errors.New400Response("旧密码不正确")
	}

	params.NewPassword = hash.SHA1String(params.NewPassword)
	return a.UserRepo.UpdatePassword(ctx, userID, params.NewPassword)
}
