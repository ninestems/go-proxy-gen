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
				out.AddContext(extractContextIOTag(info, proxyType))
			case "input":
				out.AddInput(extractInputIOTag(info, proxyType))
			case "output":
				out.AddOutput(extractOutputIOTag(info, proxyType))
			case "retry":
				out.AddRetry(extractRetryTag(info, proxyType))
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
	case 2:
		alias = in[1]
		key = in[1]
	case 3:
		alias = in[1]
		key = in[2]
	}

	return entity.NewIOContextTag(alias, "", "context.Context", key, ptype)
}

func extractInputIOTag(in []string, ptype entity.ProxyType) *entity.InputIO {
	var (
		alias string
		key   string
		name  string
		path  string
	)

	switch len(in) {
	case 3:
		alias = in[2]
		name, path = extractPath(in[1])
		key = in[2]
	case 4:
		alias = in[1]
		name, path = extractPath(in[2])
		key = in[3]
	}

	return entity.NewIOInputTag(alias, name, path, key, ptype)
}

func extractOutputIOTag(in []string, ptype entity.ProxyType) *entity.OutputIO {
	var (
		alias string
		key   string
		name  string
		path  string
	)

	switch len(in) {
	case 3:
		alias = in[2]
		name, path = extractPath(in[1])
		key = in[2]
	case 4:
		alias = in[1]
		name, path = extractPath(in[2])
		key = in[3]
	}

	return entity.NewIOOutputTag(alias, name, path, key, ptype)
}

func extractRetryTag(_ []string, _ entity.ProxyType) *entity.Retry {
	panic("not yet implemented")
}
