package main

import "unicode"

func maxSentenceWords(s string) int {
	maxWord, nWord := 0,0
	word := false

	for _, r:= range s{
		switch r {
		case '.', '?','!':
			word = false
			if maxWord < nWord{
				maxWord = nWord
			}
			nWord = 0
		default:
			//check is '\t', '\n', '\v', '\f', '\r' return true
			if unicode.IsSpace(r){
				word =false
			}else if word == false{
				word = true
				nWord++
			}
		}
		if maxWord < nWord{
			maxWord = nWord
		}
	}
	return maxWord
}

func main()  {
	var s string
	s = "A B C. A B ? 1 2 3 t 5 1 !"
	maxSentenceWords(s)
	println("s: ",maxSentenceWords(s))
	s1 := "Forget    CVs.. Save time . x x"
	println("s1:",maxSentenceWords(s1))
}