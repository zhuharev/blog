// Copyright 2018 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package context contain middleware what registre special Context
// for macaron framework
package context

import (
	"github.com/zhuharev/go-macaron-jade"
	macaron "gopkg.in/macaron.v1"
)

// Context wrap *macaron.Context
type Context struct {
	*macaron.Context
	jade.Render
}

// Contexter middleware
func Contexter() macaron.Handler {
	return func(c *macaron.Context, r jade.Render) {
		ctx := &Context{
			Context: c,
			Render:  r,
		}
		c.Map(ctx)
	}
}
