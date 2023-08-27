package types

type message struct {
	MTI         mti
	Bitmap      string
	Fields      map[string]field
	Raw         string
	LogFileName string
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

func (m *message) GetField(field string) (Field, bool) {
	f, ok := m.Fields[field]
	return f, ok
}

func (m *message) SetField(field string, value string) *message {
	m.Fields[field] = Field{Length: int64(len(value)), Value: value, Raw: []byte(value)}
	return m
}

func (m *message) GetRaw() string {
	return m.Raw
}

func (m *message) GetResponseCodeMessage() (string, bool) {
	f, ok := m.Fields["39"]
	if !ok {
		return "", false
	}
	return ResponseCodeMap[f.Value], true
}

type Message = message
