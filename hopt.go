// A docopt(854c423c810880e30b9fecdabb12d54f4a92f9bb) wrapper that adds:
//   * type checking
//   * 64 exit code on usge error.
package hopt

import "os"
import "fmt"
import "errors"
import "regexp"
import "strings"
import "strconv"
import "github.com/docopt/docopt-go"

// Aliases

var compile = regexp.MustCompile
var puts, sprintf = fmt.Println, fmt.Sprintf

// Globals

const VERSION = "0.0.1"

var Version = "0.0.0"
var Help = `Usage:
  %s [options] [<arg>...]
Options:
  -h --help
  -v --version
  -q --quiet
  -V --verbose
  -T --trace`
var First = true
var OptionsFirst = true
var DocOptHelp = true
var DocOptExit = false
var Exit = true
var Argv []string = nil
var Err error = nil
var TypeMap = make(map[string]string)

var Options map[string]interface{} = nil

func Parse() bool {
  if First { initialize() }
  Options, Err = docopt.Parse(Help, Argv, DocOptHelp, Version, OptionsFirst, DocOptExit)
  if Err == nil {
    Err = type_check()
    if Err != nil { puts(Err) }
  }
  if Exit {
    if Err != nil { os.Exit(64) }
    for _, opt := range(os.Args[1:]) {
      if opt == "--version" || opt == "-h" || opt == "--help" {
        os.Exit(0)
      }
    }
  }
  return Err == nil
}

func initialize() {
  compiles()  // Must be done first
  format()
  type_map()
  First = false
}

func format() {
  if is_format.MatchString(Help) {
    name := os.Args[0]
    path := strings.Split(name, "/")
    name = path[len(path)-1]
    Help = sprintf(Help, name)
  }
}


func type_map() {
  ks := is_typemap.FindAllString(Help, -1)
  if ks != nil {
    for _, k := range ks {
      ab := strings.SplitN(k, "=", 2)
      TypeMap[ab[0]] = ab[1]
    }
  }
}

var FloatX = `^\d+\.\d+$`
var IntX   = `^\d+$`
var DateX  = `\d\d\d\d-\d\d-\d\d`
var WordX  = `^\w+$`
var FileX  = `^[^*&%\s]+$`
var CsvX   = `^\w+(,\w+)*$`

var is_float *regexp.Regexp
var is_int   *regexp.Regexp
var is_date  *regexp.Regexp
var is_word  *regexp.Regexp
var is_file  *regexp.Regexp
var is_csv   *regexp.Regexp

var FormatX  = `\s%s\s`
var TypeMapX = `--\w+=\w+`

var is_format  *regexp.Regexp
var is_typemap *regexp.Regexp

func compiles() {
  is_float = compile(FloatX)
  is_int   = compile(IntX)
  is_date  = compile(DateX)
  is_word  = compile(WordX)
  is_file  = compile(FileX)
  is_csv   = compile(CsvX)

  is_format  = compile(FormatX)
  is_typemap = compile(TypeMapX)
}

func type_check() error {
  var ok bool
  for key, kind := range TypeMap {
    value := Options[key]
    if value != nil {
      ok = true
      switch kind {
      case "FLOAT":
        ok = is_float.MatchString(value.(string))
      case "INT":
        ok = is_int.MatchString(value.(string))
      case "DATE":
        ok = is_date.MatchString(value.(string))
      case "FILE":
        ok = is_file.MatchString(value.(string))
      case "WORD":
        ok = is_word.MatchString(value.(string))
      case "CSV":
        ok = is_csv.MatchString(value.(string))
      }
      if !ok {
        return errors.New(
          sprintf("Expected a %s for %s, but got \"%v\".", kind, key, value))
      }
    }
  }
  return nil
}

func oops(k string, t string) {
  fmt.Fprintf(
    os.Stderr,
    "Could not parse %s as a %s.",
    k, t)
  os.Exit(65)
}

func Tos(k string) string {
  a := Options[k]
  if a == nil { return "" }
  return a.(string)
}

func Tob(k string) bool {
  a := Options[k]
  if a == nil { return false }
  return a.(bool)
}

func Toi(k string) int {
  a := Options[k]
  if a == nil { return 0 }
  f, e := strconv.Atoi(a.(string))
  if e != nil { oops(k, "Int") }
  return f
}

func Tof(k string) float64 {
  a := Options[k]
  if a == nil { return 0.0 }
  f, e := strconv.ParseFloat(a.(string), 64)
  if e != nil { oops(k, "Float") }
  return f
}

func Destroy() {
  // Not sure if Destroy is usefull,
  // but after hopt gets used
  // why keep these around?  :-??
  compile      = nil
  puts         = nil
  sprintf      = nil
  Version      = ""
  Help         = "" // Specially this one as it can be quite long.
//First        = nil
//OptionsFirst = nil
//DocOptHelp   = nil
//DocOptExit   = nil
//Exit         = nil
  Argv         = nil
  Err          = nil
  TypeMap      = nil
  Options      = nil
  FloatX       = ""
  IntX         = ""
  DateX        = ""
  WordX        = ""
  FileX        = ""
  CsvX         = ""
  is_float     = nil
  is_int       = nil
  is_date      = nil
  is_word      = nil
  is_file      = nil
  is_csv       = nil
  FormatX      = ""
  TypeMapX     = ""
  is_format    = nil
  is_typemap   = nil
}
