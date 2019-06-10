package main

import (
	"fmt"
	"math"
	"reflect"
	"time"
)

func split(s []string) int {
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
	//mapvalue := make(map[string][]string)
	mang := []string{}
	for _, v := range sa {
		day, firtTime, endTime := SplipStr(v)

		mang = append(mang, day)
		mang = append(mang, firtTime)
		mang = append(mang, endTime)
	}
	fmt.Println(mang)

	lenMang := len(mang)
	for i := 2; i<lenMang; i=i+3{
		end := mang[i]
		conut := i+2
		if conut > lenMang{
			break
		}
		start := mang[i+2]

		fmt.Println(reflect.TypeOf(start))

		fmt.Println(end)
		fmt.Println(start)
		a:=TimeMeeting(end,start)
		fmt.Println(a)
		fmt.Println("***********************")
	}
	return 1
}

func solution(str []string) int  {
	//max := 0

	return 1
}

func TimeMeeting(firtTime, endTime string) int {
	t1,_ := time.Parse("15:04", firtTime)
	t2,_ := time.Parse("15:04", endTime)
	t := int(t2.Sub(t1).Minutes())
	if math.Signbit(float64(t)){
		if(firtTime=="24:00"){
			fmt.Println("errrrrrrrrrrrrrr")
			
		}else {
			t3,_ := time.Parse("15:04", "23:59")
			t13 := int(t3.Sub(t1).Minutes())
			t13= t13 +1

			t4, _ := time.Parse("15:04", "00:00")
			t14 := int(t2.Sub(t4).Minutes())
			fmt.Println("t14: ", t14)
			t = t13 +t14
			fmt.Println("t13: " ,t13)
		}

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
	//fmt.Println(TimeMeeting("10:00","16:30"))
	//fmt.Println(TimeMeeting("23:50","10:00"))
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
