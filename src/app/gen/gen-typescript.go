package gen

import (
	"fmt"
	"marcel-kerth/protobuf-generator/util"
	"os"
	"os/exec"
	"path/filepath"
)

func GenTypescript(srcDir, outDir string) error {
	if err := os.MkdirAll(outDir, 0755); err != nil {
		return fmt.Errorf("failed to create parent dir: %w", err)
	}

	var cmds []*exec.Cmd

	// find all .proto files in srcDir
	pattern := filepath.Join(srcDir, "*.proto")
	protoFiles, err := filepath.Glob(pattern)
	if err != nil {
		return fmt.Errorf("glob failed: %w", err)
	}
	if len(protoFiles) == 0 {
		return fmt.Errorf("no .proto files found in %s", srcDir)
	}

	args := []string{
		fmt.Sprintf("--plugin=protoc-gen-ts=%s", "./node_modules/.bin/protoc-gen-ts"),
		fmt.Sprintf("--ts_out=%s", outDir),
		"--ts_opt=long_type_string",
		fmt.Sprintf("--proto_path=%s", srcDir),
	}
	args = append(args, protoFiles...)

	cmds = append(cmds, exec.Command("protoc", args...))

	if err := util.RunCmds(cmds); err != nil {
		return fmt.Errorf("protoc TS generation failed: %w", err)
	}

	return nil
}
