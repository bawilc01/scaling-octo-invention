package main

// import block
import "fmt"

/*
gRPC API: Attempt 1; On Mac
Helpful links: https://grpc.io/docs/languages/go/basics/
*/
func main() {
	fmt.Println("Hello there.")
	startWebServer()
}

func startWebServer() {
	fmt.Println(("Starting server..."))
	// do server stuff
	fmt.Println("Server started")
}
