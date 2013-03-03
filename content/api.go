package content

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"io"
	"reflect"
	"runtime"
	"strings"
)

func NewAPI(i interface{}) error {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)

	// TODO return error if not a func

	fn := func(w io.Writer, intent Intent) {

		if len(intent.Data) != t.NumIn() {
			panic("Array length of intent data does not match function signature.")
		}

		var in []reflect.Value

		// TODO order of intent.Data shouldn't matter since values are named
		// and can be matched against func sig args.
		for i, data := range intent.Data {
			val := reflect.New(t.In(i))

			// TODO code block here is questionable, addressing the fact that
			// binary.Read only works with fixed-size values.
			switch t.In(i).Kind() {
			case reflect.String:
				val.Elem().SetString(string(data.Value))
			default:
				buf := bytes.NewBuffer(data.Value)
				// TODO always binary.LittleEndian ??
				err := binary.Read(buf, binary.LittleEndian, val)
				if err != nil {
					panic(fmt.Sprintf("Failed to read value: %s", err))
				}
			}

			// TODO hack just for tests, fix this
			if val.Kind() == reflect.Ptr && t.In(i).Kind() != reflect.Ptr {
				val = val.Elem()
			}

			//
			in = append(in, val)
		}

		out := v.Call(in)

		var data []Data
		for i, val := range out {
			buf := new(bytes.Buffer)

			// TODO see comments on `switch t.Out(i).Kind()` code-block above.
			switch t.Out(i).Kind() {
			case reflect.String:
				buf.WriteString(val.String())
			default:
				binary.Write(buf, binary.LittleEndian, val)
			}

			// TODO api funcs should likely require named value returns for
			// constructing a proper type Data. Not actually sure if this is
			// possible or how to get it. The below returns a string representing
			// the type.
			name := t.Out(i).String()
			data = append(data, Data{name, buf.Bytes()})
		}

		resp := Response{1, data}

		gob.NewEncoder(w).Encode(resp)

		// not working
		// binary.Write(w, binary.LittleEndian, &resp)
		/*
			respBuf := new(bytes.Buffer)
			if err := binary.Write(respBuf, binary.LittleEndian, &resp); err != nil {
				panic(err)
			}
		*/
	}

	// TODO this isn't enough to properly scope a method. For example, java/parse.lex
	// would result (from the below) as registering as parse.lex
	//
	// It may be safe to assume that since the package exists under `github.com/quarstar/completion`
	// then we can strip this from the beginning and further format however desired.
	seg := strings.Split(runtime.FuncForPC(v.Pointer()).Name(), "/")
	name := seg[len(seg)-1]
	AddHandler(name, fn)
	return nil
}
