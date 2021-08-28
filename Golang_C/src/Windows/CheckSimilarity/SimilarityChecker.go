package CheckSimilarity

//#include "simil_check.h"
//import "C"
import (
	"io/ioutil"
	"os"
)

/*
 *  by Rei-chi Lin
 */

func Version() string {
	return "1.3.0"
}

func CheckSimilarity(f1 *os.File, f2 *os.File) (float64, float64, int) {
	b1, err := ioutil.ReadAll(f1)
	Check(err)
	b2, err := ioutil.ReadAll(f2)
	Check(err)

	//index := C.int(-1)
	index := -1
	count := -1
	diff := len(b1) - len(b2)

	//count = int(C.check_similarity(C.CString(string(b1)), &index, C.CString(string(b2)), C.int(len(b2)), C.int(diff)))
	length1 := len(b1)
	length2 := len(b2)
	maxCount := -1
	for startIndex := 0; startIndex <= diff; startIndex += 1 {
		count := 0
		for i := 0; i < length2; i += 1 {
			if b1[startIndex+i] == b2[i] {
				count += 1
			}
		}
		if count > maxCount {
			maxCount = count
			index = startIndex
		}
	}

	ratioToSmall := float64(count) * 100.0 / float64(length2)
	ratioToLarge := float64(count) * 100.0 / float64(length1)
	return ratioToSmall, ratioToLarge, int(index)
}

func Author() string {
	return "Rei-chi Lin"
}
