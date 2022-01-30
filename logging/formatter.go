package logging

// // Formatter implements logrus.Formatter interface.
// type SimpleFormatter struct {
// 	// Timestamp format
// 	TimestampFormat string
// 	// Available standard keys: time, msg, lvl
// 	// Also can include custom fields but limited to strings.
// 	// All of fields need to be wrapped inside %% i.e %time% %msg%
// 	LogFormat string
// }

// func (f *SimpleFormatter) Format(entry *logrus.Entry) ([]byte, error) {
// 	var name string
// 	if n, ok := entry.Data["name"]; ok {
// 		name = fmt.Sprintf(" %s", entry.Data["name"].(string))
// 	}

// 	output := fmt.Sprintf("[%s] harness-go-sdk%s %s\n", strings.ToUpper(entry.Level.String()), name, entry.Message)

// 	return []byte(output), nil
// }
