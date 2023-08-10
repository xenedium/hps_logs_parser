package types

type message struct {
	MTI    mti
	Bitmap []byte
	Fields map[int]field
	Raw    []byte
}

func (m *message) String() string {
	return string(m.Raw)
}

func (m *message) SetMTI(mti string) *message {
	m.MTI.setMTI(mti)
	return m
}

func (m *message) GetMTIClassName() string {
	return m.MTI.getMTIClassName()
}

func (m *message) GetMTIFunctionName() string {
	return m.MTI.getMTIFunctionName()
}

func (m *message) GetMTIOriginName() string {
	return m.MTI.getMTIOriginName()
}

func (m *message) IsRequest() bool {
	return m.MTI.isRequest()
}

func (m *message) IsResponse() bool {
	return m.MTI.isResponse()
}

func (m *message) GetField(field int) (Field, bool) {
	f, ok := m.Fields[field]
	return f, ok
}

func (m *message) SetField(field int, value string) *message {
	m.Fields[field] = Field{Length: len(value), Value: value}
	return m
}

func (m *message) GetRaw() []byte {
	return m.Raw
}

type Message = message
