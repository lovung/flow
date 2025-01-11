package flow_test

import (
	"context"
	"log"
	"testing"

	"github.com/lovung/flow"
	"github.com/stretchr/testify/assert"
)

func TestSteps(t *testing.T) {
	handler := LoginHandler{}
	// Replace with net or API framework context
	ctx := context.Background()
	apiCtx := &LoginContext{
		Context:  ctx,
		username: "username", // Get from ctx instead
		password: "password", // Get from ctx instead
	}

	err := flow.Steps(
		apiCtx,
		handler.step1,
		handler.step2,
		handler.step2,
	)
	assert.NoError(t, err)
}

func TestGo(t *testing.T) {
	handler := LoginHandler{}
	// Replace with net or API framework context
	ctx := context.Background()
	apiCtx := &LoginContext{
		Context:  ctx,
		username: "username", // Get from ctx instead
		password: "password", // Get from ctx instead
	}

	// Running concurrently
	err := flow.Go(
		apiCtx,
		handler.step1,
		handler.step2,
		handler.step2,
	)
	assert.NoError(t, err)
}

type LoginContext struct {
	context.Context

	username string
	password string
}

func (c *LoginContext) Embed(ctx context.Context) {
	c.Context = ctx
}

type LoginHandler struct{}

func (h *LoginHandler) step1(ctx *LoginContext) error {
	log.Println(ctx.username)
	return nil
}

func (h *LoginHandler) step2(ctx *LoginContext) error {
	log.Println(ctx.password)
	return nil
}

func (h *LoginHandler) step3(ctx *LoginContext) error {
	log.Println(ctx.username, ctx.password)
	return nil
}
