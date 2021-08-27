package CheckSimilarity

//#include "simil_check.h"
import "C"

/*
 *  by Rei-chi Lin
 */

import (
	"io/ioutil"
	"os"
)

func Version() string {
	return "1.0.1"
}

func CheckSimilarity(f1 *os.File, f2 *os.File) (float64, int) {
	b1, err := ioutil.ReadAll(f1)
	Check(err)
	b2, err := ioutil.ReadAll(f2)
	Check(err)

	index := -1
	maxRatio := -1.0
	diff := len(b1) - len(b2)
	for startIndex := 0; startIndex <= diff; startIndex += 1 {
		ratio := float64(C.check_similarity(C.CString(string(b1)), C.int(startIndex), C.CString(string(b2)), C.int(len(b2))))
		if ratio > maxRatio {
			maxRatio = ratio
			index = startIndex
		}
	}
	return maxRatio, index
}

func Author() string {
	return "Rei-chi Lin"
}
