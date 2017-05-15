package test

import (
	"io/ioutil"

	"github.com/Sirupsen/logrus"
)

// mapTest.Hook is a hook designed for dealing with logs in mapTest scenarios.
type Hook struct {
	Entries []*logrus.Entry
}

// Installs a mapTest hook for the global logger.
func NewGlobal() *Hook {

	hook := new(Hook)
	logrus.AddHook(hook)

	return hook

}

// Installs a mapTest hook for a given local logger.
func NewLocal(logger *logrus.Logger) *Hook {

	hook := new(Hook)
	logger.Hooks.Add(hook)

	return hook

}

// Creates a discarding logger and installs the mapTest hook.
func NewNullLogger() (*logrus.Logger, *Hook) {

	logger := logrus.New()
	logger.Out = ioutil.Discard

	return logger, NewLocal(logger)

}

func (t *Hook) Fire(e *logrus.Entry) error {
	t.Entries = append(t.Entries, e)
	return nil
}

func (t *Hook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// LastEntry returns the last entry that was logged or nil.
func (t *Hook) LastEntry() (l *logrus.Entry) {

	if i := len(t.Entries) - 1; i < 0 {
		return nil
	} else {
		return t.Entries[i]
	}

}

// Reset removes all Entries from this mapTest hook.
func (t *Hook) Reset() {
	t.Entries = make([]*logrus.Entry, 0)
}
