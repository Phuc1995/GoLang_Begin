package main

import (
	"fmt"
	"golang.org/x/net/context"
	"time"
)

func main()  {
	ctx,_ :=context.WithTimeout(context.Background(), time.Second*3	)
	doSomeThing(ctx)
}

func doSomeThing(ctx context.Context) {
	canceledChannel := make(chan bool)
	go func() {
		select {
		case <- ctx.Done():
			fmt.Println(ctx.Err())
			canceledChannel <- true
			return
		}
	}()

	isCanceledChannel := <- canceledChannel
	if(isCanceledChannel){
		close(canceledChannel)
		fmt.Println("end")
		return
	}

	time.Sleep(time.Second*10)

}