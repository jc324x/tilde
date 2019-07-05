// +build !windows

// Package tilde implements shorthand "~/"
package tilde

import (
	"log"
	u "os/user"
	p "path"
	"strings"
)

// Abs replaces "~/" with "/User/$user/" and returns a clean path.
func Abs(path string) string {

	var user *u.User
	var err error

	if user, err = u.Current(); err != nil {
		log.Fatalf("Unable to identify current user")
	}

	if !p.IsAbs(path) {
		return p.Join(user.HomeDir, strings.TrimPrefix(path, "~/"))
	}

	return p.Clean(path)
}