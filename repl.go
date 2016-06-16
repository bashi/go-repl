package repl

import (
	"github.com/peterh/liner"
)

var ErrPromptAborted = liner.ErrPromptAborted

func Run(callback func(line string) error) error {
	line := liner.NewLiner()
	defer line.Close()

	line.SetCtrlCAborts(true)
	for {
		l, err := line.Prompt("> ")
		if err != nil {
			return err
		}
		if err := callback(l); err != nil {
			return err
		}
	}
	return nil
}

type Handler interface {
	Handle(line string) error
}

func RunHandler(h Handler) error {
	return Run(func(line string) error {
		return h.Handle(line)
	})
}
