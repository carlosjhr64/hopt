package main

import "fmt"
import "github.com/carlosjhr64/hopt"

var puts, printf = fmt.Println, fmt.Printf

func main(){
  hopt.Help = `Usage:
  simple [options]
Options:
  -v --verbose
  --x=FLOAT
  --y=number    Try putting not a number here.
  --int=INT
  --file=FILE
  --wd=WORD
  --start=DATE`
  hopt.Parse()
  puts(hopt.Options)
  printf("verbose: %v\n", hopt.Tob("--verbose"))
  printf("x:       %f\n", hopt.Tof("--x"))
  printf("y:       %f\n", hopt.Tof("--y"))
  printf("int:     %d\n", hopt.Toi("--int"))
  printf("file:    %s\n", hopt.Tos("--file"))
  printf("wd:      %s\n", hopt.Tos("--wd"))
  printf("start:   %s\n", hopt.Tos("--start"))
  hopt.Destroy()
  puts("Done!")
}
