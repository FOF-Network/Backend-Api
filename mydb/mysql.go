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
	fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8&collation=utf8_unicode_ci&readTimeout=%s",
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

func (db *MySql) GetContacts(cellphone string) ([]*models.ContactModel, error) {
	stmt, err := db.Connection.Prepare(`select first_name, last_name, cellphone, email, birth_day, birth_month, job, interests, city_name from contacts where user_cellphone = ?`)
	if err != nil {
		return nil, err
	}

	contacts := make([]*models.ContactModel, 0, 10)
	rows, err := stmt.Query(cellphone)
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
			&contact.BirthMonth,
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

func (db *MySql) GetContactUserCell(id uint) (*string, error) {
	stmt, err := db.Connection.Prepare(`select user_cellphone from contacts where id = ?`)
	if err != nil {
		return nil, err
	}

	str := new(string)
	err = stmt.QueryRow(id).Scan(str)
	if err != nil {
		return nil, err
	}

	return str, nil
}


func (db *MySql) InsertContact(userCellphone string, contact *models.ContactModel) error {
	stmt, err := db.Connection.Prepare(`insert into contacts (user_cellphone, first_name, last_name, cellphone, birth_day, birth_month, email, job, interests, city_name, created_at, updated_at) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, current_timestamp, current_timestamp)`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(userCellphone, contact.FirstName, contact.LastName, contact.Cellphone, contact.BirthDay, contact.BirthMonth, contact.Email, contact.Job, contact.Interests, contact.CityName)
	return err
}

func (db *MySql) UpdateContact(contactID uint, contact *models.ContactModel) error {
	stmt, err := db.Connection.Prepare(`update contacts set first_name = ?, last_name = ?, cellphone = ?, birth_day = ?, birth_month = ?, email = ?, job = ?, interests = ?, city_name = ?, updated_at = current_timestamp where id = ?`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(contact.FirstName, contact.LastName, contact.Cellphone, contact.BirthDay, contact.BirthMonth, contact.Email, contact.Job, contact.Interests, contact.CityName, contactID)
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
	stmt, err := db.Connection.Prepare(`insert into users (first_name, last_name, cellphone, password, birth_day, birth_month, email, job, interests, city_name, created_at, updated_at) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, current_timestamp, current_timestamp)`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Cellphone, user.Password, user.BirthDay, user.BirthMonth, user.Email, user.Job, user.Interests, user.CityName)
	return err
}

func (db *MySql) GetUser(userID uint) (*models.User, error) {
	stmt, err := db.Connection.Prepare(`select cellphone from users where id = ?`)
	if err != nil {
		return nil, err
	}

	user := new(models.User)
	err = stmt.QueryRow(userID).Scan(&user.Cellphone)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (db *MySql) GetUserWithCellphone(cellphone string) (*models.User, error) {
	stmt, err := db.Connection.Prepare(`select id, password from users where cellphone = ?`)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	user := new(models.User)
	err = stmt.QueryRow(cellphone).Scan(&user.ID, &user.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (db *MySql) UpdateUser(userID uint, user *models.User) error {
	stmt, err := db.Connection.Prepare(`update users set first_name = ?, last_name = ?, cellphone = ?, password = ?, birth_day = ?, birth_month = ?, email = ?, job = ?, interests = ?, city_name = ?, updated_at = current_timestamp where id = ?`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Cellphone,  user.Password, user.BirthDay, user.BirthMonth, user.Email, user.Job, user.Interests, user.CityName, userID)
	return err
}