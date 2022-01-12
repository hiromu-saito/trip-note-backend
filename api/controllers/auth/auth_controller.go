package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiromu-saito/trip-note-backend/database"
	"github.com/hiromu-saito/trip-note-backend/form/request"
	"github.com/hiromu-saito/trip-note-backend/models/user"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var body request.UserRequest

	if err := c.BindJSON(&body); err != nil {
		c.Status(http.StatusBadRequest)
	}
	//passwordのhash化
	password, _ := bcrypt.GenerateFromPassword([]byte(body.Password), 14)
	entity := user.User{
		Email:    body.Email,
		Password: password,
	}

	tx := database.Db.MustBegin()
	user.Insert(entity, *tx)
	tx.Commit()

	c.Status(http.StatusOK)
}
