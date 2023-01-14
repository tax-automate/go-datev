package datev

type errToIndex = map[int]error

type BookingLogger interface {
	SetIdentifierForNextErrors(s string)
	addError(err error)
}

type Logger struct {
	identifier string
	values     map[string][]error
}

func (lgr *Logger) SetIdentifierForNextErrors(s string) {
	lgr.identifier = s
}

func (lgr *Logger) addError(index int, err error) {
	lgr.values[lgr.identifier] = append(lgr.values[lgr.identifier], err)
}

func NewBookingLogger() BookingLogger {
	return &Logger{
		values: make(map[string][]error, 0),
	}
}
