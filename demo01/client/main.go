package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	log.Println("begin dial...")
	conn, err := net.DialTimeout("tcp", ":8888", 2*time.Second)
	if err != nil {
		fmt.Println("dial error:", err)
	}
	log.Println("dial ok")

	var data = make([]byte, 65536)
	total := 0
	for {
		n, err := conn.Write(data)
		if err != nil {
			log.Println("write %d bytes error, error: %s", n, err)
			break
		}
		total += n
		log.Println("write %d bytes this time, %d byte in total", n, total)
	}

	log.Println("[total] write %d bytes in total", total)
}
