// Copyright 2018 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package controllers contains no-categorized controllers
package controllers

import (
	"log"

	"github.com/zhuharev/blog/models"
	"github.com/zhuharev/blog/pkg/context"
)

// Create handler shows creating form
func Create(ctx *context.Context) {
	ctx.HTML(200, "create")
}

// PostCreate handle post request with form
func PostCreate(ctx *context.Context, post models.Post) {
	err := models.Create(&post)
	if err != nil {
		ctx.Error(200, err.Error())
		return
	}
	log.Printf("%+v", post)
	ctx.Redirect("/" + post.Slug)
}

// Show shows single post
func Show(ctx *context.Context) {
	var (
		slug = ctx.Params(":slug")
	)
	post, err := models.Get(slug)
	if err != nil {
		ctx.Error(404, "Страница не найдена")
		return
	}
	ctx.Data["post"] = post
	ctx.HTML(200, "show")
}

// List show recent blog posts
func List(ctx *context.Context) {
	posts, err := models.List()
	if err != nil {
		ctx.Error(404, "Страница не найдена")
		return
	}
	ctx.Data["posts"] = posts
	ctx.HTML(200, "list")
}
