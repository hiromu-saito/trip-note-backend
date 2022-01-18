package request

import (
	"time"

	"github.com/hiromu-saito/trip-note-backend/models/memory"
)

type MemoryRequest struct {
	Id               int    `json:"id"`
	UserId           int    `json:"userId"`
	HotelName        string `json:"hotelName"`
	HotelImage       string `json:"hotelImage"`
	Impression       string `json:"impression"`
	AccomodationDate string `json:"accommodationDate"`
	DetailUrl        string `json:"detailUrl"`
}

func (request *MemoryRequest) ToMemory() memory.Memory {
	t, _ := time.Parse("2006-01-02", request.AccomodationDate)
	return memory.Memory{
		Id:               request.Id,
		UserId:           request.UserId,
		HotelName:        request.HotelName,
		HotelImage:       request.HotelImage,
		Impression:       request.Impression,
		AccomodationDate: t,
		DetailUrl:        request.DetailUrl,
	}
}
