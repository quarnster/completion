package java

import (
	"archive/zip"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var (
	re_classpath = regexp.MustCompile(` {4}([^=]+)=((?:[^\n]+\n)(?: {8}[^\n]+\n)*)`)
	re_pathsplit = regexp.MustCompile(`\s*(\S+)\s*`)
)

func properties() (ret map[string]string, err error) {
	// Adding -version just so that we get a better exit code
	if out, err := exec.Command("java", "-XshowSettings:properties", "-version").CombinedOutput(); err != nil {
		return nil, errors.New(fmt.Sprintf("Couldn't execute java: %s", err))
	} else {
		sub := re_classpath.FindAllStringSubmatch(string(out), -1)
		ret = make(map[string]string)
		for i := range sub {
			ret[strings.TrimSpace(sub[i][1])] = strings.TrimSpace(sub[i][2])
		}
	}
	return ret, nil
}

func DefaultClasspath() ([]string, error) {
	if props, err := properties(); err != nil {
		return nil, err
	} else if v, ok := props["sun.boot.class.path"]; ok {
		matches := re_pathsplit.FindAllStringSubmatch(v, -1)
		ret := make([]string, len(matches))
		for i := range matches {
			ret[i] = matches[i][1]
		}
		return ret, nil
	} else {
		err := "Couldn't get the default classpath. Available properties:\n"
		for k, v := range props {
			err += fmt.Sprintf("%s = %s\n", k, v)
		}
		return nil, errors.New(err)
	}
	panic("unreachable")
}

func ClassesInPath(path string) (ret []Classname, err error) {
	if strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".jar") {
		if z, err := zip.OpenReader(path); err != nil {
			return nil, err
		} else {
			defer z.Close()
			for _, zf := range z.File {
				if strings.HasSuffix(zf.Name, ".class") {
					ret = append(ret, Filename(zf.Name).Classname())
				}
			}
		}
	} else {
		if f, err := os.Open(path); err != nil {
			return nil, err
		} else {
			defer f.Close()
			if fi, err := f.Readdir(-1); err != nil {
				return nil, err
			} else {
				for _, f := range fi {
					if strings.HasSuffix(f.Name(), ".class") {
						ret = append(ret, Filename(f.Name()).Classname())
					}
				}
			}
		}
	}
	return ret, err
}

// Returns the name of all classes found in the specified class paths.
// The returned string array may be populated even if err != nil.
func Classes(classpath []string) (ret []Classname, err error) {
	var errstr string
	for i := range classpath {
		paths, err := ClassesInPath(classpath[i])
		if err != nil {
			errstr += err.Error() + "\n"
		}
		ret = append(ret, paths...)
	}
	if len(errstr) != 0 {
		err = errors.New(errstr)
	}
	return ret, err
}

func ClasspathMap(classpath []string) map[string][]Classname {
	ret := make(map[string][]Classname)
	for i := range classpath {
		if paths, _ := ClassesInPath(classpath[i]); len(paths) != 0 {
			ret[classpath[i]] = paths
		}
	}
	return ret
}

func LoadAllClassesInPath(path string) (ret [][]byte) {
	if strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".jar") {
		if z, err := zip.OpenReader(path); err != nil {
			return nil
		} else {
			defer z.Close()
			for _, zf := range z.File {
				if strings.HasSuffix(zf.Name, ".class") {
					if f, err := zf.Open(); err == nil {
						defer f.Close()
						if d, err := ioutil.ReadAll(f); err == nil {
							ret = append(ret, d)
						}
					}
				}
			}
		}
	} else {
		if f, err := os.Open(path); err != nil {
			return nil
		} else {
			defer f.Close()
			if fi, err := f.Readdir(-1); err != nil {
				return nil
			} else {
				for _, f := range fi {
					if strings.HasSuffix(f.Name(), ".class") {
						if d, err := ioutil.ReadFile(f.Name()); err == nil {
							ret = append(ret, d)
						}
					}
				}
			}
		}
	}
	return ret
}

func LoadClassEx(path string, class Classname) ([]byte, error) {
	fn := class.Filename()
	if strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".jar") {
		if z, err := zip.OpenReader(path); err != nil {
			return nil, err
		} else {
			defer z.Close()
			for _, zf := range z.File {
				if fn == Filename(zf.Name) {
					if f, err := zf.Open(); err != nil {
						return nil, errors.New(fmt.Sprintf("Couldn't open file: %s", err))
					} else {
						defer f.Close()
						if data, err := ioutil.ReadAll(f); err != nil {
							return nil, err
						} else {
							return data, nil
						}
					}
				}
			}
		}
		return nil, errors.New(fmt.Sprintf("Couldn't find class %s in class path %s", fn, path))
	} else {
		path = path + string(os.PathSeparator) + string(fn)
		data, err := ioutil.ReadFile(path)
		return data, err
	}
	panic("unreachable")
}

func LoadClass(classpath []string, class Classname) ([]byte, error) {
	for _, cp := range classpath {
		if data, err := LoadClassEx(cp, class); err == nil {
			return data, err
		}
	}
	return nil, errors.New(fmt.Sprintf("Couldn't find class: %s", class))
}
