package service

import (
    "context"
    v1 "github.com/BingguWang/bingfood-client-micro/api/bingfood/service/v1/pbgo/v1"
    "github.com/BingguWang/bingfood-client-micro/app/bingfood/service/internal/utils"
    "github.com/go-kratos/kratos/v2/log"
)

func (s *BingfoodServiceImpl) UserLoginOrRegister(ctx context.Context, in *v1.UserLoginOrRegisterRequest) (*v1.UserLoginOrRegisterReply, error) {
    token, err := s.ac.LoginOrRegister(ctx, in)
    if err != nil {
        return &v1.UserLoginOrRegisterReply{RetMsg: err.Error()}, err
    }
    log.Infof(utils.ToJsonString(token))
    return &v1.UserLoginOrRegisterReply{RetCode: 200, RetMsg: "登录成功 : ", Token: token}, nil
}
