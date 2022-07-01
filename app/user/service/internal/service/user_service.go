package service

import (
	"context"
	"fmt"
    v12 "github.com/go-kratos/bingfood-client-micro/api/user/service/v1"
	"github.com/go-kratos/bingfood-client-micro/app/user/service/internal/biz/user"
)

type BingfoodServiceImpl struct {
    v12.UnimplementedBingfoodServiceServer

	uc *user.UserUsecase
}

func NewBingfoodService(uc *user.UserUsecase) *BingfoodServiceImpl {
	return &BingfoodServiceImpl{uc: uc}
}

func (s *BingfoodServiceImpl) LoginOrRegister(ctx context.Context, in *v12.UserLoginOrRegisterRequest) (*v12.UserLoginOrRegisterReply, error) {
	token, err := s.uc.LoginOrRegisterUser(ctx, in)
	if err != nil {
		return &v12.UserLoginOrRegisterReply{RetMsg: err.Error()}, err
	}
	fmt.Println(token)
	return &v12.UserLoginOrRegisterReply{RetCode: 200, RetMsg: "login successfully : " + in.UserMobile, Token: token}, nil
}

func (s *BingfoodServiceImpl) SetUserPassword(ctx context.Context, in *v12.SetUserPasswordRequest) (*v12.SetUserPasswordReply, error) {
	if err := s.uc.SetPassword(ctx, in); err != nil {
		return nil, err
	}
	return &v12.SetUserPasswordReply{
		RetCode: 200,
		RetMsg:  "用户密码修改成功",
	}, nil

}
