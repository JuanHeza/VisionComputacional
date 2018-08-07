package main

import (
    "archive/zip"
    "fmt"
    "io"
    "log"
    "os"
)

func main() {

    // Files to Zip
    files := []string{"example.csv", "data.csv"}
    output := "done.zip"

    err := ZipFiles(output, files)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Zipped File: " + output)
}

// ZipFiles compresses one or many files into a single zip archive file
func ZipFiles(filename string, files []string) error {

    newfile, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer newfile.Close()

    zipWriter := zip.NewWriter(newfile)
    defer zipWriter.Close()

    // Add files to zip
    for _, file := range files {

        zipfile, err := os.Open(file)
        if err != nil {
            return err
        }
        defer zipfile.Close()

        // Get the file information
        info, err := zipfile.Stat()
        if err != nil {
            return err
        }

        header, err := zip.FileInfoHeader(info)
        if err != nil {
            return err
        }

        // Change to deflate to gain better compression
        // see http://golang.org/pkg/archive/zip/#pkg-constants
        header.Method = zip.Deflate

        writer, err := zipWriter.CreateHeader(header)
        if err != nil {
            return err
        }
        _, err = io.Copy(writer, zipfile)
        if err != nil {
            return err
        }
    }
    return nil
}

====================================

package main

import (
    "io"
    "net/http"
    "os"
)

func main() {

    fileUrl := "https://golangcode.com/images/avatar.jpg"

    err := DownloadFile("avatar.jpg", fileUrl)
    if err != nil {
        panic(err)
    }

}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

    // Create the file
    out, err := os.Create(filepath)
    if err != nil {
        return err
    }
    defer out.Close()

    // Get the data
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    // Write the body to file
    _, err = io.Copy(out, resp.Body)
    if err != nil {
        return err
    }

    return nil
}

==================================

func zipit(source, target string) error {
	zipfile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	info, err := os.Stat(source)
	if err != nil {
		return nil
	}

	var baseDir string
	if info.IsDir() {
		baseDir = filepath.Base(source)
	}

	filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		if baseDir != "" {
			header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, source))
		}

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(writer, file)
		return err
	})

	return err
}

func unzip(archive, target string) error {
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}

	for _, file := range reader.File {
		path := filepath.Join(target, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			return err
		}
	}

	return nil
}

zipit("/tmp/documents", "/tmp/backup.zip")
zipit("/tmp/report.txt", "/tmp/report-2015.zip")
unzip("/tmp/report-2015.zip", "/tmp/reports/")

=================================================

package main

import (
    "archive/zip"
    "fmt"
    "io"
    "log"
    "os"
    "path/filepath"
    "strings"
)

func main() {

    files, err := Unzip("done.zip", "output")

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Unzipped: " + strings.Join(files, ", "))

}

// Unzip will un-compress a zip archive,
// moving all files and folders to an output directory
func Unzip(src, dest string) ([]string, error) {

    var filenames []string

    r, err := zip.OpenReader(src)
    if err != nil {
        return filenames, err
    }
    defer r.Close()

    for _, f := range r.File {

        rc, err := f.Open()
        if err != nil {
            return filenames, err
        }
        defer rc.Close()

        // Store filename/path for returning and using later on
        fpath := filepath.Join(dest, f.Name)
        filenames = append(filenames, fpath)

        if f.FileInfo().IsDir() {

            // Make Folder
            os.MkdirAll(fpath, os.ModePerm)

        } else {

            // Make File
            var fdir string
            if lastIndex := strings.LastIndex(fpath, string(os.PathSeparator)); lastIndex > -1 {
                fdir = fpath[:lastIndex]
            }

            err = os.MkdirAll(fdir, os.ModePerm)
            if err != nil {
                log.Fatal(err)
                return filenames, err
            }
            f, err := os.OpenFile(
                fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
            if err != nil {
                return filenames, err
            }
            defer f.Close()

            _, err = io.Copy(f, rc)
            if err != nil {
                return filenames, err
            }

        }
    }
    return filenames, nil
}

=============================================

package main

import (
  "fmt"
)

func main() {

  messages := make(chan string)

  go work(messages)

  msg := <-messages
  fmt.Println(msg)
}

func work(messages chan<- string) {
  messages <- "done"
}

================================================