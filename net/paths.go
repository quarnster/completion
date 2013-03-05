package net

import (
	"os/exec"
	"regexp"
)

func WindowsFrameworks() (paths []string, err error) {
	re, err := regexp.Compile(`InstallPath\s*REG_SZ\s*([^\n\r]+)`)
	if err != nil {
		return nil, err
	}
	c := exec.Command("reg", "query", "HKLM\\SOFTWARE\\Microsoft\\NET Framework Setup\\NDP", "/s", "/f", "InstallPath", "/e")
	if d, err := c.CombinedOutput(); err != nil {
		return nil, err
	} else {
		m := re.FindAllSubmatch(d, -1)
		for i := range m {
			paths = append(paths, string(m[i][1]))
		}
	}
	return
}

func MonoDefaultPath() string {
	c := exec.Command("mono", "hack.exe")
	c.Env = append(c.Env, "MONO_LOG_MASK=asm", "MONO_LOG_LEVEL=info")
	out, _ := c.Output()
	so := string(out)
	re := regexp.MustCompile(`image (.*)(mscorlib\.dll)`)
	strs := re.FindStringSubmatch(so)
	if len(strs) > 0 {
		return strs[1]
	}
	return ""
}

func DefaultPaths() (paths []string) {
	if mp := MonoDefaultPath(); mp != "" {
		paths = append(paths, mp)
	}
	if p2, err := WindowsFrameworks(); err == nil {
		paths = append(paths, p2...)
	}
	return
}
