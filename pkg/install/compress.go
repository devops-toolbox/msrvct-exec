package install

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/ulikunitz/xz"
)

func CompressTar(src, dst string) (err error) {

	tarFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer tarFile.Close()

	tr := tar.NewReader(tarFile)
	// os.Chdir(dst)
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		target := filepath.Join(dst, header.Name)
		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(target, 0755); err != nil {
				return err
			}
		case tar.TypeReg:
			ow, err := os.OpenFile(target, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.FileMode(header.Mode))
			if err != nil {
				return err
			}
			if _, err := io.Copy(ow, tr); err != nil {
				return err
			}
			ow.Close()
		case tar.TypeSymlink:
			if err := os.Symlink(header.Linkname, target); err != nil {
				return err
			}
		default:
			fmt.Printf("Can't: %c, %s\n", header.Typeflag, target)
		}
	}
	return nil
}
func CompressTarGz(src, dst string) (err error) {

	tarFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer tarFile.Close()

	gzr, err := gzip.NewReader(tarFile)
	if err != nil {
		return err
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)
	// os.Chdir(dst)
	// wd, _ := os.Getwd()
	// fmt.Println(wd)
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		target := filepath.Join(dst, header.Name)
		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(target, 0755); err != nil {
				return err
			}
		case tar.TypeReg:
			ow, err := os.OpenFile(target, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.FileMode(header.Mode))
			if err != nil {
				return err
			}
			if _, err := io.Copy(ow, tr); err != nil {
				return err
			}
			ow.Close()
		case tar.TypeSymlink:
			if err := os.Symlink(header.Linkname, target); err != nil {
				return err
			}
		default:
			fmt.Printf("Can't: %c, %s\n", header.Typeflag, target)
		}
	}
	return nil
}

func CompressTarXz(src, dst string) (err error) {

	tarFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer tarFile.Close()

	xzr, err := xz.NewReader(tarFile)
	if err != nil {
		return err
	}
	// defer xzr.Close()

	tr := tar.NewReader(xzr)
	// err = os.Chdir(dst)
	// if err != nil {
	// 	return err
	// }
	// wd, _ := os.Getwd()
	// fmt.Println(wd)
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		target := filepath.Join(dst, header.Name)
		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(target, os.FileMode(header.Mode)); err != nil {
				return err
			}
		case tar.TypeReg:
			ow, err := os.OpenFile(target, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.FileMode(header.Mode))
			if err != nil {
				return err
			}
			if _, err := io.Copy(ow, tr); err != nil {
				return err
			}
			ow.Close()
		case tar.TypeSymlink:
			if err := os.Symlink(header.Linkname, target); err != nil {
				return err
			}
		default:
			fmt.Printf("Can't: %c, %s\n", header.Typeflag, target)
		}
	}
	return nil
}
