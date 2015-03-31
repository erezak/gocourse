package c11file

import (
	"os"
)

func SumFilesSizes(folderName string) (int64, error) {
	var totalBytes int64
	folder, err := os.Open(folderName)
	if err != nil {
		return totalBytes, err
	}
	
	// Note the defered call to file.close
	defer folder.Close()

	files, err := folder.Readdir(0)
	if err != nil {
		return totalBytes, err
	}

	for _, file := range files {
		if !file.IsDir() {
			totalBytes += file.Size()
		}
	}

	return totalBytes, err
}
