package main

import (
	"log"
	"os"

	"github.com/hashicorp/logutils"
	"github.com/jessevdk/go-flags"

	"github.com/alaija/rock_n_walk_bot/bot/tg"
)

var opts struct {
	Dbg   bool   `long:"dbg" env:"DEBUG" description:"debug mode"`
	Token string `short:"t" env:"TOKEN" description:"telegram bot token"`
}

func main() {
	log.Printf("Rock-n-Walk Bot")

	if _, err := flags.Parse(&opts); err != nil {
		os.Exit(1)
	}

	setupLog(opts.Dbg)

	tg.RunRNWBot(opts.Token, opts.Dbg)
}

func setupLog(dbg bool) {
	filter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "INFO", "WARN", "ERROR"},
		MinLevel: logutils.LogLevel("INFO"),
		Writer:   os.Stdout,
	}

	log.SetFlags(log.Ldate | log.Ltime)

	if dbg {
		log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
		filter.MinLevel = logutils.LogLevel("DEBUG")
	}

	log.SetOutput(filter)
}
