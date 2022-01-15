package memory

import (
	"log"
	"time"

	"github.com/hiromu-saito/trip-note-backend/database"
)

type Memory struct {
	Id               int       `db:"id"`
	UserId           int       `db:"user_id"`
	HotelName        string    `db:"hotel_name"`
	HotelImage       string    `db:"hotel_image"`
	Impression       string    `db:"impression"`
	AccomodationDate time.Time `db:"accommodation_date"`
	DetailUrl        string    `db:"detail_url"`
	DeleteFlag       int       `db:"delete_flag"`
}

const selectByUserIdSql = `
select
	*
from
	memories
where
		user_id     = ?
and delete_flag = 0
`

func SelectByUserId(userId int) ([]Memory, error) {
	var records []Memory
	err := database.Db.Select(&records, selectByUserIdSql, userId)
	if err != nil {
		log.Printf("memory selectByUserId Error:%s", err)
		return nil, err
	}
	return records, nil
}
