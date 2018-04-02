// Copyright 2018 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"github.com/zhuharev/blog/pkg/setting"

	"github.com/asdine/storm"
)

var db *storm.DB

// NewContext connect to db
func NewContext() (err error) {
	db, err = storm.Open(setting.App.Data.KvPath)
	return
}
