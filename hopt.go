// A docopt(854c423c810880e30b9fecdabb12d54f4a92f9bb) wrapper that adds:
//   * type checking
//   * 64 exit code on usge error.
package hopt

import "os"
import "fmt"
import "errors"
import "regexp"
import "strings"
import "github.com/docopt/docopt-go"

// Aliases

var compile = regexp.MustCompile
var puts, sprintf = fmt.Println, fmt.Sprintf

// Globals

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
  if compile(`\s%s\s`).MatchString(Help) {
    name := os.Args[0]
    path := strings.Split(name, "/")
    name = path[len(path)-1]
    Help = sprintf(Help, name)
  }
  type_map()
  First = false
}

func type_map() {
  kx := compile(`--\w+=\w+`)
  ks := kx.FindAllString(Help, -1)
  if ks != nil {
    for _, k := range ks {
      ab := strings.SplitN(k, "=", 2)
      TypeMap[ab[0]] = ab[1]
    }
  }
}

func type_check() error {
  is_float := compile(`^\d+\.\d+$`)
  is_int   := compile(`^\d+$`)
  is_date  := compile(`\d\d\d\d-\d\d-\d\d`)
  is_word  := compile(`^\w+$`)
  is_file  := compile(`^[^*&%\s]+$`)
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
      }
      if !ok {
        return errors.New(
          sprintf("Expected a %s for %s, but got \"%v\".", kind, key, value))
      }
    }
  }
  return nil
}
