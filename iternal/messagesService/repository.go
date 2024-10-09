package messagesService

import (

	"gorm.io/gorm"
)

type MessageRepository interface {
	CreateMessage(message Message) (Message, error)
	GetAllMessages() ([]Message, error)
	UpdateMessage(id uint, message Message) (Message, error)
	DeleteMessage(id uint) error
}

type messagesRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *messagesRepository {
	return &messagesRepository{db: db}
}

func (r *messagesRepository) CreateMessage(message Message) (Message, error) {
	result := r.db.Create(&message)
	if result.Error != nil {
		return Message{}, result.Error
	}
	return message, nil
}

func (r *messagesRepository) GetAllMessages() ([]Message, error) {
	var messages []Message
	err := r.db.Find(&messages).Error
	return messages, err
}

func (r *messagesRepository) UpdateMessage(id uint, message Message) (Message, error) {
	var existingMessage Message

	if err := r.db.First(&existingMessage, id).Error; err != nil {
		return Message{}, err
	}
	if err := r.db.Model(&existingMessage).Updates(message).Error; err != nil {
		return Message{}, err
	}
	return existingMessage, nil
}

func (r *messagesRepository) DeleteMessage(id uint) error {
	var message Message
	if err := r.db.Delete(&message, id).Error; err != nil {
		return err
	}
	return nil
}
