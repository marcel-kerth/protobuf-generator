package gen

import (
	"fmt"
	"marcel-kerth/protobuf-generator/util"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func GenGo(srcDir, outDir string) error {
	// convert outDir path
	parent := filepath.Dir(outDir)
	base := filepath.Base(outDir)
	safeName := strings.ToLower(strings.ReplaceAll(base, "-", "_"))
	outDir = filepath.Join(parent, safeName)

	parentDir := filepath.Dir(outDir)
	if err := os.MkdirAll(parentDir, 0755); err != nil {
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
		fmt.Sprintf("-I=%s", srcDir),
		fmt.Sprintf("--go_out=%s", parentDir),
		fmt.Sprintf("--go-grpc_out=%s", parentDir),
	}
	args = append(args, protoFiles...)

	cmds = append(cmds, exec.Command("protoc", args...))

	if err := util.RunCmds(cmds); err != nil {
		return fmt.Errorf("protoc Go generation failed: %w", err)
	}

	return nil
}
