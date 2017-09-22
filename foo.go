// Package vendoroverride in the gopath defines foo and bar with some docs.
// These will be overridden in the vendor directory of the main package. We
// could probably do the same with branches or something, but this seems a bit
// more controllable.
package vendoroverride

// Foo returns the string "foo" (the gopath version)
func Foo() string {
	return "foo"
}

// Bar returns the string "bar" (the gopath version)
func Bar() string {
	return "bar"
}
