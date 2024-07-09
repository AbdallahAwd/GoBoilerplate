package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"plate/internal/utils"
)

func UltimateProjectStructure(projectName string) error {
	// Create main directories
	utils.RunGoCommand("mod", "init", projectName)
	directories := []string{
		"cmd", "internal", "pkg", "api", "web", "scripts", "configs", "tests", "docs",
	}
	for _, dir := range directories {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return fmt.Errorf("error creating directory %s: %v", dir, err)
		}
	}

	// Create subdirectories and files
	if err := createCmdStructure(projectName); err != nil {
		return err
	}
	if err := createInternalStructure(); err != nil {
		return err
	}
	if err := createPkgStructure(); err != nil {
		return err
	}
	if err := createApiStructure(); err != nil {
		return err
	}
	if err := createWebStructure(); err != nil {
		return err
	}
	if err := createScriptsStructure(); err != nil {
		return err
	}
	if err := createConfigsStructure(); err != nil {
		return err
	}
	if err := createTestsStructure(); err != nil {
		return err
	}

	// Create root files

	rootFiles := []string{".gitignore", "README.md"}
	for _, file := range rootFiles {
		if err := createFile(file); err != nil {
			return err
		}
	}

	return nil
}

func createCmdStructure(projectName string) error {
	cmdDirs := []string{projectName, "another-app"}
	for _, dir := range cmdDirs {
		cmdPath := filepath.Join("cmd", dir)
		if err := os.MkdirAll(cmdPath, os.ModePerm); err != nil {
			return fmt.Errorf("error creating cmd subdirectory %s: %v", dir, err)
		}
		if err := createFile(filepath.Join(cmdPath, "main.go")); err != nil {
			return err
		}
	}
	return nil
}

func createInternalStructure() error {
	internalDirs := []string{"config", "database"}
	for _, dir := range internalDirs {
		internalPath := filepath.Join("internal", dir)
		if err := os.MkdirAll(internalPath, os.ModePerm); err != nil {
			return fmt.Errorf("error creating internal subdirectory %s: %v", dir, err)
		}
		if err := createFile(filepath.Join(internalPath, dir+".go")); err != nil {
			return err
		}
	}
	return nil
}

func createPkgStructure() error {
	pkgPath := filepath.Join("pkg", "mypackage")
	if err := os.MkdirAll(pkgPath, os.ModePerm); err != nil {
		return fmt.Errorf("error creating pkg subdirectory: %v", err)
	}
	return createFile(filepath.Join(pkgPath, "mypackage.go"))
}

func createApiStructure() error {
	apiDirs := []string{"handler", "middleware"}
	for _, dir := range apiDirs {
		apiPath := filepath.Join("api", dir)
		if err := os.MkdirAll(apiPath, os.ModePerm); err != nil {
			return fmt.Errorf("error creating api subdirectory %s: %v", dir, err)
		}
		if err := createFile(filepath.Join(apiPath, dir+".go")); err != nil {
			return err
		}
	}
	return nil
}

func createWebStructure() error {
	webDirs := []string{"static/css", "static/js", "templates"}
	for _, dir := range webDirs {
		webPath := filepath.Join("web", dir)
		if err := os.MkdirAll(webPath, os.ModePerm); err != nil {
			return fmt.Errorf("error creating web subdirectory %s: %v", dir, err)
		}
	}
	return createFile(filepath.Join("web", "templates", "index.html"))
}

func createScriptsStructure() error {
	scripts := []string{"build.sh", "deploy.sh"}
	for _, script := range scripts {
		if err := createFile(filepath.Join("scripts", script)); err != nil {
			return err
		}
	}
	return nil
}

func createConfigsStructure() error {
	configs := []string{"development.yaml", "production.yaml"}
	for _, config := range configs {
		if err := createFile(filepath.Join("configs", config)); err != nil {
			return err
		}
	}
	return nil
}

func createTestsStructure() error {
	testDirs := []string{"unit", "integration"}
	for _, dir := range testDirs {
		testPath := filepath.Join("tests", dir)
		if err := os.MkdirAll(testPath, os.ModePerm); err != nil {
			return fmt.Errorf("error creating test subdirectory %s: %v", dir, err)
		}
	}
	return nil
}

func createFile(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("error creating file %s: %v", path, err)
	}
	defer file.Close()

	// Write package declaration if it's a .go file
	if filepath.Ext(path) == ".go" {
		packageName := filepath.Base(filepath.Dir(path))
		if packageName == "cmd" {
			packageName = "main"
		}
		_, err = file.WriteString(fmt.Sprintf("package %s\n", packageName))
		if err != nil {
			return fmt.Errorf("error writing to file %s: %v", path, err)
		}
	}

	return nil
}
