package wordcase

import (
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

func TestNotLowerOrDigit(t *testing.T) {
	tests := []struct {
		name string
		r    rune
		want bool
	}{
		{
			name: "a lowercase ascii letter",
			r:    'l',
			want: false,
		},
		{
			name: "an uppercase ascii letter",
			r:    'L',
			want: true,
		},
	}
	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := NotLowerOrDigit(tt.r)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestNotLetterOrDigit(t *testing.T) {
	tests := []struct {
		name string
		r    rune
		want bool
	}{
		{
			name: "a lowercase ascii letter",
			r:    'l',
			want: false,
		},
		{
			name: "an uppercase ascii letter",
			r:    'L',
			want: false,
		},
	}
	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := NotLetterOrDigit(tt.r)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_SimpleSeparator(t *testing.T) {
	type args struct {
		word string
		idx  int
		sep  IsRuneSeparator
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "aaa",
			args: args{
				word: "aaa",
				idx:  1,
				sep:  unicode.IsUpper,
			},
			want: false,
		},
		{
			name: "Aaa",
			args: args{
				word: "Aaa",
				idx:  1,
				sep:  unicode.IsUpper,
			},
			want: false,
		},
		{
			name: "aAa",
			args: args{
				word: "aAa",
				idx:  1,
				sep:  unicode.IsUpper,
			},
			want: true,
		},
		{
			name: "AAa",
			args: args{
				word: "AAa",
				idx:  1,
				sep:  unicode.IsUpper,
			},
			want: true,
		},
		{
			name: "aaA",
			args: args{
				word: "aaA",
				idx:  1,
				sep:  unicode.IsUpper,
			},
			want: false,
		},
		{
			name: "AaA",
			args: args{
				word: "AaA",
				idx:  1,
				sep:  unicode.IsUpper,
			},
			want: false,
		},
		{
			name: "aAA",
			args: args{
				word: "aAA",
				idx:  1,
				sep:  unicode.IsUpper,
			},
			want: true,
		},
		{
			name: "AAA",
			args: args{
				word: "AAA",
				idx:  1,
				sep:  unicode.IsUpper,
			},
			want: true,
		},
		{
			name: "aa at end",
			args: args{
				word: "aa",
				idx:  1,
				sep:  unicode.IsUpper,
			},
			want: false,
		},
		{
			name: "Aa at end",
			args: args{
				word: "Aa",
				idx:  1,
				sep:  unicode.IsUpper,
			},
			want: false,
		},
		{
			name: "aA at end",
			args: args{
				word: "aA",
				idx:  1,
				sep:  unicode.IsUpper,
			},
			want: true,
		},
		{
			name: "AA at end",
			args: args{
				word: "AA",
				idx:  1,
				sep:  unicode.IsUpper,
			},
			want: true,
		},
		{
			name: "aa at beginning",
			args: args{
				word: "aa",
				idx:  0,
				sep:  unicode.IsUpper,
			},
			want: false,
		},
		{
			name: "Aa at beginning",
			args: args{
				word: "Aa",
				idx:  0,
				sep:  unicode.IsUpper,
			},
			want: true,
		},
		{
			name: "aA at beginning",
			args: args{
				word: "aA",
				idx:  0,
				sep:  unicode.IsUpper,
			},
			want: false,
		},
		{
			name: "AA at beginning",
			args: args{
				word: "AA",
				idx:  0,
				sep:  unicode.IsUpper,
			},
			want: true,
		},
	}
	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := SimpleCategorizer(tt.args.word, tt.args.idx, tt.args.sep)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_LookAroundSeparator(t *testing.T) {
	type args struct {
		word string
		idx  int
		sep  IsRuneSeparator
	}

	// all possibles cases are considered
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "aaa",
			args: args{
				word: "aaa",
				idx:  1,
				sep:  unicode.IsUpper,
			},
			want: false,
		},
		{
			name: "Aaa",
			args: args{
				word: "Aaa",
				idx:  1,
				sep:  unicode.IsUpper,
			},
			want: false,
		},
		{
			name: "aAa",
			args: args{
				word: "aAa",
				idx:  1,
				sep:  unicode.IsUpper,
			},
			want: true,
		},
		{
			name: "AAa",
			args: args{
				word: "AAa",
				idx:  1,
				sep:  unicode.IsUpper,
			},
			want: true,
		},
		{
			name: "aaA",
			args: args{
				word: "aaA",
				idx:  1,
				sep:  unicode.IsUpper,
			},
			want: false,
		},
		{
			name: "AaA",
			args: args{
				word: "AaA",
				idx:  1,
				sep:  unicode.IsUpper,
			},
			want: false,
		},
		{
			name: "aAA",
			args: args{
				word: "aAA",
				idx:  1,
				sep:  unicode.IsUpper,
			},
			want: true,
		},
		{
			name: "AAA",
			args: args{
				word: "AAA",
				idx:  1,
				sep:  unicode.IsUpper,
			},
			want: false,
		},
		{
			name: "aa at end",
			args: args{
				word: "aa",
				idx:  1,
				sep:  unicode.IsUpper,
			},
			want: false,
		},
		{
			name: "Aa at end",
			args: args{
				word: "Aa",
				idx:  1,
				sep:  unicode.IsUpper,
			},
			want: false,
		},
		{
			name: "aA at end",
			args: args{
				word: "aA",
				idx:  1,
				sep:  unicode.IsUpper,
			},
			want: true,
		},
		{
			name: "AA at end",
			args: args{
				word: "AA",
				idx:  1,
				sep:  unicode.IsUpper,
			},
			want: false,
		},
		{
			name: "aa at beginning",
			args: args{
				word: "aa",
				idx:  0,
				sep:  unicode.IsUpper,
			},
			want: false,
		},
		{
			name: "Aa at beginning",
			args: args{
				word: "Aa",
				idx:  0,
				sep:  unicode.IsUpper,
			},
			want: true,
		},
		{
			name: "aA at beginning",
			args: args{
				word: "aA",
				idx:  0,
				sep:  unicode.IsUpper,
			},
			want: false,
		},
		{
			name: "AA at beginning",
			args: args{
				word: "AA",
				idx:  0,
				sep:  unicode.IsUpper,
			},
			want: false,
		},
	}
	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := LookAroundCategorizer(tt.args.word, tt.args.idx, tt.args.sep)
			assert.Equal(t, tt.want, got)
		})
	}
}
