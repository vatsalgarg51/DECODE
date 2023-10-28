package main

import (
	"encoding/binary"
	"fmt"
)

// DecodedMessage represents decoded message structure
type DecodedMessage struct {
	ShortNumber1 int16
	CharBytes1   string
	Byte1        byte
	CharBytes2   string
	ShortNumber2 int16
	CharBytes3   string
	LongBytes1   int32
}

func decodePacket(packet []byte) DecodedMessage {
	if len(packet) != 44 {
		panic("Invalid packet")
	}

	data := DecodedMessage{}
	data.ShortNumber1 = int16(binary.BigEndian.Uint16(packet[0:2]))
	data.CharBytes1 = string(packet[2:14])
	data.Byte1 = packet[14]
	data.CharBytes2 = string(packet[15:23])
	data.ShortNumber2 = int16(binary.BigEndian.Uint16(packet[23:25]))
	data.CharBytes3 = string(packet[25:40])
	data.LongBytes1 = int32(binary.BigEndian.Uint32(packet[40:44]))

	return data
}

func main() {
	packet := []byte{
		0x04, 0xD2, 0x6B, 0x65, 0x65, 0x70, 0x64, 0x65,
		0x63, 0x6F, 0x64, 0x69, 0x6E, 0x67, 0x38, 0x64,
		0x6F, 0x6E, 0x74, 0x73, 0x74, 0x6F, 0x70, 0x03,
		0x15, 0x63, 0x6F, 0x6E, 0x67, 0x72, 0x61, 0x74,
		0x75, 0x6C, 0x61, 0x74, 0x69, 0x6F, 0x6E, 0x73,
		0x07, 0x5B, 0xCD, 0x15,
	}

	decodedStruct := decodePacket(packet)
	fmt.Printf("Decoded struct is: %+v\n", decodedStruct)
}
