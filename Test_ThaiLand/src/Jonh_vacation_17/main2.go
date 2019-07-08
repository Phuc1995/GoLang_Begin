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
	totalMonthBegin, dayMonthBegin := DayOfMonthBegin(year,monthBegin,day)
	totalMonthEnd, dayMonthEnd := DayOfMonthEnd(year,monthEnd,day)
	week := ((totalMonthEnd - totalMonthBegin)-(dayMonthBegin+dayMonthEnd))/7
	return week
}

func DayOfMonthBegin(y int, month,day string) (total, intDay int) {
	var totalDay int
	if isLeap(y){
		totalDayOfMonth[2] = 29
	}
	for i := 1; i < stringToInt[month]; i++ {
		totalDay = totalDay + totalDayOfMonth[i]
	}
	fmt.Println("totalDayBegin : ", totalDay)
	intDay = totalDay % 7

	intDay = (dayToint[day] + intDay)
	if intDay <=9{
		intDay = 9-(intDay)
	}else {
		intDay = 8- (intDay -7)
	}
	fmt.Println("intDayyBegin : ", intDay)
	return totalDay,intDay
}

func DayOfMonthEnd(y int, monthEnd,day string) (total, intDay int) {
	var totalDay int
	if isLeap(y){
		totalDayOfMonth[2] = 29
	}
	for i := 1; i <= stringToInt[monthEnd]; i++ {
		totalDay = totalDay + totalDayOfMonth[i]

	}
	fmt.Println("totalDayEnd : ", totalDay)
	intDay = totalDay % 7
	fmt.Println("A1: ",intDay)
	intDay = dayToint[day] + intDay
	if intDay >= 8 {
		intDay = intDay -8
	}else {
		intDay = intDay -2
	}
	fmt.Println("intDayyEnd : ", intDay)
	return totalDay,intDay
}

func isLeap(year int) bool {
	fmt.Println("Nam Nhuan")
	return year%400 == 0 || year%4 == 0 && year%100 != 0
}

func main() {
	fmt.Println("Solution: ",Solution(2015, "February", "April", "Thursday"))

}
