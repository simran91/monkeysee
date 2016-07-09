package monkey

import "io/ioutil"
import "../util"

//
// LoadImageFromFile takes a filename and returns the contents of that file as a string...
//
func LoadImageFromFile(filename string) *Monkey {
	sliceOfBytes, err := ioutil.ReadFile(filename)
	util.CheckError(err)
	data := string(sliceOfBytes)
	monkey := &Monkey{}
	monkey.SetRawData(data)
	return monkey
}
