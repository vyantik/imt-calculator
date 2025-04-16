package main

import (
	"fmt"
)

func main() {
	fmt.Println("__ Калькулятор ИМТ __")
	userChoise := 1
	for {
		if userChoise == 1 {
			userHeight, userKg := getUserInput()

			imt, err := calculateIMT(userHeight, userKg)
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Printf("Ваш ИМТ равен: %.2f\n", imt)
			fmt.Println(getIMTInfo(imt))
		} else if userChoise == 2 {
			break
		} else {
			fmt.Println("Неверный ввод...")
		}
		fmt.Print("Хотите ли вы попробовать еще раз?\n1. Повторить рассчет.\n2. Выход\n")
		fmt.Scan(&userChoise)
	}
}
