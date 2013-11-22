package java

import (
	"errors"
	"fmt"
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
	// TODO(d) this isn't compatible with 1.6 jvm, openjdk. `javac -verbose EmptyFile.java` provides whats needed
	// but 100x slower and more apt to change than the no-guarantee -X settings.
	out, err := exec.Command("java", "-XshowSettings:properties", "-version").CombinedOutput()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Couldn't execute java: %s", err))
	}
	sub := re_classpath.FindAllStringSubmatch(string(out), -1)
	ret = make(map[string]string)
	for i := range sub {
		ret[strings.TrimSpace(sub[i][1])] = strings.TrimSpace(sub[i][2])
	}
	return ret, nil
}

func DefaultClasspath() ([]string, error) {
	props, err := properties()
	if err != nil {
		return nil, err
	}
	if v, ok := props["sun.boot.class.path"]; ok {
		matches := re_pathsplit.FindAllStringSubmatch(v, -1)
		ret := make([]string, len(matches))
		for i := range matches {
			ret[i] = matches[i][1]
		}
		return ret, nil
	}
	var errs = "Couldn't get the default classpath. Available properties:\n"
	for k, v := range props {
		errs += fmt.Sprintf("%s = %s\n", k, v)
	}
	return nil, errors.New(errs)
}
