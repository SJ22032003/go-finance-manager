# Go Finance Manager

This project is a small application that demonstrates essential Go implementations, file handling, and concurrency.

## Features

### File Handling

The application uses the `FileManager` struct from the [`file_manager`](file_manager/file_manager.go) package to handle file operations. It provides methods to read from a file and write data to a JSON file. The `ReadFileManager` method reads lines from a file and returns them as a slice of strings.

### Concurrency

The application uses Goroutines and channels to concurrently read files and process data. The `ProcessFile` function reads data from a file and sends it to a channel. The `ProcessData` function receives data from the channel, processes it, and writes the results to a file.

## Usage

To run the application, execute the following command:

```sh
go run main.go
```

The application reads data from the `data.txt` file, processes the data concurrently, and writes the results to the `result.json` file.
