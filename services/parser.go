package services

import (
	"fmt"
	"math/rand"
	"time"
)

func StartSimulator()  {
	go func ()  {
		time.Sleep(2 * time.Second)

		now := time.Now()
		datestr := now.Format("2006-01-02")
		timestr := now.Format("15:03:11")

		distance := rand.Intn(80) +20

		action := "ok"
		status := "ok"
		if distance < 30 { action = "alarm"}

		err := SaveRangeLog( datestr,timestr, distance, action, status)
		if err != nil{
			fmt.Println("Ошибка симулятора" , err)
		}
	}()
}