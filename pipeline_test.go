package wordcase

import (
	"strings"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

func TestNewPipeline(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want Tokens
	}{
		{
			name: "empty",
			s:    "sample",
			want: Tokens{"sample"},
		},
	}
	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := NewPipeline()
			assert.Equal(t, tt.want, got(tt.s))
		})
	}
}

func TestPipeline_TokenizeUsing(t *testing.T) {
	type Args struct {
		test SeparatorTest
		sep  IsRuneSeparator
		del  bool
	}
	tests := []struct {
		name string
		args Args
		str  string
		want Tokens
	}{
		{
			name: "basic",
			args: Args{
				test: LookAroundCategorizer,
				sep:  unicode.IsUpper,
				del:  false,
			},
			str:  "basic",
			want: Tokens{"basic"},
		},
		{
			name: "basic - 2 tokens",
			args: Args{
				test: LookAroundCategorizer,
				sep:  unicode.IsUpper,
				del:  false,
			},
			str:  "oneTwo",
			want: Tokens{"one", "Two"},
		},
		{
			name: "multiple separators at the end",
			args: Args{
				test: LookAroundCategorizer,
				sep:  NotLetterOrDigit,
				del:  false,
			},
			str:  "Add$$",
			want: Tokens{"Add", "$$"},
		},
	}
	pl := NewPipeline()

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := pl.TokenizeUsing(tt.args.test, tt.args.sep, tt.args.del)
			assert.Equal(t, tt.want, got(tt.str))
		})
	}
}

func TestPipeline_WithFormatter(t *testing.T) {
	tests := []struct {
		name string
		fmt  Formatter
		sel  TokenSelector
		s    string
		want Tokens
	}{
		{
			name: "basic",
			fmt:  strings.ToUpper,
			sel:  ToFirst,
			s:    "basic",
			want: Tokens{"BASIC"},
		},
	}
	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			p := NewPipeline()
			got := p.WithFormatter(tt.fmt, tt.sel)
			assert.Equal(t, tt.want, got(tt.s))
		})
	}
}

func TestPipeline_WithAllFormatter(t *testing.T) {
	tests := []struct {
		name string
		fmt  Formatter
		s    string
		want Tokens
	}{
		{
			name: "basic",
			fmt:  strings.ToUpper,
			s:    "basic",
			want: Tokens{"BASIC"},
		},
	}
	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			p := NewPipeline()
			got := p.WithAllFormatter(tt.fmt)
			assert.Equal(t, tt.want, got(tt.s))
		})
	}
}

// TestPipeline_JoinWith provides unit test coverage for Pipeline.JoinWith()
func TestPipeline_JoinWith(t *testing.T) {
	tests := []struct {
		name string
		sep  string
		str  string
		want string
	}{
		{
			name: "basic",
			sep:  "-",
			str:  "one two three",
			want: "one-two-three",
		},
	}
	pl := NewPipeline().TokenizeUsing(SimpleCategorizer, unicode.IsSpace, true)

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := pl.JoinWith(tt.sep)
			assert.Equal(t, tt.want, got(tt.str))
		})
	}
}
