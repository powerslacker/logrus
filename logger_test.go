package logrus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFieldValueError(t *testing.T) {
	buf := &bytes.Buffer{}
	l := &Logger{
		Out:       buf,
		Formatter: new(JSONFormatter),
		Hooks:     make(LevelHooks),
		Level:     DebugLevel,
	}
	l.WithField("func", func() {}).Info("test")
	fmt.Println(buf.String())
	var data map[string]interface{}
	json.Unmarshal(buf.Bytes(), &data)
	_, ok := data[FieldKeyLogrusError]
	require.True(t, ok)
}

func TestNoFieldValueError(t *testing.T) {
	buf := &bytes.Buffer{}
	l := &Logger{
		Out:       buf,
		Formatter: new(JSONFormatter),
		Hooks:     make(LevelHooks),
		Level:     DebugLevel,
	}
	l.WithField("str", "str").Info("test")
	fmt.Println(buf.String())
	var data map[string]interface{}
	json.Unmarshal(buf.Bytes(), &data)
	_, ok := data[FieldKeyLogrusError]
	require.False(t, ok)
}

func TestWarninglnNotEqualToWarning(t *testing.T) {
	buf := &bytes.Buffer{}
	bufln := &bytes.Buffer{}

	formatter := new(TextFormatter)
	formatter.DisableTimestamp = true
	formatter.DisableLevelTruncation = true

	l := &Logger{
		Out:       buf,
		Formatter: formatter,
		Hooks:     make(LevelHooks),
		Level:     DebugLevel,
	}
	l.Warning("hello,", "world")

	l.SetOutput(bufln)
	l.Warningln("hello,", "world")
	assert.NotEqual(t, buf.String(), bufln.String(), "Warning() and Wantingln() should not be equal")
}

func TestNotice(t *testing.T) {
	buf := &bytes.Buffer{}

	formatter := new(TextFormatter)
	formatter.DisableTimestamp = true
	formatter.DisableLevelTruncation = true
	l := &Logger{
		Out:       buf,
		Formatter: formatter,
		Hooks:     make(LevelHooks),
		Level:     NoticeLevel,
	}
	l.Notice("hello", "world")

	assert.Equal(t, "level=notice message=helloworld\n", buf.String())
}

func TestLoggerDefaultLevel(t *testing.T) {
	l := New()
	assert.True(t, l.IsLevelEnabled(NoticeLevel))
}
