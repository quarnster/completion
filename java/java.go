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
		if archive, err = NewCompositeArchive(cp); err != nil {
			return err
		} else if session != nil {
			session.Set("java_archive", archive)
		}
	}

	className, err := fqnToClassname(args.Location)
	if err != nil {
		return err
	}

	data, err := archive.LoadClass(className)
	if err != nil {
		return err
	}

	class, err := NewClass(bytes.NewReader(data))
	if err != nil {
		return err
	}

	ct, err := class.ToContentType()
	if err != nil {
		return err
	}

	// TODO(q): Inherited fields and methods?
	// 			I see value in being able to "just" get the smaller set,
	//			but getting the full set should definitely be possible "server side"
	cmp.Fields = ct.Fields
	cmp.Types = ct.Types
	cmp.Methods = ct.Methods

	return nil
}

// TODO(d) provide generic completions based on context. New lines should complete
// types, methods, and members, inherited or otherwise. Argument position should show
// results only for the type accepted, etc.
func (c *Java) CompleteAt(a *content.CompleteAtArgs, ret *content.CompletionResult) error {
	return nil
}
