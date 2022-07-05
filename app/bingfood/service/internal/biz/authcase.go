package biz

import (
    "context"
    v1 "github.com/go-kratos/bingfood-client-micro/api/bingfood/service/v1/pbgo/v1"
    v13 "github.com/go-kratos/bingfood-client-micro/api/user/service/v1/pbgo/v1"
    "github.com/go-kratos/bingfood-client-micro/app/bingfood/service/internal/utils"
    "github.com/go-kratos/kratos/v2/log"
)

// AuthCase用于解处理认证逻辑
type AuthCase struct {
    key string
    //uc  UserCase // 依赖其他的case是不行的，因为注入的同时其他的case也还没有注入进来
    usInter UserSrvInterface

    log *log.Helper
}

func NewAuthCase(usInter UserSrvInterface, logger log.Logger) *AuthCase {
    return &AuthCase{usInter: usInter, log: log.NewHelper(logger)}
}

func (ac *AuthCase) LoginOrRegister(ctx context.Context, req *v1.UserLoginOrRegisterRequest) (token string, err error) {
    ac.log.WithContext(ctx).Infof("LoginOrRegisterUser args: %v", utils.ToJsonString(req))

    if req.UserMobile == "" {
        err = v1.ErrorInvalidArgument("手机号不能为空")
        return
    }
    rq := &v13.GetUsersByCondRequest{UserCond: &v13.User{
        UserMobile: req.UserMobile,
    }}
    callRes, err := ac.usInter.GetUsersByCond(ctx, rq)
    if err != nil {
        return "", err
    }
    log.Infof("调用服务bingfood.user.service/GetUsersByCond, 得到结果: %v ", utils.ToJsonString(callRes))

    if len(callRes.UserList) == 0 { // 不存在就注册
        log.Infof("手机未注册,将自动注册并登录")
        if _, err = ac.usInter.RegisterUser(ctx, &v13.RegisterUserRequest{UserMobile: req.UserMobile}); err != nil {
            return "", err
        }
        token, err = ac.usInter.Login(ctx, &v1.UserLoginOrRegisterRequest{
            LoginType:  0,
            UserMobile: req.UserMobile,
            ValidCode:  req.ValidCode,
        })
        if err != nil {
            return "", err
        }
        return token, nil
    }

    return ac.usInter.Login(ctx, req)
}
