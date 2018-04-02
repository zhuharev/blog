// Copyright 2018 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"time"

	"github.com/asdine/storm"
)

// Post model for blog post
type Post struct {
	ID int64 `storm:"id,increment"`

	Title string
	Slug  string `storm:"unique"`
	Body  string

	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

// Create insert post in database (obtain ID) and return *Post
func Create(p *Post) (err error) {
	p.CreatedAt = time.Now()
	err = db.Save(p)
	return
}

// Get returns post by slug
func Get(slug string) (*Post, error) {
	p := new(Post)
	err := db.One("Slug", slug, p)
	return p, err
}

// List return posts in descending order
func List() (posts []Post, err error) {
	//query := q.
	err = db.All(&posts, storm.Reverse())
	return
}
