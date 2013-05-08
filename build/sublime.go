package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
)

var verbose = true

func copy(a, b string) {
	f1, err := os.Open(a)
	if err != nil {
		panic(err)
	}
	defer f1.Close()
	f2, err := os.Create(b)
	if err != nil {
		panic(err)
	}
	defer f2.Close()
	if verbose {
		fmt.Printf("Copying %s to %s\n", f1.Name(), f2.Name())
	}
	if _, err := io.Copy(f2, f1); err != nil {
		panic(err)
	}
}

func main() {
	var (
		st_paths []string
		files    = []string{
			"../3rdparty/jsonrpc.py",
			"../editor/sublime/plugin.py",
		}
	)

	if u, err := user.Current(); err != nil {
		log.Fatal(err)
	} else {
		switch runtime.GOOS {
		case "darwin":
			st_paths = append(st_paths, filepath.Join(u.HomeDir, "Library", "Application Support", "Sublime Text 2", "Packages"))
			st_paths = append(st_paths, filepath.Join(u.HomeDir, "Library", "Application Support", "Sublime Text 3", "Packages"))
		case "linux":
			st_paths = append(st_paths, filepath.Join(u.HomeDir, ".config", "sublime-text-2", "Packages"))
			st_paths = append(st_paths, filepath.Join(u.HomeDir, ".config", "sublime-text-3", "Packages"))
		default:
			log.Fatalf("Don't know where to install Sublime Text files on OS %s", runtime.GOOS)
		}
	}
	for i := range st_paths {
		if fi, err := os.Stat(st_paths[i]); err != nil || !fi.IsDir() {
			continue
		} else {
			for _, f := range files {
				copy(f, filepath.Join(st_paths[i], filepath.Base(f)))
			}
		}
	}
}
