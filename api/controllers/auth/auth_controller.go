package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiromu-saito/trip-note-backend/database"
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

	tx, err := database.Db.Beginx()
	if err != nil {
		apiErr := errors.ApiErr{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
		c.JSON(apiErr.Status, apiErr)
		return
	}

	defer tx.Rollback()

	err = user.Insert(entity, *tx)
	if err != nil {
		apiErr := errors.ApiErr{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
		c.JSON(apiErr.Status, apiErr)
		return
	}

	tx.Commit()
	c.Status(http.StatusOK)
}
