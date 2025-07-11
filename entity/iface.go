package entity

type InputTag interface {
	ApplyParameter(param *Parameter)
}

type OutputTag interface {
	ApplyParameter(param *Parameter)
}

type LogContextTag interface {
	Alias() string
	Name() string
	Key() string
}

type LogInputTag interface {
	Alias() string
	Name() string
	Key() string
	Source() string
}

type LogOutputTag interface {
	Alias() string
	Name() string
	Key() string
	Source() string
}

type TraceContextTag interface {
	Alias() string
	Name() string
	Key() string
}

type TraceInputTag interface {
	Alias() string
	Name() string
	Key() string
	Source() string
}

type TraceOutputTag interface {
	Alias() string
	Name() string
	Key() string
	Source() string
}

type RetryTag interface {
	// define later
}
