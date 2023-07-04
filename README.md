Just a demonstration of some weird behavior I'm experiencing with the openapi-generated SLO models

Run:
```
make generate-slo-client
go run main.go
```

Errors:

```
❯ go run main.go 
matched IndicatorPropertiesApmAvailability
matched IndicatorPropertiesApmLatency
panic: unable to unmarshal data into slo.SloResponse: data matches more than one schema in oneOf(SloResponseIndicator)

goroutine 1 [running]:
main.main()
        /Users/nick/work/busted-slo-parsing/main.go:149 +0x148
exit status 2
```