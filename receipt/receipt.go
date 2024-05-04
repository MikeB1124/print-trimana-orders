package receipt

import (
	"bytes"

	"github.com/MikeB1124/print-trimana-orders/esc"
)

type EscFormattedReceipts struct {
	ID          string
	EscCommands []byte
}

func EscFormatReceipts(orders []CustomOrder) []EscFormattedReceipts {
	var receipts []EscFormattedReceipts
	for _, o := range orders {
		var receipt EscFormattedReceipts
		var escCmdBuffer bytes.Buffer
		escCmdBuffer = ReceiptInit(escCmdBuffer)
		escCmdBuffer = ReceiptBusinessInfoHeader(escCmdBuffer)
		escCmdBuffer = ReceiptOrderDetails(o, escCmdBuffer)
		escCmdBuffer = ReceiptItems(o, escCmdBuffer)
		escCmdBuffer = ReceiptFooter(o, escCmdBuffer)
		escCmdBuffer.Write(esc.FeedPaperAndCut)

		receipt.ID = o.ID
		receipt.EscCommands = escCmdBuffer.Bytes()
		receipts = append(receipts, receipt)
		escCmdBuffer.Reset()
	}
	return receipts
}
