package java

import (
	"archive/zip"
	"errors"
	"fmt"
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

func ClassesInPath(path string) (ret []string, err error) {
	if strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".jar") {
		if z, err := zip.OpenReader(path); err != nil {
			return nil, err
		} else {
			defer z.Close()
			for _, zf := range z.File {
				if strings.HasSuffix(zf.Name, ".class") {
					ret = append(ret, FileToClass(zf.Name))
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
						ret = append(ret, FileToClass(f.Name()))
					}
				}
			}
		}
	}
	return ret, err
}

// Returns the name of all classes found in the specified class path.
// The returned string array may be populated even if err != nil.
func Classes(classpath []string) (ret []string, err error) {
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
