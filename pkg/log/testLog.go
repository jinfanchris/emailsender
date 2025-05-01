package log

import (
	"bytes"
	"testing"

	"github.com/sirupsen/logrus"
)

// testLoggerHook is a custom hook for logrus to capture logs and redirect them to testing.T.
type testLoggerHook struct {
	t *testing.T
}

func (h *testLoggerHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *testLoggerHook) Fire(entry *logrus.Entry) error {
	buffer := new(bytes.Buffer)
	entry.Logger.SetOutput(buffer)
	entry.Logger.Formatter.Format(entry)
	return nil
}

// SetupTestLogger sets logrus to redirect logs to testing.T.
func SetupTestLogger(t *testing.T) {
	logrus.SetFormatter(&logrus.TextFormatter{}) // Adjust formatter if needed
	logrus.SetLevel(logrus.DebugLevel)           // Ensure all levels are captured
	logrus.AddHook(&testLoggerHook{t: t})
}
