package prettylog

import (
	"bytes"
	"log/slog"
	"testing"
)

type captureStream struct {
	lines [][]byte
}

func (cs *captureStream) Write(bytes []byte) (int, error) {
	cs.lines = append(cs.lines, bytes)
	return len(bytes), nil
}

func Test_WritesToProvidedStream(t *testing.T) {
	cs := &captureStream{}
	handler := New(cs, nil)
	logger := slog.New(handler)
	msg := "testing logger"

	logger.Info(msg)

	if len(cs.lines) != 1 {
		t.Errorf("expected 1 lines logged, got: %d", len(cs.lines))
	}

	out := bytes.Buffer{}
	out.WriteString(colorize(cyan, "INFO:"))
	out.WriteString(" ")
	out.WriteString(colorize(white, msg))
	out.WriteString(" ")
	out.WriteString(colorize(darkGray, "{}"))

	line := cs.lines[0]
	if bytes.Contains(line, out.Bytes()) == false {
		t.Errorf("expected `testing logger` but found `%s`", line)
	}
	if !bytes.HasSuffix(line, []byte("\n")) {
		t.Errorf("exected line to be terminated with `\\n` but found `%s`", line[len(line)-1:])
	}
}
