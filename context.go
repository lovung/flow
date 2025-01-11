// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package errgroup provides synchronization, error propagation, and Context
// cancelation for groups of goroutines working on subtasks of a common task.
//
// [errgroup.Group] is related to [sync.WaitGroup] but adds handling of tasks
// returning errors.
package flow

import (
	"context"
)

type EmbedableContext interface {
	context.Context
	Embed(context.Context)
}

func WithCancelCause[T EmbedableContext](parent T) (T, func(error)) {
	_ctx, cancel := context.WithCancelCause(parent)
	parent.Embed(_ctx)
	return parent, cancel
}
