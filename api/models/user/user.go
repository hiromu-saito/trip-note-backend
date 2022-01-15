package user

import (
	"github.com/jmoiron/sqlx"
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

func Insert(user User, tx sqlx.Tx) error {
	if _, err := tx.NamedExec(userInsert, user); err != nil {
		return err
	}
	return nil
}
