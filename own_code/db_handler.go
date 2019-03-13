package own_code

import (
	"SzerfoldAPI/models"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)
import "strconv"

var psqlInfo string

func dbConnect() *sql.DB {
	if psqlInfo == "" {
		conf := ReadConfiguration()
		port, err := strconv.Atoi(conf["port"])
		if err != nil {
			panic(err)
		}
		psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			conf["host"], port, conf["DbUser"], conf["DbPassword"], conf["DbName"])
	}
	db, openErr := sql.Open("postgres", psqlInfo)
	if openErr != nil {
		panic(openErr)
	}
	return db
}

func AddNewDaily(newDaily models.Daily) models.Daily {
	addedDaily := models.Daily{}
	db := dbConnect()
	SqlStatement := `
			INSERT INTO dailies (message, verse, pray, title)
			VALUES ($1, $2, $3, $4)
			RETURNING id, message, verse, pray, title`
	err := db.QueryRow(SqlStatement, newDaily.Message, newDaily.Verse, newDaily.Pray, newDaily.Title).Scan(&addedDaily.ID, &addedDaily.Message, &addedDaily.Verse, &addedDaily.Pray, &addedDaily.Title)
	if err != nil {
		panic(err)
	}
	return addedDaily
}
