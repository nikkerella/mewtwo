# once

[sync package](https://pkg.go.dev/sync@go1.23.0#Once)

``` go
type Once struct {}

func (o *Once) Do(f func())
```

> Once is an object that will perform exactly one action. In the terminology of the Go memory model, the return from f “synchronizes before” the return from any call of once.Do(f).

## perform exactly one action?

sync.Once ensures that the function passed to its Do method is executed only once.

## Go memory model?