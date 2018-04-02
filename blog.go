// Copyright 2018 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Blog is web server primary hosted on blog.zhuharev.ru
package main

import (
	"os"

	"github.com/zhuharev/blog/cmd"
	"gopkg.in/urfave/cli.v2"
)

var version = "1.0.0"

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			cmd.Blog,
		},
		Version: version,
	}
	app.Run(os.Args)
}
