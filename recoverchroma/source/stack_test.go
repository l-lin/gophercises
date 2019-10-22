package source

import (
	"strings"
	"testing"
)

func TestRenderStackLine(t *testing.T) {
	var tests = map[string]struct {
		given    string
		expected string
	}{
		"basic": {
			given: `	/home/llin/apps/go/src/runtime/debug/stack.go:24 +0x9d`,
			expected: `	<a href="/debug/?path=/home/llin/apps/go/src/runtime/debug/stack.go">/home/llin/apps/go/src/runtime/debug/stack.go:24 +0x9d</a>`,
		},
		"no link": {
			given:    `goroutine 19 [running]:`,
			expected: `goroutine 19 [running]:`,
		},
		"empty string": {
			given:    ``,
			expected: ``,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual := renderStackLine(tt.given)
			if actual != tt.expected {
				t.Errorf("expected %v, actual %v", tt.expected, actual)
			}
		})
	}
}

func TestRenderStack(t *testing.T) {
	var tests = map[string]struct {
		given    string
		expected string
	}{
		"basic": {
			given: `goroutine 19 [running]:
runtime/debug.Stack(0xc000055ae0, 0x1, 0x1)
	/home/llin/apps/go/src/runtime/debug/stack.go:24 +0x9d
github.com/l-lin/gophercises/recoverchroma/cmd.devMw.func1.1(0xbd1c80, 0xc000110000)
	/home/llin/perso/gophercises/recoverchroma/cmd/serve.go:43 +0xa3
panic(0x996a60, 0xbbb9a0)
	/home/llin/apps/go/src/runtime/panic.go:522 +0x1b5
github.com/l-lin/gophercises/recoverchroma/cmd.funcThatPanics(...)
	/home/llin/perso/gophercises/recoverchroma/cmd/serve.go:63
github.com/l-lin/gophercises/recoverchroma/cmd.panicDemo(0xbd1c80, 0xc000110000, 0xc000098200)
	/home/llin/perso/gophercises/recoverchroma/cmd/serve.go:54 +0x3a`,
			expected: `goroutine 19 [running]:
runtime/debug.Stack(0xc000055ae0, 0x1, 0x1)
	<a href="/debug/?path=/home/llin/apps/go/src/runtime/debug/stack.go">/home/llin/apps/go/src/runtime/debug/stack.go:24 +0x9d</a>
github.com/l-lin/gophercises/recoverchroma/cmd.devMw.func1.1(0xbd1c80, 0xc000110000)
	<a href="/debug/?path=/home/llin/perso/gophercises/recoverchroma/cmd/serve.go">/home/llin/perso/gophercises/recoverchroma/cmd/serve.go:43 +0xa3</a>
panic(0x996a60, 0xbbb9a0)
	<a href="/debug/?path=/home/llin/apps/go/src/runtime/panic.go">/home/llin/apps/go/src/runtime/panic.go:522 +0x1b5</a>
github.com/l-lin/gophercises/recoverchroma/cmd.funcThatPanics(...)
	<a href="/debug/?path=/home/llin/perso/gophercises/recoverchroma/cmd/serve.go">/home/llin/perso/gophercises/recoverchroma/cmd/serve.go:63</a>
github.com/l-lin/gophercises/recoverchroma/cmd.panicDemo(0xbd1c80, 0xc000110000, 0xc000098200)
	<a href="/debug/?path=/home/llin/perso/gophercises/recoverchroma/cmd/serve.go">/home/llin/perso/gophercises/recoverchroma/cmd/serve.go:54 +0x3a</a>`,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual := RenderStack(tt.given)
			if strings.Trim(actual, "\n") != strings.Trim(tt.expected, "\n") {
				t.Errorf("expected:\n%v\nactual:\n%v", tt.expected, actual)
			}
		})
	}
}
