package port

import (
	"net"
	"strconv"
	"time"
)

type ScanResult struct {
	Port  int
	Stage string
}

func ScanPort(protocol, hostname string, port int) ScanResult {
	result := ScanResult{Port: port}

	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		result.Stage = "closed"
		return result
	}
	defer conn.Close()

	result.Stage = "Open"
	return result
}

func InitialScan(hostname string) []ScanResult {
	var results []ScanResult

	for i := 1; i <= 1024; i++ {
		results = append(results, ScanPort("tcp", hostname, i))
	}

	return results
}
