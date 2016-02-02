package hopt // import "github.com/carlosjhr64/hopt"

A docopt(854c423c810880e30b9fecdabb12d54f4a92f9bb) wrapper that adds:
    * type checking
    * 64 exit code on usage error.
    * 65 exit code on bad user input (but don't see how it can normally happen).

const VERSION = "0.1.0"

var Argv []string = nil
var CsvX = `^\w+(,\w+)*$`
var DateX = `\d\d\d\d-\d\d-\d\d`
var DocOptExit = false
var DocOptHelp = true
var Err error = nil
var Exit = true
var FileX = `^[^*&%\s]+$`
var First = true
var FloatX = `^\d+\.\d+$`
var FormatX = `\s%s\s`
var Help = `Usage:
  %s [options] [<arg>...]
Options:
  -h --help
  -v --version
  -q --quiet
  -V --verbose
  -T --trace`
var IntX = `^\d+$`
var Options map[string]interface{} = nil
var OptionsFirst = true
var TypeMap = make(map[string]string)
var TypeMapX = `--\w+=\w+`
var Version = "0.0.0"
var WordX = `^\w+$`

func Destroy()
func Parse() bool
func Tob(k string) bool
func Tof(k string) float64
func Toi(k string) int
func Tos(k string) string
