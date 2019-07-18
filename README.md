# Benchmark Notes

[Profile your benchmark](https://medium.com/@felipedutratine/profile-your-benchmark-with-pprof-fb7070ee1a94)

``` bash
go test -bench=. -benchmem -cpuprofile profile.out
go tool pprof -http=:8080 profile.out
```
