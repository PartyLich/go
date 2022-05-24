package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"
)

type data struct {
	InType  string
	OutType string
	Name    string
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var d data

	flag.StringVar(&d.InType, "itype", "T", "The type alias for the iterators contained values.")
	flag.StringVar(&d.OutType, "otype", "T", "The type alias for the iterators output value, if it differs from the contained value")
	flag.StringVar(&d.Name, "name", "", "The name used for the adapter type being generated. This should start with a capital letter so that it is exported.")
	fname := flag.String("output", fmt.Sprintf("%v_ext_gen.go", strings.ToLower(d.Name)), "The output file name")
	flag.Parse()

	if d.InType != d.OutType {
		d.InType = fmt.Sprintf("%v, %v", d.InType, d.OutType)
	}

	const tmplName = "adapter_ext.tmpl"

	_, callFile, _, _ := runtime.Caller(0)
	path := filepath.Dir(callFile)
	t := template.Must(template.New(tmplName).ParseFiles(path + "/" + tmplName))

	file, err := os.Create(*fname)
	handleErr(err)
	defer file.Close()

	cmds := []*exec.Cmd{
		exec.Command("gofmt", "-s"),
	}

	inPipe, err := cmds[0].StdinPipe()
	handleErr(err)

	// hook up command pipeline
	for i := 1; i < len(cmds); i += 2 {
		cmds[i].Stdin, err = cmds[i-1].StdoutPipe()
		handleErr(err)
	}

	// cmds[len(cmds)-1].Stdout = os.Stdout
	cmds[len(cmds)-1].Stdout = file

	// start command pipeline
	for _, cmd := range cmds {
		handleErr(cmd.Start())
	}

	// execute template
	handleErr(t.Execute(inPipe, d))
	handleErr(inPipe.Close())

	// wait for command pipeline
	for _, cmd := range cmds {
		handleErr(cmd.Wait())
	}
}
