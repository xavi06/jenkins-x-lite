package util

import (
	"archive/tar"
	"compress/flate"
	"compress/gzip"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/mholt/archiver"
)

// DownloadFile func
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

// DeCompress func
func DeCompress(tarFile, dest string) error {
	srcFile, err := os.Open(tarFile)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	gr, err := gzip.NewReader(srcFile)
	if err != nil {
		return err
	}
	defer gr.Close()
	tr := tar.NewReader(gr)
	for {
		hdr, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		filename := dest + hdr.Name
		file, err := createFile(filename)
		if err != nil {
			return err
		}
		io.Copy(file, tr)
	}
	return nil
}

func createFile(name string) (*os.File, error) {
	err := os.MkdirAll(string([]rune(name)[0:strings.LastIndex(name, "/")]), 0755)
	if err != nil {
		return nil, err
	}
	return os.Create(name)
}

// ReadFile func
func ReadFile(name string) (string, error) {
	b, err := ioutil.ReadFile(name)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// ReadFileBytes func
func ReadFileBytes(name string) ([]byte, error) {
	b, err := ioutil.ReadFile(name)
	if err != nil {
		return []byte{}, err
	}
	return b, nil
}

// WriteFileString func
func WriteFileString(fname, content string) error {
	f, err := os.Create(fname)
	defer f.Close()
	if err != nil {
		return err
	}
	_, err = f.Write([]byte(content))
	return err
}

// WriteFileBytes func
func WriteFileBytes(fname string, content []byte) error {
	f, err := os.Create(fname)
	defer f.Close()
	if err != nil {
		return err
	}
	_, err = f.Write(content)
	return err
}

// Compress func
func Compress(fname, path string) error {
	/*
		z := archiver.Zip{
			CompressionLevel:       flate.DefaultCompression,
			MkdirAll:               true,
			SelectiveCompression:   true,
			ContinueOnError:        false,
			OverwriteExisting:      false,
			ImplicitTopLevelFolder: false,
		}
	*/
	tar := &archiver.Tar{
		OverwriteExisting: true,
		MkdirAll:          true,
	}
	z := archiver.TarGz{
		Tar:              tar,
		CompressionLevel: flate.DefaultCompression,
	}
	err := z.Archive([]string{path}, fname)
	return err
}

// CheckFileExsit func
func CheckFileExsit(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
