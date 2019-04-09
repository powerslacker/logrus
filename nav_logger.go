package logrus

// Base is used to create new NavLogger entries.
type Base struct {
	Application string
	Habitat     string
	Host        string
	PrettyPrint bool
}

func (b *Base) setDefaults() {
	if b.Application == "" {
		b.Application = "unknown"
	}
	if b.Habitat == "" {
		b.Habitat = "development"
	}
	if b.Host == "" {
		b.Host = "localhost"
	}
}

// NewLogger creates a new logger instance
func (b *Base) NewLogger() NavLogger {
	b.setDefaults()
	logger := New()
	logger.SetFormatter(&JSONFormatter{
		PrettyPrint: b.PrettyPrint,
	})
	logger.ReportCaller = true
	e := logger.WithFields(Fields{
		"application": b.Application,
		"habitat":     b.Habitat,
		"host":        b.Host,
		"trace_id":    "reserved",
		"span_id":     "reserved",
	})
	return e
}

// NavLogger is a logger  that implements RFC 12:
// https://git.nav.com/engineering/Nav-Engineering-Standards/blob/master/proposals/012-logging-standardization.md
// Also fufills the `Logger` interface.
type NavLogger interface {
	Emergencyf(format string, args ...interface{})
	Alertf(format string, args ...interface{})
	Criticalf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Noticef(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Printf(format string, args ...interface{})
	Emergency(args ...interface{})
	Alert(args ...interface{})
	Critical(args ...interface{})
	Error(args ...interface{})
	Warning(args ...interface{})
	Notice(args ...interface{})
	Info(args ...interface{})
	Debug(args ...interface{})
	WithFields(fields Fields) *Entry
	WithField(key string, value interface{}) *Entry
}
