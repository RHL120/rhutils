package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

type config struct {
	number bool
	files  []string
}

func loadConfig() *config {
	var cfg config
	flag.BoolVar(&cfg.number, "number", false, "number all output lines")
	flag.Parse()
	cfg.files = flag.Args()
	return &cfg
}

func main() {
	cfg := loadConfig()
	var err error
	if len(cfg.files) == 0 {
		var buff string
		for {
			b := make([]byte, 1)
			os.Stdin.Read(b)
			if b[0] == '\n' {
				fmt.Println(buff)
				buff = ""
			} else if b[0] == 0 {
				fmt.Print(buff)
				return
			} else {
				buff += string(b)
			}
		}
	}
	if cfg.number {
		lineNum := 1
		fmt.Printf("    1  ")
		for fileIdx, file := range cfg.files {
			content, e := ioutil.ReadFile(file)
			if e != nil {
				err = e
				break
			}
			for charIdx, char := range content {
				if char == '\n' && !(charIdx == len(content)-1 &&
					fileIdx == len(cfg.files)-1) {
					lineNum++
					fmt.Printf("\n    %d  ", lineNum)
					continue
				}
				fmt.Printf("%c", char)
			}
		}
	} else {
		for _, i := range cfg.files {
			content, e := ioutil.ReadFile(i)
			if e != nil {
				err = e
				break
			}
			fmt.Print(string(content))
		}
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
