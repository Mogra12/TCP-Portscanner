package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

type ServiceInfo struct {
	Type 	string `json:"type"`
	Version string `json:"version"`
}

func main() {
        var hostname string
        var allPorts bool
        var Espport int

        // go run main.go -h <HOSTNAME> <OPTIONS>
        flag.StringVar(&hostname, "h", "", "Hostname")
        flag.BoolVar(&allPorts, "Ap", false, "Enable scan all ports")
        flag.IntVar(&Espport, "p", 80, "Verify if especified port is open")
        flag.Parse()

        if hostname == "" {
            fmt.Println("Hostname not defined!")
			os.Exit(1)
        }

		if allPorts {
			// verify is it localhost scan
			if hostname == "127.0.0.1" {
				fmt.Printf("Scan report from localhost %s\n", hostname)
			} else {
				fmt.Printf("Scan report from %s\n", hostname)
			}

			// print colunms name
			fmt.Println("Port	 Status  Service  Version")

			// define ports in range 1 - 65535
			for port := 1; port < 65535; port++ {
				// try all ports in range 65535
				inf := hostname + ":" + strconv.Itoa(port)
				conn, err := net.DialTimeout("tcp", inf, 3*time.Second)
				if err != nil {
					continue	
				}
				defer conn.Close()

				// read timeout
				conn.SetReadDeadline(time.Now().Add(3 * time.Second))
				reader := bufio.NewReader(conn)

				banner, err := reader.ReadString('\n')
				if err != nil {
					fmt.Printf("%d    OPEN    Unknown\n", port)
                	continue
				}

				// remove blank spaces
				banner = strings.TrimSpace(banner)

				// try convert json banner 
				var service ServiceInfo
				err = json.Unmarshal([]byte(banner), &service)
				
				if err != nil {
					fmt.Printf("%d    OPEN    Unknown\n", port)
				} else {
					fmt.Printf("%d    OPEN    %v    %v\n", port, service.Type, service.Version)
				}
				// close connection
				conn.Close()
			}
			return
		}

		// verify an specific port in host
		inf := hostname + ":" + strconv.Itoa(Espport)
		conn, err := net.Dial("tcp", inf)
		if err != nil {
			log.Printf("Connection error to %s: %v\n", inf, err)
			os.Exit(1)		
		} else {
			fmt.Println("<========================>")
			fmt.Printf("%v:%v\n", hostname, Espport)
			fmt.Println("<========================>")
			fmt.Printf("Port %v OPEN\n", Espport)
			fmt.Println("<========================>")
			// close connetion
			conn.Close()
		}
}