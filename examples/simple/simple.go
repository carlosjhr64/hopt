package main

import "fmt"
import "github.com/carlosjhr64/hopt"

var puts, printf = fmt.Println, fmt.Printf

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
  printf("x:     %f\n", hopt.Tof("--x"))
  printf("int:   %d\n", hopt.Toi("--int"))
  printf("file:  %s\n", hopt.Tos("--file"))
  printf("wd:    %s\n", hopt.Tos("--wd"))
  printf("start: %s\n", hopt.Tos("--start"))
}
