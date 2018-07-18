// Copyright 2018 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package base

import "testing"

func TestSlug(t *testing.T) {
	var dataTable = []struct {
		In, Out string
	}{
		{"привет", "privet"},
		{"Привет, друзья", "privet-druzya"},
	}
	for _, v := range dataTable {
		if Slug(v.In) != v.Out {
			t.Errorf("slug expected: %v, got: %v", v.Out, Slug(v.In))
		}
	}
}
