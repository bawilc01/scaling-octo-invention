package main

// import block
import "fmt"

/*
gRPC API: Attempt 1; On Mac
Helpful links: https://grpc.io/docs/languages/go/basics/
*/
func main() {
	fmt.Println("Hello there.")
	//create port variable and set port on which to listen
	port := 3000

	// pass in port variable as teh function argument
	startWebServer(port)
}

// declares variable and type for port as function argument
func startWebServer(port int) {
	fmt.Println(("Starting server..."))
	// do server stuff
	// print text with port that is being listened
	fmt.Println("Server started on port", port)
}
