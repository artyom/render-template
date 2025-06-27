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
	var tPath, vPath string
	flag.StringVar(&tPath, "t", tPath, "path to the template, see: https://pkg.go.dev/text/template")
	flag.StringVar(&vPath, "v", vPath, "path to JSON mapping of variables to use in template, see:\nhttps://pkg.go.dev/encoding/json#Unmarshal")
	flag.Parse()
	if err := run(tPath, vPath); err != nil {
		log.Fatal(err)
	}
}

func run(tPath, vPath string) error {
	if tPath == "" || vPath == "" {
		return errors.New("need both template and variables")
	}
	tpl, err := template.ParseFiles(tPath)
	if err != nil {
		return err
	}
	tpl = tpl.Option("missingkey=error")
	data, err := os.ReadFile(vPath)
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
