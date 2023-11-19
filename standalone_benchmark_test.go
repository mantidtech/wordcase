package wordcase

import (
	"bytes"
	"math/rand"
	"testing"
)

// noinspection GoUnusedGlobalVariable
var result string // to avoid compiler optimisations

func helperGenTestString(size int, pctRm, pctSep float64) string {
	var b bytes.Buffer
	for i := 0; i < size; i++ {
		p := rand.Float64()
		if p < pctRm {
			b.WriteRune(' ')
		} else if p < pctRm+pctSep {
			b.WriteRune('A')
		} else {
			b.WriteRune('a')
		}
	}
	return b.String()
}

var benchmarkStrings = map[string]string{
	"1k char  10% rm  10% sep":  helperGenTestString(1000, 0.1, 0.1),
	"1k char  10% rm  20% sep":  helperGenTestString(1000, 0.1, 0.2),
	"1k char  20% rm  10% sep":  helperGenTestString(1000, 0.2, 0.1),
	"1k char  20% rm  20% sep":  helperGenTestString(1000, 0.2, 0.2),
	"1k char  50% rm  50% sep":  helperGenTestString(1000, 0.5, 0.5),
	"5k char  20% rm  20% sep":  helperGenTestString(5000, 0.2, 0.2),
	"10k char  20% rm  20% sep": helperGenTestString(10000, 0.2, 0.2),
}

func BenchmarkCamelCase(b *testing.B) {
	for name, str := range benchmarkStrings {
		b.Run(name, func(b *testing.B) {
			var r string
			for n := 0; n < b.N; n++ {
				r = CamelCase(str)
			}
			result = r
		})
	}
}

func BenchmarkSnakeCase(b *testing.B) {
	for name, str := range benchmarkStrings {
		b.Run(name, func(b *testing.B) {
			var r string
			for n := 0; n < b.N; n++ {
				r = SnakeCase(str)
			}
			result = r
		})
	}
}

func BenchmarkKebabCase(b *testing.B) {
	for name, str := range benchmarkStrings {
		b.Run(name, func(b *testing.B) {
			var r string
			for n := 0; n < b.N; n++ {
				r = KebabCase(str)
			}
			result = r
		})
	}
}

//                                                          Inlined         Pipelined
//  BenchmarkSnakeCase/#00-16                 20459869        58.6 ns/op      66.3 ns/op
//  BenchmarkSnakeCase/simple-16               2649813       434 ns/op       457 ns/op
//  BenchmarkSnakeCase/two-tokens-16           1296022       919 ns/op      1005 ns/op
//  BenchmarkSnakeCase/a-16                    4428896       237 ns/op       253 ns/op
//  BenchmarkSnakeCase/A-16                    4206933       262 ns/op       263 ns/op
//  BenchmarkSnakeCase/a---------b-16          1431070       838 ns/op       864 ns/op
//  BenchmarkSnakeCase/dooker-16               2533857       443 ns/op       441 ns/op
//  BenchmarkSnakeCase/dookerSpam99_rawr-16     737336      1405 ns/op      1430 ns/op
//  BenchmarkSnakeCase/IDOne_XMLHttp_ON-16      594961      1959 ns/op      2025 ns/op
//  BenchmarkSnakeCase/maxID-16                1764518       700 ns/op       669 ns/op
//  BenchmarkSnakeCase/maxId-16                1717924       673 ns/op       670 ns/op
//  BenchmarkSnakeCase/ENV_VAR-16              1286334       961 ns/op       922 ns/op
//  BenchmarkSnakeCase/snake_case-16           1201674       923 ns/op       911 ns/op
//  BenchmarkSnakeCase/kebab-case-16           1294059       917 ns/op       920 ns/op
//  BenchmarkSnakeCase/IDOne-16                1518310       780 ns/op       829 ns/op
//  BenchmarkSnakeCase/99two-16                2754214       425 ns/op       423 ns/op
//  BenchmarkSnakeCase/99Two-16                1641892       709 ns/op       717 ns/op
//  BenchmarkSnakeCase/*a.B-16                 1618341       713 ns/op       709 ns/op
