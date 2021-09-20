package main

// import block
import (
	"fmt"
)

/*
gRPC API: Attempt 1; On Mac
Helpful links: https://grpc.io/docs/languages/go/basics/
*/
func main() {
	fmt.Println("Hello there.")
	//create port variable and set port on which to listen
	port := 3000

	// pass in port variable as the function argument
	// ignore returned port; only get actual error if error happens
	_, err := startWebServer(port)
	fmt.Println(err)

}

// declares variable and type for port as function argument
// returns as error (for now)
func startWebServer(port int) (int, error) {
	fmt.Println(("Starting server..."))
	// do server stuffß
	// print text with port that is being listened
	fmt.Println("Server started on port", port)

	//port number and error returned from second var in function (not arg)
	return port, nil
}
