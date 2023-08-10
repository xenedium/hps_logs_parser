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

type Message = message
