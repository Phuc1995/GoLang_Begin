package main

func main() {
	doneChan := make(chan string)
	go func() {
		// Do some work…
		doneChan <- "I’m all done!"
	}()

	<-doneChan // block until go routine signals work is done
}