package user

import (
	"context"
	v12 "github.com/go-kratos/bingfood-client-micro/api/user/service/v1"
	utils2 "github.com/go-kratos/bingfood-client-micro/app/user/service/internal/utils"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"

	"golang.org/x/crypto/bcrypt"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v12.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

type UserRepo interface {
    Save(context.Context, *User) error
    Update(context.Context, *User) error
    //FindByID(context.Context, int64) (*User, error)
    //ListByHello(context.Context, string) ([]*User, error)
    //ListAll(context.Context) ([]*User, error)
    GetUserByCond(ctx context.Context, user *User) ([]*User, error)
}

type UserUsecase struct {
	repo UserRepo

	log *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) LoginOrRegisterUser(ctx context.Context, req *v12.UserLoginOrRegisterRequest) (token string, err error) {
	uc.log.WithContext(ctx).Infof("LoginOrRegisterUser args: %v", utils2.ToJsonString(req))

	if req.UserMobile == "" {
		err = v12.ErrorInvalidArgument("手机号不能为空")
		return
	}
	u := &User{UserMobile: req.UserMobile}

	var users []*User
	if users, err = uc.repo.GetUserByCond(ctx, u); err != nil {
		return
	}
	uc.log.WithContext(ctx).Infof("search ret: %v", utils2.ToJsonString(users))

	if len(users) == 0 {
		return uc.Register(ctx, req)
	} else {
		return uc.Login(ctx, req)
	}
	return
}

func (uc *UserUsecase) Register(ctx context.Context, userParam *v12.UserLoginOrRegisterRequest) (string, error) {
	if err := uc.repo.Save(ctx, &User{UserMobile: userParam.UserMobile}); err != nil {
		return "", v12.ErrorInternal("create user failed")
	}
	return LoginByMobile(ctx, userParam)
}

func (uc *UserUsecase) Login(ctx context.Context, userParam *v12.UserLoginOrRegisterRequest) (string, error) {
	switch userParam.LoginType {
	case 0:
		return LoginByMobile(ctx, userParam)
	case 1:
		return LoginByPassword(ctx, userParam, uc)
	default:
		return "", v12.ErrorInvalidArgument("登录方式指定有错误")
	}
}

func LoginByMobile(ctx context.Context, userParam *v12.UserLoginOrRegisterRequest) (string, error) {
	// 生成token
	token, err := utils2.CreateToken(userParam.UserMobile)
	if err != nil {
		return "", err
	}
	return token, nil
}
func LoginByPassword(ctx context.Context, userParam *v12.UserLoginOrRegisterRequest, uc *UserUsecase) (string, error) {
	if userParam.Password == "" {
		return "", v12.ErrorInvalidArgument("请输入密码")
	}
	// 生成token
	ur, err := uc.repo.GetUserByCond(ctx, &User{
		UserMobile: userParam.UserMobile,
	})
	if err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(ur[0].LoginPassword), []byte(userParam.Password)); err != nil {
		return "", v12.ErrorPasswordFalse("密码错误")
	}
	token, err := utils2.CreateToken(userParam.UserMobile)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (uc *UserUsecase) SetPassword(ctx context.Context, param *v12.SetUserPasswordRequest) (err error) {
	uc.log.WithContext(ctx).Infof("SetPassword args: %v", utils2.ToJsonString(param))
	if param.Password == "" || param.UserMobile == "" {
		err = v12.ErrorInvalidArgument("用户号码和密码都不能为空")
		return
	}
	var pwd []byte
	if pwd, err = bcrypt.GenerateFromPassword([]byte(param.Password), bcrypt.DefaultCost); err != nil {
		return err
	}

	if err = uc.repo.Update(ctx, &User{UserMobile: param.UserMobile, LoginPassword: string(pwd)}); err != nil {
		return
	}
	return
}
