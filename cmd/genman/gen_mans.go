// This tool helps to generate Linux man docs.
// About Linux man command, please see: https://blog.csdn.net/shuizhizhiyin/article/details/51668962
package main

import (
	"bytes"
	"fmt"
	iam_apiserver "github.com/mingyuans/go-layout/internal/iam-apiserver"
	"io"
	"os"
	"strings"

	mangen "github.com/cpuguy83/go-md2man/v2/md2man"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/marmotedu/iam/pkg/util/genutil"
)

type genFunc func() (string, *cobra.Command)

func main() {
	modules := []genFunc{
		genApiServerCmdDoc,
	}

	// use os.Args instead of "flags" because "flags" will mess up the man pages!
	path := "docs/man/man1"
	if len(os.Args) == 2 {
		path = os.Args[1]
	} else {
		_, _ = fmt.Fprintf(os.Stderr, "usage: %s [output directory] [module] \n", os.Args[0])
		os.Exit(1)
	}

	outDir, err := genutil.OutDir(path)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to get output directory: %v\n", err)
		os.Exit(1)
	}

	// Set environment variables used by command so the output is consistent,
	// regardless of where we run.
	_ = os.Setenv("HOME", "/home/username")

	for _, genFunc := range modules {
		module, cmd := genFunc()
		genMarkdown(cmd, "", outDir)
		for _, c := range cmd.Commands() {
			genMarkdown(c, module, outDir)
		}
	}
}

func genApiServerCmdDoc() (string, *cobra.Command) {
	module := "iam-apiserver"
	cmd := iam_apiserver.NewApp(module).Command()
	return module, cmd
}

func preamble(out *bytes.Buffer, name, short, long string) {
	out.WriteString(`% IAM(1) iam User Manuals
% Eric Paris
% Jan 2015
# NAME
`)
	_, _ = fmt.Fprintf(out, "%s \\- %s\n\n", name, short)
	_, _ = fmt.Fprintf(out, "# SYNOPSIS\n")
	_, _ = fmt.Fprintf(out, "**%s** [OPTIONS]\n\n", name)
	_, _ = fmt.Fprintf(out, "# DESCRIPTION\n")
	_, _ = fmt.Fprintf(out, "%s\n\n", long)
}

func printFlags(out io.Writer, flags *pflag.FlagSet) {
	flags.VisitAll(func(flag *pflag.Flag) {
		format := "**--%s**=%s\n\t%s\n\n"
		if flag.Value.Type() == "string" {
			// put quotes on the value
			format = "**--%s**=%q\n\t%s\n\n"
		}

		// Todo, when we mark a shorthand is deprecated, but specify an empty message.
		// The flag.ShorthandDeprecated is empty as the shorthand is deprecated.
		// Using len(flag.ShorthandDeprecated) > 0 can't handle this, others are ok.
		if !(len(flag.ShorthandDeprecated) > 0) && len(flag.Shorthand) > 0 {
			format = "**-%s**, " + format
			_, _ = fmt.Fprintf(out, format, flag.Shorthand, flag.Name, flag.DefValue, flag.Usage)
		} else {
			_, _ = fmt.Fprintf(out, format, flag.Name, flag.DefValue, flag.Usage)
		}
	})
}

func printOptions(out io.Writer, command *cobra.Command) {
	flags := command.NonInheritedFlags()
	if flags.HasFlags() {
		_, _ = fmt.Fprintf(out, "# OPTIONS\n")
		printFlags(out, flags)
		_, _ = fmt.Fprintf(out, "\n")
	}
	flags = command.InheritedFlags()
	if flags.HasFlags() {
		fmt.Fprintf(out, "# OPTIONS INHERITED FROM PARENT COMMANDS\n")
		printFlags(out, flags)
		fmt.Fprintf(out, "\n")
	}
}

func genMarkdown(command *cobra.Command, parent, docsDir string) {
	dparent := strings.ReplaceAll(parent, " ", "-")
	name := command.Name()

	dname := name
	if len(parent) > 0 {
		dname = dparent + "-" + name
		name = parent + " " + name
	}

	out := new(bytes.Buffer)

	short, long := command.Short, command.Long
	if len(long) == 0 {
		long = short
	}

	preamble(out, name, short, long)
	printOptions(out, command)

	if len(command.Example) > 0 {
		_, _ = fmt.Fprintf(out, "# EXAMPLE\n")
		_, _ = fmt.Fprintf(out, "```\n%s\n```\n", command.Example)
	}

	if len(command.Commands()) > 0 || len(parent) > 0 {
		_, _ = fmt.Fprintf(out, "# SEE ALSO\n")

		if len(parent) > 0 {
			_, _ = fmt.Fprintf(out, "**%s(1)**, ", dparent)
		}

		for _, c := range command.Commands() {
			_, _ = fmt.Fprintf(out, "**%s-%s(1)**, ", dname, c.Name())
			genMarkdown(c, name, docsDir)
		}

		_, _ = fmt.Fprintf(out, "\n")
	}

	out.WriteString(`
# HISTORY
January 2015, Originally compiled by Eric Paris (eparis at redhat dot com) based on the marmotedu source material, but hopefully they have been automatically generated since!
`)

	final := mangen.Render(out.Bytes())

	filename := docsDir + dname + ".1"

	outFile, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer outFile.Close()

	_, err = outFile.Write(final)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
