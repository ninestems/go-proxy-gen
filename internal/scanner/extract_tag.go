package scanner

import (
	"go/ast"
	"strconv"
	"strings"
	"time"
)

func extractFunctionTags(in *ast.CommentGroup) *tagSpecification {
	var (
		found      = false
		lines      []string
		ctxTags    []*tagContextSpecification
		inputTags  []*tagInputSpecification
		outputTags []*tagOutputSpecification
		retryTags  []*tagRetrySpecification
	)
	for _, c := range in.List {
		line := strings.TrimSpace(strings.TrimPrefix(c.Text, "//"))
		switch {
		case strings.HasPrefix(line, "goproxygen:"):
			found = true
			continue

		case !found || line == "":
			continue

		default:
			lines = append(lines, line)
		}
	}

	for _, line := range lines {
		data := strings.Fields(line)
		info := strings.Split(data[1], "::")
		switch info[0] {
		case "ctx":
			ctxTags = append(ctxTags, extractContextTagSpecification(data[0], info[1:]))
		case "input":
			inputTags = append(inputTags, extractInputTagSpecification(data[0], info[1:]))
		case "output":
			outputTags = append(outputTags, extractOutputTagSpecification(data[0], info[1:]))
		case "retry":
			retryTags = append(retryTags, extractRetryTagSpecification(data[0], info[1:]))
		}
	}

	return &tagSpecification{
		context: ctxTags,
		input:   inputTags,
		output:  outputTags,
		retry:   retryTags,
	}
}

func extractContextLabels(in []string) (alias string, key string, source string) {
	switch len(in) {
	case 1:
		key = in[0]
		alias = in[0]
	case 2:
		key = in[0]
		alias = in[1]
	}
	source = "context.Context"

	return
}

func extractContextTagSpecification(proxy string, data []string) *tagContextSpecification {
	alias, key, source := extractContextLabels(data)
	return &tagContextSpecification{
		proxy:  proxy,
		alias:  alias,
		key:    key,
		source: source,
	}
}

func extractIOLabels(in []string) (name string, source string, accessor string, alias string) {
	data := strings.Split(in[0], ":")

	switch len(data) {
	case 1:
		source = data[0]
	case 2:
		name, source = data[0], data[1]
		accessor = name
		alias = name
	}

	switch len(in) {
	case 2:
		accessor = in[1]
		alias = accessor
	case 3:
		accessor = in[1]
		alias = in[2]
	}
	return
}

func extractRetryLabels(in []string) (start time.Duration, end time.Duration, multiplier float32, attempts uint64) {
	start, _ = time.ParseDuration(in[0])
	end, _ = time.ParseDuration(in[1])
	m, _ := strconv.ParseFloat(in[2], 32)
	multiplier = float32(m)
	attempts, _ = strconv.ParseUint(in[3], 10, 32)
	return
}

func extractInputTagSpecification(proxy string, data []string) *tagInputSpecification {
	name, source, accessor, alias := extractIOLabels(data)
	return &tagInputSpecification{
		proxy:    proxy,
		name:     name,
		source:   source,
		accessor: accessor,
		alias:    alias,
	}
}

func extractOutputTagSpecification(proxy string, data []string) *tagOutputSpecification {
	name, source, accessor, alias := extractIOLabels(data)
	return &tagOutputSpecification{
		proxy:    proxy,
		name:     name,
		source:   source,
		accessor: accessor,
		alias:    alias,
	}
}

func extractRetryTagSpecification(proxy string, data []string) *tagRetrySpecification {
	start, end, multiplier, attempts := extractRetryLabels(data)
	return &tagRetrySpecification{
		proxy:      proxy,
		start:      start,
		end:        end,
		multiplier: multiplier,
		attempts:   attempts,
	}
}
