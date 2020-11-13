package mydb

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"Backend-Api/models"
	"fmt"
)

type MySql struct {
	Connection *sql.DB
}

func New(Env map[string]string) (*MySql, error) {
	db, err := sql.Open("mysql",
	fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&collation=utf8_unicode_ci&readTimeout=%s",
	Env["DB_USER"],
	Env["DB_PASS"],
	Env["DB_HOST"],
	Env["DB_PORT"],
	Env["DB_NAME"],
	Env["DB_READ_TIMEOUT"],
	))

	if err != nil {
		return nil, err
	}

	return &MySql{Connection: db}, nil
}

func (db *MySql) SetToken(userID uint, token string) error {
	stmt, err := db.Connection.Prepare(`insert into user_tokens (user_id, token) values (?, ?)`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(userID, token)

	if err != nil {
		return err
	}
	return nil
}

func (db *MySql) GetIDFromToken(Token string) (uint, error) {
	stmt, err := db.Connection.Prepare(`select user_id from user_tokens where token = ?`)
	if err != nil {
		return 0, err
	}

	var id uint
	err = stmt.QueryRow(Token).Scan(&id)

	if err != nil {
		return 0, err
	}
	return id, nil
}

func (db *MySql) GetContacts(id uint) ([]*models.ContactModel, error) {
	stmt, err := db.Connection.Prepare(`select * from contacts where user_id = ?`)
	if err != nil {
		return nil, err
	}

	contacts := make([]*models.ContactModel, 0, 10)
	rows, err := stmt.Query(id)
	if err != nil {
		return contacts, err
	}

	for rows.Next() {
		contact := new(models.ContactModel)
		err = rows.Scan(
			&contact.FirstName, 
			&contact.LastName, 
			&contact.Cellphone,
			&contact.Email, 
			&contact.BirthDay, 
			&contact.Job, 
			&contact.Interests, 
			&contact.CityName)
		if err != nil {
			return contacts, err
		}
		contacts = append(contacts, contact)
	}

	return contacts, rows.Err()
}

func (db *MySql) GetContact(id uint) (*models.ContactModel, error) {
	stmt, err := db.Connection.Prepare(`select * from contacts where user_id = ?`)
	if err != nil {
		return nil, err
	}

	contact := new(models.ContactModel)
	err = stmt.QueryRow(id).Scan(contact)
	if err != nil {
		return nil, err
	}

	return contact, nil
}


func (db *MySql) InsertContact(userID uint, contact *models.ContactModel) error {
	stmt, err := db.Connection.Prepare(`insert into contacts (first_name, last_name, cellphone, birth_day, email, job, interests, city_name, created_at, updated_at) values (?, ?, ?, ?, ?, ?, ?, current_timestamp, current_timestamp)`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(contact.FirstName, contact.LastName, contact.Cellphone, contact.BirthDay, contact.Email, contact.Job, contact.Interests, contact.CityName)
	return err
}

func (db *MySql) UpdateContact(contactID uint, contact *models.ContactModel) error {
	stmt, err := db.Connection.Prepare(`update contacts set first_name = ?, last_name = ?, cellphone = ?, birth_day = ?, email = ?, job = ?, interests = ?, city_name = ?, updated_at = current_timestamp where id = ?`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(contact.FirstName, contact.LastName, contact.Cellphone, contact.BirthDay, contact.Email, contact.Job, contact.Interests, contact.CityName, contactID)
	return err
}

func (db *MySql) DeleteContact(contactID uint) error {
	stmt, err := db.Connection.Prepare(`delete from contacts where id = ?`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(contactID)
	return err

}

func (db *MySql) InsertUser(user *models.User) error {
	stmt, err := db.Connection.Prepare(`insert into users (first_name, last_name, cellphone, password, birth_day, email, job, interests, city_name, created_at, updated_at) values (?, ?, ?, ?, ?, ?, ?, current_timestamp, current_timestamp)`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Cellphone, user.Password, user.BirthDay, user.Email, user.Job, user.Interests, user.CityName)
	return err
}

func (db *MySql) GetUser(userID uint) (*models.User, error) {
	stmt, err := db.Connection.Prepare(`select * from users where id = ?`)
	if err != nil {
		return nil, err
	}

	user := new(models.User)
	err = stmt.QueryRow(userID).Scan(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (db *MySql) UpdateUser(userID uint, user *models.User) error {
	stmt, err := db.Connection.Prepare(`update users set first_name = ?, last_name = ?, cellphone = ?, password = ?, birth_day = ?, email = ?, job = ?, interests = ?, city_name = ?, updated_at = current_timestamp where id = ?`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Cellphone, user.BirthDay, user.Email, user.Job, user.Interests, user.CityName, userID)
	return err
}