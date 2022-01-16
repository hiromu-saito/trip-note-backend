package request

import (
	"time"

	"github.com/hiromu-saito/trip-note-backend/models/memory"
)

type MemoryRequest struct {
	Id               int       `json:"id"`
	UserId           int       `json:"userId"`
	HotelName        string    `json:"hotelName"`
	HotelImage       string    `json:"hotelImage"`
	Impression       string    `json:"impression"`
	AccomodationDate time.Time `json:"accomodationDate"`
	DetailUrl        string    `json:"detailUrl"`
}

func (request *MemoryRequest) ToMemory() memory.Memory {
	return memory.Memory{
		UserId:           request.UserId,
		HotelName:        request.HotelName,
		Impression:       request.Impression,
		AccomodationDate: time.Time(request.AccomodationDate),
		DetailUrl:        request.DetailUrl,
	}
}
