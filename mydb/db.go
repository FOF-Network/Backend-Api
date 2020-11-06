package mydb

import "Backend-Api/models"

type DB interface {
	GetIDFromToken(token string) (uint, error)
	GetContacts(id uint) ([]*models.ContactModel, error)
	GetContact(id uint) (*models.ContactModel, error)
	GetCity(id uint) ()
	SetToken(userID uint) error
	InsertContact(userID uint, contact *models.ContactModel) error
	UpdateContact(contactID uint, contact *models.ContactModel) error
	DeleteContact(contactID uint) error
	InsertUser(user *models.User) error
	UpdateUser(userID uint, user *models.User) error
}