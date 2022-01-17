package auth

import (
	"log"
	"net/http"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/hiromu-saito/trip-note-backend/form/request"
	"github.com/hiromu-saito/trip-note-backend/models/user"
	"github.com/hiromu-saito/trip-note-backend/utility"
	"golang.org/x/crypto/bcrypt"
)

type Claimes struct {
	jwt.StandardClaims
}

func Register(c *gin.Context) {
	var body request.UserRequest

	if err := c.BindJSON(&body); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	//passwordのhash化
	password, _ := bcrypt.GenerateFromPassword([]byte(body.Password), 14)
	entity := user.User{
		Email:    body.Email,
		Password: password,
	}

	err := user.Insert(entity)
	if err != nil {
		apiErr := utility.ApiErr{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
		c.JSON(apiErr.Status, apiErr)
		return
	}
	c.Status(http.StatusOK)
}

func Login(c *gin.Context) {
	var body request.UserRequest

	if err := c.BindJSON(&body); err != nil {
		log.Printf("bind json error:%s\n", err)
		c.Status(http.StatusBadRequest)
		return
	}

	user, err := user.SelectByEmail(body.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "unRegisterd email",
		})
		return
	}

	if err = bcrypt.CompareHashAndPassword(user.Password, []byte(body.Password)); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "incorrect password",
		})
		return
	}

	claimes := jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claimes)
	token, err := jwtToken.SignedString([]byte("secret"))
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.SetCookie("jwt", token, time.Now().Hour(), "/", "localhost", false, true)
	c.JSON(http.StatusOK, map[string]string{
		"message": "success",
	})
}

func Authentication(c *gin.Context) (userId int, err error) {
	cookie, err := c.Cookie("jwt")
	if err != nil {
		return
	}
	token, err := jwt.ParseWithClaims(cookie, &Claimes{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil || !token.Valid {
		return
	}

	claimes := token.Claims.(*Claimes)
	userId, _ = strconv.Atoi(claimes.Issuer)

	_, err = user.SelectById(userId)
	if err != nil {
		return
	}
	return
}

func Logout(c *gin.Context) {
	c.SetCookie("jwt", "", 0, "/", "localhost", false, true)
	c.Status(http.StatusOK)
}
