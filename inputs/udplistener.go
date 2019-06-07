package inputs

import (
	"fmt"
	"net"

	"github.com/nomenclate/forzago"
)

type UdpListener struct {
	dataOutSrc net.IP
	port       int
}

func NewUdpListener(port int) *UdpListener {
	return &UdpListener{port: port}
}

func (l *UdpListener) StartAccepting(q forzago.Queue) {
	fmt.Printf("Starting UDP listening on port %d\n", l.port)

	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", l.port))
	if err != nil {
		fmt.Println(err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	buf := make([]byte, 1024)

	for {
		n, addr, err := conn.ReadFromUDP(buf)
		if !addr.IP.Equal(l.dataOutSrc) {
			break
		}
		if err != nil {
			fmt.Println("Error: ", err)
		}

		message := buf[0:n]
		q.Enqueue(message)
	}
}
