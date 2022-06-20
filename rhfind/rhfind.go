package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

type config struct {
	name string
	sp   []string
	sep  bool
}

func loadConfig() (*config, error) {
	var cfg config
	flag.StringVar(&cfg.name, "name", "",
		"filter in files that match the regexp")
	flag.BoolVar(&cfg.sep, "sep", false,
		"Add a seprator between starting paths")
	flag.Parse()
	cfg.sp = flag.Args()
	if len(cfg.sp) < 1 {
		cfg.sp = append(cfg.sp, ".")
	}
	return &cfg, nil
}

func main() {
	cfg, err := loadConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	f := func(path string, entry os.FileInfo, err error) error {
		match, err := regexp.MatchString(cfg.name, entry.Name())
		if err != nil {
			return err
		}
		if match {
			fmt.Println(path)
		}
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
