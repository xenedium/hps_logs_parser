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
			2:  {Length: 6, Value: "123456"},
			3:  {Length: 6, Value: "123456"},
			4:  {Length: 6, Value: "123456"},
			39: {Length: 2, Value: "000"},
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
	fmt.Println("Is Request:", instance.IsRequest())
	fmt.Println("Is Response:", instance.IsResponse())

	field, ok := instance.GetField(2)
	if ok {
		fmt.Println("Field 2:", field)
	}

	field, ok = instance.GetField(3)
	if ok {
		fmt.Println("Field 3:", field)
	}

	field, ok = instance.GetField(99)
	if ok {
		fmt.Println("Field 99:", field)
	} else {
		fmt.Println("Field 99 not found")
	}

	instance.SetField(2, "654321")

	field, ok = instance.GetField(2)
	if ok {
		fmt.Println("Field 2:", field)
	}

	responseCodeMessage, ok := instance.GetResponseCodeMessage()
	if ok {
		fmt.Println("Response Code Message:", responseCodeMessage)
	}

}
