package category

import (
	"context"

	"github.com/google/wire"
	"gorm.io/gorm"

	"github.com/mlsjla/gin-nuxt/server/internal/app/dao/util"
	"github.com/mlsjla/gin-nuxt/server/internal/app/schema"
	"github.com/mlsjla/gin-nuxt/server/pkg/errors"
)

// Injection wire
var CategorySet = wire.NewSet(wire.Struct(new(CategoryRepo), "*"))

type CategoryRepo struct {
	DB *gorm.DB
}

func (a *CategoryRepo) getQueryOption(opts ...schema.CategoryQueryOptions) schema.CategoryQueryOptions {
	var opt schema.CategoryQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

func (a *CategoryRepo) Query(ctx context.Context, params schema.CategoryQueryParam, opts ...schema.CategoryQueryOptions) (*schema.CategoryQueryResult, error) {
	opt := a.getQueryOption(opts...)

	db := GetCategoryDB(ctx, a.DB)

	// TODO: Your where condition code here...

	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	if len(opt.OrderFields) > 0 {
		db = db.Order(util.ParseOrder(opt.OrderFields))
	}

	var list Categories
	pr, err := util.WrapPageQuery(ctx, db, params.PaginationParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	qr := &schema.CategoryQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaCategories(),
	}

	return qr, nil
}

func (a *CategoryRepo) Get(ctx context.Context, id uint64, opts ...schema.CategoryQueryOptions) (*schema.Category, error) {
	var item Category
	ok, err := util.FindOne(ctx, GetCategoryDB(ctx, a.DB).Where("id=?", id), &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return item.ToSchemaCategory(), nil
}

func (a *CategoryRepo) Create(ctx context.Context, item schema.Category) error {
	eitem := SchemaCategory(item).ToCategory()
	result := GetCategoryDB(ctx, a.DB).Create(eitem)
	return errors.WithStack(result.Error)
}

func (a *CategoryRepo) Update(ctx context.Context, id uint64, item schema.Category) error {
	eitem := SchemaCategory(item).ToCategory()
	result := GetCategoryDB(ctx, a.DB).Where("id=?", id).Updates(eitem)
	return errors.WithStack(result.Error)
}

func (a *CategoryRepo) Delete(ctx context.Context, id uint64) error {
	result := GetCategoryDB(ctx, a.DB).Where("id=?", id).Delete(Category{})
	return errors.WithStack(result.Error)
}
