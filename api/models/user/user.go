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

const userInsert = `
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

func Insert(user User) error {
	tx, err := database.Db.Beginx()
	defer tx.Rollback()
	if err != nil {
		log.Printf("transaction begin error%s", err)
		return err
	}

	_, err = tx.NamedExec(userInsert, user)
	if err != nil {
		return err
	}

	tx.Commit()
	return nil
}
