package biz

import (
    "context"
    v1 "github.com/BingguWang/bingfood-client-micro/api/bingfood/service/v1/pbgo/v1"
    v13 "github.com/BingguWang/bingfood-client-micro/api/user/service/v1/pbgo/v1"
    "github.com/BingguWang/bingfood-client-micro/app/bingfood/service/internal/conf"
    "github.com/BingguWang/bingfood-client-micro/app/bingfood/service/internal/utils"
    "github.com/go-kratos/kratos/v2/log"
    "github.com/go-kratos/kratos/v2/middleware/auth/jwt"
    "github.com/go-kratos/kratos/v2/middleware/recovery"
    "github.com/go-kratos/kratos/v2/registry"
    "github.com/go-kratos/kratos/v2/transport/grpc"
    jwt2 "github.com/golang-jwt/jwt/v4"
    "golang.org/x/crypto/bcrypt"
)

// UserCase 专门用于处理user相关的逻辑
type UserCase struct {
    usc v13.UserServiceClient

    log *log.Helper
}

func NewUserCase(cc v13.UserServiceClient, logger log.Logger) *UserCase {
    return &UserCase{usc: cc, log: log.NewHelper(logger)}
}

type UserSrvInterface interface {
    GetUsersByCond(ctx context.Context, user *v13.GetUsersByCondRequest) (*v13.GetUsersByCondReply, error)
    RegisterUser(ctx context.Context, user *v13.RegisterUserRequest) (*v13.RegisterUserReply, error)
    Login(ctx context.Context, req *v1.UserLoginOrRegisterRequest) (string, error)
}

func (uc *UserCase) GetUsersByCond(ctx context.Context, req *v13.GetUsersByCondRequest) (*v13.GetUsersByCondReply, error) {
    // 调用远程user服务
    return uc.usc.GetUsersByCond(ctx, req)
}

func (uc *UserCase) RegisterUser(ctx context.Context, req *v13.RegisterUserRequest) (*v13.RegisterUserReply, error) {
    // 调用远程user服务
    return uc.usc.RegisterUser(ctx, req)
}

func (uc *UserCase) Login(ctx context.Context, req *v1.UserLoginOrRegisterRequest) (string, error) {
    switch req.LoginType {
    case 0:
        log.Infof("使用的手机号登录")
        return LoginByMobile(ctx, req)
    case 1:
        log.Infof("使用的账号密码登录")
        return LoginByPassword(ctx, req, uc)
    default:
        return "", v1.ErrorInvalidArgument("登录方式指定有错误")
    }
}

func LoginByPassword(ctx context.Context, userParam *v1.UserLoginOrRegisterRequest, uc *UserCase) (string, error) {
    if userParam.Password == "" {
        return "", v1.ErrorInvalidArgument("请输入密码")
    }
    // 调用user服务 获取的登录的用户
    users, err := uc.usc.GetUsersByCond(ctx, &v13.GetUsersByCondRequest{UserCond: &v13.User{UserMobile: userParam.UserMobile}})
    if err != nil {
        return "", err
    }
    if users == nil || len(users.UserList) == 0 {
        return "", v1.ErrorInvalidArgument("此用户没注册")
    }
    if err := bcrypt.CompareHashAndPassword([]byte(users.UserList[0].LoginPassword), []byte(userParam.Password)); err != nil {
        return "", v1.ErrorInvalidArgument("密码错误")
    }
    token, err := utils.CreateToken(userParam.UserMobile)
    if err != nil {
        return "", err
    }
    return token, nil
}
func LoginByMobile(ctx context.Context, userParam *v1.UserLoginOrRegisterRequest) (string, error) {
    // 生成token
    token, err := utils.CreateToken(userParam.UserMobile)
    if err != nil {
        return "", err
    }
    return token, nil
}

// NewUserSrvInterface 给其他的case用于wire 注入的，依赖注入时case之间不应该相互依赖，所以要依赖应该以接口的形式
func NewUserSrvInterface(usc v13.UserServiceClient, logger log.Logger) UserSrvInterface {
    return &UserCase{usc: usc, log: log.NewHelper(logger)}
}

func NewUserServiceClient(r registry.Discovery, ac *conf.JWT) v13.UserServiceClient {
    conn, err := grpc.DialInsecure(
        context.Background(),
        grpc.WithEndpoint("discovery:///bingfood.user.service"),
        grpc.WithDiscovery(r),
        grpc.WithMiddleware(
            recovery.Recovery(),
            jwt.Client(func(token *jwt2.Token) (interface{}, error) {
                return []byte(ac.ServiceSecretKey), nil
            }, jwt.WithSigningMethod(jwt2.SigningMethodHS256)),
        ),
    )
    if err != nil {
        panic(err)
    }
    c := v13.NewUserServiceClient(conn)
    return c
}
