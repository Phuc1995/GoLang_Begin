package main

import (
	"fmt"

	"math"
	"time"
)

func split(s []string) int {
	arrayMax := []int{}
	mang := []string{}
	for _, v := range s {
		day, firtTime, endTime := SplipStr(v)

		mang = append(mang, day)
		mang = append(mang, firtTime)
		mang = append(mang, endTime)
	}
	fmt.Println(mang)

	lenMang := len(mang)
	arrayDay := []string{}

	for i := 2; i<lenMang; i=i+3{
		end := mang[i]
		day := mang[i-2]
		fmt.Println(day)
		count := i+2
		if count > lenMang{
			break
		}
		start := mang[i+2]
		arrayDay = append(arrayDay, day)
		bool :=newWeek(arrayDay,day)
		fmt.Println("new week:",newWeek(arrayDay,day))

		if newWeek(arrayDay,day) {
			arrayDay = nil
			arrayDay=append(arrayDay,day)
			}
		fmt.Println(arrayDay,"-")
		fmt.Println(end)
		fmt.Println(start)
		a:=TimeMeeting(day,end,start,bool)
		arrayMax = append(arrayMax,a)
		fmt.Println("Max: ",Max(arrayMax))
		fmt.Println("time: ",a)
		fmt.Println("***********************")
	}
	return 1
}

//return max minute
func Max(arrayMax []int) int {
	max := arrayMax[0]
	for _,v := range arrayMax{
		if v > max {
			max = v // found another smaller value, replace previous value in max
		}
	}
	return max
}

//result is true if this day is begin new Week
func newWeek(array []string,day string) bool {
	var istrue bool
	istrue = false
	if(day=="Sun"){
		istrue = true
		return istrue
	}else if(day=="Mon"){
		istrue = true
		return istrue
	}else if(day=="Tue") {
		for _,v := range array {
			if(v=="Mon" || v=="Tue"){
				istrue = false
				return istrue
			}else {
				istrue = true
				return istrue
			}
		}
	}else if(day=="Wed"){
		for _,v := range array {
			if(v=="Mon" || v=="Tue" || v == "Wed"){
				istrue = false
				return istrue
			}else {
				istrue = true
				return istrue
			}
		}
	}else if(day=="Wed"){
		for _,v := range array {
			if(v=="Mon" || v=="Tue" || v == "Wed"){
				istrue = false
				return istrue
			}else {
				istrue = true
				return istrue
			}
		}
	}else if(day=="Thu"){
		for _,v := range array {
			if(v=="Mon" || v=="Tue" || v == "Wed" || v == "Thu"){
				istrue = false
				return istrue
			}else {
				istrue = true
				return istrue
			}
		}
	}else if(day=="Fri"){
		for _,v := range array {
			if(v=="Mon" || v=="Tue" || v == "Wed" || v=="Thu" || v == "Fri"){
				istrue = false
				return istrue
			}else {
				istrue = true
				return istrue
			}
		}
	}else if(day=="Sat"){
		for _,v := range array {
			if(v=="Mon" || v=="Tue" || v == "Wed" || v=="Thu" || v == "Fri" || v == "Sat"){
				istrue = false
				return istrue
			}else {
				istrue = true
				return istrue
			}
		}
	}

	return istrue

}

func solution(str []string) int  {
	//max := 0

	return 1
}

func TimeMeeting(day, end, start string, bool bool) int {
	t := 0
	if(bool){
		if day == "Sun" {
			t1,_ := time.Parse("15:04", end)
			t2,_ := time.Parse("15:04", "23:59")
			t := int(t2.Sub(t1).Minutes())
			t=t+1
			return t
		}
	}else {
		t1,_ := time.Parse("15:04", end)
		t2,_ := time.Parse("15:04", start)
		t := int(t2.Sub(t1).Minutes())
		switch day {
		case "Sun":
			t1,_ := time.Parse("15:04", end)
			t2,_ := time.Parse("15:04", "23:59")
			t := int(t2.Sub(t1).Minutes())
			t=t+1
			return t
		default:
			if math.Signbit(float64(t)){
				if(end=="24:00"){
					t1,_ := time.Parse("15:04", "00:01")
					t2,_ := time.Parse("15:04", start)
					t := int(t2.Sub(t1).Minutes())
					t = t+1
					return t

				}else {
					t3,_ := time.Parse("15:04", "23:59")
					t13 := int(t3.Sub(t1).Minutes())
					t13= t13 +1

					t4, _ := time.Parse("15:04", "00:00")
					t14 := int(t2.Sub(t4).Minutes())

					t = t13 +t14
				}
			}
		}
		return t
	}
	return t

}

func SplipStr(str string) (day string, timeFirt string, timeEnd string){
	for i:= 0; i < len(str); i++ {
		if i < 3 {
			day += string(str[i])
			}
		if 3 < i && i < 9 {
			timeFirt += string(str[i])
			}
		if 9 < i && i < 15 {
			timeEnd += string(str[i])
		}
	}
	return
}

func main()  {
	t1,_ := time.Parse("4:00", "4:00")
	t2,_ := time.Parse("15:00", "24:00")
	t := int(t2.Sub(t1).Minutes())
	//t=t+1
	//fmt.Println(TimeMeeting("10:00","16:30"))
	fmt.Println("Sun: ",t)
	sa := []string{"Sun 10:00-20:00",
		"Fri 05:00-10:00",
		"Fri 16:30-23:50",
		"Sat 10:00-24:00",
		"Sun 01:00-04:00",
		"Sat 02:00-06:00",
		"Tue 03:30-18:15",
		"Tue 19:00-20:00",
		"Wed 04:25-15:14",
		"Wed 15:14-22:40",
		"Thu 00:00-23:59",
		"Mon 05:00-13:00",
		"Mon 15:00-21:00"}
	split(sa)

}
