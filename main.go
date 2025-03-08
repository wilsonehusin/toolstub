package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path"
)

var (
	optToolImport string
	optExe        string
	optModDir     string
	optOutDir     string
	optPrint      bool
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s: create shorthand executable for build-time dependencies / third-party Go tools\n", "toolstub")
		fmt.Fprintf(flag.CommandLine.Output(), "Usage:\n")
		flag.PrintDefaults()
	}
}

func main() {
	flag.StringVar(&optToolImport, "tool", "", "[required] go tool path, e.g. github.com/golangci/golangci-lint/cmd/golangci-lint")
	flag.StringVar(&optExe, "exe", "", "[optional] executable name (defaults to tool's basename)")
	flag.StringVar(&optModDir, "moddir", "_tools", "[optional] directory to store go.mod files (defaults to _tools/)")
	flag.StringVar(&optOutDir, "outdir", "bin", "[optional] directory to store generated executables (defaults to bin/)")
	flag.BoolVar(&optPrint, "print", false, "[optional] print the generated stub and do not create file")
	flag.Parse()

	if optToolImport == "" {
		fmt.Fprintf(os.Stderr, "error: flag -tool is required but not provided\n")
		flag.Usage()
		os.Exit(1)
	}

	if optExe == "" {
		optExe = path.Base(optToolImport)
	}

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err.Error())
		os.Exit(1)
	}
}

func run() error {
	var w io.Writer

	if optPrint {
		w = os.Stdout
	} else {

		if err := os.MkdirAll(optOutDir, 0755); err != nil {
			return fmt.Errorf("ensuring directory %q exists: %w", optOutDir, err)
		}

		fp := path.Join(optOutDir, optExe)

		f, err := os.OpenFile(fp, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil {
			return fmt.Errorf("opening file %q to write: %w", fp, err)
		}
		defer f.Close()

		w = f
	}

	g := &Generator{
		ToolImport: optToolImport,
		Exe:        optExe,
		ModDir:     optModDir,
	}
	if _, err := g.WriteTo(w); err != nil {
		return fmt.Errorf("writing to file: %w", err)
	}
	return nil
}
