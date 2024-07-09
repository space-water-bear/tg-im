package model

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

var _ UsersModel = (*customUsersModel)(nil)

type (
	// UsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUsersModel.
	UsersModel interface {
		usersModel
		customUsersLogicModel
	}

	customUsersModel struct {
		*defaultUsersModel
	}

	customUsersLogicModel interface {
		FindListByPage(ctx context.Context, conditions Users, pageIndex, pageSize int32) ([]Users, int64, error)
		FindOneByEmail(ctx context.Context, email string) (*Users, error)
	}
)

// NewUsersModel returns a model for the database table.
func NewUsersModel(conn *gorm.DB) UsersModel {
	return &customUsersModel{
		defaultUsersModel: newUsersModel(conn),
	}
}

func (m *customUsersModel) FindListByPage(ctx context.Context, conditions Users, page, pageSize int32) ([]Users, int64, error) {
	var total int64
	var results []Users

	if err := m.conn.Model(&Users{}).Where(conditions).Count(&total).Offset(int((page - 1) * pageSize)).Limit(int(pageSize)).Find(&results).Error; err != nil {
		return nil, 0, err
	}

	return results, total, nil
}

func (m *customUsersModel) FindOneByEmail(ctx context.Context, email string) (*Users, error) {
	var res Users
	err := m.conn.Where("email = ?", email).First(&res).Error
	switch {
	case err == nil:
		return &res, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
