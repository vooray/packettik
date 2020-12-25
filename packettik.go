/* Try to connect to port on host periodically
 */

package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	var tikTime int // connect interval sec.
	var timeout int // connect timeout sec.
	var destination string
	var port int
	const statCycles = 10 // print stats every X cycles
	var counterSuccess, counterFail, counterCycles = 0, 0, 0
	var chanSuccess = make(chan bool)
	var logFilename string // log to file named logFilename

	flag.StringVar(&destination, "d", "", "destination host <mandatory>")
	flag.IntVar(&port, "p", 0, "destination port number <mandatory>")
	flag.IntVar(&tikTime, "i", 0, "check interval (sec) <mandatory>")
	flag.IntVar(&timeout, "t", 0, "session timeout (sec) <mandatory>")
	flag.StringVar(&logFilename, "l", "", "log to file [filename] <optional>")
	flag.Parse()

	if destination == "" || port == 0 || tikTime == 0 || timeout == 0 {
		fmt.Println("\n\t* Mandatory arguments missing")
		fmt.Println("\n\t* Example: packettik.exe -d google.com -p 443 -i 1 -t 1 -l google_com_443.log\n")
		fmt.Println(" Help:\n")
		flag.PrintDefaults()
		fmt.Println("\n")
		os.Exit(1)
	}

	if logFilename == "" {
		log.Println("Logging to console only")
	} else {
		logFile, err := os.OpenFile(logFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 666)
		if err != nil {
			log.Fatal(err)
		}

		log.SetOutput(io.MultiWriter(os.Stdout, logFile))
		log.SetFlags(log.LstdFlags)
		log.Println("Logging to console and file: ", logFilename)
	}

	connStr := destination + ":" + strconv.Itoa(port)

	for {
		counterCycles++
		go checkTcp(connStr, timeout, chanSuccess, counterCycles)
		time.Sleep(time.Duration(tikTime) * time.Second)

		if <-chanSuccess == true {
			counterSuccess++
		} else {
			counterFail++
		}

		if counterCycles%statCycles == 0 {
			log.Println("--- Stats: < Success: ", counterSuccess, "> / < Fail: ", counterFail, ">")
		}
	}
}

func checkTcp(connStr string, timeout int, chanSuccess chan bool, counterCycles int) {
	conn, err := net.DialTimeout("tcp", connStr, time.Duration(timeout)*time.Second)
	if err != nil {
		log.Println("[", counterCycles, "] Fail! - Failed to connect: <", connStr, "> Error:", err)
		chanSuccess <- false
	} else {
		log.Println("[", counterCycles, "] OK! - Connected to: <", connStr, " >")
		err := conn.Close()
		if err != nil {
			log.Println("Can not close connection! SNAFU!")
		}
		chanSuccess <- true
	}
}
