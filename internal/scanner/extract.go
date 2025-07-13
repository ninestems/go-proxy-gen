package scanner

import (
	"strings"
)

// extractPath extract parameters name and path to type.
func extractPath(in string) (name, path string) {
	data := strings.Split(in, ":")
	path = data[0]
	if len(data) > 1 {
		name, path = data[0], data[1]
	}

	return
}
