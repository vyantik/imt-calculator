package main

import (
	"errors"
	"fmt"
	"math"
)

const IMTPower = 2

func calculateIMT(heightCm float64, weightKg float64) (float64, error) {
	if weightKg <= 0 || heightCm <= 0 {
		return 0, errors.New("Не указан вес или высота")
	}
	return weightKg / math.Pow(heightCm/100, IMTPower), nil
}

func getIMTInfo(imt float64) string {
	switch {
	case imt < 16:
		return "Выраженный дефицит массы тела"
	case imt < 18.5:
		return "Недостаточная (дефицит) масса тела"
	case imt < 25:
		return "Норма"
	case imt < 30:
		return "Избыточная масса тела (предожирение)"
	case imt < 35:
		return "Ожирение первой степени"
	case imt < 40:
		return "Ожирение второй степени"
	default:
		return "Ожирение третьей степени (морбидное)"
	}
}

func getUserInput() (float64, float64) {
	var userHeight, userKg float64

	fmt.Print("Введите ваш рост в сантиметрах: ")
	fmt.Scan(&userHeight)

	fmt.Print("Введите ваш вес: ")
	fmt.Scan(&userKg)

	return userHeight, userKg
}
