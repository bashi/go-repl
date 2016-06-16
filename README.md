# go-repl

A simple library for running a read-eval-print-loop. Just a wrapper of [peterh/liner](https://github.com/peterh/liner) for now.

## Example

Using callback:
```go
repl.Run(func(line string) error) {
  fmt.Printf("read: %s\n", line)
})
```

Using handler:
```go
type handler struct {
  c int
}
func (h *handler) Handle(line string) error {
  fmt.Printf("%d: %s\n", h.c, line)
  h.c += 1
  if (h.c >= 10) {
    return io.EOF
  }
  return nil
}

func main() {
  repl.RunHandler(&handler{1})
}
```
