go-windows-subst
================

go-windows-subst is API wrapper for subst.exe.

example:

```go
package main

import (
	"fmt"
	"os"

	"github.com/zetamatta/go-windows-subst"
)

func mains(args []string) error {
	switch len(args) {
	case 1:
		target, err := subst.Query(args[0])
		if err != nil {
			return err
		}
		fmt.Printf("%s => %s\n", args[0], target)
		return nil
	case 2:
		if args[1] == "/D" {
			return subst.Remove(args[0])
		} else {
			return subst.Define(args[0], args[1])
		}
	default:
		return fmt.Errorf("Usage: %s DRIVE TARGET", os.Args[0])
	}
}

func main() {
	if err := mains(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
```
