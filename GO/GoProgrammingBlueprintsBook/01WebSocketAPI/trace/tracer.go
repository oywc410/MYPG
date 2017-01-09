package trace

import (
	"io"
	"fmt"
)

type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
	t.out.Write([]byte(fmt.Sprintln(a...)))
}

func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

type nilTracer struct {
}

func (t *nilTracer) Trace(a ...interface{}) {
}

func Off() Tracer {
	return &nilTracer{}
}