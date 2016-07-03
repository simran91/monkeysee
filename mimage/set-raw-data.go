package mimage

//
// SetRawData is a exported function so that in just like "LoadImageFromFile" is calling it now,
// in the future we might call it from imagedata we might already have in memory...
// TODO: Longer term, i want to write some gimp plugins, and i suspect we can just get the data
//       from GIMP in memory and return it in memory, so we won't have to use
//       files / temporary files...
//
func (i *MImage) SetRawData(data string) {
	i.rawdata = data
}
