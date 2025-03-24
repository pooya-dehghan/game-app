package mysql

import (
	"database/sql"
	"fmt"

	"github.com/pooya-dehghan/entity"
)

func (d *MySQLDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	row := d.db.QueryRow(`select * from users where phone_number = ?`, phoneNumber)

	user, err := scanUser(row)

	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}
	}

	if user.ID != 0 {
		return false, nil
	}

	return true, nil
}

func (d *MySQLDB) RegisterUser(user entity.User) (createdUser entity.User, err error) {
	res, err := d.db.Exec(`insert into users(name , phone_number, hashed_password) values(? , ? , ?)`, user.Name, user.PhoneNumber, user.HashedPassword)
	if err != nil {
		return entity.User{}, fmt.Errorf("cant execute command: %w", err)
	}

	id, _ := res.LastInsertId()

	user.ID = uint(id)

	return user, nil

}

func (d *MySQLDB) FindUserByPhoneNumber(phoneNumber string) (entity.User, error) {
	row := d.db.QueryRow(`select * from users where phone_number = ?`, phoneNumber)

	user, err := scanUser(row)
	fmt.Println(user)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, nil
		}
	}

	return user, nil
}

func (d *MySQLDB) FindUserByID(id uint) (entity.User, error) {
	row := d.db.QueryRow(`select * from users where id = ?`, id)

	user, err := scanUser(row)
	fmt.Println(user)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, nil
		}
	}

	return user, nil
}

func scanUser(row *sql.Row) (entity.User, error) {
	var createdAt []uint8
	var user entity.User
	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.HashedPassword, &createdAt)

	return user, err
}
