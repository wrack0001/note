package 性能篇

import (
	"fmt"
	"strconv"
	"testing"
)

// Benchmark_IToa-8   	536403448	         2.250 ns/op
func Benchmark_IToa(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = strconv.Itoa(0)
	}
}

// Benchmark_FormatFloat-8   	19182428	        62.59 ns/op
func Benchmark_FormatFloat(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = strconv.FormatFloat(float64(i), 'g', 1, 64)
	}
}

// Benchmark_interface_int-8   	48251550	        25.47 ns/op
func Benchmark_interface_int(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = numberToStr(i)
	}
}

// Benchmark_interface_float-8   	18578846	        63.94 ns/op
func Benchmark_interface_float(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = numberToStr(float64(i))
	}
}

// Benchmark_interface_string-8   	519241784	         2.309 ns/op
func Benchmark_interface_string(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = numberToStr("1")
	}
}

// numberToStr interface 转 string
func numberToStr(num interface{}) string {
	switch num.(type) {
	case int:
		return strconv.Itoa(num.(int))
	case int32:
		return strconv.Itoa(int(num.(int32)))
	case float64:
		return strconv.FormatFloat(num.(float64), 'g', 1, 64)
	case string:
		return num.(string)
	default:

	}
	fmt.Println("de")
	return ""
}

// Benchmark_fmt-8   	18345724	        64.90 ns/op
func Benchmark_fmt(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%d", i)
	}
}

// Benchmark_fmt_v-8   	17300934	        67.23 ns/op
func Benchmark_fmt_v(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%v", i)
	}
}

// Benchmark_fmt_float64-8   	 7647301	       156.2 ns/op
func Benchmark_fmt_float64(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%f", float64(i))
	}
}

// Benchmark_fmt_float64_v-8   	12051979	       104.7 ns/op
func Benchmark_fmt_float64_v(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%v", float64(i))
	}
}

// Benchmark_fmt_string-8   	37324408	        31.97 ns/op
func Benchmark_fmt_string(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s", "1")
	}
}

// Benchmark_fmt_string_v-8   	37517586	        32.76 ns/op
func Benchmark_fmt_string_v(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%v", "i")
	}
}

// Benchmark_StrAndStr-8   	1000000000	         0.3179 ns/op
func Benchmark_StrAndStr(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = "i" + "i"
	}
}

// Benchmark_fmt_string_20-8   	 3258351	       372.4 ns/op
func Benchmark_fmt_string_20(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%v:%v:%v:%v:%v:%v:%v:%v:%v:%v:%v:%v:%v:%v:%v:%v:%v:%v:%v:%v:",
			"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9")
	}
}

// Benchmark_fmt_string_20_s-8   	 3328732	       369.7 ns/op
func Benchmark_fmt_string_20_s(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s:%s:%s:%s:%s:%s:%s:%s:%s:%s:%s:%s:%s:%s:%s:%s:%s:%s:%s:%s:",
			"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9")
	}
}

// Benchmark_StrAndStr_20-8   	1000000000	         0.3165 ns/op
func Benchmark_StrAndStr_20(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = "0" + "1" + "2" + "3" + "4" + "5" + "6" + "7" + "8" + "9" + "0" + "1" + "2" + "3" + "4" + "5" + "6" + "7" + "8" + "9"
	}
}
