package main

import (
	"fmt"
	"log"
	"marcel-kerth/protobuf-generator/gen"
	"os"
	"path/filepath"
	"strings"
)

const genDir string = "/src/generated"
const srcDir string = "/src/source"

// key (e.g. go) : src (e.g. /path/to/dir)
var genDirs map[string]string

func init() {
	genDirs = make(map[string]string)

	genDirs["go"] = filepath.Join(genDir, "go")
	genDirs["python"] = filepath.Join(genDir, "python")
	genDirs["typescript"] = filepath.Join(genDir, "typescript")

	// create the directories in generated
	for _, v := range genDirs {
		if err := os.MkdirAll(v, 0755); err != nil {
			log.Fatalf("[ERROR] %v", err)
		}
	}
}

func main() {
	if err := generate(); err != nil {
		log.Fatalf("[ERROR] %v", err)
	}
}

func generate() error {
	srcFiles, err := os.ReadDir(srcDir)
	if err != nil {
		return err
	}

	for _, f := range srcFiles {
		var cursor string = filepath.Join(srcDir, f.Name())

		// if the entry isn't a folder, return error
		if !f.IsDir() {
			return fmt.Errorf("%s isn't a folder", cursor)
		}

		// read the protobuf folder
		protoFiles, err := os.ReadDir(filepath.Join(srcDir, f.Name()))
		if err != nil {
			return err
		}

		// check if all files are protobuf files
		for _, f := range protoFiles {
			cursor = filepath.Join(cursor, f.Name())

			if !strings.HasSuffix(f.Name(), ".proto") {
				return fmt.Errorf("%s is not a .proto file", cursor)
			}
		}

		cursor = filepath.Join(srcDir, f.Name())

		// generate protobuf files for python
		if err := gen.GenPython(cursor, filepath.Join(genDirs["python"], f.Name())); err != nil {
			return err
		}

		// generate protobuf files for go
		if err := gen.GenGo(cursor, filepath.Join(genDirs["go"], f.Name())); err != nil {
			return err
		}

		// generate protobuf files for typescript
		if err := gen.GenTypescript(cursor, filepath.Join(genDirs["typescript"], f.Name())); err != nil {
			return err
		}
	}

	return nil
}
