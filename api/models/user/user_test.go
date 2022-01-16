package user

import (
	"log"
	"testing"

	"github.com/hiromu-saito/trip-note-backend/database"
	"github.com/hiromu-saito/trip-note-backend/utility"
	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	utility.DataSetupBySqlFile("/go/src/github.com/hiromu-saito/trip-note-backend/testdata/models/user/insert.sql")

	insertUser := User{
		Id:       1,
		Email:    "foo@example.com",
		Password: []byte("12345"),
	}
	if err := Insert(insertUser); err != nil {
		log.Fatal(err)
	}

	var user User
	database.Db.Get(&user, "select * from users")

	assert.Equal(t, user.Id, user.Id)
	assert.Equal(t, user.Email, "foo@example.com")
	assert.Equal(t, user.Password, []byte("12345"))

}

func TestSelectById(t *testing.T) {
	utility.DataSetupBySqlFile("/go/src/github.com/hiromu-saito/trip-note-backend/testdata/models/user/selectById.sql")

	user, err := SelectById(1)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, user.Id, 1)
	assert.Equal(t, user.Email, "hoge@example.com")
	assert.Equal(t, user.Password, []byte("12345"))
}
