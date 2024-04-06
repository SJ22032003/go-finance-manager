package file_manager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func New(inputFilePath, outputFilePath string) *FileManager {
	return &FileManager{
		InputFilePath:  inputFilePath,
		OutputFilePath: outputFilePath,
	}
}

func (fl FileManager) ReadFileManager() ([]string, error) {
	file, err := os.Open(fl.InputFilePath)
	if err != nil {
		return nil, errors.New("file not found")
	}

	defer file.Close() // Close the file when the function returns

	scanner := bufio.NewScanner(file)

	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return lines, nil
}

func (fl FileManager) WriteFileToJSON(data interface{}) error {
	file, err := os.Create(fl.OutputFilePath)
	if err != nil {
		file.Close()
		return errors.New("file not found")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		file.Close()
		return errors.New("error encoding json")
	}

	file.Close()
	return nil

}

func (fl *FileManager) SetOutputPath(path string) {
	fl.OutputFilePath = path
}

func (fl *FileManager) SetInputPath(path string) {
	fl.InputFilePath = path
}