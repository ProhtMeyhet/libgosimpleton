package logging

import(
	"fmt"
	"time"
)

type abstractLogger struct{
	name string
	lastE error
	interfaceConfig LogConfigInterface
}

func (logger *abstractLogger) initialise(config LogConfigInterface) {
	logger.interfaceConfig = config
}

func (logger *abstractLogger) ShouldLog(level uint8) bool {
	return level == level & logger.interfaceConfig.GetLevel()
}

func (logger *abstractLogger) SetName(to string) {
	logger.name = to
}

func (logger *abstractLogger) GetName() string {
	return logger.name
}

func (logger *abstractLogger) SetLevel(to uint8) {
	logger.interfaceConfig.SetLevel(to)
}

func (logger *abstractLogger) GetLevel() uint8 {
	return logger.interfaceConfig.GetLevel()
}

func (logger *abstractLogger) NowClock() string {
	return logger.ClockTime(time.Now())
}

func (logger *abstractLogger) ClockTime(t time.Time) string {
	return t.Format(TIME_FORMAT)
}

/* go */ func (abstract *abstractLogger) run() {
infinite:
	for {
		select {
		case logEntry, ok := <-logging:
			if !ok { break infinite }
			formatted := ""
			for _, message := range logEntry.message {
				formatted = fmt.Sprintf(logEntry.format, message)
			}
			logger.Log(logEntry.level, formatted)
		}
	}
}
