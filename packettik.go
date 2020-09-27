/* Try to connect to port on host periodically
 */

package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func main() {
	const tikTime = 1 // connect interval sec.
	const timeout = 1 // connect timeout sec.
	const hostname = "golang.org"
	const port = 443
	const statCycles = 10 // print stats every X cycles
	var counterSuccess, counterFail, counterCycles = 0, 0, 0
	var chanSuccess = make(chan bool)

	connStr := hostname + ":" + strconv.Itoa(port)

	for {
		go checkTcp(connStr, timeout, chanSuccess)
		counterCycles++
		time.Sleep(time.Duration(tikTime) * time.Second)

		if <-chanSuccess == true {
			counterSuccess++
		} else {
			counterFail++
		}
		if counterCycles%statCycles == 0 {
			fmt.Println("--- Stats: < Success: ", counterSuccess, "> / < Fail: ", counterFail, ">")
		}
	}
}

func checkTcp(connStr string, timeout int, chanSuccess chan bool) {
	conn, err := net.DialTimeout("tcp", connStr, time.Duration(timeout)*time.Second)
	if err != nil {
		fmt.Println(time.Now().Format(time.ANSIC), "Fail! - Can not connect <", connStr, "> Error:", err)
		chanSuccess <- false
	} else {
		fmt.Println(time.Now().Format(time.ANSIC), "OK! - Connected <", connStr, " >")
		err := conn.Close()
		if err != nil {
			fmt.Println("Can not close connection! SNAFU!")
		}
		chanSuccess <- true
	}
}
