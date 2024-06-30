package models

type Data []byte

// Set implements flag.Value.
func (m *Data) Set(value string) error {
	*m = []byte(value)
	return nil
}

// String implements flag.Value.
func (m *Data) String() string {
	return string(*m)
}
