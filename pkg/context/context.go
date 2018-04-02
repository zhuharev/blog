// Copyright 2018 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package context contain middleware what registre special Context
// for macaron framework
package context

import macaron "gopkg.in/macaron.v1"

// Context wrap *macaron.Context
type Context struct {
	*macaron.Context
}

// Contexter middleware
func Contexter() macaron.Handler {
	return func(c *macaron.Context) {
		ctx := &Context{
			Context: c,
		}

		c.Map(ctx)
	}
}
