//go:build tools
// +build tools

package tools

import (
	_ "bitbucket.org/liamstask/goose/cmd/goose"
	_ "golang.org/x/lint/golint"
)
