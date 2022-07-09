package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/BossHobby/fc-db/pkg/betaflight"
	"github.com/BossHobby/fc-db/pkg/fc"
	"github.com/BossHobby/fc-db/pkg/quicksilver"
)

func parseUnified(source string) error {
	log.Printf("processing %s\n", source)

	t, err := betaflight.ParseConfig(source)
	if err != nil {
		return err
	}

	targetFolder := filepath.Join("target", t.Manufacturer+"-"+t.Board)
	if err := os.MkdirAll(targetFolder, 0755); err != nil && !os.IsExist(err) {
		return err
	}

	filename := filepath.Join(targetFolder, t.Manufacturer+"-"+t.Board+".json")
	if err := fc.WriteTarget(filename, t); err != nil {
		return err
	}

	return nil
}

func targetFromUnified(filename string) error {
	source := filepath.Join(filename, "configs", "default")

	files, err := ioutil.ReadDir(source)
	if err != nil {
		return err
	}

	for _, f := range files {
		path := filepath.Join(source, f.Name())
		if f.IsDir() {
			continue
		}

		if err := parseUnified(path); err != nil {
			return err
		}
	}

	return nil
}

func targetToQuicksilver(source string) error {
	files, err := ioutil.ReadDir(source)
	if err != nil {
		return err
	}

	for _, f := range files {
		if !f.IsDir() {
			continue
		}

		folder := filepath.Join(source, f.Name())

		log.Printf("processing %s\n", folder)
		t, err := fc.ReadTarget(filepath.Join(folder, f.Name()+".json"))
		if err != nil {
			return err
		}

		if err := quicksilver.WriteHeader(t, filepath.Join(folder, "target.h")); err != nil {
			return err
		}
	}

	return nil
}

func targetFromDump(source string, dest string) error {
	t, err := betaflight.ParseConfig(source)
	if err != nil {
		return err
	}

	if err := quicksilver.WriteHeader(t, dest); err != nil {
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
	case "from-unified":
		if flag.NArg() < 3 {
			log.Println("missing <filename>")
			return nil
		}
		return targetFromUnified(flag.Arg(2))
	case "to-quicksilver":
		if flag.NArg() < 3 {
			log.Println("missing <filename>")
			return nil
		}
		return targetToQuicksilver(flag.Arg(2))
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
	}
}
