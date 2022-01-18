package response

import (
	"time"

	"github.com/hiromu-saito/trip-note-backend/models/memory"
)

type MemoryResponse struct {
	Id               int      `json:"id"`
	UserId           int      `json:"userId"`
	HotelName        string   `json:"hotelName"`
	HotelImage       string   `json:"hotelImage"`
	Impression       string   `json:"impression"`
	AccomodationDate jsonTime `json:"accommodationDate"`
	DetailUrl        string   `json:"detailUrl"`
}

// 独自のjsonTimeを作成
type jsonTime struct {
	time.Time
}

// formatを設定
func (j jsonTime) format() string {
	return j.Time.Format("2006/01/02")
}

// MarshalJSON() の実装
func (j jsonTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + j.format() + `"`), nil
}

func CreateMemoryResponse(memory memory.Memory) MemoryResponse {
	return MemoryResponse{
		Id:               memory.Id,
		UserId:           memory.UserId,
		HotelName:        memory.HotelName,
		HotelImage:       memory.HotelImage,
		Impression:       memory.Impression,
		AccomodationDate: jsonTime{memory.AccomodationDate},
		DetailUrl:        memory.DetailUrl,
	}
}
