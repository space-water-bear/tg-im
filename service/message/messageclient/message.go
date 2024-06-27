// Code generated by goctl. DO NOT EDIT.
// Source: message.proto

package messageclient

import (
	"context"

	"im2/service/message/message"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AudioMsg              = message.AudioMsg
	CommandMsg            = message.CommandMsg
	FileMsg               = message.FileMsg
	ImageMsg              = message.ImageMsg
	LocationMsg           = message.LocationMsg
	MarkdownMsg           = message.MarkdownMsg
	MergeForwardMsg       = message.MergeForwardMsg
	MessageContent        = message.MessageContent
	MessageInfo           = message.MessageInfo
	ReceiveMessageRequest = message.ReceiveMessageRequest
	RichTextMsg           = message.RichTextMsg
	SendMessageRequest    = message.SendMessageRequest
	SendMessageResponse   = message.SendMessageResponse
	TextMsg               = message.TextMsg
	VideoMsg              = message.VideoMsg

	Message interface {
		SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error)
		ReceiveMessage(ctx context.Context, in *ReceiveMessageRequest, opts ...grpc.CallOption) (message.Message_ReceiveMessageClient, error)
	}

	defaultMessage struct {
		cli zrpc.Client
	}
)

func NewMessage(cli zrpc.Client) Message {
	return &defaultMessage{
		cli: cli,
	}
}

func (m *defaultMessage) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error) {
	client := message.NewMessageClient(m.cli.Conn())
	return client.SendMessage(ctx, in, opts...)
}

func (m *defaultMessage) ReceiveMessage(ctx context.Context, in *ReceiveMessageRequest, opts ...grpc.CallOption) (message.Message_ReceiveMessageClient, error) {
	client := message.NewMessageClient(m.cli.Conn())
	return client.ReceiveMessage(ctx, in, opts...)
}
