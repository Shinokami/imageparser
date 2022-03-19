package loader

type image struct {
	name      string
	url       string
	localDir  string
	localPath string
}

type imageList []image

func (images *imageList) Paths() []byte {
	data := []byte{}
	for _, image := range *images {
		data = append(data, image.localPath+"\n"...)
	}
	return data
}
