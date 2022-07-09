package main

import (
	"flag"
	"log"
	"os"

	"github.com/BossHobby/fc-db/pkg/betaflight"
	"github.com/BossHobby/fc-db/pkg/quicksilver"
)

func targetFromDump(source string, dest string) error {
	t, err := betaflight.ParseConfig(source)
	if err != nil {
		return err
	}

	if err := quicksilver.WriteHeader(*t, dest); err != nil {
		return err
	}

	return nil
}

func targetSubcommands() error {
	if flag.NArg() < 2 {
		log.Println("missing target command")
		return nil
	}

	switch flag.Arg(1) {

	case "from-dump":
		if flag.NArg() < 3 {
			log.Println("missing <source>")
			return nil
		}
		if flag.NArg() < 4 {
			log.Println("missing <dest>")
			return nil
		}
		return targetFromDump(flag.Arg(2), flag.Arg(3))
	}

	return nil
}

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		log.Println("missing subcommand")
		os.Exit(1)
	}

	switch flag.Arg(0) {
	case "target":
		if err := targetSubcommands(); err != nil {
			log.Fatal(err)
		}
	case "web":
		if err := generateWeb(); err != nil {
			log.Fatal(err)
		}
	}
}
