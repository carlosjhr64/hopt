package hopt // import "github.com/carlosjhr64/hopt"

A docopt(854c423c810880e30b9fecdabb12d54f4a92f9bb) wrapper that adds:
    * type checking
    * 64 exit code on usge error.

var Version = "0.0.0"
var Help = `Usage:
  %s [options] [<arg>...]
Options:
  -h --help
  -v --version
  -q --quiet
  -V --verbose
  -T --trace`

var Argv []string = nil
var Options map[string]interface{} = nil

var DocOptHelp = true
var DocOptExit = false
var OptionsFirst = true

var Err error = nil
var Exit = true

var CsvX = `^\w+(,\w+)*$`
var DateX = `\d\d\d\d-\d\d-\d\d`
var FileX = `^[^*&%\s]+$`
var First = true
var FloatX = `^\d+\.\d+$`
var IntX = `^\d+$`
var WordX = `^\w+$`

var FormatX = `\s%s\s`

var TypeMapX = `--\w+=\w+`
var TypeMap = make(map[string]string)

func Parse() bool
func Tof(k string) float64
func Toi(k string) int
func Tos(k string) string
