package models

import (
	"errors"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type UserFriends struct {
	ID      string `gorm:"primary_key;" json:"id"`
	UserOne string `gorm:"not null" json:"user_one"`
	UserTwo string `gorm:"not null" json:"user_two"`
}

func (u *UserFriends) SaveUserFriend(db *gorm.DB) (*UserFriends, error) {

	var err error
	err = db.Create(&u).Error
	if err != nil {
		return &UserFriends{}, err
	}
	return u, nil
}

func FindFriendsByID(db *gorm.DB, uid string) ([]UserFriends, error) {
	var err error
	var uf []UserFriends
	err = db.Debug().Model(&UserFriends{}).Where("user_one = ? OR user_two = ?", uid, uid).Find(&uf).Error
	if err != nil {
		return []UserFriends{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return []UserFriends{}, errors.New("friends record not found")
	}
	return uf, err
}

func (u *UserFriends) PopulateUserFriend() {
	u.ID = uuid.New().String()
}
