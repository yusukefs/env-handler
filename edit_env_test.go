package main

import (
  "testing"
)

func TestGetKeyFromSingleLine(t *testing.T) {
  input := `TEST_1='test===tesrefe'`
  actual := ConvertSingleLineTosample(input)
  expected := `TEST_1=`

  if actual != expected {
    t.Errorf("Failed on removing value from single line. actual: %v, expected: %v", actual, expected)
  }
}
