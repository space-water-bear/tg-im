package message

import (
	"context"
	"encoding/json"
	"fmt"
	"im2/internal/errno"
	"net/http"
	"time"

	"im2/restful/gateway/internal/svc"
	"im2/restful/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MessageFileSendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

const maxFileSize = 10 << 20

// Send file type message
func NewMessageFileSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageFileSendLogic {
	return &MessageFileSendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessageFileSendLogic) MessageFileSend(r *http.Request) (resp *types.MessageInfoResponse, err error) {
	userIdJson := l.ctx.Value("id")
	userId, _ := userIdJson.(json.Number).Int64()

	_ = r.ParseMultipartForm(maxFileSize)
	file, handler, err := r.FormFile("file")
	if err != nil {
		return nil, errno.NewDefaultError(errno.ErrFileTooLarge)
	}
	defer file.Close()

	fmt.Println("upload File: ", handler.Filename)
	fmt.Println("upload File size: ", handler.Size)
	fmt.Println("upload File ContentType: ", handler.Header)

	filename := fmt.Sprintf("%d_%s_%s", userId, time.Now().Format("20060102150405"), handler.Filename)

	if err := l.svcCtx.Minio.Upload(l.ctx, filename, file, handler.Size); err != nil {
		return nil, errno.NewDefaultError(errno.ErrFileSave)
	}

	return &types.MessageInfoResponse{
		Success: true,
	}, nil
}
