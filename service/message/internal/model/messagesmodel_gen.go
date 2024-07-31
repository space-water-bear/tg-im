// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"github.com/SpectatorNan/gorm-zero/gormc"
	"time"

	"gorm.io/gorm"
)

type (
	messagesModel interface {
		Insert(ctx context.Context, tx *gorm.DB, data *Messages) error

		FindOne(ctx context.Context, id int64) (*Messages, error)
		FindOneBySenderIdTimestamp(ctx context.Context, senderId int64, timestamp time.Time) (*Messages, error)
		Update(ctx context.Context, tx *gorm.DB, data *Messages) error

		Delete(ctx context.Context, tx *gorm.DB, id int64) error
		Transaction(ctx context.Context, fn func(db *gorm.DB) error) error
	}

	defaultMessagesModel struct {
		conn  *gorm.DB
		table string
	}

	Messages struct {
		Id         int64          `gorm:"column:id"`          // 消息ID，自增主键
		SenderId   int64          `gorm:"column:sender_id"`   // 发送者用户ID
		ReceiverId int64          `gorm:"column:receiver_id"` // 接收者用户ID
		Content    sql.NullString `gorm:"column:content"`     // 消息内容
		FileUrl    sql.NullString `gorm:"column:file_url"`    // 文件URL
		FileType   sql.NullString `gorm:"column:file_type"`   // 文件类型，如audio、video、image等
		Timestamp  time.Time      `gorm:"column:timestamp"`   // 消息时间戳
		Status     int64          `gorm:"column:status"`      // 消息状态，0：已发送，1：已送达，2：已读
	}
)

func (Messages) TableName() string {
	return `"public"."messages"`
}

func newMessagesModel(conn *gorm.DB) *defaultMessagesModel {
	return &defaultMessagesModel{
		conn:  conn,
		table: `"public"."messages"`,
	}
}

func (m *defaultMessagesModel) Insert(ctx context.Context, tx *gorm.DB, data *Messages) error {
	db := m.conn
	if tx != nil {
		db = tx
	}
	err := db.WithContext(ctx).Save(&data).Error
	return err
}

func (m *defaultMessagesModel) FindOne(ctx context.Context, id int64) (*Messages, error) {
	var resp Messages
	err := m.conn.WithContext(ctx).Model(&Messages{}).Where("id = ?", id).Take(&resp).Error
	switch err {
	case nil:
		return &resp, nil
	case gormc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultMessagesModel) FindOneBySenderIdTimestamp(ctx context.Context, senderId int64, timestamp time.Time) (*Messages, error) {
	var resp Messages
	err := m.conn.WithContext(ctx).Model(&Messages{}).Where("sender_id = $1 and timestamp = $2", senderId, timestamp).Take(&resp).Error
	switch err {
	case nil:
		return &resp, nil
	case gormc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultMessagesModel) Update(ctx context.Context, tx *gorm.DB, data *Messages) error {
	db := m.conn
	if tx != nil {
		db = tx
	}
	err := db.WithContext(ctx).Save(data).Error
	return err
}

func (m *defaultMessagesModel) Delete(ctx context.Context, tx *gorm.DB, id int64) error {
	db := m.conn
	if tx != nil {
		db = tx
	}
	err := db.WithContext(ctx).Delete(&Messages{}, id).Error

	return err
}

func (m *defaultMessagesModel) Transaction(ctx context.Context, fn func(db *gorm.DB) error) error {
	return m.conn.WithContext(ctx).Transaction(fn)
}
