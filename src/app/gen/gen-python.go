package gen

import (
	"fmt"
	"marcel-kerth/protobuf-generator/util"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

func GenPython(srcDir, outDir string) error {
	// convert outDir path
	parent := filepath.Dir(outDir)
	base := filepath.Base(outDir)
	safeName := strings.ToLower(strings.ReplaceAll(base, "-", "_"))
	outDir = filepath.Join(parent, safeName)

	if err := os.MkdirAll(outDir, 0755); err != nil {
		return err
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
		"-m", "grpc_tools.protoc",
		fmt.Sprintf("-I=%s", srcDir),
		fmt.Sprintf("--python_out=%s", outDir),
		fmt.Sprintf("--pyi_out=%s", outDir),
		fmt.Sprintf("--grpc_python_out=%s", outDir),
	}
	args = append(args, protoFiles...)

	cmds = append(cmds, exec.Command("python3", args...))

	if err := util.RunCmds(cmds); err != nil {
		return err
	}

	// fix relative imports in generated Python files
	if err := fixPythonImports(outDir); err != nil {
		return fmt.Errorf("import fix failed: %w", err)
	}

	return nil
}

// fixPythonImports adjusts 'import xxx_pb2' → 'from . import xxx_pb2'
func fixPythonImports(dir string) error {
	files, err := filepath.Glob(filepath.Join(dir, "**/*.py"))
	if err != nil {
		return err
	}

	re := regexp.MustCompile(`^import (.*_pb2(?:_grpc)?)`)

	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			return err
		}

		lines := strings.Split(string(content), "\n")
		for i, line := range lines {
			lines[i] = re.ReplaceAllString(line, "from . import $1")
		}

		if err := os.WriteFile(file, []byte(strings.Join(lines, "\n")), 0644); err != nil {
			return err
		}
	}

	return nil
}
