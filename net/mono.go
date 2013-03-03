package net

import (
	"os/exec"
	"regexp"
)

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
