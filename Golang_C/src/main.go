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
		tempF := f1
		f1 = f2
		f2 = tempF
		//tempFI := fi1
		//fi1 = fi2
		//fi2 = tempFI
	}

	fmt.Println("comparing ...\nThis process may take a few 'minutes'.\nPlease wait.\n")
	ratioToSmall, ratioToLarge, index := simich.CheckSimilarity(f1, f2)
	text := fmt.Sprintf("similarity:\n %v %% similar to the smaller one and\n %v %% similar to the larger one \n  at byte number %v in the larger one.\n", ratioToSmall, ratioToLarge, index+1)
	fmt.Println(text)
}
