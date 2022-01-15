package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiromu-saito/trip-note-backend/form/request"
	"github.com/hiromu-saito/trip-note-backend/models/user"
	"github.com/hiromu-saito/trip-note-backend/utils/errors"
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

	err := user.Insert(entity)
	if err != nil {
		apiErr := errors.ApiErr{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
		c.JSON(apiErr.Status, apiErr)
		return
	}

	c.Status(http.StatusOK)
}
