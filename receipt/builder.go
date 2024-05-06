package receipt

import (
	"bytes"
	"fmt"
	"time"

	"github.com/MikeB1124/escpos"
	"github.com/MikeB1124/print-trimana-orders/wix"
)

type CustomOrder struct {
	ID              string       `json:"id"`
	CustomerName    string       `json:"customerName"`
	CustomerNumber  string       `json:"customerNumber"`
	Fulfillment     string       `json:"fulfillment"`
	DueDate         string       `json:"dueDate"`
	PaymentType     string       `json:"paymentType"`
	OrderComment    string       `json:"orderComment"`
	Tax             string       `json:"tax"`
	Tip             string       `json:"tip"`
	SubTotal        string       `json:"subTotal"`
	Total           string       `json:"total"`
	Items           []CustomItem `json:"items"`
	DeliveryAddress string       `json:"deliveryAddress"`
}

type CustomItem struct {
	Item    string   `json:"item"`
	Options []string `json:"options"`
	Comment string   `json:"comment"`
}

var paymentTypeMap = map[string]string{
	"offline":    "CASH",
	"creditCard": "CREDIT CARD",
}

func FormatOrdersForPrinting(orders wix.WixOrdersResponse) []CustomOrder {
	var formattedOrders []CustomOrder

	for _, o := range orders.Orders {
		var formattedOrder CustomOrder
		formattedOrder.ID = o.ID
		formattedOrder.CustomerName = fmt.Sprintf("%s %s", o.Customer.FirstName, o.Customer.LastName)
		formattedOrder.CustomerNumber = formatPhoneNumber(o.Customer.Phone)
		formattedOrder.Fulfillment = o.Fulfillment.Type
		formattedOrder.DueDate = formatDateTime(o.Fulfillment.PromisedTime)
		formattedOrder.PaymentType = paymentTypeMap[o.Payments[0].Method]
		formattedOrder.OrderComment = o.Comment
		formattedOrder.Tax = o.Totals.Tax
		formattedOrder.Tip = o.Totals.Tip
		formattedOrder.SubTotal = o.Totals.Subtotal
		formattedOrder.Total = o.Totals.Total
		if o.Fulfillment.Type == "DELIVERY" {
			formattedOrder.DeliveryAddress = fmt.Sprintf(
				"%s %s #%s",
				o.Fulfillment.DeliveryDetails.Address.StreetNumber,
				o.Fulfillment.DeliveryDetails.Address.Street,
				o.Fulfillment.DeliveryDetails.Address.Apt,
			)
		}

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
	buf.Write(escpos.Init)
	buf.Write(escpos.DoubleHeightMode)
	return buf
}

func ReceiptBusinessInfoHeader(buf bytes.Buffer) bytes.Buffer {
	buf.Write(escpos.LineFeed)
	buf.Write(escpos.CenterAlign)
	buf.Write(escpos.StringToHexBytes("Trimana Grill"))
	buf.Write(escpos.LineFeed)
	buf.Write(escpos.StringToHexBytes("10920 Wilshire Blvd"))
	buf.Write(escpos.LineFeed)
	buf.Write(escpos.StringToHexBytes("Los Angeles, CA 90024"))
	buf.Write(escpos.LineFeed)
	buf.Write(escpos.StringToHexBytes("(310) 208-2946"))
	buf.Write(escpos.FeedPaper)
	return buf
}

func ReceiptOrderDetails(order CustomOrder, buf bytes.Buffer) bytes.Buffer {
	buf.Write(escpos.LeftAlign)
	buf.Write(escpos.StringToHexBytes("------------------------------------------"))
	buf.Write(escpos.LineFeed)
	buf.Write(escpos.StringToHexBytes(fmt.Sprintf("Order number: %s", order.ID)))
	buf.Write(escpos.LineFeed)
	buf.Write(escpos.StringToHexBytes(fmt.Sprintf("Name: %s", order.CustomerName)))
	buf.Write(escpos.LineFeed)
	buf.Write(escpos.StringToHexBytes(fmt.Sprintf("Phone Number: %s", order.CustomerNumber)))
	buf.Write(escpos.LineFeed)
	buf.Write(escpos.StringToHexBytes(fmt.Sprintf("Payment: %s", order.PaymentType)))
	buf.Write(escpos.LineFeed)
	buf.Write(escpos.StringToHexBytes(fmt.Sprintf("Due Date: %s", order.DueDate)))
	if order.DeliveryAddress != "" {
		buf.Write(escpos.LineFeed)
		buf.Write(escpos.StringToHexBytes(fmt.Sprintf("Delivery Address: %s", order.DeliveryAddress)))
	}
	buf.Write(escpos.LineFeed)
	buf.Write(escpos.StringToHexBytes(order.Fulfillment))
	buf.Write(escpos.LineFeed)
	buf.Write(escpos.StringToHexBytes("------------------------------------------"))
	buf.Write(escpos.FeedPaper)
	return buf
}

func ReceiptItems(order CustomOrder, buf bytes.Buffer) bytes.Buffer {
	buf.Write(escpos.LeftAlign)
	for _, item := range order.Items {
		buf.Write(escpos.StringToHexBytes(item.Item))
		buf.Write(escpos.LineFeed)
		for _, option := range item.Options {
			buf.Write(escpos.StringToHexBytes(fmt.Sprintf("    %s", option)))
			buf.Write(escpos.LineFeed)
		}
		if item.Comment != "" {
			buf.Write(escpos.StringToHexBytes(fmt.Sprintf("    Comment: %s", item.Comment)))
			buf.Write(escpos.LineFeed)
		}
		buf.Write(escpos.LineFeed)
	}
	return buf
}

func ReceiptTotals(order CustomOrder, buf bytes.Buffer) bytes.Buffer {
	buf.Write(escpos.RightAlign)
	buf.Write(escpos.LineFeed)
	buf.Write(escpos.StringToHexBytes(fmt.Sprintf("Subtotal: $%s", order.SubTotal)))
	buf.Write(escpos.LineFeed)
	buf.Write(escpos.StringToHexBytes(fmt.Sprintf("Tax: $%s", order.Tax)))
	buf.Write(escpos.LineFeed)
	if order.Tip != "" {
		buf.Write(escpos.StringToHexBytes(fmt.Sprintf("Tip: $%s", order.Tip)))
		buf.Write(escpos.LineFeed)
	}
	buf.Write(escpos.StringToHexBytes(fmt.Sprintf("Total: $%s", order.Total)))
	buf.Write(escpos.FeedPaper)
	return buf
}

func ReceiptFooter(order CustomOrder, buf bytes.Buffer) bytes.Buffer {
	buf.Write(escpos.LineFeed)
	buf.Write(escpos.LeftAlign)
	if order.OrderComment != "" {
		buf.Write(escpos.StringToHexBytes(fmt.Sprintf("Order Comment: %s", order.OrderComment)))
		buf.Write(escpos.LineFeed)
	}
	buf.Write(escpos.CenterAlign)
	buf.Write(escpos.StringToHexBytes("Thank you!"))
	buf.Write(escpos.FeedPaper)
	return buf
}

func formatPhoneNumber(number string) string {
	return fmt.Sprintf("(%s) %s-%s", number[2:5], number[5:8], number[8:])
}

func formatDateTime(dueDate string) string {
	t, _ := time.Parse(time.RFC3339Nano, dueDate)
	pst, _ := time.LoadLocation("America/Los_Angeles")
	t = t.In(pst)
	formattedTime := t.Format("January 2, 2006 at 3:04 PM")
	return formattedTime
}
