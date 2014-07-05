package logging

import(
	"time"
)

type abstractLogger struct{
	name string
	logLevel uint8
	lastE error
}

func (logger *abstractLogger) ShouldLog(level uint8) bool {
	return IsLevel(level) && level == level & logger.logLevel
}

func (logger *abstractLogger) HasError() (bool, error) {
	if logger.lastE != nil {
		return true, logger.lastE
	} else {
		return false, nil
	}
}

func (logger *abstractLogger) SetName(to string) {
	logger.name = to
}

func (logger *abstractLogger) GetName() string {
	return logger.name
}

func (logger *abstractLogger) SetLevel(to uint8) {
	logger.logLevel = to
}

func (logger *abstractLogger) GetLevel() uint8 {
	return logger.logLevel
}

func (logger *abstractLogger) NowClock() string {
	return logger.ClockTime(time.Now())
}

func (logger *abstractLogger) ClockTime(t time.Time) string {
	return t.Format(TIME_FORMAT)
}
