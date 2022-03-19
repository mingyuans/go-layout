// Generate cmd docs.
package main

import (
	"fmt"
	"github.com/mingyuans/go-layout/internal/apiserver"
	"github.com/spf13/cobra"
	"os"

	"github.com/spf13/cobra/doc"

	"github.com/marmotedu/iam/pkg/util/genutil"
)

type genFunc func() (string, *cobra.Command)

func main() {
	modules := []genFunc{
		genApiServerCmdDoc,
	}

	// use os.Args instead of "flags" because "flags" will mess up the man pages!
	path := ""
	if len(os.Args) == 2 {
		path = os.Args[1]
	} else {
		_, _ = fmt.Fprintf(os.Stderr, "usage: %s [output directory] \n", os.Args[0])
		os.Exit(1)
	}

	outDir, err := genutil.OutDir(path)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to get output directory: %v\n", err)
		os.Exit(1)
	}

	for _, genFunc := range modules {
		_, cmd := genFunc()
		_ = doc.GenMarkdownTree(cmd, outDir)
	}
}

func genApiServerCmdDoc() (string, *cobra.Command) {
	module := "apiserver"
	cmd := apiserver.NewApp(module).Command()
	return module, cmd
}
