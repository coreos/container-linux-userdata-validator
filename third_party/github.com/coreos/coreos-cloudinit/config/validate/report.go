package validate

import (
	"encoding/json"
	"fmt"
)

type entryKind int

const (
	entryError   entryKind = iota
	entryWarning entryKind = iota
)

var (
	entryKindStrings = map[entryKind]string{
		entryError:   "error",
		entryWarning: "warning",
	}
)

type Entry struct {
	kind    entryKind
	message string
	line    int
}

func (e Entry) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"kind":    entryKindStrings[e.kind],
		"message": e.message,
		"line":    e.line,
	})
}

func (e Entry) String() string {
	return fmt.Sprintf("line %d: %s", e.line, e.message)
}

func (e Entry) IsError() bool {
	return (e.kind == entryError)
}

func (e Entry) IsWarning() bool {
	return (e.kind == entryWarning)
}

type Report struct {
	entries []Entry
}

func (r *Report) Error(line int, message string) {
	r.entries = append(r.entries, Entry{entryError, message, line})
}

func (r *Report) Warning(line int, message string) {
	r.entries = append(r.entries, Entry{entryWarning, message, line})
}

func (r *Report) Entries() []Entry {
	return r.entries
}
