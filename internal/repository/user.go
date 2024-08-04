package repository

import (
	"github.com/DavidEsdrs/go-mercado/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(conn *gorm.DB) *UserRepository {
	return &UserRepository{
		db: conn,
	}
}

func (ur *UserRepository) Insert(product *model.User) error {
	return ur.db.Create(product).Error
}

func (ur *UserRepository) Read(id uint) (model.User, error) {
	var user model.User
	if err := ur.db.First(&user, id).Error; err != nil { // First founds first match
		return user, err
	}
	return user, nil
}

func (ur *UserRepository) Update(user *model.User) error {
	return ur.db.Save(user).Error // upsert
}

func (ur *UserRepository) Delete(id uint) error {
	return ur.db.Delete(&model.User{}, id).Error // deletes by id
}

func (ur *UserRepository) FindAll() (user []model.User, err error) {
	tx := ur.db.Find(&user)
	return user, tx.Error
}

func (ur *UserRepository) FindUserByEmail(email string) (model.User, error) {
	var user model.User
	tx := ur.db.Find(&user, "email = ?", email)
	return user, tx.Error
}
