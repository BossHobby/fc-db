package main

import (
	"encoding/json"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/NotFastEnuf/fc-db/pkg/betaflight"
)

func parseConfig(filename string, w io.Writer) error {
	t, err := betaflight.ParseConfig(filename)
	if err != nil {
		return err
	}

	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	if err := enc.Encode(t); err != nil {
		return err
	}

	return nil
}

func main() {
	flag.Parse()

	if flag.NArg() != 1 {
		log.Println("missing <filename>")
		os.Exit(1)
	}

	filename := flag.Arg(0)

	s, err := os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}

	if s.IsDir() {
		files, err := ioutil.ReadDir(filename)
		if err != nil {
			log.Fatal(err)
		}

		for _, f := range files {
			path := filepath.Join(filename, f.Name())
			if f.IsDir() {
				continue
			}

			log.Printf("processing %s\n", path)
			w, err := os.Create(f.Name() + ".json")
			if err != nil {
				log.Fatal(err)
			}
			defer w.Close()

			if err := parseConfig(path, w); err != nil {
				log.Fatal(err)
			}
		}
	} else {
		if err := parseConfig(filename, os.Stdout); err != nil {
			log.Fatal(err)
		}
	}

}
