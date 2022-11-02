package util

type InvalidLocationId struct{}

func (m *InvalidLocationId) Error() string {
	return "invalid location id"
}
