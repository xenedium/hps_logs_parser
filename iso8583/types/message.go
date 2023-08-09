package types

type Message struct {
	MTI    string
	Bitmap []byte
	Fields map[int]string
	Raw    []byte
}

func IsRequest(m *Message) bool {
	return m.MTI[2] == '0' || m.MTI[2] == '2' || m.MTI[2] == '4' || m.MTI[2] == '6' || m.MTI[2] == '8'
}

func IsResponse(m *Message) bool {
	return m.MTI[2] == '1' || m.MTI[2] == '3' || m.MTI[2] == '5' || m.MTI[2] == '7' || m.MTI[2] == '9'
}

func GetOrigin(m *Message) uint8 {
	return m.MTI[3] - '0'
}

func GetOriginName(m *Message) string {
	return (&message_origin_map{}).init().getOrigin(m.MTI[3] - '0')
}
