package user

import (
	"github.com/jmoiron/sqlx"
)

type User struct {
	Id       int
	Email    string
	Password []byte
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
 ,?
 ,?
from
    users;
`

func Insert(user User, tx sqlx.Tx) {
	tx.MustExec(userInsert, user.Email, user.Password)
}
