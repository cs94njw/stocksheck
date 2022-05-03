package types

type Requirements struct {
	name  string
	alert Alert
}

type Alert struct {
	risesTo     bool
	fallsTo     bool
	targetPrice float32
}
