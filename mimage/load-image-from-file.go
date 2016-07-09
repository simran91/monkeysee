package mimage

import "io/ioutil"
import "../util"

//
// LoadImageFromFile takes a filename and returns the contents of that file as a string...
//
func LoadImageFromFile(filename string) *MImage {
	sliceOfBytes, err := ioutil.ReadFile(filename)
	util.CheckError(err)
	data := string(sliceOfBytes)
	mimage := &MImage{}
	mimage.SetRawData(data)
	return mimage
}
