package printer

import (
	"fmt"
	"net"
)

type PrinterConfig struct {
	PrinterAddr   string
	PrinterPort   string
	NetConnection net.Conn
}

func (p *PrinterConfig) InitPrinter(ip string, port string) error {
	conn, err := connectToPrinter(ip, port)
	if err != nil {
		panic(err)
	}
	p.PrinterAddr = ip
	p.PrinterPort = port
	p.NetConnection = conn
	return nil
}

func connectToPrinter(ip string, port string) (net.Conn, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", ip, port))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to printer: %+v", err)
	}
	return conn, nil
}

func (p PrinterConfig) WriteToPrinter(cmd []byte) error {
	_, err := p.NetConnection.Write(cmd)
	if err != nil {
		return fmt.Errorf("failed to write to printer: %+v", err)
	}
	p.NetConnection.Close()
	return nil
}
