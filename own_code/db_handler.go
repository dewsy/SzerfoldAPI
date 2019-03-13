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
	defer db.Close()
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

func GetLatestDailies(since *int64) (dailies []*models.Daily) {
	db := dbConnect()
	defer db.Close()
	rows, err := db.Query("SELECT id, message, pray, title, verse FROM dailies WHERE id > $1 ORDER BY id DESC LIMIT 20", since)
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		daily := models.Daily{}
		err = rows.Scan(&daily.ID, &daily.Message, &daily.Pray, &daily.Title, &daily.Verse)
		if err != nil {
			// handle this error
			panic(err)
		}
		dailies = append(dailies, &daily)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return
}

func UpdateDaily(dailyToUpdate models.Daily, id int64) models.Daily {
	db := dbConnect()
	sqlStatement := `
			UPDATE dailies
			SET message = $2, pray = $3, title = $4, verse = $5
			WHERE id = $1;`
	_, err := db.Exec(sqlStatement, id, dailyToUpdate.Message, dailyToUpdate.Pray, dailyToUpdate.Title, dailyToUpdate.Verse)
	if err != nil {
		panic(err)
	}
	db.Close()
	freshDaily := GetDailyByID(id)
	return freshDaily
}

func GetDailyByID(id int64) (resultDaily models.Daily) {
	db := dbConnect()
	defer db.Close()
	sqlStatement := `SELECT id, message, pray, title, verse FROM dailies WHERE id = $1;`
	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&resultDaily.ID, &resultDaily.Message, &resultDaily.Pray, &resultDaily.Title, &resultDaily.Verse)
	if err != nil {
		panic(err)
	}
	return
}

func DeleteDailyByID(id int64) {
	db := dbConnect()
	defer db.Close()
	sqlStatement := `
			DELETE FROM dailies
			WHERE id = $1;`
	_, err := db.Exec(sqlStatement, id)
	if err != nil {
		panic(err)
	}
}
