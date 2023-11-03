package jsonl

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Writer struct {
	w io.Writer
}

func NewWriter(w io.Writer) Writer {
	return Writer{
		w: w,
	}
}

func (w Writer) Close() error {
	if c, ok := w.w.(io.WriteCloser); ok {
		return c.Close()
	}
	return fmt.Errorf("given writer is no WriteCloser")
}

func (w Writer) Write(data interface{}) error {
	j, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("could not json marshal data: %w", err)
	}

	_, err = w.w.Write(j)
	if err != nil {
		return fmt.Errorf("could not write json data to underlying io.Writer: %w", err)
	}

	_, err = w.w.Write([]byte("\n"))
	if err != nil {
		return fmt.Errorf("could not write newline to underlying io.Writer: %w", err)
	}

	if f, ok := w.w.(http.Flusher); ok {
		// If http writer, flush as well
		f.Flush()
	}
	return nil
}
