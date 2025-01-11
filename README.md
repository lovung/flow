# Flow

`flow` is a lightweight Golang package designed to simplify the execution of sequential and concurrent steps. It leverages contexts for dependency management, making it easy to build robust and efficient workflows.

## Features

- **Sequential Execution**: Run a series of steps consecutively using `Seq`.
- **Concurrent Execution**: Run multiple steps concurrently with `Go`.
- **Context Integration**: Full support for `context.Context` and custom embedded contexts.

## Installation

Install the package via `go get`:

```bash
go get github.com/lovung/flow
```

## Usage

### Sequential Execution (`Seq`)

The `Seq` function runs a series of steps in order. Each step is executed one after the other, and the execution stops if any step returns an error.

```go
package main

import (
	"context"
	"log"

	"github.com/lovung/flow"
)

type MyContext struct {
	context.Context
	Data string
}

type Handler struct{}

func (h *Handler) step1(ctx *MyContext) error {
	log.Println("Step 1:", ctx.Data)
	return nil
}

func (h *Handler) step2(ctx *MyContext) error {
	log.Println("Step 2:", ctx.Data)
	return nil
}

func main() {
	handler := &Handler{}
	ctx := &MyContext{
		Context: context.Background(),
		Data:    "Example Data",
	}

	err := flow.Seq(handler.step1, handler.step2)(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Concurrent Execution (`Go`)

The `Go` function runs multiple steps concurrently. It waits for all steps to complete and returns the first error encountered, if any.

```go
package main

import (
	"context"
	"log"

	"github.com/lovung/flow"
)

type MyContext struct {
	context.Context
	Data string
}

// Embed implements EmbedableContext interface
func (c *MyContext) Embed(ctx context.Context) {
	c.Context = ctx
}

type Handler struct{}

func (h *Handler) step1(ctx *MyContext) error {
	log.Println("Step 1:", ctx.Data)
	return nil
}

func (h *Handler) step2(ctx *MyContext) error {
	log.Println("Step 2:", ctx.Data)
	return nil
}

func main() {
	handler := &Handler{}
	ctx := &MyContext{
		Context: context.Background(),
		Data:    "Example Data",
	}

	err := flow.Go(handler.step1, handler.step2)(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
```

### Combine Execution (`Go`)

The `Go` function runs multiple steps sequentially or concurrently. 

```go
package main

import (
	"context"
	"log"

	"github.com/lovung/flow"
)

type MyContext struct {
	context.Context
	Data string
}

// Embed implements EmbedableContext interface
func (c *MyContext) Embed(ctx context.Context) {
	c.Context = ctx
}

type Handler struct{}

func (h *Handler) step1(ctx *MyContext) error {
	log.Println("Step 1:", ctx.Data)
	return nil
}

func (h *Handler) step2(ctx *MyContext) error {
	log.Println("Step 2:", ctx.Data)
	return nil
}

func (h *Handler) step3(ctx *MyContext) error {
	log.Println("Step 3:", ctx.Data)
	return nil
}

func main() {
	handler := &Handler{}
	ctx := &MyContext{
		Context: context.Background(),
		Data:    "Example Data",
	}

	err := flow.Seq(
        handler.step1, 
        // step2 and step3 will be run concurrently but after step1
        flow.Go(
            handler.step2,
            handler.step3,
        ),
    )(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
```

## Testing

`flow` includes built-in tests to demonstrate its functionality. Refer to `flow_test.go` for example test cases:

```go
github.com/lovung/flow/flow_test.go
```

Run tests using:

```bash
go test ./...
```

## License

This package is open source and available under the [MIT License](LICENSE).

