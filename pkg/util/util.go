package util

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strconv"
)

func MustParseInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}

func WriteJson(filename string, val interface{}) error {
	if err := os.MkdirAll(filepath.Dir(filename), 0755); err != nil && !os.IsExist(err) {
		return err
	}

	w, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer w.Close()

	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	if err := enc.Encode(val); err != nil {
		return err
	}

	return nil
}

func ReadJson(filename string, val interface{}) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := json.NewDecoder(f).Decode(val); err != nil {
		return err
	}

	return nil
}
