// Package project_version 这里给 cmd 添加 -v, --version 的 flag, 用于打印 version 信息；
package project_version

import (
	"fmt"
	"os"
	"strconv"

	flag "github.com/spf13/pflag"
)

type versionValue int

// Define some const.
const (
	VersionFalse versionValue = 0
	VersionTrue  versionValue = 1
	VersionRaw   versionValue = 2
)

const strRawVersion string = "raw"

func (v *versionValue) IsBoolFlag() bool {
	return true
}

func (v *versionValue) Get() interface{} {
	return v
}

func (v *versionValue) Set(s string) error {
	if s == strRawVersion {
		*v = VersionRaw
		return nil
	}
	boolVal, err := strconv.ParseBool(s)
	if boolVal {
		*v = VersionTrue
	} else {
		*v = VersionFalse
	}
	return err
}

func (v *versionValue) String() string {
	if *v == VersionRaw {
		return strRawVersion
	}
	return fmt.Sprintf("%v", bool(*v == VersionTrue))
}

// Type The type of the flag as required by the pflag.Value interface.
func (v *versionValue) Type() string {
	return "version"
}

// VersionVar defines a flag with the specified name and usage string.
func VersionVar(p *versionValue, name string, shorthand string, value versionValue, usage string) {
	*p = value
	flag.VarP(p, name, shorthand, usage)
	// "--version" will be treated as "--version=true"
	flag.Lookup(name).NoOptDefVal = "true"
}

// Version wraps the VersionVar function.
func Version(name string, shorthand string, value versionValue, usage string) *versionValue {
	p := new(versionValue)
	VersionVar(p, name, shorthand, value, usage)
	return p
}

const versionFlagName = "version"
const versionShortName = "v"

var versionFlag = Version(versionFlagName, versionShortName, VersionFalse, "Print version information and quit.")

// AddVersionFlags registers this package's flags on arbitrary FlagSets, such that they point to the
// same value as the global flags.
func AddVersionFlags(fs *flag.FlagSet) {
	fs.AddFlag(flag.Lookup(versionFlagName))
}

// PrintAndExitIfRequested will check if the -version flag was passed
// and, if so, print the version and exit.
func PrintAndExitIfRequested() {
	if *versionFlag == VersionRaw {
		fmt.Printf("%#v\n", Get())
		os.Exit(0)
	} else if *versionFlag == VersionTrue {
		fmt.Printf("%s\n", Get())
		os.Exit(0)
	}
}
