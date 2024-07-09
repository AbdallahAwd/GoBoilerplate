package commands

import (
	"fmt"
	"os"
	"plate/internal/utils"
)

func APIProject(projectName string) error {
	// Define the directory structure
	utils.RunGoCommand("mod", "init", projectName)

	dirs := []string{
		"cmd/" + projectName,
		"config",
		"internal/handlers",
		"internal/models",
		"internal/repositories",
		"internal/services",
		"internal/router",
		"pkg/logger",
		"pkg/middleware",
		"scripts",
	}

	// Create directories
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return fmt.Errorf("error creating directory %s: %v", dir, err)
		}
	}

	// Define files to create
	files := map[string]string{
		"cmd/" + projectName + "/main.go":          "package main",
		"config/config.go":                         "package config",
		"internal/handlers/user_handler.go":        "package handlers",
		"internal/models/user.go":                  "package models",
		"internal/repositories/user_repository.go": "package repositories",
		"internal/services/user_service.go":        "package services",
		"internal/router/router.go":                "package router",
		"pkg/logger/logger.go":                     "package logger",
		"pkg/middleware/auth_middleware.go":        "package middleware",
		"scripts/migrate.sh":                       "#!/bin/bash\n# Migration script",
		".env":                                     "# Environment variables",
		".gitignore":                               "# Git ignore file",
		"Dockerfile":                               "# Dockerfile",
	}

	// Create files
	for filePath, content := range files {
		if err := createFiles(filePath, content); err != nil {
			return fmt.Errorf("error creating file %s: %v", filePath, err)
		}
	}

	fmt.Println("API structure generated successfully.")
	return nil
}

func createFiles(path string, content string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}
