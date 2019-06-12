# golang-note

#### golang currency

- 常用并发处理机制性能分析

```bash
    go test -bench=. -benchmem -benchtime=3s -cpuprofile=profile.out
```


- 并发控制 

- chan
- [semaphore](https://godoc.org/golang.org/x/sync/semaphore#example-package--WorkerPool)

![bench result](https://github.com/g-airport/tech-stack/blob/master/golang/currency-pattern/benchmark.jpg)
