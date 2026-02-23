package utils

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/asekirieshka/luarpc/assets"
)

const (
	dirMode  fs.FileMode = 0755
	fileMode fs.FileMode = 0644
)

func InitProject(targetDir string) error {
	mainTemplateFile, err := assets.TemplateFS.ReadFile("lua-template/main.lua")
	if err != nil {
		return fmt.Errorf("rror reading embedded main.lua template file %v", err)
	}

	stubsTemplateFile, err := assets.TemplateFS.ReadFile("lua-template/definitions.lua")
	if err != nil {
		return fmt.Errorf("error reading embedded definitions.lua template file: %v", err)
	}

	if targetDir != "." {
		if err := os.MkdirAll(targetDir, dirMode); err != nil {
			return fmt.Errorf("failed to create folder: %v", err)
		}
	}

	mainPath := filepath.Join(targetDir, "main.lua")
	stubsPath := filepath.Join(targetDir, "definitions.lua")

	if _, err := os.Stat(mainPath); os.IsNotExist(err) {
		if err := os.WriteFile(mainPath, mainTemplateFile, fileMode); err != nil {
			return fmt.Errorf("failed to write main.lua file: %v", err)
		}
	} else {
		fmt.Printf("main.lua already exists in %s, skipping...\n", targetDir)
	}

	if err := os.WriteFile(stubsPath, stubsTemplateFile, fileMode); err != nil {
		return fmt.Errorf("failed to write definitions.lua file: %v", err)
	}

	var projectInitPlace = targetDir

	if targetDir == "." {
		projectInitPlace = "the current folder"
	}

	fmt.Printf("Project initialized in %s! Open main.lua to get started!\n", projectInitPlace)
	return nil
}
