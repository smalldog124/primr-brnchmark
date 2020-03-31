# BANMACH-GO
## primr-brnchmark
[เขียน Benchmark อย่างไรใน GO](https://siamchamnankit.co.th/%E0%B9%80%E0%B8%82%E0%B8%B5%E0%B8%A2%E0%B8%99-banmacharks-%E0%B8%AD%E0%B8%A2%E0%B9%88%E0%B8%B2%E0%B8%87%E0%B9%84%E0%B8%A3%E0%B9%83%E0%B8%99-go-f3f05c118be3)

run brnchmark
```
$ go test ./prime_number -bench=.
$ go test ./prime_number -bench=. -benchmem
$ go test ./prime_number -bench=. -benchtime=20s
$ go test ./prime_number -bench=. -memprofile mem.out -cpuprofile cpu.out
```

## seal service
run seal service
```
$ docker-compose up -d
$ ab -k -c 10 -n 1000000 "http://127.0.0.1:3030/api/v1/health"
$ go tool pprof http://localhost:8080/debug/pprof/profile

```
open browser: http://localhost:8080/debug/pprof

## pprof tool
```
$ go tool pprof mem.out
$ go tool pprof cpu.out
```

## reference
https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go

https://golang.org/pkg/net/http/pprof/

https://blog.golang.org/pprof

https://jvns.ca/blog/2017/09/24/profiling-go-with-pprof/
