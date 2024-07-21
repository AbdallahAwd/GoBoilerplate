package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/AbdallahAwd/GoBoilerplate/internal/utils"
)

func Cmd(projectName string) error {
	start := time.Now()

	utils.RunGoCommand("mod", "init", projectName)
	mainFolders := []string{
		"cmd",
		"internal",
		"pkg",
	}
	for _, folder := range mainFolders {
		if err := os.MkdirAll(folder, os.ModePerm); err != nil {
			return fmt.Errorf("error found %v", err)
		}
	}

	if err := subfolders(); err != nil {
		return fmt.Errorf(err.Error())
	}

	// creating main.go
	if err := createFileWithPackage("main.go", "main"); err != nil {
		return fmt.Errorf("error creating main.go: %v", err)
	}

	// creating root.go
	fileName := "root.go"
	root, err := utils.FilePath(&mainFolders[0], &fileName)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if err := createFileWithPackage(root, "cmd"); err != nil {
		return fmt.Errorf("error creating root.go: %v", err)
	}

	fmt.Println(time.Since(start))
	return nil
}

func subfolders() error {
	subFolders := []string{
		"internal/commands",
		"internal/utils",
		"pkg/shared",
	}
	fileNames := []string{
		"command.go",
		"helper.go",
		"const.go",
	}

	for i, subFolder := range subFolders {
		if err := os.MkdirAll(subFolder, os.ModePerm); err != nil {
			return fmt.Errorf("error creating subfolder %v: %v", subFolder, err)
		}

		filePath, err := utils.FilePath(&subFolder, &fileNames[i])
		if err != nil {
			return fmt.Errorf(err.Error())
		}

		packageName := filepath.Base(subFolder)
		if err := createFileWithPackage(filePath, packageName); err != nil {
			return fmt.Errorf("error creating file %s: %v", filePath, err)
		}
	}

	return nil
}

func createFileWithPackage(filePath, packageName string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("package %s\n", packageName))
	return err
}
