package util

type InvalidLocationId struct{}

func (m *InvalidLocationId) Error() string {
	return "invalid location id"
}

type InvalidEventId struct{}

func (m *InvalidEventId) Error() string {
	return "invalid event id"
}
