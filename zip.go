package zipFile

import (
	"time"
	"archive/zip"
	"os"
)

type FileContent struct {
	Name string
	Content string
}

func ZipFolder(files []FileContent, dest string) error {

	zf, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer zf.Close()

	archive := zip.NewWriter(zf)
	defer archive.Close()

	// archive.RegisterCompressor(zip.Deflate, func(out io.Writer) (io.WriteCloser, error) {
	// 	return flate.NewWriter(out, flate.BestCompression)
	// })

	for _, file := range files {
		header := zip.FileHeader{
			Name : file.Name,
			Modified : time.Now(),
		}
	
		z, err := archive.CreateHeader(&header)
		if err != nil {
			return err
		}
		_, err = z.Write([]byte(file.Content))
		if err != nil {
			return err
		}
	
		
	}

	err = archive.Close()
	if err != nil {
		return err
	}
	

	return nil
}