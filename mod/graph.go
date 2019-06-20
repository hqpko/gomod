package mod

import (
	"fmt"
	"sort"
	"strings"

	"github.com/hqpko/gosh"
	"github.com/urfave/cli"
)

func ActionGoModGraph(c *cli.Context) error {
	session := gosh.NewSession()
	session.SetHandlerIn(func(s *gosh.Session, cmd string) {})
	session.SetHandlerOut(handlerFormatGraph(c))
	return session.Run("go mod graph")
}

func handlerFormatGraph(c *cli.Context) func(s *gosh.Session, out []byte) {
	deep := getDeep(c)
	return func(s *gosh.Session, out []byte) {
		if root := getRootDep(out); root != nil {
			fmt.Println(root.String("", false, deep))
		}
	}
}

func getRootDep(out []byte) *dep {
	var root *dep
	cache := newDep("")
	for _, line := range strings.Split(string(out), "\n") {
		if path, dep, ok := getDepStr(line); ok {
			if root == nil && isRoot(path) {
				root = newDep(path)
			}
			cache.cache(path, dep)
		}
	}
	if root != nil {
		root.join(cache)
	}
	return root
}

type dep struct {
	path string
	deps map[string]*dep
}

func newDep(path string) *dep {
	return &dep{path: path, deps: map[string]*dep{}}
}

func (d *dep) cache(path, dep string) {
	if d.deps[path] == nil {
		d.deps[path] = newDep(path)
	}
	d.deps[path].deps[dep] = newDep(dep)
}

func (d *dep) join(d2 *dep) {
	if d2.deps[d.path] != nil {
		d.deps = d2.deps[d.path].deps
		delete(d2.deps, d.path)
	}
	for _, dep := range d.deps {
		dep.join(d2)
	}
}

/**
github.com/hqpko/gomod
├─ golang.org/x/crypto@v0.0.0-20190618222545-ea8f1a30c443
│  ├─ golang.org/x/net@v0.0.0-20190404232315-eb5bcb51f2a3
│  │  ├─ golang.org/x/crypto@v0.0.0-20190308221718-c2843e01d9a2
│  │  │  └─ golang.org/x/sys@v0.0.0-20190215142949-d0b11bdaac8a
│  │  └─ golang.org/x/text@v0.3.0
│  └─ golang.org/x/sys@v0.0.0-20190412213103-97732733099d
├─ github.com/hqpko/gosh@v0.0.2
└─ github.com/urfave/cli@v1.20.0
*/

func (d *dep) String(preStr string, isLast bool, deep int) string {
	if deep < 0 {
		return ""
	}
	out := ""
	if isRoot(d.path) {
		out += d.path
		preStr = ""
	} else {
		if isLast {
			out += "\n" + preStr + "└─ " + d.path
			preStr += "   "
		} else {
			out += "\n" + preStr + "├─ " + d.path
			preStr += "│  "
		}
	}

	sortedSlice := d.sort()
	for i, dep := range sortedSlice {
		out += dep.String(preStr, i == len(sortedSlice)-1, deep-1)
	}
	return out
}

func (d *dep) sort() []*dep {
	ds := make([]*dep, 0)
	for _, dep := range d.deps {
		ds = append(ds, dep)
	}
	sort.Slice(ds, func(i, j int) bool {
		return ds[i].path < ds[j].path
	})
	return ds
}
