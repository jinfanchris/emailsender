package log

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var (
	// logrus   = logrus.New()
	RUNTIME  = "./runtime"
	logFn    = "app.log"
	logfiles = []string{}
	loglevel = logrus.TraceLevel
)

func Join(p ...string) string {
	fp := filepath.Join(p...)
	dir := filepath.Dir(fp)
	os.Mkdir(dir, os.ModePerm)
	return fp
}

type nullWriter struct{}

func (*nullWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

// showLog == false remove log information to `stdout`
func Setup(showLog bool) error {
	// Mkdir & Return error if it exists
	os.Mkdir(RUNTIME, os.ModePerm)

	if !showLog {
		logrus.SetOutput(&nullWriter{})
	}

	logrus.SetLevel(loglevel)
	logrus.SetReportCaller(true)
	logrus.SetFormatter(
		&logrus.TextFormatter{
			DisableColors:   false,
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				// Get the relative file path from the current working directory
				relPath := filepath.Base(f.File)

				// Extract the full function name (package + function)
				funcName := f.Function
				// Extract just the function name (last part after ".")
				funcParts := strings.Split(funcName, ".")
				shortFuncName := funcParts[len(funcParts)-1]

				// Format as "path.go:func:line"
				formatted := fmt.Sprintf(" [%s:%s:%d]", relPath, shortFuncName, f.Line)
				return "", formatted
			},
		},
	)

	// Define hooker writes to local file
	fn := filepath.Join(RUNTIME, logFn)
	logFile, err := os.OpenFile(fn, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	logfiles = append(logfiles, fn)

	fileHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.TraceLevel: logFile,
		logrus.DebugLevel: logFile,
		logrus.InfoLevel:  logFile,
		logrus.WarnLevel:  logFile,
		logrus.ErrorLevel: logFile,
		logrus.FatalLevel: logFile,
		logrus.PanicLevel: logFile,
	}, &logrus.TextFormatter{
		DisableColors:   false,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			// Get the relative file path from the current working directory
			relPath := filepath.Base(f.File)

			// Extract the full function name (package + function)
			funcName := f.Function
			// Extract just the function name (last part after ".")
			funcParts := strings.Split(funcName, ".")
			shortFuncName := funcParts[len(funcParts)-1]

			// Format as "path.go:func:line"
			formatted := fmt.Sprintf(" [%s:%s:%d]", relPath, shortFuncName, f.Line)
			return "", formatted
		},
	},
	)

	logrus.AddHook(fileHook)
	return nil
}

// New a log handler
//
// If fileName != "", Add hook to output log to file
//
// If showLog == true, SetOutput(os.Stdout)
func NewLogger(fileName string, showLog bool) (ret *logrus.Logger, err error) {
	os.Mkdir(RUNTIME, os.ModePerm)

	ret = logrus.New()

	if showLog {
		ret.SetOutput(os.Stdout)
		ret.SetLevel(loglevel)
		ret.SetFormatter(
			&prefixed.TextFormatter{
				DisableColors:   false,
				TimestampFormat: "2006-01-02 15:04:05",
				FullTimestamp:   true,
				ForceFormatting: true,
			},
		)
	} else {
		ret.SetOutput(&nullWriter{})
	}

	if fileName != "" {
		var logFile *os.File
		fn := filepath.Join(RUNTIME, fileName)
		logFile, err = os.OpenFile(fn, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return
		}

		logfiles = append(logfiles, fn)

		fileHook := lfshook.NewHook(lfshook.WriterMap{
			logrus.InfoLevel:  logFile,
			logrus.WarnLevel:  logFile,
			logrus.ErrorLevel: logFile,
			logrus.FatalLevel: logFile,
			logrus.PanicLevel: logFile,
		}, &logrus.JSONFormatter{})

		ret.AddHook(fileHook)
	}

	return
}

func ShowLogfile() []string {
	return logfiles
}
