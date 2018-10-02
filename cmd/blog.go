// Copyright 2018 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package cmd used for include cli commands from main project file
package cmd

import (
	"html/template"
	"log"
	"ug/pkg/base"

	"github.com/go-macaron/binding"
	"github.com/zhuharev/blog/controllers"
	"github.com/zhuharev/blog/models"
	"github.com/zhuharev/blog/pkg/auth"
	"github.com/zhuharev/blog/pkg/context"
	"github.com/zhuharev/blog/pkg/setting"
	jade "github.com/zhuharev/go-macaron-jade"

	macaron "gopkg.in/macaron.v1"
	"gopkg.in/urfave/cli.v2"
)

// Blog is cli command for start blog server
var Blog = &cli.Command{
	Name:  "web",
	Usage: "Start blog web server",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "config, c",
			Usage: "Config file path",
			Value: "conf/app.yaml",
		},
	},
	Action: runWeb,
}

func runWeb(ctx *cli.Context) error {
	if err := setting.NewContext(ctx.String("config")); err != nil {
		log.Fatalf("err loading conf file: %v", err)
	}
	if err := models.NewContext(); err != nil {
		log.Fatalf("err init db: %v", err)
	}

	m := macaron.Classic()
	// m.Use(macaron.Renderer(macaron.RenderOptions{
	// 	Layout: "layout",
	// 	Funcs: []template.FuncMap{
	// 		{"markdown": func(in string) template.HTML { return template.HTML(base.RenderMarkdownString(in)) }},
	// 	},
	// }))
	m.Use(jade.Renderer(jade.Options{
		Directory:  "templates",       // Specify what path to load the templates from.
		Extensions: []string{".jade"}, // Specify extensions to load for templates.
		Funcs: []template.FuncMap{map[string]interface{}{
			"markdown": func(in string) template.HTML { return template.HTML(base.RenderMarkdownString(in)) },
		}}, // Specify helper function maps for templates to access.
		Charset:    "UTF-8", // Sets encoding for json and html content-types. Default is "UTF-8".
		IndentJSON: true,    // Output human readable JSON
	}))
	m.Use(context.Contexter())

	m.Group(setting.App.Admin.Path, func() {
		m.Get("/create", controllers.Create)
		m.Post("/create", binding.Bind(models.Post{}), controllers.PostCreate)
	}, auth.SigninRequired())

	m.Get("/:slug", controllers.Show)
	m.Get("/", controllers.List)

	m.Run()
	return nil
}
