package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

type config struct {
	char  bool
	lines bool
	words bool
	files []string
}

func loadCfg() *config {
	var cfg config
	flag.BoolVar(&cfg.char, "chars", false, "print the char count")
	flag.BoolVar(&cfg.lines, "lines", false, "print the line count")
	flag.BoolVar(&cfg.lines, "words", false, "print the words count")
	flag.Parse()
	cfg.files = flag.Args()
	if !cfg.char && !cfg.lines && !cfg.words {
		cfg.char = true
		cfg.lines = true
		cfg.words = true
	}
	return &cfg
}

func countChar(f []byte) int {
	return len(f)
}

func countLines(f []byte) (ret int) {
	for _, i := range f {
		if i == '\n' {
			ret += 1
		}
	}
	return ret
}

func isWordSep(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func countWords(f []byte) (ret int) {
	for idx, c := range f {
		if isWordSep(c) {
			continue
		}
		if idx >= len(f)-1 || isWordSep(f[idx+1]) {
			ret += 1
		}
	}
	return ret
}

func readInput() []byte {
	buff := make([]byte, 0)
	for {
		i := make([]byte, 1)
		os.Stdin.Read(i)
		if i[0] == 0 {
			break
		}
		buff = append(buff, i[0])
	}
	return buff
}

func main() {
	cfg := loadCfg()
	var files [][]byte
	if len(cfg.files) > 0 {
		files = make([][]byte, len(cfg.files))
		for idx, i := range cfg.files {
			content, err := ioutil.ReadFile(i)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			files[idx] = content
		}
	} else {
		files = make([][]byte, 1)
		files[0] = readInput()
	}
	total := struct {
		lines uint
		words uint
		char  uint
	}{0, 0, 0}
	for idx, i := range files {
		// These if statements are duped bellow. They could
		// be extracted to function but that would run the functions
		// even if their corresponding bool is false. For example, if
		// cfg.lines was false countLines would be called even if it
		// was not needed. I guess the function could take another function
		// and call it if cfg.x is true but that seems a bit hackey.
		// I don't feel like go works well with functional programming.
		if cfg.lines {
			total.lines += 1
			fmt.Print(" ", countLines(i))
		}
		if cfg.words {
			total.words += 1
			fmt.Print(" ", countWords(i))
		}
		if cfg.char {
			total.words += 1
			fmt.Print(" ", countChar(i))
		}
		if len(cfg.files) >= 1 {
			fmt.Println(" ", cfg.files[idx])
		}
	}
	if len(files) > 1 {
		if cfg.lines {
			fmt.Print(" ", total.lines)
		}
		if cfg.words {
			fmt.Print(" ", total.words)
		}
		if cfg.char {
			fmt.Print(" ", total.char)
		}
		fmt.Println(" total")
	}
}
