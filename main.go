package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"log"
	"os"
	"text/template"
)

func main() {
	log.SetFlags(0)
	args := runArgs{}
	flag.StringVar(&args.t, "t", args.t, "path to the template, see: https://pkg.go.dev/text/template")
	flag.StringVar(&args.v, "v", args.v, "path to JSON mapping of variables to use in template, see:\nhttps://pkg.go.dev/encoding/json#Unmarshal")
	flag.Parse()
	if err := run(args); err != nil {
		log.Fatal(err)
	}
}

type runArgs struct {
	t string
	v string
}

func run(args runArgs) error {
	if args.t == "" || args.v == "" {
		return errors.New("need both template and variables")
	}
	tpl, err := template.ParseFiles(args.t)
	if err != nil {
		return err
	}
	tpl = tpl.Option("missingkey=error")
	data, err := os.ReadFile(args.v)
	if err != nil {
		return err
	}
	var vars any
	if err := json.Unmarshal(data, &vars); err != nil {
		return err
	}
	var buf bytes.Buffer
	if err := tpl.Execute(&buf, vars); err != nil {
		return err
	}
	_, err = os.Stdout.ReadFrom(&buf)
	return err
}
