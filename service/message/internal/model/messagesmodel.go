package model

import (
	"gorm.io/gorm"
)

var _ MessagesModel = (*customMessagesModel)(nil)

type (
	// MessagesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMessagesModel.
	MessagesModel interface {
		messagesModel
		customMessagesLogicModel
	}

	customMessagesModel struct {
		*defaultMessagesModel
	}

	customMessagesLogicModel interface {
	}
)

// NewMessagesModel returns a model for the database table.
func NewMessagesModel(conn *gorm.DB) MessagesModel {
	return &customMessagesModel{
		defaultMessagesModel: newMessagesModel(conn),
	}
}
