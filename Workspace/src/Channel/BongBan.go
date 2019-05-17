package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// wg được sử dụng để đợi các goroutine kết thúc
var wg sync.WaitGroup

// Hàm này chạy trước khi các lệnh hàm main được thực thi
func init()  {
	rand.Seed(time.Now().UnixNano())
}

//Ham mo ta mot nguoi choi quan vot
func player(name string, court chan int)  {
//thong bao nguoi nay choi xong van dau
	defer wg.Done()

	for{
		//Doi banh tu doi thu
		ball, ok := <-court
		if !ok{
			//thang neu kenh da dong
			fmt.Printf("%s thang!\n", name)
		}

		//lay mot gia tri ngau nhien 0-99
		n := rand.Intn(100)
		if n%13==0{
			fmt.Printf("%s đánh hỏng ở lượt đánh thứ %d!\n", name, ball)

			//Dong kenh khi danh hong
			close(court)
			return
		}
		fmt.Printf("Lượt đánh bóng thành công thứ %d: %s\n", ball, name)
		ball++

		// Đánh banh về lại đối thủ
		court <- ball
	}

}

func main()  {
	court := make(chan int)

	//If the counter becomes zero, all goroutines blocked on Wait are released
	//If the counter goes negative, Add panic
	wg.Add(2)

	go player("Federer", court)

}