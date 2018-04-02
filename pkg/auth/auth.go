// Copyright 2018 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/zhuharev/blog/pkg/setting"

	"github.com/go-macaron/auth"
	macaron "gopkg.in/macaron.v1"
)

// SigninRequired defense write APIs from others
func SigninRequired() macaron.Handler {
	return auth.Basic(setting.App.Admin.Login, setting.App.Admin.Password)
}
