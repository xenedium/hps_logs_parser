package types

type message_origin_map map[int]string

func (m *message_origin_map) init() *message_origin_map {
	(*m)[0] = "Acquirer"
	(*m)[1] = "Acquirer Repeat"
	(*m)[2] = "Issuer"
	(*m)[3] = "Issuer Repeat"
	(*m)[4] = "Other"
	(*m)[6] = "Reserved for ISO use"
	return m
}

func (m *message_origin_map) getOrigin(origin uint8) string {
	return (*m)[int(origin)]
}
