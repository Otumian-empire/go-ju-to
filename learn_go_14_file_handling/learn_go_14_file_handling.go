package main

// https://pkg.go.dev/os@go1.20.7
import (
	"fmt"
)

func main() {
	fmt.Println("File Handling")

	// create folder
	// https://www.redhat.com/sysadmin/linux-file-permissions-explained
	// https://www.guru99.com/file-permissions.html
	// r=4, w=2, x=1, user_group_others
	/* mkdirError := os.Mkdir("some_folder", 0664)
	if mkdirError != nil {
		fmt.Println("An error occurred")
		fmt.Println(mkdirError)
	} */

	// remove folder
	/* removeErr := os.Remove("some_folder")
	if removeErr != nil {
		fmt.Println("An error occurred")
		fmt.Println(removeErr)
	} */

	// create file and write into it
	/* writeErr := os.WriteFile("hello.txt", []byte("Hello there\n"), 0664)

	if writeErr != nil {
		fmt.Println(writeErr)
	} */

	// read file: whole, line by line, word by word
	/* content, readFileErr := os.ReadFile("hello.txt")

	if readFileErr != nil {
		fmt.Println(readFileErr)
	} else {
		fmt.Println(string(content))
	} */

	// remove file
	/* removeFileError := os.Remove("hello.txt")
	if removeFileError != nil {
		fmt.Println(removeFileError)
	} else {
		fmt.Println("File removed")
	} */
}
