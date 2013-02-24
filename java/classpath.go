package java

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
	"testing"
)

var (
	re_classpath = regexp.MustCompile(` {4}([^=]+)=((?:[^\n]+\n)(?: {8}[^\n]+\n)*)`)
	re_pathsplit = regexp.MustCompile(`\s*\S+\s*`)
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
		matches := re_pathsplit.FindAllString(v, -1)
		return matches, nil
	} else {
		err := "Couldn't get the default classpath. Available properties:\n"
		for k, v := range props {
			err += fmt.Sprintf("%s = %s\n", k, v)
		}
		return nil, errors.New(err)
	}
	panic("unreachable")
}
