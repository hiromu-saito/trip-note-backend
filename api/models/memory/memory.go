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

const updateSql = `
update
    memories
set
    impression         = :impression
	 ,accommodation_date = :accommodation_date
`

const insertSql = `
INSERT INTO memories values (
	id
 ,user_id
 ,hotel_name
 ,hotel_image
 ,impression
 ,accommodation_date
 ,delete_flag
) 
select
 case
    when max(id) is null then 1
    else max(id)+1
  end
 ,:user_id
 ,:hotel_name
 ,:hotel_image
 ,:impression
 ,:accommodation_date
 ,0
from
  memories;
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

func Update(memory Memory) error {
	tx, err := database.Db.Beginx()
	defer tx.Rollback()
	if err != nil {
		log.Printf("transaction begin error%s", err)
		return err
	}
	_, err = tx.NamedExec(updateSql, &memory)
	if err != nil {
		log.Printf("memory update Error:%s", err)
		return err
	}
	tx.Commit()
	return nil
}

func Insert(memory Memory) error {
	tx, err := database.Db.Beginx()
	defer tx.Rollback()
	if err != nil {
		log.Printf("transaction begin error%s", err)
		return err
	}

	_, err = tx.NamedExec(insertSql, memory)
	if err != nil {
		log.Printf("memory insert error%s", err)
		return err
	}

	tx.Commit()
	return nil
}

func Delete(id int) error {
	tx, err := database.Db.Beginx()
	defer tx.Rollback()
	if err != nil {
		log.Printf("transaction begin error%s", err)
		return err
	}
	if _, err := tx.Exec("delete from memories where id = ?", id); err != nil {
		log.Printf("memory delete error:%s", err)
	}
	tx.Commit()
	return nil
}
