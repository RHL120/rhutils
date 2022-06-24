package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"
)

type config struct {
	cmd      string
	interval int
	limit    int
	title    bool
}

func loadConfig() *config {
	var cfg config
	flag.IntVar(&cfg.interval, "interval", 2,
		"how much time should I sleep between runs")
	flag.IntVar(&cfg.limit, "limit", 0,
		"how many times should I run the command. Will run forever if 0")
	flag.BoolVar(&cfg.title, "title", false,
		"Should I display the title")
	flag.Parse()
	cfg.cmd = flag.Arg(0)
	if cfg.cmd == "" {
		fmt.Fprint(os.Stderr, "expected command")
		os.Exit(1)
	}
	return &cfg
}

func main() {
	cfg := loadConfig()
	cursesInit()
	sleepC := make(chan bool)
	charC := make(chan bool)
	go func() {
		for {
			time.Sleep(time.Duration(cfg.interval) * time.Second)
			sleepC <- true
		}
	}()
	go func() {
		for {
			c := cursesGetChar()
			if c == 'q' {
				charC <- true
			}
		}
	}()
	for i := 0; i < cfg.limit || cfg.limit == 0; i++ {
		cursesClear()
		if cfg.title {
			if i == 0 {
				cursesWritef("executed %s %d time\n\n",
					cfg.cmd, i+1)
			} else {
				cursesWritef("executed %s %d times\n\n",
					cfg.cmd, i+1)
			}
		}
		cmd := exec.Command("sh", "-c", cfg.cmd)
		out, err := cmd.Output()
		_ = out
		if err != nil {
			cursesWritef("error: %v", err)
		} else {
			cursesWrite(string(out))
		}
		select {
		case _ = <-sleepC:
			continue
		case _ = <-charC:
			cursesClean()
			return
		}
	}
	cursesClean()
}
