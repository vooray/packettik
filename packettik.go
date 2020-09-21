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
	const tik_time = 1
	const timeout = 1
	const hostname = "golang.org"
	const port = 443
	connStr := hostname + ":" + strconv.Itoa(port)
	//var success int = 0

	for {
		go checkTcp(connStr, timeout)
		time.Sleep(time.Duration(tik_time) * time.Second)
	}
}

func checkTcp(connStr string, timeout int) {
	conn, err := net.DialTimeout("tcp", connStr, time.Duration(timeout)*time.Second)
	if err != nil {
		fmt.Println(time.Now().Format(time.ANSIC), "Fail! - Can not connect: ", err)
	} else {
		fmt.Println(time.Now().Format(time.ANSIC), "OK! - Connected!")
		conn.Close()
	}
}
