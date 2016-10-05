package stop

//go:generate stringer -type=EvalType eval_type.go

// EvalType represents the type of the output returned from a STOP
// evaluation.
type EvalType uint32

const (
	TUnsupported EvalType = 0
	TString      EvalType = 1 << iota
	TList
	TMap
	TBool
)
