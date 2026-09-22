# Exponential moving average

Raw metric streams are noisy. An EWMA smooths this by weighting recent data heavier than older data. Picture a sudden spike. The trend line follows it up, but slower. You get a clean signal.

This Go package processes streaming data on the fly. It is dependency-free. Go Ewma uses only the standard library.

```
ewma.go
```

Check the test file next to the implementation for concrete examples.