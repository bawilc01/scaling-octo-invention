package main

// import block
import (
	"errors"
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
	// bool returned from function whether server started or not; printed here as err if error is returned
	err := startWebServer(port, 2)
	fmt.Println(err)

}

// declares variable and type for port as function argument
func startWebServer(port int, numberOfRetries int) error {
	fmt.Println(("Starting server..."))
	// do server stuff
	// print text with port that is being listened
	fmt.Println("Server started on port", port)

	//number of retries if port doesn't start on first try
	fmt.Println("Number of retries", numberOfRetries)
	return errors.New("Something went wrong.")
}
