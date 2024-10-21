package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	// Mengatur alamat IP dan port untuk NTP server
	addr, err := net.ResolveUDPAddr("udp", ":123")
	if err != nil {
		fmt.Println("Error resolving address:", err)
		os.Exit(1)
	}

	// Membuat socket UDP
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("NTP server berjalan di", addr.String())

	for {
		// Menerima data dari client
		buf := make([]byte, 1024) // Menggunakan buffer yang cukup besar untuk menerima data
		n, clientAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error reading from UDP:", err)
			continue
		}
		fmt.Printf("Permintaan diterima dari %s, %d bytes\n", clientAddr, n)

		// Mengambil waktu sistem saat ini
		currentTime := time.Now().Unix()

		// Mengirimkan waktu ke client
		response := []byte(fmt.Sprintf("%d", currentTime))
		_, err = conn.WriteToUDP(response, clientAddr)
		if err != nil {
			fmt.Println("Error sending response:", err)
		} else {
			fmt.Printf("Mengirimkan waktu ke %s: %d\n", clientAddr, currentTime)
		}
	}
}
