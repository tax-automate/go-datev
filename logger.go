package datev

type BookingLogger struct {
	values map[int][]string
}

func NewBookingLogger() *BookingLogger {
	return &BookingLogger{}
}

func (lgr *BookingLogger) addError(index int, errs []string) {
	lgr.values[index] = errs
}
