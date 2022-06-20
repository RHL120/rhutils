package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type config struct {
	sp  []string
	sep bool
}

func loadConfig() (*config, error) {
	var cfg config
	flag.BoolVar(&cfg.sep, "sep", false,
		"Add a seprator between starting paths")
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		cfg.sp = make([]string, 1)
		cfg.sp[0] = "."
	} else {
		cfg.sp = flag.Args()
	}
	return &cfg, nil
}

func main() {
	cfg, err := loadConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	f := func(path string, entry os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	}
	for idx, i := range cfg.sp {

		if err := filepath.Walk(i, filepath.WalkFunc(f)); err != nil {
			fmt.Printf("Failed to walk through directory %s, error: %v\n",
				i, err)

		}
		if idx != len(cfg.sp)-1 && cfg.sep {
			fmt.Println("----")
		}
	}
}
