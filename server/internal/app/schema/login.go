package schema

type LoginParam struct {
	Username    string `json:"username" binding:"required"`     // 用户名
	Password    string `json:"password" binding:"required"`     // 密码(md5加密)
	CaptchaID   string `json:"captcha_id" binding:"required"`   // 验证码ID
	CaptchaCode string `json:"captcha_code" binding:"required"` // 验证码
}

type SmsRegisterParam struct {
	Value       string `json:"value"`
	Mobile      string `json:"mobile"`
	Username    string `json:"username"`     // 用户名
	Password    string `json:"password"`     // 密码(md5加密)
	CaptchaID   string `json:"captcha_id"`   // 验证码ID
	CaptchaCode string `json:"captcha_code"` // 验证码
}

type UserLoginInfo struct {
	UserID        uint64 `json:"user_id,string"` // 用户ID
	Username      string `json:"username"`       // 用户名
	Realname      string `json:"realname"`       // 真实姓名
	Roles         Roles  `json:"roles"`          // 角色列表
	Nickname      string `json:"nickname"`       // 昵称
	ThreadCount   int    `json:"thread_count"`   //
	FollowCount   int    `json:"follow_count"`   //
	FansCount     int    `json:"fans_count"`     //
	LickdCount    int    `json:"lickd_count"`    //
	QuestionCount int    `json:"question_count"` //
	Avatar        string `json:"avatar"`
	Introduce     string `json:"introduce"` // 个人介绍
}

type UpdateInfoParam struct {
	Nickname  string `json:"nickname"` // 昵称
	Avatar    string `json:"avatar"`
	Introduce string `json:"introduce"` // 个人介绍
}

type UpdatePasswordParam struct {
	OldPassword string `json:"old_password" binding:"required"` // 旧密码(md5加密)
	NewPassword string `json:"new_password" binding:"required"` // 新密码(md5加密)
}

type LoginCaptcha struct {
	CaptchaID string `json:"captcha_id"` // 验证码ID
}

type LoginTokenInfo struct {
	AccessToken string `json:"access_token"` // 访问令牌
	TokenType   string `json:"token_type"`   // 令牌类型
	ExpiresAt   int64  `json:"expires_at"`   // 过期时间戳
}
