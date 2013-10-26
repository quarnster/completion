package main

import (
	"encoding/json"
	"flag"
	"github.com/quarnster/completion/content"
	"github.com/robmerrell/comandante"
	"io/ioutil"
	"os"
)

var completeAtFlags = struct {
	content.SourceLocation
	contents bool
}{}

func init() {
	completeAt := comandante.NewCommand("completeat", "Run a completeat operation", runCompleteAt)
	completeAt.FlagInit = completeAtFlagInit
	bin.RegisterCommand(completeAt)
}

func completeAtFlagInit(fs *flag.FlagSet) {
	fs.StringVar(&completeAtFlags.File.Name, "file", completeAtFlags.File.Name, "The file to use")
	fs.UintVar(&completeAtFlags.Line, "line", completeAtFlags.Line, "The line to use")
	fs.UintVar(&completeAtFlags.Column, "column", completeAtFlags.Column, "The column to use")
	fs.BoolVar(&completeAtFlags.contents, "contents", completeAtFlags.contents, "If true, the contents of the file should be written to this process's stdin")
}

func runCompleteAt() error {
	if completeAtFlags.contents {
		if data, err := ioutil.ReadAll(os.Stdin); err != nil {
			return err
		} else {
			completeAtFlags.File.Contents = string(data)
		}
	}
	var (
		req = content.CompleteAtArgs{Location: completeAtFlags.SourceLocation}
		ifs content.CompleteAt
		res content.CompletionResult
	)
	// TODO: "driver" lookup, which will overlap(?) with the daemon code
	if err := ifs.CompleteAt(&req, &res); err != nil {
		return err
	}
	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(res); err != nil {
		return err
	}

	return nil
}
