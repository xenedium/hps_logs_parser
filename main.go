package main

import (
	"fmt"

	"github.com/xenedium/hps_logs_parser/iso8583/types"
)

func main() {
	instance := types.Message{
		MTI:    types.MTI{Version: 2, Class: 0, Function: 0, Origin: 0},
		Bitmap: []byte{0x00, 0x00, 0x00, 0x00, 0x00},
		Fields: map[int]types.Field{
			2: {Length: 6, Value: "123456"},
			3: {Length: 6, Value: "123456"},
			4: {Length: 6, Value: "123456"},
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
