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
