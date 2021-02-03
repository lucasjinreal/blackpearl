package main

// "./picbed"

func UploadFromLocal(imgFilePath string) string {
	// if imgFilePath != "" {
	// 	imgBytes, err := ioutil.ReadFile(imgFilePath)
	// 	if(err != nil) {
	// 		fmt.Println(err)
	// 	}
	// 	// fmt.Println(imgBytes)

	// 	imgParam := picbed.ImageParam{
	// 		Name:    filepath.Base(imgFilePath),
	// 		Type:    "jpg",
	// 		Content: &imgBytes,
	// 	}

	// 	client := picbed.Ali{FileLimit: nil, MaxSize: 5024}
	// 	res, _ := client.Upload(&imgParam)
	// 	return res.Url
	// } else {
	// 	fmt.Println("provide an image path.")
	// 	return ""
	// }
}
