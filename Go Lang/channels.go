package main

import (
  "fmt"
  "time"
)

func pinger(c chan string) {
  for i := 0; ; i++ {
    c <- "ping"+string(i)
    time.Sleep(time.Second * 1)
  }
}

func printer(c chan string) {
  for {
    msg := <- c
    fmt.Println(msg)
    time.Sleep(time.Second * 1)
  }
}

func ponger(c chan string) {
  for i := 0; ; i++ {
    c <- "pong"
  }
}

func goChannels() {
  var c chan string = make(chan string)

  go pinger(c)
  go ponger(c)
  go printer(c)

  var input string
  fmt.Scanln(&input)
}