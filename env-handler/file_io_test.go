package main

import (
  "testing"
  "reflect"
)

func TestReadEnvFile(t *testing.T) {
  actual := ReadEnvFile("./testdata/envfile")
  expected := []string{"# test", `TEST_1='test_1 val'`, `TEST_2='test2 val'`}

  if !reflect.DeepEqual(actual, expected) {
    t.Errorf("not equal slice")
  }
}

func TestWriteEnvFile(t *testing.T) {
  texts := ReadEnvFile("./testdata/envfile")
  WriteEnvFile("./testdata/write_test_envfile", texts)

  actual := ReadEnvFile("./testdata/envfile")
  expected := ReadEnvFile("./testdata/write_test_envfile")

  if !reflect.DeepEqual(actual, expected) {
    t.Errorf("not equal slice")
  }
}
