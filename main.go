package main

import (
	"encoding/json"

	"github.com/MikeB1124/escpos"
	"github.com/MikeB1124/print-trimana-orders/configuration"
	"github.com/MikeB1124/print-trimana-orders/logger"
	"github.com/MikeB1124/print-trimana-orders/receipt"
	"github.com/MikeB1124/print-trimana-orders/wix"
)

func main() {
	logger.InfoLogger.Println("----------------------------------------------")
	logger.InfoLogger.Println("START PRINTING")
	configuration.Init()

	//Get Wix Orders
	orders, err := wix.GetWixOrders()
	if err != nil {
		logger.ErrorLogger.Printf("Failed to get wix orders: %+v\n", err)
		return
	}
	logger.InfoLogger.Printf("%d Wix orders found.\n", len(orders.Orders))

	//Check if any order available for printing
	if len(orders.Orders) == 0 {
		logger.InfoLogger.Println("0 orders available for printing.")
		logger.InfoLogger.Println("END PRINTING")
		logger.InfoLogger.Println("----------------------------------------------")
		return
	}

	//Parse and format orders
	formattedOrders := receipt.FormatOrdersForPrinting(orders)
	jsonFormat, _ := json.MarshalIndent(formattedOrders, "", "\t")
	logger.InfoLogger.Printf("%+v", string(jsonFormat))

	//Get esc commands from formatted orders
	escFormattedReceipts := receipt.EscFormatReceipts(formattedOrders)

	//Init printers
	printerConfigs := []escpos.PrinterConfig{}
	for _, p := range configuration.Config.Printers {
		printer := &escpos.PrinterConfig{}
		printer.InitPrinter(p.IP, p.Port)
		printerConfigs = append(printerConfigs, *printer)
	}
	logger.InfoLogger.Printf("Configured all printers %+v\n", printerConfigs)

	//Print Orders
	for _, p := range printerConfigs {
		for _, o := range escFormattedReceipts {
			err := p.WriteToPrinter(o.EscCommands)
			if err != nil {
				logger.ErrorLogger.Printf("Printer %s:%s failed to print order# %s: %+v\n", p.PrinterAddr, p.PrinterPort, o.ID, err)
			} else {
				logger.InfoLogger.Printf("Order# %s has been printed by %s:%s\n", o.ID, p.PrinterAddr, p.PrinterPort)
				_, err := wix.AcceptWixOrder(o.ID)
				if err != nil {
					logger.ErrorLogger.Printf("Could not accept wix order# %s, Error: %+v\n", o.ID, err)
				} else {
					logger.InfoLogger.Printf("Order# %s has been accepted.\n", o.ID)
				}
			}
		}
		p.NetConnection.Close()
	}
	logger.InfoLogger.Println("END PRINTING")
	logger.InfoLogger.Println("----------------------------------------------")
}
