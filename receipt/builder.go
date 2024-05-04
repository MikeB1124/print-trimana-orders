package receipt

import (
	"bytes"
	"fmt"
	"time"

	"github.com/MikeB1124/escpos/esc"
	"github.com/MikeB1124/escpos/wix.go"
)

type CustomOrder struct {
	ID             string       `json:"id"`
	CustomerName   string       `json:"customerName"`
	CustomerNumber string       `json:"customerNumber"`
	Fulfillment    string       `json:"fulfillment"`
	DueDate        time.Time    `json:"dueDate"`
	PaymentType    string       `json:"paymentType"`
	OrderComment   string       `json:"orderComment"`
	Tax            string       `json:"tax"`
	Tip            string       `json:"tip"`
	SubTotal       string       `json:"subTotal"`
	Total          string       `json:"total"`
	Items          []CustomItem `json:"items"`
}

type CustomItem struct {
	Item    string   `json:"item"`
	Options []string `json:"options"`
	Comment string   `json:"comment"`
}

func FormatOrdersForPrinting(orders wix.WixOrdersResponse) []CustomOrder {
	var formattedOrders []CustomOrder

	for _, o := range orders.Orders {
		var formattedOrder CustomOrder
		formattedOrder.ID = o.ID
		formattedOrder.CustomerName = fmt.Sprintf("%s %s", o.Customer.FirstName, o.Customer.LastName)
		formattedOrder.CustomerNumber = o.Customer.Phone
		formattedOrder.Fulfillment = o.Fulfillment.Type
		formattedOrder.DueDate = o.Fulfillment.PromisedTime
		formattedOrder.PaymentType = o.Payments[0].Method
		formattedOrder.OrderComment = o.Comment
		formattedOrder.Tax = o.Totals.Tax
		formattedOrder.Tip = o.Totals.Tip
		formattedOrder.SubTotal = o.Totals.Subtotal
		formattedOrder.Total = o.Totals.Total

		//Parse all items and options
		for _, item := range o.LineItems {
			var formattedItem CustomItem
			formattedItem.Item = fmt.Sprintf("%d x %s $%s", item.Quantity, item.CatalogReference.CatalogItemName, item.Price)
			formattedItem.Comment = item.Comment

			var formattedOptions []string
			for _, option := range item.DishOptions {
				for _, selected := range option.SelectedChoices {
					formattedOptions = append(formattedOptions, fmt.Sprintf("%s: %s +%s", option.Name, selected.CatalogReference.CatalogItemName, selected.Price))
				}
			}
			formattedItem.Options = append(formattedItem.Options, formattedOptions...)
			formattedOrder.Items = append(formattedOrder.Items, formattedItem)
		}
		formattedOrders = append(formattedOrders, formattedOrder)
	}

	return formattedOrders
}

func ReceiptInit(buf bytes.Buffer) bytes.Buffer {
	buf.Write(esc.Init)
	buf.Write(esc.DoubleHeightMode)
	return buf
}

func ReceiptBusinessInfoHeader(buf bytes.Buffer) bytes.Buffer {
	buf.Write(esc.LineFeed)
	buf.Write(esc.CenterAlign)
	buf.Write(esc.StringToHexBytes("Trimana Grill"))
	buf.Write(esc.LineFeed)
	buf.Write(esc.StringToHexBytes("10920 Wilshire Blvd"))
	buf.Write(esc.LineFeed)
	buf.Write(esc.StringToHexBytes("Los Angeles, CA 90024"))
	buf.Write(esc.LineFeed)
	buf.Write(esc.StringToHexBytes("(310) 208-2946"))
	buf.Write(esc.FeedPaper)
	return buf
}

func ReceiptOrderDetails(order CustomOrder, buf bytes.Buffer) bytes.Buffer {
	buf.Write(esc.LeftAlign)
	buf.Write(esc.StringToHexBytes("------------------------------------------"))
	buf.Write(esc.LineFeed)
	buf.Write(esc.StringToHexBytes(fmt.Sprintf("Order number: %s", order.ID)))
	buf.Write(esc.LineFeed)
	buf.Write(esc.StringToHexBytes(fmt.Sprintf("Name: %s", order.CustomerName)))
	buf.Write(esc.LineFeed)
	buf.Write(esc.StringToHexBytes(fmt.Sprintf("Phone Number: %s", order.CustomerNumber)))
	buf.Write(esc.LineFeed)
	buf.Write(esc.StringToHexBytes(fmt.Sprintf("Payment: %s", order.PaymentType)))
	buf.Write(esc.LineFeed)
	buf.Write(esc.StringToHexBytes(fmt.Sprintf("Due Date: %s", order.DueDate)))
	buf.Write(esc.LineFeed)
	buf.Write(esc.StringToHexBytes(order.Fulfillment))
	buf.Write(esc.LineFeed)
	buf.Write(esc.StringToHexBytes("------------------------------------------"))
	buf.Write(esc.FeedPaper)
	return buf
}

func ReceiptItems(order CustomOrder, buf bytes.Buffer) bytes.Buffer {
	buf.Write(esc.LeftAlign)
	for _, item := range order.Items {
		buf.Write(esc.StringToHexBytes(item.Item))
		buf.Write(esc.LineFeed)
		for _, option := range item.Options {
			buf.Write(esc.StringToHexBytes(fmt.Sprintf("   %s", option)))
			buf.Write(esc.LineFeed)
		}
		buf.Write(esc.StringToHexBytes(item.Comment))
		buf.Write(esc.LineFeed)
	}
	buf.Write(esc.FeedPaper)
	return buf
}
