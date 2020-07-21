package main

import (
	"fmt"

	"RESTapi_go/portscaneer/port"
	// "github.com/elliotforbes/athena/port"
)

func main() {
	fmt.Println("port scaner in go")

	open := port.ScanPort("tcp", "localhost", 1314)
	fmt.Printf("port open: %t\n", open)

	results := port.InitialScan("localhost")
	fmt.Println("port scaneer results", results)
}
