package wordcase

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestFirst provides unit test coverage for ToFirst()
func TestFirst(t *testing.T) {
	type Args struct {
		t Tokens
	}

	tests := []struct {
		name string
		args Args
		want []int
	}{
		{
			name: "empty",
			args: Args{t: Tokens{}},
			want: []int{0},
		},
		{
			name: "one",
			args: Args{t: Tokens{"one"}},
			want: []int{0},
		},
		{
			name: "two",
			args: Args{t: Tokens{"one", "two"}},
			want: []int{0},
		},
		{
			name: "three",
			args: Args{t: Tokens{"one", "two", "three"}},
			want: []int{0},
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := ToFirst(tt.args.t)
			assert.Equal(t, tt.want, got)
		})
	}
}

// TestLast provides unit test coverage for ToLast()
func TestLast(t *testing.T) {
	type Args struct {
		t Tokens
	}

	tests := []struct {
		name string
		args Args
		want []int
	}{
		{
			name: "empty",
			args: Args{t: Tokens{}},
			want: []int{-1},
		},
		{
			name: "one",
			args: Args{t: Tokens{"one"}},
			want: []int{0},
		},
		{
			name: "two",
			args: Args{t: Tokens{"one", "two"}},
			want: []int{1},
		},
		{
			name: "three",
			args: Args{t: Tokens{"one", "two", "three"}},
			want: []int{2},
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := ToLast(tt.args.t)
			assert.Equal(t, tt.want, got)
		})
	}
}

// TestRest provides unit test coverage for ToRest()
func TestRest(t *testing.T) {
	type Args struct {
		t Tokens
	}

	tests := []struct {
		name string
		args Args
		want []int
	}{
		{
			name: "empty",
			args: Args{t: Tokens{}},
			want: nil,
		},
		{
			name: "one",
			args: Args{t: Tokens{"one"}},
			want: nil,
		},
		{
			name: "two",
			args: Args{t: Tokens{"one", "two"}},
			want: []int{1},
		},
		{
			name: "three",
			args: Args{t: Tokens{"one", "two", "three"}},
			want: []int{1, 2},
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := ToRest(tt.args.t)
			assert.Equal(t, tt.want, got)
		})
	}
}

// TestAll provides unit test coverage for ToAll()
func TestAll(t *testing.T) {
	type Args struct {
		t Tokens
	}

	tests := []struct {
		name string
		args Args
		want []int
	}{
		{
			name: "empty",
			args: Args{t: Tokens{}},
			want: []int{},
		},
		{
			name: "one",
			args: Args{t: Tokens{"one"}},
			want: []int{0},
		},
		{
			name: "two",
			args: Args{t: Tokens{"one", "two"}},
			want: []int{0, 1},
		},
		{
			name: "three",
			args: Args{t: Tokens{"one", "two", "three"}},
			want: []int{0, 1, 2},
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := ToAll(tt.args.t)
			assert.Equal(t, tt.want, got)
		})
	}
}

// TestKeyWordFn provides unit test coverage for KeyWordFn()
func TestKeyWordFn(t *testing.T) {

	tests := []struct {
		name     string
		keywords []string
		tokens   Tokens
		want     []int
	}{
		{
			name:     "nil matcher",
			keywords: []string{},
			tokens:   Tokens{"one", "two", "three"},
			want:     nil,
		},
		{
			name:     "one",
			keywords: []string{"one"},
			tokens:   Tokens{"one", "two", "three"},
			want:     []int{0},
		},
		{
			name:     "some",
			keywords: []string{"one", "three"},
			tokens:   Tokens{"one", "two", "three"},
			want:     []int{0, 2},
		},
		{
			name:     "many",
			keywords: []string{"one", "two", "three"},
			tokens:   Tokens{"one", "two", "three"},
			want:     []int{0, 1, 2},
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			kw := KeyWordFn(tt.keywords)
			got := kw(tt.tokens)
			assert.Equal(t, tt.want, got)
		})
	}
}

// TestNot provides unit test coverage for Not()
func TestNot(t *testing.T) {
	tests := []struct {
		name   string
		tokens Tokens
		idx    []int
		want   []int
	}{
		{
			name:   "nothing",
			tokens: Tokens{},
			idx:    []int{},
			want:   nil,
		},
		{
			name:   "one from one",
			tokens: Tokens{"one"},
			idx:    []int{0},
			want:   nil,
		},
		{
			name:   "none from one",
			tokens: Tokens{"one"},
			idx:    []int{},
			want:   []int{0},
		},
		{
			name:   "none from two",
			tokens: Tokens{"one", "two"},
			idx:    []int{},
			want:   []int{0, 1},
		},
		{
			name:   "one from two",
			tokens: Tokens{"one", "two"},
			idx:    []int{1},
			want:   []int{0},
		},
		{
			name:   "other from two",
			tokens: Tokens{"one", "two"},
			idx:    []int{0},
			want:   []int{1},
		},
		{
			name:   "two from two",
			tokens: Tokens{"one", "two"},
			idx:    []int{0, 1},
			want:   nil,
		},
		{
			name:   "middle from three",
			tokens: Tokens{"one", "two", "three"},
			idx:    []int{1},
			want:   []int{0, 2},
		},
		{
			name:   "ends from three",
			tokens: Tokens{"one", "two", "three"},
			idx:    []int{0, 2},
			want:   []int{1},
		},
	}

	tsGen := func(i []int) TokenSelector {
		return func(t Tokens) []int {
			return i
		}
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ts := tsGen(tt.idx)
			not := Not(ts)
			got := not(tt.tokens)
			assert.Equal(t, tt.want, got)
		})
	}
}

// TestAnd provides unit test coverage for And()
func TestAnd(t *testing.T) {
	tests := []struct {
		name   string
		tokens Tokens
		one    []int
		two    []int
		want   []int
	}{
		{
			name:   "nothing",
			tokens: Tokens{},
			one:    []int{},
			two:    []int{},
			want:   nil,
		},
		{
			name:   "true and true",
			tokens: Tokens{"one"},
			one:    []int{0},
			two:    []int{0},
			want:   []int{0},
		},
		{
			name:   "true and false",
			tokens: Tokens{"one"},
			one:    []int{0},
			two:    []int{},
			want:   nil,
		},
		{
			name:   "false and true",
			tokens: Tokens{"one"},
			one:    []int{},
			two:    []int{0},
			want:   nil,
		},
		{
			name:   "false and false",
			tokens: Tokens{"one"},
			one:    []int{},
			two:    []int{},
			want:   nil,
		},
		{
			name:   "more",
			tokens: Tokens{"one", "two", "three"},
			one:    []int{0, 1},
			two:    []int{1, 2},
			want:   []int{1},
		},
	}

	tsGen := func(i []int) TokenSelector {
		return func(t Tokens) []int {
			return i
		}
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s1 := tsGen(tt.one)
			s2 := tsGen(tt.two)
			and := And(s1, s2)
			got := and(tt.tokens)
			if tt.want == nil {
				assert.Empty(t, got)
			} else {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

// TestOr provides unit test coverage for Or()
func TestOr(t *testing.T) {
	tests := []struct {
		name   string
		tokens Tokens
		one    []int
		two    []int
		want   []int
	}{
		{
			name:   "nothing",
			tokens: Tokens{},
			one:    []int{},
			two:    []int{},
			want:   nil,
		},
		{
			name:   "true or true",
			tokens: Tokens{"one"},
			one:    []int{0},
			two:    []int{0},
			want:   []int{0},
		},
		{
			name:   "true or false",
			tokens: Tokens{"one"},
			one:    []int{0},
			two:    []int{},
			want:   []int{0},
		},
		{
			name:   "false or true",
			tokens: Tokens{"one"},
			one:    []int{},
			two:    []int{0},
			want:   []int{0},
		},
		{
			name:   "false or false",
			tokens: Tokens{"one"},
			one:    []int{},
			two:    []int{},
			want:   nil,
		},
		{
			name:   "more",
			tokens: Tokens{"one", "two", "three"},
			one:    []int{0, 1},
			two:    []int{1, 2},
			want:   []int{0, 1, 2},
		},
	}

	tsGen := func(i []int) TokenSelector {
		return func(t Tokens) []int {
			return i
		}
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s1 := tsGen(tt.one)
			s2 := tsGen(tt.two)
			or := Or(s1, s2)
			got := or(tt.tokens)
			if tt.want == nil {
				assert.Empty(t, got)
			} else {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
