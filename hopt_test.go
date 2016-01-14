package hopt

import "testing"

func TestParse(test *testing.T) {
  var bad = test.Error // Alias
  if Version      != "0.0.0" { bad("1. Version should be 0.0.0.") }
  if First        != true    { bad("2. First should be true.") }
  if OptionsFirst != true    { bad("3. OptionsFirst should be true.") }
  if DocOptHelp   != true    { bad("4. DocOptHelp should be true.") }
  if DocOptExit   != false   { bad("5. DockOptExit should be false.") }
  if Exit         != true    { bad("6. Exit should be true.") }
  if Argv         != nil     { bad("7. Argv should be nil.") }
  if Err          != nil     { bad("8. Err should be nil.") }

  if Options      != nil     { bad("9. Options should be nil.") }
  if Parse()      != true    { bad("10.Parse() should not Err, got", Err) }
  if Options      == nil     { bad("11. Options should now not be nil.") }

  if Err          != nil     { bad("12. Err should still be nil.") }
  if First        != false   { bad("13.First should now be false.") }

  First, Help = true, " %s "
  initialize()
  if Help != " hopt.test " {
    bad("14. Expected hopt.test in Help, got", Help)
  }

  First, Help = true, " --abc=xyz "
  initialize()
  if Help != " --abc=xyz " {
    bad("15. Help should not have changed, got", Help)
  }
  if TypeMap["--abc"] != "xyz" {
    bad("16. TypeMap did not work.")
  }
  TypeMap["--float"] = "FLOAT"
  Options["--float"] = "Not A Float"
  err := type_check()
  if err == nil {
    bad("17. Expected a type error.")
  }
  Options["--float"] = "3.14"
  err = type_check()
  if err != nil {
    bad("18. Did not expect a type error.")
  }

  Destroy()
}
