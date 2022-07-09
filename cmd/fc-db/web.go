package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/BossHobby/fc-db/pkg/betaflight"
	"github.com/BossHobby/fc-db/pkg/fc"
	"github.com/BossHobby/fc-db/pkg/quicksilver"
	"github.com/BossHobby/fc-db/pkg/util"
)

type WebIndexEntry struct {
	MCU          string `json:"mcu"`
	Board        string `json:"board"`
	Manufacturer string `json:"manufacturer"`
}

type WebTarget struct {
	fc.Target
}

func generateWeb() error {
	unifiedFolder := filepath.Join("betaflight", "configs", "default")
	files, err := ioutil.ReadDir(unifiedFolder)
	if err != nil {
		return err
	}

	index := make([]WebIndexEntry, 0)
	for _, f := range files {
		if f.IsDir() {
			continue
		}

		path := filepath.Join(unifiedFolder, f.Name())
		log.Printf("processing %s\n", path)

		t, err := betaflight.ParseConfig(path)
		if err != nil {
			return err
		}

		index = append(index, WebIndexEntry{
			MCU:          t.MCU,
			Board:        t.Board,
			Manufacturer: t.Manufacturer,
		})

		targetFolder := filepath.Join("target", t.Manufacturer+"-"+t.Board)

		webTarget := WebTarget{}
		targetFile := filepath.Join(targetFolder, t.Manufacturer+"-"+t.Board+".json")
		if err := util.ReadJson(targetFile, &webTarget); err != nil && !os.IsNotExist(err) {
			return err
		}

		webTarget.Target = *t

		if err := util.WriteJson(targetFile, webTarget); err != nil {
			return err
		}

		if err := quicksilver.WriteHeader(*t, filepath.Join(targetFolder, "target.h")); err != nil {
			return err
		}
	}

	if err := util.WriteJson(filepath.Join("target", "index.json"), index); err != nil {
		return err
	}

	return nil
}
