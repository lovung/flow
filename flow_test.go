package flow_test

import (
	"context"
	"errors"
	"log"
	"testing"
	"time"

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

	t.Run("normal", func(t *testing.T) {
		err := flow.Seq(
			handler.step1,
			handler.step2,
			handler.step3,
		)(apiCtx)
		assert.NoError(t, err)
	})
	t.Run("error", func(t *testing.T) {
		err := flow.Seq(
			handler.step1,
			handler.step2Error,
			handler.step3,
		)(apiCtx)
		assert.Error(t, err)
	})
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

	t.Run("normal", func(t *testing.T) {
		err := flow.Go(
			handler.step1,
			handler.step2,
			handler.step3,
		)(apiCtx)
		assert.NoError(t, err)
	})
	t.Run("error", func(t *testing.T) {
		err := flow.Go(
			handler.step1,
			handler.step2Error,
			handler.step3,
		)(apiCtx)
		assert.Error(t, err)
	})
	t.Run("cancel", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		apiCtx := &LoginContext{
			Context:  ctx,
			username: "username", // Get from ctx instead
			password: "password", // Get from ctx instead
		}
		cancel()
		err := flow.Go(
			handler.step1,
			handler.step2,
			handler.step3,
		)(apiCtx)
		assert.NoError(t, err)
	})
	t.Run("timeout", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(
			context.Background(),
			1*time.Second,
		)
		apiCtx := &LoginContext{
			Context:  ctx,
			username: "username", // Get from ctx instead
			password: "password", // Get from ctx instead
		}
		cancel()
		err := flow.Go(
			handler.step1,
			handler.step2,
			handler.step3,
		)(apiCtx)
		assert.NoError(t, err)
	})
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
	time.Sleep(1 * time.Second)
	return nil
}

func (h *LoginHandler) step2(ctx *LoginContext) error {
	log.Println(ctx.password)
	time.Sleep(1 * time.Second)
	return nil
}
func (h *LoginHandler) step2Error(ctx *LoginContext) error {
	log.Println(ctx.username, ctx.password)
	time.Sleep(1 * time.Second)
	return errors.New("some random error")
}

func (h *LoginHandler) step3(ctx *LoginContext) error {
	log.Println(ctx.username, ctx.password)
	time.Sleep(1 * time.Second)
	return nil
}
