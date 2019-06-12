# golang-note

#### golang currency

- 常用并发处理机制性能分析

```bash
    go test -bench=. -benchmem -benchtime=3s -cpuprofile=profile.out
```

![bench result](https://github.com/g-airport/golang-note/blob/master/currency-pattern/benchmark.jpg)

- 并发控制 

- chan
- [semaphore](https://godoc.org/golang.org/x/sync/semaphore#example-package--WorkerPool)
