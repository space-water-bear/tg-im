package message

import (
	"context"
	"encoding/json"
	"im2/internal/errno"
	"im2/internal/utils"
	"im2/restful/gateway/internal/svc"
	"im2/restful/gateway/internal/types"
	"im2/service/message/message"
	"strconv"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type MessageSendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Send message
func NewMessageSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageSendLogic {
	return &MessageSendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessageSendLogic) MessageSend(req *types.MessageInfoRequest) (resp *types.MessageInfoResponse, err error) {

	userIdJson := l.ctx.Value("id")
	userId, _ := userIdJson.(json.Number).Int64()

	var msgContent = new(message.MessageContent)

	switch req.Types {
	case "text":
		msgContent.MessageType = message.MessageType_TEXT
		msgContent.Content = &message.MessageContent_TextMsg{
			TextMsg: &message.TextMsg{
				Content: req.Content,
			},
		}

	case "rich_text":
		msgContent.MessageType = message.MessageType_RICH_TEXT
		msgContent.Content = &message.MessageContent_RichTextMsg{
			RichTextMsg: &message.RichTextMsg{
				HtmlContent: req.Content,
			},
		}

	case "markdown":
		msgContent.MessageType = message.MessageType_MARKDOWN
		msgContent.Content = &message.MessageContent_MarkdownMsg{
			MarkdownMsg: &message.MarkdownMsg{
				MarkdownContent: req.Content,
			},
		}

	case "location":
		parts := strings.Split(req.Content, ",")
		longitude, _ := strconv.ParseFloat(parts[0], 64)
		latitude, _ := strconv.ParseFloat(parts[1], 64)
		msgContent.MessageType = message.MessageType_LOCATION
		msgContent.Content = &message.MessageContent_LocationMsg{
			LocationMsg: &message.LocationMsg{
				Latitude:  latitude,
				Longitude: longitude,
				Address:   parts[2],
			},
		}
	case "command":
		parts := strings.Split(req.Content, ",")
		msgContent.MessageType = message.MessageType_COMMAND
		msgContent.Content = &message.MessageContent_CommandMsg{
			CommandMsg: &message.CommandMsg{
				Command: parts[0],
				Parameters: utils.TernaryIF(len(parts) == 3, map[string]string{
					"param1": parts[1],
					"param2": parts[2],
				}, nil),
			},
		}
	case "merge":
		return nil, nil
	default:
		return nil, errno.NewDefaultError(errno.InternalServerError)
	}

	msg := &message.SendMessageRequest{
		SenderId:       userId,
		ReceiverId:     req.ReceiverId,
		ChatType:       utils.TernaryIF(req.ChatType == 0, message.ChatType_PRIVATE, message.ChatType_GROUP),
		MessageContent: msgContent,
	}
	rest, err := l.svcCtx.MessageRpc.SendMessage(l.ctx, msg)
	if err != nil {
		return nil, errno.NewDefaultError(errno.InternalServerError)
	} else if rest.Code != errno.OK {
		return nil, errno.NewDefaultError(rest.Code)
	}

	return &types.MessageInfoResponse{
		Success: true,
	}, nil
}
