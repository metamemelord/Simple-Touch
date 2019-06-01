package windowstouch

import (
	"fmt"
	"os"
)

func main() {
	dirname, err := os.Getwd()
	if err != nil {
		fmt.Println(fmt.Errorf("Something went wrong while reading the filesystem"))
		os.Exit(0)
	}
	commandLineArgs := os.Args
	if len(commandLineArgs) < 2 {
		fmt.Println("Please provide a file to create")
		os.Exit(0)
	}
	filesToCreate := commandLineArgs[1:]
	for _, filename := range filesToCreate {
		finalPath := dirname + string(os.PathSeparator) + filename
		if _, err := os.Stat(finalPath); err != nil {
			switch err {
			case os.ErrInvalid:
				fmt.Println("Invalid argument:", filename)
				os.Exit(0)
			case os.ErrPermission:
				fmt.Println("Permission denied")
				os.Exit(0)
			default:
				_, err := os.Create(finalPath)
				if err != nil {
					switch err {
					case os.ErrInvalid:
						fmt.Println("Invalid argument", filename)
					case os.ErrPermission:
						fmt.Println("Permission denied")
					default:
						fmt.Println("Something went wrong")
					}
					os.Exit(0)
				}
			}
		}
	}
}
