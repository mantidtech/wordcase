package wordcase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokens_Join(t *testing.T) {
	tests := []struct {
		name string
		t    Tokens
		sep  string
		want string
	}{
		{
			name: "empty",
			t:    Tokens{},
			sep:  "-",
			want: "",
		},
		{
			name: "simple",
			t:    Tokens{"simple"},
			sep:  "-",
			want: "simple",
		},
		{
			name: "one two",
			t:    Tokens{"one", "two"},
			sep:  "-",
			want: "one-two",
		},
	}
	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.t.Join(tt.sep)
			assert.Equal(t, tt.want, got)
		})
	}
}
