package main

import (
  "strings"
)

func ConvertSingleLineTosample(line string) string {
  line_split := strings.SplitN(line, "=", 2)
  if len(line_split) == 1 {
    return line
  } else {
    return line_split[0] + "="
  }
}

func BuildValueRemovedLines(original_lines []string) []string{
  lines := make([]string, 0)
  for _, line := range original_lines {
    lines = append(lines, ConvertSingleLineTosample(line))
  }

  return lines
}
