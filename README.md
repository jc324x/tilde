# tilde # 

`brf` is a small utility package, used in a few of my projects: 

- [git-in-sync](https://github.com/jychri/git-in-sync): Keep all of
  your git repos in sync across multiple computers
- [js2x](https://github.com/jychri/js2x): Convert Javascript to Markdown and beyond with js2x

`tilde`'s only function, `Abs`, expands `~/` to `/Users/current`,
providing an absolute path to the current user's home directory. 

```go
// Package tilde implements shorthand "~/"
package tilde

import (
	"log"
	u "os/user"
	p "path"
	"strings"
)


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
