package user

import (
	"testing"

	"github.com/hiromu-saito/trip-note-backend/database"
	"github.com/hiromu-saito/trip-note-backend/utility"
	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	utility.DataSetupBySqlFile("/go/src/testdata/models/user/insert.sql")

	var user User
	database.Db.Get(&user, "select * from users")

	assert.Equal(t, user.Id, user.Id)
	assert.Equal(t, user.Email, "hoge@example.com")
	assert.Equal(t, user.Password, []byte("12345"))

}
