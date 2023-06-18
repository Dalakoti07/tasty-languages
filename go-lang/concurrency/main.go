package main

func main() {
	done := make(chan bool)

	go func() {
		done <- true
	}()

	println(<-done)
}
