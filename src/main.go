package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 5 {
		fmt.Println("Usage: string-replacer <path to folder> <.ext> <search string> <replace string>")
		os.Exit(1)
	}

	folderPath := os.Args[1]
	fileExt := os.Args[2]
	searchString := os.Args[3]
	replaceString := os.Args[4]

	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == fileExt {
			if err := replaceInFile(path, searchString, replaceString); err != nil {
				fmt.Printf("Ошибка при замене в файле %s: %v\n", path, err)
			} else {
				fmt.Printf("Заменено в файле %s\n", path)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Ошибка при поиске файлов: %v\n", err)
		os.Exit(1)
	}
}

func replaceInFile(filePath, oldStr, newStr string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, strings.ReplaceAll(scanner.Text(), oldStr, newStr))
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	newContent := strings.Join(lines, "\n")
	if err := os.WriteFile(filePath, []byte(newContent), 0644); err != nil {
		return err
	}

	return nil
}
