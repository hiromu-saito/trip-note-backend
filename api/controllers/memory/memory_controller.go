package memory

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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
