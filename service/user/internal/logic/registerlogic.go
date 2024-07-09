package logic

import (
	"context"
	"fmt"
	"im2/internal/errno"
	"im2/internal/utils"
	"im2/service/user/internal/model"

	"im2/service/user/internal/svc"
	"im2/service/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {

	_, err := l.svcCtx.Model.FindOneByUsername(l.ctx, in.Username)
	if err == nil {
		return &user.RegisterResponse{Code: errno.ErrUserExist}, nil
	}

	_, err = l.svcCtx.Model.FindOneByEmail(l.ctx, in.Email)
	if err == nil {
		return &user.RegisterResponse{Code: errno.ErrEmailExist}, nil
	}

	encodePass, _ := utils.Encrypt(in.Password)

	u := model.Users{Username: in.Username, Password: encodePass, Email: in.Email}
	err = l.svcCtx.Model.Insert(l.ctx, nil, &u)
	if err != nil {
		fmt.Println("err: ", err)
		return &user.RegisterResponse{Code: errno.ErrDatabase}, nil
	}

	logx.Infof("user %s register success", in.Username)
	return &user.RegisterResponse{Code: errno.OK}, nil
}
