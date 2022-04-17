package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/BossHobby/fc-db/pkg/fc"
)

func readFile(filename string) (*fc.Target, error) {
	var target fc.Target

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(f).Decode(&target); err != nil {
		return nil, err
	}

	return &target, nil
}

func main() {
	flag.Parse()

	if flag.NArg() != 1 {
		log.Println("missing <folder>")
		os.Exit(1)
	}

	filename := flag.Arg(0)
	files, err := ioutil.ReadDir(filename)
	if err != nil {
		log.Fatal(err)
	}

	w := csv.NewWriter(os.Stdout)
	if err := w.Write([]string{"MCU", "MGFR", "BOARD", "SPI1", "SPI2", "SPI3", "SPI4", "SPI5", "SPI6"}); err != nil {
		log.Fatal(err)
	}

	for _, f := range files {

		t, err := readFile(filepath.Join(filename, f.Name()))
		if err != nil {
			log.Fatal(err)
		}

		offset := 2
		entries := make([]string, 8)
		addEntry := func(d *fc.SPIDevice, name string) {
			if d == nil || d.Port == 0 {
				return
			}

			index := offset + d.Port
			if entries[index] == "" {
				entries[index] = name
			} else {
				entries[index] += ", " + name
			}
		}

		entries[0] = t.MCU
		entries[1] = t.Manufacturer
		entries[2] = t.Board

		if len(t.Gyros) >= 1 && t.Gyros[0].Port != 0 {
			entries[offset+t.Gyros[0].Port] = "GYRO1"
		}
		if len(t.Gyros) >= 2 && t.Gyros[1].Port != 0 {
			entries[offset+t.Gyros[1].Port] = "GYRO2"
		}

		addEntry(t.OSD, "OSD")
		addEntry(t.DataFlash, "DATA_FLASH")
		addEntry(t.SDCard, "SD_CARD")
		addEntry(&t.RX.SPIDevice, "RX")

		if err := w.Write(entries); err != nil {
			log.Fatal(err)
		}
	}

	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
