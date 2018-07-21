package main

import (
  "testing"
  "reflect"
)

func TestReadEnvFile(t *testing.T) {
  actual := ReadEnvFile("./testdata/envfile")
  expected := []string{"# test", `TEST_1='test_1 val'`, "", `TEST_2='test2 val'`, ""}

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

func TestIsLastLineEmpty(t *testing.T) {
  actual := IsLastLineEmpty("./testdata/envfile")
  expected := true

  if actual != expected {
    t.Errorf("Last line is empty but IsLastLineEmpty returned false")
  }
}

func TestAppendLineToFile(t *testing.T) {
  AppendLineToFile("./testdata/envfile-append", `APPEND='append val'`)

  actual := ReadEnvFile("./testdata/envfile-append")
  expected := ReadEnvFile("./testdata/envfile-append.expected")

  if !reflect.DeepEqual(actual, expected) {
    t.Errorf("failed at appending line to file")
  }
}

