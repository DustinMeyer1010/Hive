package types

type Database interface {
	BuildScanArgs([]string) ([]any, error)
	Fields() []string
}
