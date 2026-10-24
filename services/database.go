package services

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite" // Драйвер
)

type SQLite struct{
	Id int
	Date string
	Time string
	Distance int
	Action string
	Status string
}

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "rangelog.db")

	if err != nil {
		log.Fatal("Ошибка открытия Базы данных", err)
	}

	query := `create table if not exists rangelog (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		date text,
		time text,
		distance integer,
		action text,
		status text);`

	_ , err = DB.Exec(query)

	if err != nil{
		log.Fatal("Ошибка создания таблицы", err)
	}


	log.Print("БД успешно запущена")
}

func SaveMeasurement(date, timeStr string, distance int, action, status string) error{
	query := `INSERT INTO rangelog (
	date, 
	time, 
	distance, 
	action, 
	status) 
	VALUES (?, ?, ?, ?, ?)`

	_,err := DB.Exec(query, date,timeStr,distance,action,status)
	return  err
}


type DatabaseService struct{}


func (d *DatabaseService) GetMeasurements() ([]SQLite, error) {
	

	rows,err := DB.Query("SELECT id, date, time, distance, action, status FROM rangelog")
	if err != nil{
		return  nil,err
	}
	defer rows.Close()


	var results []SQLite
	for rows.Next() {
		var temp SQLite

		

		err := rows.Scan(&temp.Id, &temp.Date, &temp.Time, &temp.Distance, &temp.Action, &temp.Status)
		if err != nil {return nil,err}
		
		results = append(results, temp)
	}
	return results,nil
}