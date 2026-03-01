package proxier

import "text/template"

var funcMap = template.FuncMap{
	"sub":  func(a, b int) int { return a - b },
	"ge":   func(a, b int) bool { return a >= b },
	"not":  func(b bool) bool { return !b },
	"dict": dict,
	"list": func(vals ...interface{}) []interface{} { return vals },
}

func dict(values ...interface{}) map[string]interface{} {
	m := make(map[string]interface{}, len(values)/2)
	if len(values)%2 != 0 {
		panic("invalid dict call, must have even args")
	}
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			panic("dict keys must be strings")
		}
		m[key] = values[i+1]
	}
	return m
}
