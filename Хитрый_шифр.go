package main

import (
	"encoding/csv"
	"fmt"
	"strings"
)

type trainees struct {
	LastName 		string	// Фамилия
	FirstName		string	// Имя
	MiddleName		string	// Отчество
	DayOfBirth		string	// День рождения
	MonthOfBirth	string	// Месяц рождения
	YearOfBirth		int		// Код10
}

func ListOfTrainees(len int) []trainees {
	ListOfTrainees := make([]trainees, 0, len)
	for len > 0 {
		var tmp string
		fmt.Scanf("%s", &tmp)
		str := strings.NewReader(tmp)
		reader := csv.NewReader(str)
		v, _ := reader.Read()
		trainee := trainees{
			LastName: v[0],
			FirstName: v[1],
			MiddleName: v[2],
			DayOfBirth: v[3],
			MonthOfBirth: v[4],
		}
		ListOfTrainees = append(ListOfTrainees, trainee)
		len--
	}
	return ListOfTrainees
}

func encodeTrainee10(t trainees) int {
	encodeFullName := make(map[rune]int)
	var encodeDateOfBirth int
	var encodeLetterOfFirstName int
	for _, v2 := range t.FirstName {
		encodeFullName[v2] = 1
	}
	for _, v3 := range t.LastName {
		encodeFullName[v3] = 1
	}
	for _, v4 := range t.MiddleName {
		encodeFullName[v4] = 1
	}
	for j := 0; j < len(t.DayOfBirth); j++ {
		encodeDateOfBirth += int(t.DayOfBirth[j]) - '0'
	}
	for k := 0; k < len(t.MonthOfBirth); k++ {
		encodeDateOfBirth += int(t.MonthOfBirth[k]) - '0'
	}
	encodeLetterOfFirstName = int(t.LastName[0]) - 64
	return len(encodeFullName) + encodeDateOfBirth * 64 + encodeLetterOfFirstName * 256
}

func encodeTrainees10(list []trainees) {
	for i := range list {
		list[i].YearOfBirth = encodeTrainee10(list[i])
		for list[i].YearOfBirth >= 4096 {
			list[i].YearOfBirth -= 4096	
		}
	}
}

func PrintEncodedTrainees(list []trainees) {
	amount := len(list)
	for i := range list {
		fmt.Printf("%03X", list[i].YearOfBirth)
		if i != amount - 1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

func main() {
	var amount int
	fmt.Scanf("%d", &amount)
	list := ListOfTrainees(amount)
	encodeTrainees10(list)
	PrintEncodedTrainees(list)
}