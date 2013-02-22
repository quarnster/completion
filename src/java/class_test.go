package java

import (
	"archive/zip"
	//"os"
	"strings"
	"testing"
)

const rt_jar = "/Library/Java/JavaVirtualMachines/jdk1.7.0_09.jdk/Contents/Home/jre/lib/rt.jar"

func TestLoadAllClasses(t *testing.T) {
	if z, err := zip.OpenReader(rt_jar); err != nil {
		t.Fatal(err)
	} else {
		defer z.Close()
		for _, zf := range z.File {
			if strings.HasSuffix(zf.Name, ".class") {
				t.Log(zf.Name)
				if f, err := zf.Open(); err != nil {
					t.Fatal(err)
				} else {
					defer f.Close()
					if _, err := NewClass(f); err != nil {
						t.Fatal(err)
					}
				}
			}
		}
	}
}
