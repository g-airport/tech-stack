# golang-note

#### golang currency

- 常用并发处理机制性能分析

```bash
    go test -bench=. -benchmem -benchtime=3s -cpuprofile=profile.out
```

![bench result](https://github.com/g-airport/tech-stack/blob/master/golang/currency-pattern/benchmark.jpg)
