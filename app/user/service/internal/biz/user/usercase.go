package user

import (
    "context"
    v12 "github.com/go-kratos/bingfood-client-micro/api/user/service/v1/pbgo/v1"
    utils2 "github.com/go-kratos/bingfood-client-micro/app/user/service/internal/utils"
    "github.com/go-kratos/kratos/v2/errors"
    "github.com/go-kratos/kratos/v2/log"
    "github.com/jinzhu/copier"
)

var (
    // ErrUserNotFound is user not found.
    ErrUserNotFound = errors.NotFound(v12.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

type UserRepo interface {
    Save(context.Context, *User) error
    Update(context.Context, *User) error
    GetUsersByCond(ctx context.Context, user *User) ([]*User, error)
}

type UserUsecase struct {
    repo UserRepo

    log *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
    return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) GetUsersByCondHandler(ctx context.Context, req *v12.GetUsersByCondRequest) (userList []*v12.User, err error) {
    uc.log.WithContext(ctx).Infof("GetUsersByCondHandler args: %v", utils2.ToJsonString(req))

    var cond User
    copier.Copy(&cond, req.UserCond)
    users, err := uc.repo.GetUsersByCond(ctx, &cond)
    if err != nil {
        return nil, err
    }
    copier.CopyWithOption(&userList, &users, copier.Option{
        IgnoreEmpty: false,
        DeepCopy:    true,
    })
    uc.log.WithContext(ctx).Infof("GetUsersByCondHandler ret: %v", utils2.ToJsonString(userList))
    return
}
func (uc *UserUsecase) RegisterUserHandler(ctx context.Context, req *v12.RegisterUserRequest) (err error) {
    uc.log.WithContext(ctx).Infof("RegisterUserHandler args: %v", utils2.ToJsonString(req))

    if err = uc.repo.Save(ctx, &User{UserMobile: req.UserMobile}); err != nil {
        return v12.ErrorInternal("create user failed")
    }
    return nil
}
