package main

import "fmt"

var stringToInt = map[string]int{
	"January":   1,
	"February":  2,
	"March":     3,
	"April":     4,
	"May":       5,
	"June":      6,
	"July":      7,
	"August":    8,
	"September": 9,
	"October":   10,
	"November":  11,
	"December":  12,
}

var totalDayOfMonth = map[int]int{
	1:  31,
	2:  28,
	3:  31,
	4:  30,
	5:  31,
	6:  30,
	7:  31,
	8:  31,
	9:  30,
	10: 31,
	11: 30,
	12: 31,
}

var dayToint = map[string]int{
	 "Monday" : 2,
	 "Tuesday" : 3,
	 "Wednesday" : 4,
	 "Thursday" : 5,
	 "Friday" : 6,
	 "Saturday" : 7,
	 "Sunday" :8,
}

func Solution(year int, monthBegin, monthEnd, day string) int {
	totalMonthBegin, dayMonthBegin := DayOfMonthBegin(year, monthBegin, day)
	totalMonthEnd, dayMonthEnd := DayOfMonthEnd(year, monthEnd, day)
	week := ((totalMonthEnd - totalMonthBegin) - (dayMonthBegin + dayMonthEnd))/7
	return week
}

func DayOfMonthBegin(y int, monthBegin, day string) (total, kipDayBegin int) {
	var totalDayBegin int
	if isLeap(y){
		totalDayOfMonth[2] = 29
	}
	for i := 1; i < stringToInt[monthBegin]; i++ {
		totalDayBegin = totalDayBegin + totalDayOfMonth[i]
	}
	fmt.Println("totalDayBegin : ", totalDayBegin)
	kipDayBegin = totalDayBegin % 7
	kipDayBegin = (dayToint[day] + kipDayBegin)
	if kipDayBegin <=9 {
		kipDayBegin = 9 - (kipDayBegin)
	}else {
		kipDayBegin = 16 - kipDayBegin
	}
	fmt.Println("kipDayBegin : ", kipDayBegin)
	return totalDayBegin, kipDayBegin
}

func DayOfMonthEnd(year int, monthEnd, day string) (total, kipDayEnd int) {
	var totalDay int
	if isLeap(year){
		fmt.Println("Nam Nhuan")
		totalDayOfMonth[2] = 29
	}
	for i := 1; i <= stringToInt[monthEnd]; i++ {
		totalDay = totalDay + totalDayOfMonth[i]

	}
	fmt.Println("totalDayEnd : ", totalDay)
	kipDayEnd = totalDay % 7
	kipDayEnd = dayToint[day] + kipDayEnd
	if kipDayEnd > 8 {
		kipDayEnd = kipDayEnd -8
	}else {
		kipDayEnd = kipDayEnd -2
	}
	fmt.Println("kipDayEnd : ", kipDayEnd)
	return totalDay, kipDayEnd
}

func isLeap(year int) bool {

	return year % 400 == 0 || year % 4 == 0 && year % 100 != 0
}

func main() {
	fmt.Println("Solution: ",Solution(2014, "September", "May", "Wednesday"))

}
