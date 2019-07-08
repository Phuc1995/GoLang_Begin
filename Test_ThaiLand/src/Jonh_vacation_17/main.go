package main

import (
	"fmt"

	"time"
)

const (
	January = iota + 1
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

func countWeek(year int, begin, end, day string) int {
	beginMonthConvert := convertMonth(begin)
	dayBeginConvert := convertDayBegin(day, beginMonthConvert, year)
	//fmt.Println("Test countWeek: ",dayBeginConvert)
	startVacation := Date(dayBeginConvert, beginMonthConvert, year)
	fmt.Println("Test Date Begin: ", startVacation)

	endMonthConvert := convertMonth(end)
	dayEndConvert := convertDayEnd(day, endMonthConvert, year)
	endVacation := Date(dayEndConvert, endMonthConvert, year)
	fmt.Println("Test Date End: ", endVacation)

	countWeek := subtractTime(startVacation, endVacation)
	countWeek = countWeek/24/7
	return countWeek
}

func subtractTime(time1,time2 time.Time) int{
	diff := int(time2.Sub(time1).Hours())
	return diff
}

func convertMonth(s string) int {
	n := 0
	switch s {
	case "January":
		n = January
		return n
	case "February":
		n = February
		return n
	case "March":
		n = March
		return n
	case "April":
		n = April
		return n
	case "May":
		n = May
		return n
	case "June":
		n = June
		return n
	case "July":
		n = July
		return n
	case "August":
		n = August
		return n
	case "September":
		n = September
		return n
	case "October":
		n = October
		return n
	case "November":
		n = November
		return n
	case "December":
		n = December
		return n
	default:
		fmt.Println("Month not exist")
	}
	return n
}

func convertDayBegin(day string, month, year int) int {
	firtDayOfMonth := Date(1, month, year).Weekday().String()
	//fmt.Println("firt: ",firtDayOfMonth)
	n := 1
	switch firtDayOfMonth {
	case "Monday":
		n = n
	case "Tuesday":
		n = n + 6
	case "Wednesday":
		n = n + 5
	case "Thursday":
		n = n + 4
	case "Friday":
		n = n + 3
	case "Saturday":
		n = n + 2
	case "Sunday":
		n = n + 1
	}
	return n
}

func convertDayEnd(day string, month, year int) int {
	n := 0
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		endDayOfMonth := Date(31, month, year).Weekday().String()
		n := 31
		switch endDayOfMonth {
		case "Monday":
			n = n
		case "Tuesday":
			n = n - 1
		case "Wednesday":
			n = n - 2
		case "Thursday":
			n = n - 3
		case "Friday":
			n = n - 4
		case "Saturday":
			n = n - 5
		case "Sunday":
			n = n - 6
		}
		return n
	case 4, 6, 9, 11:
		endDayOfMonth := Date(30, month, year).Weekday().String()
		n := 30
		switch endDayOfMonth {
		case "Monday":
			n = n
		case "Tuesday":
			n = n - 1
		case "Wednesday":
			n = n - 2
		case "Thursday":
			n = n - 3
		case "Friday":
			n = n - 4
		case "Saturday":
			n = n - 5
		case "Sunday":
			n = n - 6
		}
		return n
	case 2:
		if isLeap(year) {
			endDayOfMonth := Date(29, month, year).Weekday().String()
			n := 29
			switch endDayOfMonth {
			case "Monday":
				n = n
			case "Tuesday":
				n = n - 1
			case "Wednesday":
				n = n - 2
			case "Thursday":
				n = n - 3
			case "Friday":
				n = n - 4
			case "Saturday":
				n = n - 5
			case "Sunday":
				n = n - 6
			}
			return n
		} else {
			endDayOfMonth := Date(28, month, year).Weekday().String()
			n := 28
			switch endDayOfMonth {
			case "Monday":
				n = n
			case "Tuesday":
				n = n - 1
			case "Wednesday":
				n = n - 2
			case "Thursday":
				n = n - 3
			case "Friday":
				n = n - 4
			case "Saturday":
				n = n - 5
			case "Sunday":
				n = n - 6
			}
			return n
		}
	default:
		fmt.Println("Error.....")
	}
	return n
}

func Date(day, month, year int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func isLeap1(year int) bool {
	return year%400 == 0 || year%4 == 0 && year%100 != 0
}

func main() {
	fmt.Println(countWeek(2014, "April", "May", "Wednesday"))

	/*convertDayEnd("aa", 5, 2014)
	fmt.Println("start: ", convertDayBegin("aa", 2, 2019))
	fmt.Println("end: ", convertDayEnd("aa", 2, 2019))
	fmt.Println(isLeap(2019))*/

	/*fmt.Println(convertMonth("January"))
	fmt.Println(convertDay("a",6,2019))
	fmt.Println(convertDay("Monday",4,2014))*/
	/*fmt.Println(February)
	t := Date(8,6,2019)
	fmt.Println(t.ISOWeek())
	fmt.Println(t.Weekday())*/

}
