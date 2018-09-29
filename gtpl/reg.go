//go:generate go-bindata -ignore=\.go -pkg=gtpl -o=bindata.go ./...

package gtpl

//Templates to be used in the response of suggester
var (
	Greeting         = NewFormatTemplate("greeting.md", MustAsset)
	NewcomerGreeting = NewStringTemplate("newcomer_greeting.md", MustAsset)
	Improvement      = NewFormatTemplate("improvement.md", MustAsset)
	Todo             = NewFormatTemplate("todo.md", MustAsset)
	Praise           = NewFormatTemplate("praise.md", MustAsset)
	NotLinted        = NewStringTemplate("not_linted.md", MustAsset)
	NotFormatted     = NewStringTemplate("not_formatted.md", MustAsset)
	Questions        = NewStringTemplate("questions.md", MustAsset)
	Benchmarking     = NewStringTemplate("benchmarking.md", MustAsset)
)