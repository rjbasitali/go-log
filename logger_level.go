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

// Level returns a leveled logger, level can be from 1 to 6.
//		Level 1 — Alert or Error
//		Level 2 — Warn
//		Level 3 — Highlight
//		Level 4 — Inform
//		Level 5 — Log
//		Level 6 — Trace
// Anything above 6 as level will be considered Level 6.
// Pass 0 to output no logs.
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
	case 6:
		logger.level = alertFlag | errorFlag | warnFlag | highlightFlag | informFlag | logFlag | traceFlag
	default:
		logger.level = 0
	}
	return logger
}

func hasLevel(b uint8, flag uint8) bool {
	return b&flag != 0
}
