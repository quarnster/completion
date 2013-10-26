package main

import (
	"code.google.com/p/log4go"
	"flag"
	"github.com/quarnster/completion/editor"
	_ "github.com/quarnster/completion/editor/sublime"
	"github.com/robmerrell/comandante"
)

var (
	instArgs []*bool
)

func init() {
	install := comandante.NewCommand("install", "Install editor hooks", installHooks)
	install.FlagInit = installFlagInit
	bin.RegisterCommand(install)
}

func installFlagInit(fs *flag.FlagSet) {
	for _, ed := range editor.List() {
		instArgs = append(instArgs, fs.Bool(ed.Name(), false, ed.Description()))
	}
}

func installHooks() error {
	for i, ed := range editor.List() {
		if *instArgs[i] {
			if err := ed.Install(); err != nil {
				log4go.Error("Failed to install editor plugin for %s: %s", ed.Description(), err)
			}
		}
	}
	return nil
}

// fs.StringVar(&install, "install", install, fmt.Sprintf("A comma separated list of editors for which to install the plugin. Valid editors are:%s", func() string {
// 	b := bytes.NewBuffer(nil)
// 	for _, p := range editor.List() {
// 		b.WriteString(fmt.Sprintf("\n\t%8s - %s", p.Name(), p.Description()))
// 	}
// 	return b.String()
// }()))
