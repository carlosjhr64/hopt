package hopt // import "github.com/carlosjhr64/hopt"

A docopt(854c423c810880e30b9fecdabb12d54f4a92f9bb) wrapper that adds:

    * type checking
    * 64 exit code on usage error.

var Argv []string = nil
var DocOptExit = false
var DocOptHelp = true
var Err error = nil
var Exit = true
var First = true
var Help = `Usage:
  %s [options] [<arg>...]
Options:
  -h --help
  -v --version
  -q --quiet
  -V --verbose
  -T --trace`
var Options map[string]interface{} = nil
var OptionsFirst = true
var TypeMap = make(map[string]string)
var Version = "0.0.0"

func Parse() bool
