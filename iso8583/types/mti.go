package types

type mti struct {
	Version  uint8
	Class    uint8
	Function uint8
	Origin   uint8
}

var MessageOriginMap = map[int]string{
	0: "Acquirer",
	1: "Acquirer Repeat",
	2: "Issuer",
	3: "Issuer Repeat",
	4: "Other",
	6: "Reserved for ISO use",
}
var MessageFunctionMap = map[int]string{
	0: "Request",
	1: "Request Response",
	2: "Advice",
	3: "Advice Response",
	4: "Notification",
	5: "Notification Acknowledgement",
	6: "Instruction",
	7: "Instruction Acknowledgement",
	8: "Reserved for ISO use",
	9: "Reserved for ISO use",
}
var MessageClassMap = map[int]string{
	0: "Reserved for ISO use",
	1: "Authorization Message",
	2: "Financial Message",
	3: "File Actions Message",
	4: "Reversal and Chargeback Message",
	5: "Reconciliation Message",
	6: "Administrative Message",
	7: "Fee Collection Message",
	8: "Network Management Message",
	9: "Reserved for ISO use",
}

func (mti *mti) String() string {
	return string([]byte{mti.Version + '0', mti.Class + '0', mti.Function + '0', mti.Origin + '0'})
}

func (m *mti) setMTI(mti string) *mti {
	m.Version = mti[0] - '0'
	m.Class = mti[1] - '0'
	m.Function = mti[2] - '0'
	m.Origin = mti[3] - '0'
	return m
}

func (m *mti) getMTIBytes() []byte {
	return []byte{m.Version + '0', m.Class + '0', m.Function + '0', m.Origin + '0'}
}

func (m *mti) getMTIClassName() string {
	return MessageClassMap[int(m.Class)]
}

func (m *mti) getMTIFunctionName() string {
	return MessageFunctionMap[int(m.Function)]
}

func (m *mti) getMTIOriginName() string {
	return MessageOriginMap[int(m.Origin)]
}

func (m *mti) isRequest() bool {
	return m.Function%2 == 0
}

func (m *mti) isResponse() bool {
	return m.Function%2 == 1
}

type MTI = mti
