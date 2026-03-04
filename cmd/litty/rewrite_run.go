package main

// NewRunRewriter creates a rewriter for go run output 🏃
// go run output IS the programs output bestie — we only rewrite compile errors
// which the build rewriter already handles, so we just delegate 💅
func NewRunRewriter() Rewriter {
	return NewBuildRewriter()
}
