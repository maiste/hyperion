/*
 * Hyperion
 * Étienne "Maiste" Marais
 */

package main

import (
  "os"
  "hyperion/api"
)

func main() {
  api.Display()
  if len(os.Args) == 2 {
    api.ConnectionHandler(os.Args[1])
  } else {
    api.ConnectionHandler("8080")
  }
}
