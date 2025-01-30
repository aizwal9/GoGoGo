package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)


func searchFiles(directory ,query string, ignoreCase bool) ([]string,error){
	var files []string

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		
		if err !=nil {
			return err
		}

		if ignoreCase{
			if strings.Contains(strings.ToLower(info.Name()),strings.ToLower(query)){
				files = append(files, path)
			}

		} else{
			if strings.Contains(info.Name(),query){
				files = append(files, path)
			}
		}
		return nil
	})

	return files,err

}

func main(){
	directory := flag.String("dir",".","directory to search in")
	query := flag.String("query","","file name or extension to search for")
	ignoreCase := flag.Bool("ignore-case",false,"ignore case while searching")
	
	flag.Parse()

	files,err := searchFiles(*directory,*query,*ignoreCase)

	if err != nil{
		log.Fatal(err)
	}

	for _,file := range files {
		fmt.Println(file)
	}

}