package services

import (
	"errors"
	"strings"

	"go.bug.st/serial"
)

// Глобальная переменная для хранения открытого порта
var ActivePort serial.Port

func ConnectBluetooth() error {
	ports,err := serial.GetPortsList()

	if err != nil{
		return  err
	}
	for _, portName := range ports{
		if strings.Contains(portName, "HC-05"){
			ActivePort, err = serial.Open(portName, &serial.Mode{BaudRate: 9600})
			if err != nil{
				return err
			}
			return  nil
		}

	}
	return errors.New("HC-05 не найден")
}