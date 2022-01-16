package memory

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hiromu-saito/trip-note-backend/form/request"
	"github.com/hiromu-saito/trip-note-backend/form/response"
	"github.com/hiromu-saito/trip-note-backend/models/memory"
	"github.com/hiromu-saito/trip-note-backend/utility"
)

func GetMemories(c *gin.Context) {

	//TODO jewtからユーザーIDを取得
	userId := 1
	fmt.Println(userId)

	memories, err := memory.SelectByUserId(userId)
	if err != nil {
		apiErr := utility.ApiErr{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}
		c.JSON(apiErr.Status, apiErr)
		return
	}

	r := []response.MemoryResponse{}
	for _, m := range memories {
		r = append(r, response.CreateMemoryResponse(m))
	}
	c.JSON(http.StatusOK, r)
}

func UpdateMemories(c *gin.Context) {
	var body request.MemoryRequest
	if err := c.BindJSON(&body); err != nil {
		log.Printf("bind json error:%s", err)
		apiErr := utility.ApiErr{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}
		c.JSON(apiErr.Status, apiErr)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		apiErr := utility.ApiErr{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}
		c.JSON(apiErr.Status, apiErr)
		return
	}

	body.Id = id
	if err := memory.Update(body.ToMemory()); err != nil {
		apiErr := utility.ApiErr{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}
		c.JSON(apiErr.Status, apiErr)
		return
	}

	c.Status(http.StatusOK)
}

func InsertMemories(c *gin.Context) {
	var body request.MemoryRequest

	if err := c.BindJSON(&body); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if err := memory.Insert(body.ToMemory()); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	c.Status(http.StatusOK)
}

func DeleteMemories(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		apiErr := utility.ApiErr{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}
		c.JSON(apiErr.Status, apiErr)
		return
	}

	if err := memory.Delete(id); err != nil {
		apiErr := utility.ApiErr{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}
		c.JSON(apiErr.Status, apiErr)
		return
	}
	c.Status(http.StatusOK)
}
