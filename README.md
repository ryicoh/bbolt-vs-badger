# bbolt vs badger

```
ryicoh@ryicohs-air bbolt-vs-badger % go test -bench .
goos: darwin
goarch: arm64
pkg: github.com/ryicoh/bbolt-vs-badger
BenchmarkBadgerPut-8         343           9243819 ns/op
BenchmarkBadgerGet-8        3572            324151 ns/op
BenchmarkBboltPut-8            1        3627057875 ns/op
BenchmarkBboltGet-8        23989             46146 ns/op
PASS
ok      github.com/ryicoh/bbolt-vs-badger       11.555s
```
