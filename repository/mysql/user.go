package mysql

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/pooya-dehghan/entity"
)

func (d *MySQLDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	user := entity.User{}
	var createdAt time.Time
	row := d.db.QueryRow(`select * from users where phone_number = ?`, phoneNumber)

	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.HashedPassword, &createdAt)
	fmt.Println(user)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
	}

	return true, nil
}

func (d *MySQLDB) RegisterUser(user entity.User) (createdUser entity.User, err error) {
	res, err := d.db.Exec(`insert into users(name , phone_number, hashed_password) valuee(? , ? , ?)`, user.Name, user.PhoneNumber, user.HashedPassword)
	if err != nil {
		return entity.User{}, fmt.Errorf("cant execute command: %w", err)
	}

	id, _ := res.LastInsertId()

	user.ID = uint(id)

	return user, nil

}

func (d *MySQLDB) FindUserByPhoneNumber(phoneNumber string) (entity.User, error) {
	user := entity.User{}
	var createdAt time.Time
	row := d.db.QueryRow(`select * from users where phone_number = ?`, phoneNumber)

	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.HashedPassword, &createdAt)
	fmt.Println(user)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, nil
		}
	}

	return user, nil
}
