package wordcase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TesUppercaseFirst provides unit test coverage for UppercaseFirst()
func TestUppercaseFirst(t *testing.T) {
	type Args struct {
		s string
	}

	tests := []struct {
		name string
		args Args
		want string
	}{
		{
			name: "empty",
			args: Args{
				s: "",
			},
			want: "",
		},
		{
			name: "one ascii char",
			args: Args{
				s: "a",
			},
			want: "A",
		},
		{
			name: "word",
			args: Args{
				s: "word",
			},
			want: "Word",
		},
		{
			name: "already done",
			args: Args{
				s: "Spam",
			},
			want: "Spam",
		},
		{
			name: "multi words",
			args: Args{
				s: "multi words",
			},
			want: "Multi words",
		},
		{
			name: "shouting",
			args: Args{
				s: "SHOUTING",
			},
			want: "SHOUTING",
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := UppercaseFirst(tt.args.s)
			assert.Equal(t, tt.want, got)
		})
	}
}
