package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/MikeB1124/escpos"
)

func TestLocalPrinting(t *testing.T) {
	ip := "192.168.254.22"
	port := "9100"
	p := escpos.PrinterConfig{}
	p.InitPrinter(ip, port)

	err := p.WriteToPrinter(getEscCommands(ip, port))
	if err != nil {
		t.Errorf("error trying to pring: %+v", err)
		return
	}
	p.NetConnection.Close()
	t.Log("Printed Receipt!")
}

func TestRemotePrinting(t *testing.T) {
	ip := "47.147.239.47"
	port := "7100"
	p := escpos.PrinterConfig{}
	p.InitPrinter(ip, port)

	err := p.WriteToPrinter(getEscCommands(ip, port))
	if err != nil {
		t.Errorf("error trying to pring: %+v", err)
		return
	}
	p.NetConnection.Close()
	t.Log("Printed Receipt!")
}

func getEscCommands(ip string, port string) []byte {
	var escCmdBuffer bytes.Buffer
	escCmdBuffer.Write(escpos.Init)
	escCmdBuffer.Write(escpos.CenterAlign)
	escCmdBuffer.Write(escpos.DoubleHeightMode)
	escCmdBuffer.Write(escpos.LineFeed)
	escCmdBuffer.Write(escpos.StringToHexBytes(fmt.Sprintf("IPv4 Address: %s", ip)))
	escCmdBuffer.Write(escpos.LineFeed)
	escCmdBuffer.Write(escpos.StringToHexBytes(fmt.Sprintf("Port: %s", port)))
	escCmdBuffer.Write(escpos.LineFeed)
	escCmdBuffer.Write(escpos.LineFeed)
	escCmdBuffer.Write(escpos.LineFeed)
	escCmdBuffer.Write(escpos.FeedPaperAndCut)
	return escCmdBuffer.Bytes()
}
