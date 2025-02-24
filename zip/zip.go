
package ziper

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
)

func Zipper() {
	// Open a zip archive for reading.
	r, err := zip.OpenReader("zip/testdata/readme.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	for _, f := range r.File {
		fmt.Printf("Contents of %s:\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.CopyN(os.Stdout, rc, 68)
		if err != nil {
			log.Fatal(err)
		}
		rc.Close()
		fmt.Println()
	}
}

func main() {
	Zipper()
}
