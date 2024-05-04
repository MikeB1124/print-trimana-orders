package esc

var (
	Init               = []byte{0x1B, 0x40}
	LineFeed           = []byte{0x0A}
	FullCut            = []byte{0x1B, 0x6D}
	RegUnderline       = []byte{0x1B, 0x2D, 0x01}
	BoldUnderline      = []byte{0x1B, 0x2D, 0x02}
	DoubleHeightMode   = []byte{0x1B, 0x21, 0x10}
	DoubleWidthMode    = []byte{0x1B, 0x21, 0x10}
	DisableDoubleWidth = []byte{0x1B, 0x21, 0x00}
	LeftAlign          = []byte{0x1B, 0x61, 0x0}
	CenterAlign        = []byte{0x1B, 0x61, 0x1}
	RightAlign         = []byte{0x1B, 0x61, 0x2}
	FeedPaper          = []byte{0x1B, 0x64, 0x3}
	FeedPaperAndCut    = []byte{0x1B, 0x64, 0x5, 0x1B, 0x6D}
)

func StringToHexBytes(input string) []byte {
	// Initialize a byte slice to store the hexadecimal representations
	var hexBytes []byte

	// Iterate over each character in the input string
	for _, char := range input {
		// Convert the character to its ASCII value and append it to the byte slice
		hexBytes = append(hexBytes, byte(char))
	}

	return hexBytes
}
