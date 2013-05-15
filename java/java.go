package java

import (
	"bytes"
	"code.google.com/p/log4go"
	"errors"
	"github.com/quarnster/completion/content"
	"strings"
)

type Java struct {
}

func fqnToClassname(fqn content.FullyQualifiedName) (Classname, error) {
	// TODO(d) should qualify abs or rel
	if len(fqn.Absolute) == 0 {
		return "", errors.New("Received empty fqn")
	}
	if !strings.HasPrefix(fqn.Absolute, "java://type/") {
		return "", errors.New("FQN prefixed incorrectly with " + fqn.Absolute)
	}
	return Classname(strings.Replace(fqn.Absolute[len("java://type/"):], "/", ".", -1)), nil
}

func (c *Java) Complete(args *content.CompleteArgs, cmp *content.CompletionResult) error {
	var archive Archive
	session := args.Session()
	if session != nil {
		if a, ok := session.Get("java_archive").(Archive); ok {
			archive = a
		}
	}
	if archive == nil {
		cp, err := DefaultClasspath()
		if err != nil {
			// We don't really care about not being able to get the default classpath as it could be provided manually by the user
			log4go.Warn("Couldn't get the default classpath: %s", err)
		}
		settings := args.Settings()
		if cp2, ok := settings.Get("classpath").([]string); ok {
			// TODO: do we ever want to override rather than append to the classpath?
			cp = append(cp, cp2...)
		}
		// TODO: this should probably be cached in the session...
		if archive, err = NewCompositeArchive(cp); err != nil {
			return err
		} else if session != nil {
			session.Set("java_archive", archive)
		}
	}
	if className, err := fqnToClassname(args.Location); err != nil {
		return err
	} else if data, err := archive.LoadClass(className); err != nil {
		return err
	} else if class, err := NewClass(bytes.NewReader(data)); err != nil {
		return err
	} else if ct, err := class.ToContentType(); err != nil {
		return err
	} else {
		// TODO: inherited fields and methods
		cmp.Fields = ct.Fields
		cmp.Types = ct.Types
		cmp.Methods = ct.Methods
	}
	return nil
}
