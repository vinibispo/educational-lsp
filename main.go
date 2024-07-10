package main

import (
	"bufio"
	"educationallsp/rpc"
	"log"
	"os"
)

func main() {
  logger := getLogger("educationallsp.txt")
  logger.Println("Starting up")

  scanner := bufio.NewScanner(os.Stdin)
  scanner.Split(rpc.Split)

  for scanner.Scan() {
    handleMessage(scanner.Text())
  }
}

func handleMessage(_ any) {}

func getLogger(filename string) *log.Logger {
  logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)

  if err != nil {
    panic("hey, you didn't provide a valid filename")
  }

  return log.New(logfile, "[educationallsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
