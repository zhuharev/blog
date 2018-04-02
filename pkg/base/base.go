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
	htmlFlags := 0
	//htmlFlags |= blackfriday.HTML_SKIP_STYLE
	htmlFlags |= blackfriday.HTML_OMIT_CONTENTS

	// set up the parser
	extensions := 0
	extensions |= blackfriday.EXTENSION_NO_INTRA_EMPHASIS
	extensions |= blackfriday.EXTENSION_TABLES
	extensions |= blackfriday.EXTENSION_FENCED_CODE
	extensions |= blackfriday.EXTENSION_AUTOLINK
	extensions |= blackfriday.EXTENSION_STRIKETHROUGH
	extensions |= blackfriday.EXTENSION_SPACE_HEADERS
	extensions |= blackfriday.EXTENSION_NO_EMPTY_LINE_BEFORE_BLOCK

	/*	if setting.Markdown.EnableHardLineBreak {
		extensions |= blackfriday.EXTENSION_HARD_LINE_BREAK
	}*/

	body = blackfriday.Markdown(body, blackfriday.HtmlRenderer(htmlFlags, "", ""), extensions)
	return body
}
