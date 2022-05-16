package models

import (
	"errors"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	UserDetails UserDetails `gorm:"embedded" json:"user"`
	GameState   GameState   `gorm:"embedded" json:"gameState"`
	CreatedAt   time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type UserDetails struct {
	ID   string `gorm:"primary_key;" json:"id"`
	Name string `gorm:"size:255;not null;" json:"name"`
}

type GameState struct {
	GamesPlayed int32 `json:"gamesPlayed"`
	Score       int64 `json:"score"`
}

type userfriend struct {
	Id    string `json:"id"`
	Name  string `json:"named"`
	Score int64  `json:"highScore"`
}

func (u *User) PopulateUser() {
	u.UserDetails.ID = uuid.New().String()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {

	var err error
	err = db.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func FindAllUsers(db *gorm.DB) ([]User, error) {
	var err error
	var users []User
	err = db.Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return []User{}, err
	}
	return users, err
}

func (u *User) FindUserByID(db *gorm.DB, uid string) (*User, error) {
	var err error
	err = db.Model(User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("user not found")
	}
	return u, err
}

func (u *User) UpdateAUser(db *gorm.DB, uid string) (*User, error) {

	u.UpdatedAt = time.Now()
	db = db.Model(&User{}).Where("id = ?", uid).Omit("id", "name", "created_at").Update(u)
	if db.Error != nil {
		return u, db.Error
	}

	return u.FindUserByID(db, u.UserDetails.ID)
}

func (u *User) GetAllFriends(db *gorm.DB, friends []string) ([]userfriend, error) {
	var ufs []userfriend
	//db = db.Raw("SELECT id, name, score  FROM `users`  WHERE (`users`.`id` IN (?))", friends).Scan(&ufs)
	db = db.Debug().Table("users").Select("id, name, score").Find(&ufs, friends)
	if db.Error != nil {
		return ufs, db.Error
	}
	return ufs, nil
}

func (u *User) ValidateCreateUser() error {
	var err error
	if u.UserDetails.Name == "" {
		return errors.New("user name cannot be empty")
	}
	return err
}

func (u *User) ValidateUpdateUser() error {
	var err error
	if u.UserDetails.ID == "" {
		return errors.New("user id cannot be empty")
	}
	return err
}

func (state *GameState) ValidateState() error {
	var err error

	if state.GamesPlayed < 0 {
		return errors.New("games played cannot be negative")
	}

	if state.Score < 0 {
		return errors.New("score cannot be negative")
	}

	return err
}
