package main

import (
	_ "embed"
	"fmt"
	"io"
	"runtime/debug"
	"text/template"
)

//go:embed stub.tmpl.bash
var stubTemplateText string
var stubTemplate = template.Must(template.New("toolstub").Parse(stubTemplateText))

type Generator struct {
	ToolImport string
	Exe        string
	ModDir     string
}

func (g *Generator) WriteTo(w io.Writer) (int64, error) {
	// Is there a point to conform with io.WriterTo interface?
	// Kinda silly since template.(*Template).Execute doesn't return the written size...
	if err := stubTemplate.Execute(w, g); err != nil {
		return 0, err
	}
	return 1, nil
}

func (t *Generator) ToolstubInfo() string {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		return "toolstub"
	}

	version := bi.Main.Version
	rev := ""

	for _, i := range bi.Settings {
		switch i.Key {
		case "vcs.revision":
			rev = i.Value
		}
	}

	if rev != "" {
		if version == "(devel)" {
			version = rev
		} else {
			version = fmt.Sprintf("%s-%s", version, rev)
		}
	}

	return fmt.Sprintf("%s@%s", bi.Path, version)
}
