package main

import (
	"fmt"
	"time"
)

func main() {
	tick1 := time.After(time.Second)
	tick2 := time.After(time.Second * 2)
	select {
	case <-tick1:
		fmt.Println("Получено значение из первого канала")
	case <-tick2:
		fmt.Println("Получено значение из второго канала")
	// Блок default выполнится раньше блока case - 1 секунда слишком много для Go
	default:
		fmt.Println("Действие по умолчанию")
	}
}
