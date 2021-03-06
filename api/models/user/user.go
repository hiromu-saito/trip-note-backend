package user

import (
	"log"

	"github.com/hiromu-saito/trip-note-backend/database"
)

type User struct {
	Id       int    `db:"id"`
	Email    string `db:"email"`
	Password []byte `db:"password"`
}

const userInsertSql = `
insert into users (
	id
 ,email
 ,password
)
select
 case
    when max(id) is null then 1
    else max(id)+1
  end
 ,:email
 ,:password
from
    users;
`

const selectByIdSql = `
select
	*
from
	users
where
	id = ?
`

const selectByEmailSql = `
select
	*
from
	users
where
	email = ?
`

func Insert(user User) error {
	tx, err := database.Db.Beginx()
	defer tx.Rollback()
	if err != nil {
		log.Printf("transaction begin error%s", err)
		return err
	}

	_, err = tx.NamedExec(userInsertSql, user)
	if err != nil {
		log.Printf("user insert error%s", err)
		return err
	}

	tx.Commit()
	return nil
}

func SelectById(id int) (User, error) {
	db := database.Db
	var user User

	if err := db.Get(&user, selectByIdSql, id); err != nil {
		return user, err
	}
	return user, nil
}

func SelectByEmail(email string) (User, error) {
	var user User
	if err := database.Db.Get(&user, selectByEmailSql, email); err != nil {
		log.Printf("user SelectByEmail error:%s\n", err)
		return user, err
	}
	return user, nil
}
