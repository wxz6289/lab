package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func read_file(filePath string) {
	// 读取文件内容
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// 打印文件内容
	fmt.Println("File content:")
	fmt.Println(string(data))
}

func show_file() {
	filePath := "builtin.md"
	read_file(filePath)
}

func read_file_stream(filePath string) {
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// 创建缓冲区读取器
	buffer := make([]byte, 10)

	for {
		// 读取文件内容到缓冲区
		n, err := file.Read(buffer)
		if err != nil {

			if err == io.EOF {
				break // 文件读取完毕
			}
			fmt.Println("Error reading file:", err)
			return
		}

		// 打印读取的内容
		fmt.Print(string(buffer[:n]))
	}
}

func show_file_stream() {
	filePath := "builtin.md"
	read_file_stream(filePath)
}

func read_file_line(filePath string) {
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// 创建缓冲区读取器
	buffer := bufio.NewReader(file)

	for {
		// 读取文件内容到缓冲区
		line, _, err := buffer.ReadLine()
		if err != nil {
			if err == io.EOF {
				break // 文件读取完毕
			}
			fmt.Println("Error reading file:", err)
			return
		}

		// 打印读取的内容
		fmt.Println(string(line))
	}
}

func show_file_line() {
	filePath := "builtin.md"
	read_file_line(filePath)
}

func read_file_by_scanner(filePath string) {
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// 创建扫描器
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords) // 按词扫描

	for scanner.Scan() {
		// 打印每一行内容
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func show_file_by_scanner() {
	filePath := "builtin.md"
	read_file_by_scanner(filePath)
}

func write_file(filePath string, content string) {
	// 创建或打开文件
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// 写入内容到文件
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Content written to file successfully.")
}

func show_write_file() {
	filePath := "output.txt"
	content := "Hello, this is a sample content written to the file."
	write_file(filePath, content)
}

func append_file(filePath string, content string) {
	// 创建或打开文件
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// 写入内容到文件
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Content appended to file successfully.")
}

func show_append_file() {
	filePath := "output.txt"
	content := "\nThis is an appended line."
	append_file(filePath, content)
}

func copy_file(srcPath string, destPath string) {
	// 打开源文件
	srcFile, err := os.Open(srcPath)
	if err != nil {
		fmt.Println("Error opening source file:", err)
		return
	}
	defer srcFile.Close()

	// 创建目标文件
	destFile, err := os.Create(destPath)
	if err != nil {
		fmt.Println("Error creating destination file:", err)
		return
	}
	defer destFile.Close()

	// 复制内容
	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		fmt.Println("Error copying file:", err)
		return
	}

	fmt.Println("File copied successfully.")
}

func show_copy_file() {
	srcPath := "output.txt"
	destPath := "output_copy.txt"
	copy_file(srcPath, destPath)
}

func list_files_in_directory(dirPath string) {
	// 打开目录
	dir, err := os.Open(dirPath)
	if err != nil {
		fmt.Println("Error opening directory:", err)
		return
	}
	defer dir.Close()

	// 读取目录内容
	files, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	// 打印文件列表
	fmt.Println("Files in directory:")
	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir(), file.Size())
	}
}

func show_list_files_in_directory() {
	dirPath := "."
	list_files_in_directory(dirPath)
}