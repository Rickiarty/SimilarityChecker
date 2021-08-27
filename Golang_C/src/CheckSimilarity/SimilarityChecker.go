package CheckSimilarity

//#include "simil_check.h"
import "C"
import (
	"io/ioutil"
	"os"
)

/*
 *  by Rei-chi Lin
 */

func Version() string {
	return "1.2.0"
}

func CheckSimilarity(f1 *os.File, f2 *os.File) (float64, int) {
	b1, err := ioutil.ReadAll(f1)
	Check(err)
	b2, err := ioutil.ReadAll(f2)
	Check(err)

	index := C.int(-1)
	maxRatio := -1.0
	diff := len(b1) - len(b2)
	maxRatio = float64(C.check_similarity(C.CString(string(b1)), &index, C.CString(string(b2)), C.int(len(b2)), C.int(diff)))
	return maxRatio, int(index)
}

func Author() string {
	return "Rei-chi Lin"
}
