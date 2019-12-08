/*
 * Hyperion
 * Étienne "Maiste" Marais
 */

package api

import (
    "io/ioutil"
    "os"
)

/* Return the file content */
func ImportJson(path string) ([]byte, error) {
  file, err := os.Open(path)
  if err != nil {
    return (make([]byte, 0)), err
  }
  defer file.Close()

  content, _:= ioutil.ReadAll(file)

  return content, nil
}


