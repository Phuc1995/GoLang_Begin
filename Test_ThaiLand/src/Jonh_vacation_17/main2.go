package main

import "fmt"

var totalDayOfMonth = map[string]int{

}

var dayToint = map[int]string{
	2: "Monday",
	3: "Tuesday",
	4: "Wednesday",
	5: "Thursday",
	6: "Friday",
	7: "Saturday",
	8: "Sunday",
}

/*func Solution(year int, monthBegin, monthEnd, day string) int {

}*/

/*func DayOfMonth(m, y int) int  {

}

func firstWeekdayOfMonth(x1, x2, x3) {
	days := 0
	for i:=1; i <= x2; i++{
		days += DayOfMonth(i)
	}
	w := day%7
	return x1+w
}
*/
func firtDayOfMothBegin(day string, totalDay int) string {
	n := totalDay % 7
	switch day {
	case "Monday":
		n = n + 1
		return dayToint[n]
	case "Tuesday":
		n = n + 3
		return dayToint[n]
	case "Wednesday":
		n = n + 3
		return dayToint[n]
	case "Thursday":
		n = n + 4
		return dayToint[n]
	case "Friday":
		n = n + 5
		return dayToint[n]
	case "Saturday":
		n = n + 6
		return dayToint[n]
	case "Sunday":
		n = n + 7
		return dayToint[n]
	}
	return ""
}

func isLeap(year int) bool {
	return year%400 == 0 || year%4 == 0 && year%100 != 0
}

func main() {
	//fmt.Println(Solution(2014, "April", "May", "Wednesday"))
	fmt.Println(firtDayOfMothBegin("Saturday", 30))
}
