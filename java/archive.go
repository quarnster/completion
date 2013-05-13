package java

import (
	"archive/zip"
	"code.google.com/p/log4go"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type (
	Archive interface {
		Classes() ([]Classname, error)
		LoadClass(class Classname) ([]byte, error)
		LoadAllClasses() [][]byte
		Close() error
	}

	JarArchive struct {
		path string
		zf   *zip.ReadCloser
	}

	DirectoryArchive struct {
		path string
		fi   []os.FileInfo
	}
	CompositeArchive struct {
		archives []Archive
	}
)

func (j *JarArchive) Classes() (ret []Classname, err error) {
	for _, zf := range j.zf.File {
		if strings.HasSuffix(zf.Name, ".class") {
			ret = append(ret, Filename(zf.Name).Classname())
		}
	}
	return ret, err
}

func (j *JarArchive) LoadClass(class Classname) ([]byte, error) {
	fn := class.Filename()
	for _, zf := range j.zf.File {
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
	return nil, errors.New(fmt.Sprintf("Couldn't find class %s in class path %s", fn, j.path))
}

func (j *JarArchive) LoadAllClasses() (ret [][]byte) {
	for _, zf := range j.zf.File {
		if strings.HasSuffix(zf.Name, ".class") {
			if f, err := zf.Open(); err == nil {
				defer f.Close()
				if d, err := ioutil.ReadAll(f); err == nil {
					ret = append(ret, d)
				}
			}
		}
	}
	return ret
}

func (j *JarArchive) Close() error {
	return j.zf.Close()
}

func (d *DirectoryArchive) Classes() (ret []Classname, err error) {
	for _, f := range d.fi {
		if strings.HasSuffix(f.Name(), ".class") {
			ret = append(ret, Filename(f.Name()).Classname())
		}
	}
	return ret, err
}

func (d *DirectoryArchive) LoadClass(class Classname) ([]byte, error) {
	fn := class.Filename()
	path := d.path + string(os.PathSeparator) + string(fn)
	data, err := ioutil.ReadFile(path)
	return data, err
}

func (d *DirectoryArchive) LoadAllClasses() (ret [][]byte) {
	for _, f := range d.fi {
		if strings.HasSuffix(f.Name(), ".class") {
			if d, err := ioutil.ReadFile(f.Name()); err == nil {
				ret = append(ret, d)
			}
		}
	}
	return ret
}

func (d *DirectoryArchive) Close() error {
	return nil
}

func (c *CompositeArchive) Classes() (ret []Classname, err error) {
	var errstr string
	for i := range c.archives {
		paths, err := c.archives[i].Classes()
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

func (c *CompositeArchive) LoadAllClasses() (ret [][]byte) {
	for i := range c.archives {
		ret = append(ret, c.archives[i].LoadAllClasses()...)
	}
	return ret
}

func (c *CompositeArchive) LoadClass(class Classname) ([]byte, error) {
	for i := range c.archives {
		if data, err := c.archives[i].LoadClass(class); err == nil {
			return data, err
		}
	}
	return nil, errors.New(fmt.Sprintf("Couldn't find class: %s", class))
}

func (c *CompositeArchive) Close() error {
	errorStr := ""
	for i := range c.archives {
		if err := c.archives[i].Close(); err != nil {
			if len(errorStr) > 0 {
				errorStr += "\n"
			}
			errorStr += err.Error()
		}
	}
	if len(errorStr) != 0 {
		return errors.New(errorStr)
	}
	return nil
}

func NewArchive(path string) (Archive, error) {
	if strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".jar") {
		if z, err := zip.OpenReader(path); err != nil {
			return nil, err
		} else {
			return &JarArchive{path, z}, nil
		}
	} else {
		if f, err := os.Open(path); err != nil {
			return nil, err
		} else {
			defer f.Close()
			fi, _ := f.Readdir(-1)
			return &DirectoryArchive{path, fi}, nil
		}
	}
}

func NewCompositeArchive(paths []string) (ret *CompositeArchive, err error) {
	var (
		errorStr = ""
		archives []Archive
	)

	for i := range paths {
		if a, err := NewArchive(paths[i]); err != nil {
			log4go.Warn(err)
			if len(errorStr) > 0 {
				errorStr += "\n"
			}
			errorStr += err.Error()
		} else {
			archives = append(archives, a)
		}
	}

	if len(archives) > 0 || len(paths) == 0 {
		ret = &CompositeArchive{archives}
	} else if len(errorStr) != 0 {
		err = errors.New(errorStr)
	}
	return ret, err
}
