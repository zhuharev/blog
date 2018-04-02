// Copyright 2018 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package setting contain global app settings
package setting

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

// App variable used for map config file
var App struct {
	Admin struct {
		Path     string
		Login    string
		Password string
	}
	Data struct {
		KvPath string `yaml:"kv_path"`
	}
}

// NewContext read file from confFilePath and map yaml data to App
func NewContext(confFilePath string) (err error) {
	data, err := ioutil.ReadFile(confFilePath)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(data, &App)
	if err != nil {
		return
	}
	log.Printf("%+v", App)
	return
}
