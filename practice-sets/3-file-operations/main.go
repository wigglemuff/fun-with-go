package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readAndPrintFileData(file string) []byte {
	bytes, _ := os.ReadFile(file)
	fmt.Println(string(bytes))
	return bytes
}

func main() {
	var file *os.File

	// 1. Create a file
	fmt.Println("Create a file")
	file, _ = os.Create("example.txt")
	file.Write([]byte("1.1 Some text with file.Write([]byte()) \n"))
	file.WriteString("1.2 Some text with file.WriteString() \n")
	file.Close()

	// 2. Create a file only if it does not exist
	// All writes under this are no op because example.txt is already created.
	fmt.Println("Create a file only if it does not exist")
	file, _ = os.OpenFile("example.txt", os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
	file.Write([]byte("2.1 ABC \n"))
	file.WriteString("2.2 DEF \n")
	file.Close()

	// 2b. Create a file only if it does not exist
	// Writes work with new example-two.txt file creation
	fmt.Println("Create a file only if it does not exist 2")
	file, _ = os.OpenFile("example-two.txt", os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
	file.Write([]byte("2.1 ABC \n"))
	file.WriteString("2.2 DEF \n")
	file.Close()

	// 3. Read all file data at once
	fmt.Println("Read all file data at once")
	bytes := readAndPrintFileData("example.txt")
	lines := strings.Split(string(bytes), "\n")
	for _, line := range lines {
		fmt.Println(line)
	}

	// 4. Read a large file line by line
	fmt.Println("Read a large file line by line")
	file, _ = os.Open("example.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// 5. Write data to the file [already done above]

	// 6. Append data to the file
	fmt.Println("Append data to the file")
	file, _ = os.OpenFile("example.txt", os.O_WRONLY|os.O_APPEND, 0644)
	file.WriteString("Appended string \n")
	file.Close()
	readAndPrintFileData("example.txt")

	// 7. Write data to a file line by line using buffer
	fmt.Println("Write data to a file line by line using buffer")
	file, _ = os.OpenFile("example.txt", os.O_WRONLY|os.O_APPEND, 0644)
	writer := bufio.NewWriter(file)
	writer.WriteString("Writing a line to buffer \n")
	writer.WriteString("Writing another line to buffer \n")
	writer.Flush()
	file.Close()
	readAndPrintFileData("example.txt")

	// 8. Delete a file
	fmt.Println("Delete a file")
	os.Remove("example.txt")
	os.Remove("example-two.txt")

}
