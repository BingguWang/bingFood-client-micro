package service

import (
    "context"
    v12 "github.com/BingguWang/bingfood-client-micro/api/user/service/v1/pbgo/v1"
    "github.com/BingguWang/bingfood-client-micro/app/user/service/internal/biz/user"
)

type UserServiceImpl struct {
    v12.UnimplementedUserServiceServer

    uc *user.UserUsecase
}

func NewUserServiceImpl(uc *user.UserUsecase) *UserServiceImpl {
    return &UserServiceImpl{uc: uc}
}

func (s *UserServiceImpl) GetUsersByCond(ctx context.Context, in *v12.GetUsersByCondRequest) (*v12.GetUsersByCondReply, error) {
    users, err := s.uc.GetUsersByCondHandler(ctx, in)
    if err != nil {
        return &v12.GetUsersByCondReply{RetMsg: err.Error()}, err
    }
    return &v12.GetUsersByCondReply{RetCode: 200, RetMsg: "call GetUsersByCond successfully", UserList: users}, nil
}

func (s *UserServiceImpl) RegisterUser(ctx context.Context, in *v12.RegisterUserRequest) (*v12.RegisterUserReply, error) {
    if err := s.uc.RegisterUserHandler(ctx, in); err != nil {
        return &v12.RegisterUserReply{RetMsg: err.Error()}, err
    }
    return &v12.RegisterUserReply{RetCode: 200, RetMsg: "call RegisterUserHandler successfully"}, nil
}
