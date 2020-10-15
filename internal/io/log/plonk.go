package log

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"github.com/apex/log"
	"github.com/fatih/color"
	colorable "github.com/mattn/go-colorable"
)

// Default handler outputting to stderr.
var Default = NewPlonkHandler(os.Stderr)

// start time.
var start = time.Now()

var bold = color.New(color.Bold)

// Colors mapping.
var Colors = [...]*color.Color{
	log.DebugLevel: color.New(color.FgWhite),
	log.InfoLevel:  color.New(color.FgBlue),
	log.WarnLevel:  color.New(color.FgYellow),
	log.ErrorLevel: color.New(color.FgRed),
	log.FatalLevel: color.New(color.FgRed),
}

// Strings mapping.
var Strings = [...]string{
	log.DebugLevel: "- debug\t\t",
	log.InfoLevel:  "• info\t\t",
	log.WarnLevel:  "⚠ warning\t",
	log.ErrorLevel: "x error\t\t",
	log.FatalLevel: "§ fatal\t\t",
}

// Handler implementation.
type Handler struct {
	mu      sync.Mutex
	Writer  io.Writer
	Padding int
}

// NewPlonkHandler handler.
func NewPlonkHandler(w io.Writer) *Handler {
	if f, ok := w.(*os.File); ok {
		return &Handler{
			Writer:  colorable.NewColorable(f),
			Padding: 3,
		}
	}

	return &Handler{
		Writer:  w,
		Padding: 3,
	}
}

// HandleLog implements log.Handler.
func (h *Handler) HandleLog(e *log.Entry) error {
	color := Colors[e.Level]
	level := Strings[e.Level]
	names := e.Fields.Names()

	h.mu.Lock()
	defer h.mu.Unlock()

	color.Fprintf(h.Writer, "%s %-25s", bold.Sprintf("%*s", h.Padding+1, level), e.Message)

	for _, name := range names {
		ammendedName := name
		if name == "duration" {
			ammendedName = "duration [ms]"
		}
		fmt.Fprintf(h.Writer, " %s=%v", color.Sprint(ammendedName), e.Fields.Get(name))
	}

	fmt.Fprintln(h.Writer)

	return nil
}
