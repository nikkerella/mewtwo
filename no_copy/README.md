# noCopy

``` go
type noCopy struct{}

func (*noCopy) Lock()   {}
func (*noCopy) Unlock() {}
```

Use ``go vet main.go`` to examines the source code and looks for potential bugs, or sth may lead to unexpected behavior.

## Accidental Copying

- Unintentionally make a copy of a struct instead of referencing the original.
- Structs in Go are value types, meaning that when you assign or pass a struct, it is copied by default.
- A copy will not share the same state as the original, leading to bugs that are often hard to trace.