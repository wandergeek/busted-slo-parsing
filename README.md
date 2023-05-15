Just a demonstration of some weird behavior I'm experiencing with the openapi-generated SLO models

Run:
```
make generate-slo-client
go run main.go
```

Errors:

```
❯ go run main.go 
panic: unable to unmarshal data into slo.CreateSloRequest: data failed to match schemas in oneOf(SloResponseIndicator)

goroutine 1 [running]:
main.main()
        /Users/nick/work/busted-slo-parsing/main.go:39 +0x150
exit status 2
```