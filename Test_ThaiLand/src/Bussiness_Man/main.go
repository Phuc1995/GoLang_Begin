package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var dayToIndexMap = map[string]int {
	`Mon`: 1,
	`Tue`: 2,
	`Wed`: 3,
	`Thu`: 4,
	`Fri`: 5,
	`Sat`: 6,
	`Sun`: 7,
}

type meettingTime struct {
	begin int
	end int
}

type meetingTimeArray []meettingTime

func (a meetingTimeArray) Len() int {
	return len(a)
}

func (a meetingTimeArray) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a meetingTimeArray) Less(i, j int) bool {
	return a[i].begin < a[j].begin
}

func split(s string)  {
	pattern := regexp.MustCompile(`:|\n|-| `)
	array := pattern.Split(s, -1)

	meetingTimes := []meettingTime{}
	fmt.Println(array)

	i := 0
	for i < len(array) {
		//	fmt.Println(strings.Trim(array[i], " \t"))
		dayIndex := dayToIndexMap[strings.Trim(array[i], " \t")]

		beginHour,_ := strconv.Atoi(array[i + 1])
		beginMin,_ := strconv.Atoi(array[i + 2])
		endHour,_ := strconv.Atoi(array[i + 3])
		endMin,_ := strconv.Atoi(array[i + 4])

		begin := (24 * (dayIndex - 1) + beginHour) * 60 + beginMin
		end := (24 * (dayIndex - 1) + endHour) * 60 + endMin

		mTime := meettingTime{begin, end}
		meetingTimes = append(meetingTimes, mTime)
		i += 5
	}

	fmt.Println("Befor sort: ",meetingTimes)
	sort.Sort(meetingTimeArray(meetingTimes))
	fmt.Println("Sort meetingTimes: ",meetingTimes)

	currentMax := 0
	i = 0
	for i < len(meetingTimes) - 1 {
		val := meetingTimes[i]
		if i == 0 {
			currentMax = val.begin
		}
		max := meetingTimes[i + 1].begin - meetingTimes[i].end
		if(i==len(meetingTimes) - 2){
			max = 24*7*60 - meetingTimes[len(meetingTimes)-1].end
		}
		if currentMax < max {
			currentMax = max
			fmt.Printf("currentMax at: %d\n", i)
		}
		i += 1
	}

	fmt.Println(currentMax)
}


func main() {
	test := `Sun 10:00-20:00
Fri 05:00-10:00
Fri 16:30-23:50
Sat 10:00-24:00
Sun 01:00-04:00
Sat 02:00-06:00
Tue 03:30-18:15
Tue 19:00-20:00
Wed 04:25-15:14
Wed 15:14-22:40
Thu 00:00-23:59
Mon 05:00-13:00
Mon 15:00-21:00`

	test1 := `Mon 01:00-23:00
Tue 01:00-23:00
Wed 01:00-23:00
Thu 01:00-23:00
Fri 01:00-23:00
Sat 01:00-23:00
Sun 01:00-21:00`


	split(test)
	split(test1)
}