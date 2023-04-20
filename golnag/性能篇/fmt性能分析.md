# 论：fmt 性能分析

>>> **导论：** 在写golang代码的时候我们经常会遇到：字符串拼接、float|int等类型转string
> 
> **结论：**
> 
> 1、float|int转string尽量使用strconv包
> 
> 2、string拼接尽量使用+号拼接

## 1、执行测试结果
- 通过下表可以详细看出：
  - 如果直接使用strconv.Itoa比使用fmt.Sprintf要快28倍
  - 使用+拼接字符串两个字符串拼接与20个字符串拼接基本不会影响性能
  - 使用fmt.Sprintf拼接字符串随着拼接字符串数量的增加，性能也会显著下降
  - 当2个字符串进行拼接时：+号拼接性能是fmt.Sprintf的100+倍
  - 当20个字符串进行拼接时：+号拼接性能是fmt.Sprintf的1100+倍

| 对应Benchmark方法             | 目的                        | 类型        | 执行次数       | ns/op        | 方法描述                                              |
|---------------------------|---------------------------|-----------|------------|--------------|---------------------------------------------------|
| **Benchmark_IToa**        | int->sting                | int       | 536403448  | 2.250 ns/op  | 直接使用 strconv.Itoa                                 |
| Benchmark_interface_int   | int->interface->sting     | interface | 48251550   | 25.47 ns/op  | 使用自定方法：interface switch+断言后使用 strconv.Itoa        |
| Benchmark_fmt             | int->sting                | int       | 18345724   | 64.90 ns/op  | 直接使用 fmt.Sprintf("%d", i)                         |
| Benchmark_fmt_v           | int->sting                | int       | 17300934   | 67.23 ns/op  | 直接使用 fmt.Sprintf("%v", i)                         |
| **Benchmark_FormatFloat** | float64->sting            | float64   | 19182428   | 62.59 ns/op  | 直接使用 strconv.FormatFloat                          |
| Benchmark_interface_float | float64->interface->sting | interface | 18578846   | 63.94 ns/op  | 使用自定方法：interface switch+断言后使用 strconv.FormatFloat |
| Benchmark_fmt_float64_v   | float64->sting            | float64   | 12051979   | 104.7 ns/op  | 直接使用 fmt.Sprintf("%v", float64(i))                |
| Benchmark_fmt_float64     | float64->sting            | float64   | 7647301    | 156.2 ns/op  | 直接使用 fmt.Sprintf("%f", float64(i))                |
| **Benchmark_fmt_string**  | sting->sting              | string    | 37324408   | 31.97 ns/op  | 直接使用 fmt.Sprintf("%s", "1")                       |
| Benchmark_fmt_string_v    | sting->sting              | string    | 37517586   | 32.76 ns/op  | 直接使用 fmt.Sprintf("%v", "1")                       |
| **Benchmark_StrAndStr**   | sting+sting               | string    | 1000000000 | 0.3179 ns/op | 直接使用 "i" + "i"                                    |
| Benchmark_fmt_string_20   | sting+sting               | string    | 3258351    | 372.4 ns/op  | 直接使用 fmt.Sprintf("%v") 拼接 20个字符串                  |
| Benchmark_fmt_string_20_s | sting+sting               | string    | 3328732    | 369.7 ns/op  | 直接使用 fmt.Sprintf("%s") 拼接 20个字符串                  |
| Benchmark_StrAndStr_20    | sting+sting               | string    | 1000000000 | 0.3165 ns/op | 直接使用 + 拼接 20个字符串                                  |

[UT代码地址](https://github.com/wrack0001/note/blob/master/golnag/%E6%80%A7%E8%83%BD%E7%AF%87/fmt%E6%80%A7%E8%83%BD%E5%88%86%E6%9E%90_test.go)