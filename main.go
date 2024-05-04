package main

import (
	"fmt"

	"github.com/MikeB1124/escpos"
	"github.com/MikeB1124/print-trimana-orders/configuration"
	"github.com/MikeB1124/print-trimana-orders/receipt"
	"github.com/MikeB1124/print-trimana-orders/wix.go"
)

func main() {
	configuration.Init()
	//Init printers
	printers := []map[string]string{{"ip": "192.168.86.29", "port": "9100"}}
	printerConfigs := []escpos.PrinterConfig{}
	for _, p := range printers {
		printer := &escpos.PrinterConfig{}
		printer.InitPrinter(p["ip"], p["port"])
		printerConfigs = append(printerConfigs, *printer)
	}

	//Get Wix Orders
	orders, err := wix.GetWixOrders()
	if err != nil {
		panic(fmt.Errorf("failed to get wix orders: %+v", err))
	}

	//Check if any order available for printing
	if len(orders.Orders) == 0 {
		fmt.Println("No orders available to print.")
		return
	}

	//Parse and format orders
	formattedOrders := receipt.FormatOrdersForPrinting(orders)
	// jsonFormat, _ := json.MarshalIndent(formattedOrders, "", "\t")
	// fmt.Println(string(jsonFormat))

	//Get esc commands from formatted orders
	escFormattedReceipts := receipt.EscFormatReceipts(formattedOrders)

	//Print Orders
	for _, p := range printerConfigs {
		for _, o := range escFormattedReceipts {
			err := p.WriteToPrinter(o.EscCommands)
			if err != nil {
				fmt.Printf("Printer %s:%s failed to print order# %s: %+v\n", p.PrinterAddr, p.PrinterPort, o.ID, err)
			} else {
				fmt.Printf("Order# %s has been printed by %s:%s\n", o.ID, p.PrinterAddr, p.PrinterPort)
			}
		}
		p.NetConnection.Close()
	}
}
