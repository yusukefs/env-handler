package main

import (
  "fmt"
  "os"
  "log"
  "bufio"
)

func ReadEnvFile(filepath string) []string {
  texts := make([]string, 0)

  file, err := os.Open(filepath)
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    texts = append(texts, scanner.Text())
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  return texts
}

func WriteEnvFile(filepath string, texts []string) error {
  file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
  if err != nil {
    log.Fatal(err)
  }

  w := bufio.NewWriter(file)
  for _, line := range texts {
    fmt.Fprintln(w, line)
  }
  w.Flush()

  return nil
}

func IsLastLineEmpty(filepath string) bool {
  texts := ReadEnvFile(filepath)

  if texts[len(texts) - 1] == "" {
    return true
  } else {
    return false
  }
}

func AppendLineToFile(filepath string, line string) {
  file, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  var string_to_write string
  if IsLastLineEmpty(filepath) {
    string_to_write = line + "\n" + ""
  } else {
    string_to_write = "\n" + line + "\n" + "not empty"
  }
  if _, err = file.WriteString(string_to_write); err != nil {
    log.Fatal(err)
  }
}

