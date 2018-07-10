package main

import (
  "testing"
)

func TestGetKeyFromSingleLine(t *testing.T) {
  input := `TEST_1='test===tesrefe'`
  actual := ConvertSingleLineToSample(input)
  expected := `TEST_1=`

  if actual != expected {
    t.Errorf("Failed on removing value from single line. actual: %v, expected: %v", actual, expected)
  }

  input = `#TEST_1='test===tesrefe'`
  actual = ConvertSingleLineToSample(input)
  expected = `#TEST_1='test===tesrefe'`

  if actual != expected {
    t.Errorf("Failed on not removing when line starts from #")
  }
}
