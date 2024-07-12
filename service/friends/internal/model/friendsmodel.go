package model

import (
	"context"
	"gorm.io/gorm"
)

var _ FriendsModel = (*customFriendsModel)(nil)

type (
	// FriendsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFriendsModel.
	FriendsModel interface {
		friendsModel
		customFriendsLogicModel
	}

	customFriendsModel struct {
		*defaultFriendsModel
	}

	customFriendsLogicModel interface {
		FindListByPage(ctx context.Context, conditions Friends, page, pageSize int32) ([]Friends, int64, error)
		GetFriendInfo(ctx context.Context, userId int64, friendId int64) (*Friends, error)
		PassFriend(ctx context.Context, Id int64, status int32) error
	}
)

// NewFriendsModel returns a model for the database table.
func NewFriendsModel(conn *gorm.DB) FriendsModel {
	return &customFriendsModel{
		defaultFriendsModel: newFriendsModel(conn),
	}
}

func (m *customFriendsModel) FindListByPage(ctx context.Context, conditions Friends, page, pageSize int32) ([]Friends, int64, error) {
	var total int64
	var results []Friends

	if pageSize < 1 {
		pageSize = 10
	}
	if page < 1 {
		page = 1
	}

	if err := m.conn.WithContext(ctx).Model(&Friends{}).
		Where("user_id = ?", conditions.UserId).
		Where("friend_username LIKE ? OR friend_nickname LIKE ?", "%"+conditions.FriendUsername+"%", "%"+conditions.FriendNickname+"%").Count(&total).
		Offset(int((page - 1) * pageSize)).
		Limit(int(pageSize)).
		Find(&results).Error; err != nil {
		return nil, 0, err
	}

	return results, total, nil
}

// GetFriendInfo 获取好友信息
// 参数：userId 用户ID， friendId 好友ID
func (m *customFriendsModel) GetFriendInfo(ctx context.Context, userId int64, friendId int64) (*Friends, error) {

	var result Friends
	if err := m.conn.WithContext(ctx).Model(&Friends{}).
		Where("user_id = ? and friend_id = ?", userId, friendId).
		First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (m *customFriendsModel) PassFriend(ctx context.Context, Id int64, status int32) error {

	if err := m.conn.WithContext(ctx).Model(&Friends{}).
		Where("id = ?", Id).
		Updates(map[string]interface{}{"status": status, "request_status": false}).Error; err != nil {
		return err
	}

	return nil
}
