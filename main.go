package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func scanPort(host string, port int, results *os.File) {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", address, 1*time.Second)
	if err != nil {
		results.WriteString(fmt.Sprintf("[-] Puerto %d cerrado\n", port))
		return
	}
	conn.Close()
	results.WriteString(fmt.Sprintf("[+] Puerto %d abierto\n", port))
}

func main() {

	var host string
	var startPort, endPort int

	fmt.Println("Ingresa la IP o dominio objetivo (e.g, 192.168.1.1): ")
	fmt.Scan(&host)

	fmt.Println("Ingresa el puerto inicial: ")
	fmt.Scan(&startPort)

	fmt.Println("Ingresa lel puerto final: ")
	fmt.Scan(&endPort)

	fmt.Printf("Escaneando %s desde el puerto %d al %d...\n ", host, startPort, endPort)

	results, err := os.Create("resultados.txt")
	if err != nil {
		fmt.Println("Error al crear archivo: ", err)
		return
	}
	defer results.Close()

	for port := startPort; port <= endPort; port++ {
		scanPort(host, port, results)
	}
	fmt.Println("Escaneo completado. Resultados guardados en 'resultados.txt'.")
}
