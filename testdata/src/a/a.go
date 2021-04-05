package a

import (
	"context"
)

func connection(_ context.Context, _ string) {}

func stringFirst(_ string, _ context.Context) {} // want "arguments contain context.Context but is not at the begining"

func nonNameFunc() {
	f := func(_ int, _ context.Context) {} // want "arguments contain context.Context but is not at the begining"
	_ = f
}
