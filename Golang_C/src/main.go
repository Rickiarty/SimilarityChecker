package main

/*
 *  by Rei-chi Lin
 */

import (
	"fmt"
	"os"
	"path/filepath"

	simich "./CheckSimilarity"
)

func main() {
	helpInfo := "Pass the program 2 arguments which indicate paths to 2 files to compare similarity(%)."
	if len(os.Args) == 2 {
		if os.Args[1] == "--version" {
			verInfo := fmt.Sprintf("ver. %v by %v", simich.Version(), simich.Author())
			fmt.Println(verInfo)
			return
		} else {
			fmt.Println(helpInfo)
			return
		}
	} else if len(os.Args) < 2 {
		fmt.Println(helpInfo)
		return
	}

	f1Path, _ := filepath.Abs(os.Args[1])
	f1, err := os.Open(f1Path)
	simich.Check(err)
	defer f1.Close()
	f2Path, _ := filepath.Abs(os.Args[2])
	f2, err := os.Open(f2Path)
	simich.Check(err)
	defer f2.Close()

	fi1, _ := f1.Stat()
	fi2, _ := f2.Stat()
	if fi1.Size() < fi2.Size() {
		// swap
		temp := f1
		f1 = f2
		f2 = temp
	}

	fmt.Println("comparing ...\nPlease wait.\n")
	ratio, index := simich.CheckSimilarity(f1, f2)
	text := fmt.Sprintf("similarity: %v %% at index %v.", ratio, index)
	fmt.Println(text)
}
