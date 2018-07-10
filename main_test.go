package main

import (
  "testing"
  "reflect"
)

func TestGenerateSampleEnvfile(t *testing.T) {
  filepath := "./testdata/envfile"
  GenerateSampleEnvfile(filepath)

  actual := ReadEnvFile(filepath + ".sample")
  expected := ReadEnvFile(filepath + ".sample.expected")

  if !reflect.DeepEqual(actual, expected) {
    t.Errorf("generated sample file is not as expected")
  }
}

