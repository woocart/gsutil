package pipeline

import (
	"fmt"
	"strings"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

type metadata map[string]string

func (h metadata) Set(value string) error {
	parts := strings.SplitN(value, "=", 2)
	if len(parts) != 2 {
		return fmt.Errorf("expected key=value got '%s'", value)
	}
	(map[string]string)(h)[parts[0]] = parts[1]
	return nil
}
func (h metadata) IsCumulative() bool {
	return true
}

func (h metadata) String() string {
	return ""
}

func Metadata(s kingpin.Settings) (target map[string]string) {
	target = map[string]string{}
	s.SetValue((metadata)(target))
	return
}
