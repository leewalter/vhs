package flow

import (
	"fmt"
	"io"
)

// InputReader is an input reader.
type InputReader interface {
	io.ReadCloser
	Meta() *Meta
}

// InputModifier wraps an InputReader.
type InputModifier interface {
	Wrap(InputReader) (InputReader, error)
}

// InputModifiers is a slice of InputModifier.
type InputModifiers []InputModifier

// Wrap wraps an InputReader in all modifiers.
func (is InputModifiers) Wrap(r InputReader) (InputReader, error) {
	var err error
	for i := len(is) - 1; i >= 0; i-- {
		if r, err = is[i].Wrap(r); err != nil {
			return nil, fmt.Errorf("failed to wrap input writer: %w", err)
		}
	}

	return r, nil
}

// OutputWriter is an output writer.
type OutputWriter interface {
	io.WriteCloser
}

// OutputModifier wraps an OtputWriter.
type OutputModifier interface {
	Wrap(OutputWriter) (OutputWriter, error)
}

// OutputModifiers is a slice of OutputModifier.
type OutputModifiers []OutputModifier

// Wrap wraps an OutputWrtier in all modifiers.
func (os OutputModifiers) Wrap(w OutputWriter) (OutputWriter, error) {
	var err error
	for i := len(os) - 1; i >= 0; i-- {
		if w, err = os[i].Wrap(w); err != nil {
			return nil, fmt.Errorf("failed to wrap output writer: %w", err)
		}
	}

	return w, err
}

// CloseSequentially closes a and if there is no
// error, closes b.
func CloseSequentially(a, b io.Closer) error {
	if err := a.Close(); err != nil {
		return err
	}
	if b != nil {
		return b.Close()
	}
	return nil
}
