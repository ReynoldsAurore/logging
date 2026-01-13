package logging

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func Logging(returnedErr error) {
	filename := "logs.log"
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("logging: cannot get filepath: %s\n", err.Error())
	} else {
		path := strings.Split(pwd, "/")
		fmt.Println(path)
		temp := filename
		filename = path[0]
		for i, val := range path {
			if i == 0 {
				continue
			}
			filename += val + "_"
		}
		filename += temp
	}

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		if os.IsPermission(err) {
			fmt.Printf("logging: cannot create or open logfile, premission denied: %s\n", err.Error())
		} else {
			fmt.Printf("logging: cannot create or open logfile: %s\n", err.Error())
		}
		return
	}
	defer file.Close()

	logger := log.New(file, "", log.Flags())
	logger.Printf("%s\n", returnedErr.Error())
}
