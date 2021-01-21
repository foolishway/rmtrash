package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"time"
)

func main() {
	flag.Parse()

	curUser, err := user.Current()
	if err != nil {
		panic(err)
	}

	userName := curUser.Username
	trashDir := fmt.Sprintf("/Users/%s/.Trash/", userName)
	if _, err := os.Stat(trashDir); os.IsNotExist(err) {
		if err != nil {
			panic(err)
		}
	}

	var srcs []string
	if len(os.Args) > 1 {
		srcs = os.Args[1:]
	}

	// log.Printf("%v", srcs)
	if len(srcs) == 0 {
		log.Println("srcs is empty.")
		return
	}

	for _, src := range srcs {
		var file os.FileInfo
		var err error
		if file, err = os.Lstat(src); os.IsNotExist(err) {
			continue
		}

		//get ext
		ext := filepath.Ext(file.Name())

		dist := filepath.Join(trashDir, fmt.Sprintf("%s%s%s", file.Name()[:len(file.Name())-len(ext)], time.Now().Format("20060102150405"), ext))
		// log.Println(dist)
		err = os.Rename(src, dist)
		if err != nil {
			log.Printf("Rename %s to %s error %s", src, dist, err)
		}
	}
}
