package main

import "fmt"
import "github.com/carlosjhr64/hopt"

var puts = fmt.Println

func main(){
  hopt.Help = `Usage:
  simple [options]
Options:
  --x=FLOAT
  --int=INT
  --file=FILE
  --wd=WORD
  --start=DATE`
  hopt.Parse()
  puts(hopt.Options)
}
