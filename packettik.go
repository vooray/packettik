/* Try to connect to port on host periodically
 */

package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	var tikTime int // connect interval sec.
	var timeout int // connect timeout sec.
	var hostname string
	var port int
	const statCycles = 10 // print stats every X cycles
	var counterSuccess, counterFail, counterCycles = 0, 0, 0
	var chanSuccess = make(chan bool)

	flag.StringVar(&hostname, "h", "", "destination hostname")
	flag.IntVar(&port, "p", 0, "destination port number")
	flag.IntVar(&tikTime, "i", 0, "check interval (sec)")
	flag.IntVar(&timeout, "t", 0, "timeout (sec)")
	flag.Parse()

	if hostname == "" || port == 0 || tikTime == 0 || timeout == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

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
