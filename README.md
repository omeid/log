# LOG

Minimalist log with semantic API and multi level lable.

[![GoDoc](https://godoc.org/github.com/omeid/log?status.svg)](https://godoc.org/github.com/omeid/log)  

```go
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
```

Will produce:

```
[INFO] I don't have a label. :( 
[WARN] This is a warning. 
[ERR!] This is a phony error. 
[INFO] child: Check out my label! 
[NOTE] child: grandchild: See, we can go deeper! 
```


### LICENSE
  [MIT](LICENSE).
