package log

const (
	alertFlag     = 0b00000001 // L1
	errorFlag     = 0b00000001 // L1
	warnFlag      = 0b00000010 // L2
	highlightFlag = 0b00000100 // L3
	informFlag    = 0b00001000 // L4
	logFlag       = 0b00010000 // L5
	traceFlag     = 0b00100000 // L6
)

func (l myLogger) Level(level uint8) Logger {
	logger := myLogger{Writer: l.Writer, prefix: l.prefix, begin: l.begin}
	switch level {
	case 1:
		logger.level = alertFlag | errorFlag
	case 2:
		logger.level = alertFlag | errorFlag | warnFlag
	case 3:
		logger.level = alertFlag | errorFlag | warnFlag | highlightFlag
	case 4:
		logger.level = alertFlag | errorFlag | warnFlag | highlightFlag | informFlag
	case 5:
		logger.level = alertFlag | errorFlag | warnFlag | highlightFlag | informFlag | logFlag
	default:
		logger.level = alertFlag | errorFlag | warnFlag | highlightFlag | informFlag | logFlag | traceFlag
	}
	return logger
}

func hasLevel(b uint8, flag uint8) bool {
	return b&flag != 0
}
