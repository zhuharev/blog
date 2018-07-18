// Copyright 2018 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package base

import (
	"strings"
)

// Slug generate uri (generaly use latin and replace space with `-`)
func Slug(in string) string {
	in = strings.ToLower(in)
	ruAlphabet := "абвгдеёжзийклмнопрстуфхцчшщьыъэюя, "
	transAlphabet := []string{"a", "b", "v", "g", "d", "e", "yo", "zh", "z", "i", "i", "k", "l",
		"m", "n", "o", "p", "r", "s", "t", "u", "f", "h", "c", "ch", "sh", "sh", "", "i", "-", "e", "y", "ya", "", "-"}
	replacerSlice := []string{}
	for i, rune := range []rune(ruAlphabet) {
		replacerSlice = append(replacerSlice, string(rune), transAlphabet[i])
	}

	rep := strings.NewReplacer(replacerSlice...)

	out := rep.Replace(in)
	return out
}
