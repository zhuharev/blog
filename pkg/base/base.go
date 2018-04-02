// Copyright 2018 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package base

import "github.com/russross/blackfriday"

// RenderMarkdownString render markdown to string
func RenderMarkdownString(in string) string {
	return string(RenderRawMarkdown([]byte(in)))
}

// RenderRawMarkdown render markdown to byte slice
func RenderRawMarkdown(body []byte) []byte {
	body = blackfriday.Run(body)
	return body
}
