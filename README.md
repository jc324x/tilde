# tilde # 

`tilde` is a small package, used in a few of my open source projects: 

- [git-in-sync](https://github.com/jychri/git-in-sync): Keep all of
  your git repos in sync across multiple computers
- [js2x](https://github.com/jychri/js2x): Convert Javascript to Markdown and beyond with js2x

`tilde`'s only function `Abs` returns an absolute path given
a string prefixed with `~/`. 

**Note:** The package does not support Windows at this time.

```go
// +build !windows

// Package tilde transforms tilde prefixed strings into absolute paths
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
```
