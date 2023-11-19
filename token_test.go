package wordcase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokens_String(t *testing.T) {
	tests := []struct {
		name string
		t    Tokens
		want string
	}{
		{
			name: "empty",
			t:    Tokens{},
			want: "",
		},
		{
			name: "basic",
			t:    Tokens{"basic"},
			want: "basic",
		},
		{
			name: "2 tokens",
			t:    Tokens{"one two"},
			want: "one two",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.t.String()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestTokens_FormatAll(t *testing.T) {
	type args struct {
		fn Formatter
	}
	tests := []struct {
		name string
		t    Tokens
		args args
		want Tokens
	}{
		// TODO: Add test cases.
	}
	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.t.FormatAll(tt.args.fn)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestTokenizeString(t *testing.T) {
	type args struct {
		s       string
		test    SeparatorTest
		sepRune IsRuneSeparator
		rmSep   bool
	}
	tests := []struct {
		name string
		args args
		want Tokens
	}{
		{
			name: "empty",
			args: args{
				s:       "",
				test:    LookAroundCategorizer,
				sepRune: NotLowerOrDigit,
				rmSep:   false,
			},
			want: Tokens{},
		},
		{
			name: "basic",
			args: args{
				s:       "basic",
				test:    LookAroundCategorizer,
				sepRune: NotLowerOrDigit,
				rmSep:   false,
			},
			want: Tokens{"basic"},
		},
		{
			name: "single non-separator",
			args: args{
				s:       "a",
				test:    LookAroundCategorizer,
				sepRune: NotLowerOrDigit,
				rmSep:   false,
			},
			want: Tokens{"a"},
		},
		{
			name: "single separator",
			args: args{
				s:       "A",
				test:    LookAroundCategorizer,
				sepRune: NotLowerOrDigit,
				rmSep:   false,
			},
			want: Tokens{"A"},
		},
		{
			name: "two single tokens",
			args: args{
				s:       "aB",
				test:    LookAroundCategorizer,
				sepRune: NotLowerOrDigit,
				rmSep:   false,
			},
			want: Tokens{"a", "B"},
		},
		{
			name: "one double token",
			args: args{
				s:       "Ab",
				test:    LookAroundCategorizer,
				sepRune: NotLowerOrDigit,
				rmSep:   false,
			},
			want: Tokens{"Ab"},
		},
		{
			name: "double separator",
			args: args{
				s:       "aBC",
				test:    LookAroundCategorizer,
				sepRune: NotLowerOrDigit,
				rmSep:   false,
			},
			want: Tokens{"a", "BC"},
		},
		{
			name: "double separator with trailing",
			args: args{
				s:       "aBCd",
				test:    LookAroundCategorizer,
				sepRune: NotLowerOrDigit,
				rmSep:   false,
			},
			want: Tokens{"a", "B", "Cd"},
		},
		{
			name: "quad separator",
			args: args{
				s:       "aBCDE",
				test:    LookAroundCategorizer,
				sepRune: NotLowerOrDigit,
				rmSep:   false,
			},
			want: Tokens{"a", "BCDE"},
		},
		{
			name: "quad separator with trailing",
			args: args{
				s:       "aBCDEf",
				test:    LookAroundCategorizer,
				sepRune: NotLowerOrDigit,
				rmSep:   false,
			},
			want: Tokens{"a", "BCD", "Ef"},
		},
		{
			name: "rolling",
			args: args{
				s:       "aBcDeFgH",
				test:    LookAroundCategorizer,
				sepRune: NotLowerOrDigit,
				rmSep:   false,
			},
			want: Tokens{"a", "Bc", "De", "Fg", "H"},
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := TokenizeString(tt.args.s, tt.args.test, tt.args.sepRune, tt.args.rmSep)
			assert.Equal(t, tt.want, got)
		})
	}
}
