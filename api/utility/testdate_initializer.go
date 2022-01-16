package utility

import (
	"bytes"
	"io/ioutil"
	"log"
	"strings"

	"github.com/hiromu-saito/trip-note-backend/database"
)

func DataSetupBySqlFile(path string) {
	sql, err := readSqlFile(path)
	if err != nil {
		log.Fatalf("reading sql file error:%s", err)
	}
	err = execSql(sql)
	if err != nil {
		log.Fatalf("exec sql for testdata error:%s", err)
	}
}

func readSqlFile(path string) (string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	b := bytes.NewBuffer(content)
	return b.String(), nil
}

func execSql(sql string) error {
	splitedSql := strings.Split(sql, ";")
	for _, s := range splitedSql {
		trimedSql := strings.TrimSpace(s)
		if len(trimedSql) == 0 {
			continue
		}
		_, err := database.Db.Exec(trimedSql)
		if err != nil {
			return err
		}
	}
	return nil
}
