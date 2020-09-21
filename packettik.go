/* Try to connect to host port periodically
 */

package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func main() {
	const tikTime = 1
	const timeout = 1
	const hostname = "golang.org"
	const port = 4433
	const statsCycles = 10
	connStr := hostname + ":" + strconv.Itoa(port)
	var counterSuccess, counterFail, counterCycles int = 0, 0, 0
	var chanSuccess chan bool = make(chan bool)

	for {
		go checkTcp(connStr, timeout, chanSuccess)
		counterCycles++
		time.Sleep(time.Duration(tikTime) * time.Second)

		if <-chanSuccess == true {
			counterSuccess++
		} else {
			counterFail++
		}
		if counterCycles%statsCycles == 0 {
			fmt.Println("---Stats: < Success: ", counterSuccess, "> / < Fail: ", counterFail, ">")
		}
	}
}

func checkTcp(connStr string, timeout int, chanSuccess chan bool) {
	conn, err := net.DialTimeout("tcp", connStr, time.Duration(timeout)*time.Second)
	if err != nil {
		fmt.Println(time.Now().Format(time.ANSIC), "Fail! - Can not connect!", err)
		chanSuccess <- false
	} else {
		fmt.Println(time.Now().Format(time.ANSIC), "OK! - Connected!")
		conn.Close()
		chanSuccess <- true
	}
}
