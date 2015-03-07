package main

import "github.com/omeid/log"

func main() {
  l := log.New("")

  l.Info("I don't have a label. :(")
  l.Warn("This is a warning.")
  l.Error("This is a phony error.")

  c := l.New("child: ")
  c.Info("Check out my label!")
  c = c.New("grandchild: ")
  c.Notice("See, we can go deeper!")
}
