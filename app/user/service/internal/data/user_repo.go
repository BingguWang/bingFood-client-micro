package data

import (
	"context"
    "github.com/go-kratos/bingfood-client-micro/api/user/service/v1"
	u "github.com/go-kratos/bingfood-client-micro/app/user/service/internal/biz/user"
	"github.com/go-kratos/bingfood-client-micro/app/user/service/internal/utils"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewUserRepo(data *Data, logger log.Logger) u.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) Save(ctx context.Context, u *u.User) error {
	r.log.WithContext(ctx).Infof("insert user , user: %v", u.UserMobile)

	db := r.data.db
	if err := db.Select("user_mobile").Create(&u).Error; err != nil {
		r.log.WithContext(ctx).Infof("insert into db failed : %v", err.Error())
		return err
	}
	return nil
}

func (r *userRepo) Update(ctx context.Context, user *u.User) error {
	db := r.data.db

	if err := db.Transaction(func(tx *gorm.DB) error {
		txx := tx.Model(&u.User{}).Where("user_mobile = ?", user.UserMobile).
			Updates(map[string]interface{}{"login_password": user.LoginPassword})
		if err := txx.Error; err != nil {
			return err
		}
		if rows := txx.RowsAffected; rows == 0 {
			return v1.ErrorUserNotFound("user of the userMobile is not exist")
		}
		return nil
	}); err != nil {
		r.log.WithContext(ctx).Infof("update user , user: %v", user.UserMobile)
		return err
	}
	return nil
}

func (r *userRepo) FindByID(context.Context, int64) (*u.User, error) {
	return nil, nil
}

func (r *userRepo) ListByHello(context.Context, string) ([]*u.User, error) {
	return nil, nil
}

func (r *userRepo) ListAll(context.Context) ([]*u.User, error) {
	return nil, nil
}

func (r *userRepo) GetUserByCond(ctx context.Context, user *u.User) ([]*u.User, error) {
	r.log.WithContext(ctx).Infof("exec GetUserByCond cond: %v", utils.ToJsonString(user))

	db := r.data.db

	var userEntitys []*u.User
	if err := db.Where("user_mobile = ?", user.UserMobile).Find(&userEntitys).Error; err != nil {
		return nil, errors.New(400, "select from db err", "failed")
	}
	r.log.WithContext(ctx).Infof("GetUserByCond: select result : %v", utils.ToJsonString(userEntitys))

	return userEntitys, nil
}
