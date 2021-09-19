package main

import (
	t "github.com/seyyedaghaei/ts2go/transpile"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	switch os.Args[1] {
	case "run":
		run(os.Args[2])
	case "build":
		build(os.Args[2])
	}
}


func toGo(file string) string {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	return t.Transpile(string(bytes))
}

func transpile(file string, f func(dir string)) {
	dir := newCacheDir()
	code := toGo(file)
	err := ioutil.WriteFile(filepath.Join(dir, "main.go"), []byte(code), os.ModePerm)
	handleError(err)
	f(dir)
	err = os.RemoveAll(dir)
	handleError(err)
}

func run(file string) {
	transpile(file, func(dir string) {
		err := runCommand("go", "run", filepath.Join(dir, "main.go"))
		handleError(err)
	})
}

func build(file string) {
	transpile(file, func(dir string) {
		err := runCommand("go", "build", filepath.Join(dir, "main.go"))
		handleError(err)
	})
}