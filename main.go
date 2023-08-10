package main

import (
	"fmt"

	"github.com/xenedium/hps_logs_parser/iso8583/types"
)

func main() {
	instance := types.Message{
		MTI:    types.MTI{Version: 2, Class: 0, Function: 0, Origin: 0},
		Bitmap: []byte{0x00, 0x00, 0x00, 0x00, 0x00},
		Fields: map[int]string{
			2:  "1234567890123456",
			3:  "000000",
			4:  "000000000000",
			7:  "1234567890",
			11: "000001",
			12: "000000",
			13: "0101",
			15: "0101",
			18: "0000",
			22: "012",
			25: "00",
			26: "12",
			28: "C00000000",
		},
		Raw: []byte{0x02, 0x00, 0x00, 0x00, 0x00, 0x00},
	}

	fmt.Println("MTI:", instance.MTI)
	fmt.Println("Bitmap:", instance.Bitmap)
	fmt.Println("Fields:", instance.Fields)
	fmt.Println("Raw:", instance.Raw)
	fmt.Println("MTI:", instance.MTI)

	fmt.Println("MTI Class:", instance.GetMTIClassName())
	fmt.Println("MTI Function:", instance.GetMTIFunctionName())
	fmt.Println("MTI Origin:", instance.GetMTIOriginName())

}
