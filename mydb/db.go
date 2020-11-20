package mydb

import "Backend-Api/models"

type DB interface {
	GetIDFromToken(token string) (uint, error)
	GetContacts(cellphone string) ([]*models.ContactModel, error)
	GetContactUserCell(id uint) (*string, error)
	SetToken(userID uint, token string) error
	InsertContact(userCellphone string, contact *models.ContactModel) error
	UpdateContact(contactID uint, contact *models.ContactModel) error
	DeleteContact(contactID uint) error
	InsertUser(user *models.User) error
	GetUser(userID uint) (*models.User, error)
	GetUserWithCellphone(cellphone string) (*models.User, error)
	UpdateUser(userID uint, user *models.User) error
}