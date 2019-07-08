package main

import (
	"fmt"
	"strings"
)

var indexToString = map[int]string{
	0: "ADD",
	1: "CHANGE",
	2: "MOVE",
	3: "NOTHING",
	4: "IMPOSSIBLE",
}

func solution(S string, T string) string {
	var result string
	if S == T {
		result = indexToString[3]
		return result
	}else {
		if len(S)==len(T){
			if Contains(S,T){
				for i := 0; i < len(S); i++{
					if S[i] != T[i]{
						result = indexToString[2]
						return result + " " + string(S[i])
					}
				}
			}else {
				var i int
				for i = 0; i < len(S); i++{
					if S[i] != T[i]{
						oldS := string(S[i])
						S = strings.Replace(S, string(S[i]),string(T[i]), 1 )
						if S == T{
							result = indexToString[1]
							return result + " " + oldS + " " + string(S[i])
						}else {
							result = indexToString[4]
							return  result
						}
					}
				}
			}
		}else {
			if Contains(S,T){
				var stringT string
				for i := len(S); i< len(T) ; i++  {
					stringT += string(T[i])
				}
				result := indexToString[0]
				return result +" " + stringT
			}else {
				result = indexToString[4]
				return  result
			}
		}

	}
	return result
}

func Contains(S,T string) bool{
	for i := 0; i <len(S); i++{
		if !strings.Contains(S,string(T[i])){
			return false
		}
	}
	return true
}

func main()  {
	S := "nice"
	T := "niceradasdasdsdsa"
	fmt.Println(solution(S,T))

	S1 := "test"
	T2 := "tent"
	fmt.Println(solution(S1,T2))

	S3 := "beans"
	T3 := "banes"
	fmt.Println(solution(S3,T3))

	S4 := "0"
	T4 := "odd"
	fmt.Println(solution(S4,T4))
}
