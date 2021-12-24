package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiromu-saito/trip-note-backend/database"
	"github.com/hiromu-saito/trip-note-backend/form/request"
	"github.com/hiromu-saito/trip-note-backend/models"
	"golang.org/x/crypto/bcrypt"
)

const userInsert = `
insert into users (
    id,
    email,
    password
)
select
    max(id) + 1,
    ?,
    ?
from
    users
`

func Register(c *gin.Context) {
	var body request.UserRequest

	if err := c.BindJSON(&body); err != nil {
		c.Status(http.StatusBadRequest)
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(body.Password), 14)
	user := models.User{
		Email:    body.Email,
		Password: password,
	}

	tx := database.Db.MustBegin()
	tx.MustExec(userInsert, user.Email, user.Password)
	tx.Commit()

	c.Status(http.StatusOK)
}
