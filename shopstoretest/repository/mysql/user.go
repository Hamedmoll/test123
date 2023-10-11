package mysql

import (
	"fmt"
	"shopstoretest/entity"
)

func (mysql MySQLDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	return true, nil
}

func (mysql MySQLDB) Register(user entity.User) (entity.User, error) {
	//fmt.Println("it is important", mysql.db)

	res, err := mysql.DB.Exec("insert into users(name, hashed_password, phone_number) values(?, ?, ?)",
		user.Name, user.Password, user.PhoneNumber)
	//res, err := shopstore_dbmysql.db.Exec(query, user.Name, user.Password, user.PhoneNumber)
	if err != nil {

		//fmt.Println("aaaaaaaaaaaaaaaaaaaa")
		return entity.User{}, err
	}

	id, iErr := res.LastInsertId()
	if iErr != nil {

		return entity.User{}, fmt.Errorf("cant get last id after save new user in table %w", iErr)
	}

	user.ID = uint(id)

	return user, nil
}
