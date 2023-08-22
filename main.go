package main

import (
	"fmt"
	"os"

	"github.com/xenedium/hps_logs_parser/iso8583/scanner"
	"github.com/xenedium/hps_logs_parser/iso8583/types"
)

func typest() {
	instance := types.Message{
		MTI:    types.MTI{Version: 2, Class: 0, Function: 0, Origin: 0},
		Bitmap: "F23A40010A0000000000000",
		Fields: map[int]types.Field{
			2:  {Length: 6, Value: "123456"},
			3:  {Length: 6, Value: "123456"},
			4:  {Length: 6, Value: "123456"},
			39: {Length: 2, Value: "000"},
		},
		Raw: "0200F23A40010A0000000000000",
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

	if field, ok := instance.GetField(2); ok {
		fmt.Println("Field 2:", field)
	}

	if field, ok := instance.GetField(3); ok {
		fmt.Println("Field 3:", field)
	}

	if field, ok := instance.GetField(99); ok {
		fmt.Println("Field 99:", field)
	} else {
		fmt.Println("Field 99 not found")
	}

	instance.SetField(2, "654321")

	if field, ok := instance.GetField(2); ok {
		fmt.Println("Field 2:", field)
	}

	responseCodeMessage, ok := instance.GetResponseCodeMessage()
	if ok {
		fmt.Println("Response Code Message:", responseCodeMessage)
	}
}

func main() {
	f, err := os.Open("logs/POSTILION.TRC000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	fScanner := scanner.Scanner{
		File: f,
	}

	fScanner.Scan()

	for _, dumpPostilion := range fScanner.DumpTlvBuffers {
		fmt.Println(dumpPostilion)
	}
}
