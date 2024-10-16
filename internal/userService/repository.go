package userService

import "gorm.io/gorm"

type UserRepository interface {
	CreateUser(u User) (User, error)
	GetAllUsers() ([]User, error)
	UpdateUser(id uint, u User) (User, error)
	DeleteUser(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) CreateUser(u User) (User, error) {
	err := ur.db.Create(&u).Error
	if err != nil {
		return User{}, err
	}
	return u, nil
}

func (ur *userRepository) GetAllUsers() ([]User, error) {
	var users []User
	err := ur.db.Find(&users).Error
	if err != nil {
		return []User{}, err
	}
	return users, nil
}

func (ur *userRepository) UpdateUser(id uint, u User) (User, error) {
	var existingUser User
	if err := ur.db.First(&existingUser, id).Error; err != nil {
		return User{}, err
	}
	if err := ur.db.Model(&existingUser).Updates(u).Error; err != nil {
		return User{}, err
	}
	return existingUser, nil
}

func (ur *userRepository) DeleteUser(id uint) error {
	var u User
	if err := ur.db.Delete(&u, id).Error; err != nil {
		return err
	}
	return nil
}
