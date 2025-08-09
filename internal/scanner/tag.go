package scanner

import (
	"go/ast"
	"strings"

	"github.com/ninestems/go-proxy-gen/entity"
)

// tags extracts custom tags from all lines of comments.
func tags(in *ast.CommentGroup) *entity.Tags {
	var (
		out   entity.Tags
		found = false
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

			var (
				proxyType entity.ProxyType
			)

			data := strings.Fields(line)

			switch data[0] {
			case "log":
				proxyType = entity.ProxyTypeLogger
			case "trace":
				proxyType = entity.ProxyTypeTracer
			case "retry":
				proxyType = entity.ProxyTypeRetrier
			}

			info := strings.Split(data[1], "::")
			switch info[0] {
			case "ctx":
				out.AddContext(extractContextIOTag(info[1:], proxyType))
			case "input":
				out.AddInput(extractInputIOTag(info[1:], proxyType))
			case "output":
				out.AddOutput(extractOutputIOTag(info[1:], proxyType))
			case "retry":
				out.AddRetry(extractRetryTag(info[1:], proxyType))
			}
		}
	}

	return &out
}

func extractContextIOTag(in []string, ptype entity.ProxyType) *entity.ContextIO {
	var (
		alias string
		key   string
	)

	switch len(in) {
	case 1:
		key = in[0]
		alias = in[0]
	case 2:
		key = in[0]
		alias = in[1]
	}

	return entity.NewIOContextTag(alias, "context.Context", key, ptype)
}

func extractIOLabels(in []string) (name string, source string, accessor string, alias string) {
	data := strings.Split(in[0], ":")
	source = data[0]
	if len(data) > 1 {
		name, source = data[0], data[1]
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

func extractInputIOTag(in []string, ptype entity.ProxyType) *entity.InputIO {
	name, source, accessor, alias := extractIOLabels(in)
	return entity.NewIOInputTag(alias, name, source, accessor, ptype)
}

func extractOutputIOTag(in []string, ptype entity.ProxyType) *entity.OutputIO {
	name, source, accessor, alias := extractIOLabels(in)
	return entity.NewIOOutputTag(alias, name, source, accessor, ptype)
}

func extractRetryTag(_ []string, _ entity.ProxyType) *entity.Retry {
	panic("not yet implemented")
}
