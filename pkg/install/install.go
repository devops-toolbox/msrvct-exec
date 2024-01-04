package install

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/devops-toolbox/msrvct-exec/global"
	"github.com/devops-toolbox/msrvct-exec/model"
	"github.com/spf13/cast"
)

func ExecPreScript() (err error) {
	log.Println("ExecPreScript...")
	log.Println(global.Config.Script.PreScripts)
	for _, i := range global.Config.Script.PreScripts {
		resultByte, err := ExecScript(i)
		if err != nil {
			log.Println(err, string(resultByte))
			return err
		}
		log.Println(string(resultByte))
	}
	log.Println("ExecPreScript Done")
	return nil
}

func ExecPostScript() (err error) {
	log.Println("ExecPostScript...")
	// log.Println(global.Config.Script.PostScripts)
	for _, i := range global.Config.Script.PostScripts {
		resultByte, err := ExecScript(i)
		if err != nil {
			log.Println(err, string(resultByte))
			return err
		}
		log.Println(string(resultByte))
	}
	log.Println("ExecPostScript Done")
	return nil
}

func Install() (err error) {
	for _, p := range global.Config.Packages {
		for _, f := range p.Documents {
			// log.Println(p, f)
			if f.Template {
				err = HandleTemplate(p, f)
				if err != nil {
					log.Println(err)
					return err
				}
			}
			switch f.Type {
			case "pkg":
				err = HandlePackage(p, f)
				if err != nil {
					log.Println(err)
					return err
				}
			case "env":
				err = HandleEnvironment(p, f)
				if err != nil {
					log.Println(err)
					return err
				}
			}
		}
	}
	return nil
}
func ExecScript(script string) (result []byte, err error) {
	r := exec.Command("/bin/bash", fmt.Sprintf("%s/%s/%s", global.RuntimeVariableMap["src"], global.RuntimeVariableMap["src_scr_dir"], script))
	return r.CombinedOutput()
}

func HandleTemplate(p model.Package, f model.Document) (err error) {
	global.RuntimeVariableMap["local_name"] = p.Name
	global.RuntimeVariableMap["local_version"] = p.Version
	srcPath := global.RuntimeVariableMap["src"]
	tmpPath := global.RuntimeVariableMap["src"]
	srcFileName := fmt.Sprintf("%s.%s", f.Name, global.RuntimeVariableMap["glo_tpl_fix"])
	srcFile := filepath.Join(srcPath, global.RuntimeVariableMap["src_tpl_dir"], srcFileName)
	tmpFile := filepath.Join(tmpPath, global.RuntimeVariableMap["src_pkg_dir"], f.Name)
	if srcPath != tmpPath {
		err = MakeDirAll(tmpPath, global.RuntimeVariableMap["glo_dir_per"])
		if err != nil {
			log.Println(err)
			return err
		}
	}
	t, err := template.ParseFiles(srcFile)
	if err != nil {
		log.Println(err)
		return err
	}
	file, err := os.OpenFile(tmpFile, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Println(err)
		return err
	}
	defer file.Close()
	err = t.Execute(file, global.RuntimeVariableMap)
	if err != nil {
		log.Println(err)
		return err
	}
	// err = t.Execute(os.Stdout, global.RuntimeVariableMap)
	// if err != nil {
	// 	log.Println(err)
	// 	return err
	// }
	return nil
}
func MakeDirAll(path string, perm string) error {
	p := fs.FileMode(cast.ToUint32(perm))
	return os.MkdirAll(path, p)
}
func HandlePackage(p model.Package, f model.Document) (err error) {
	srcPath := global.RuntimeVariableMap["src"]
	srcFile := filepath.Join(srcPath, global.RuntimeVariableMap["src_pkg_dir"], f.Name)
	dstPath := f.Path
	if dstPath == "" {
		dstPath = filepath.Join(global.RuntimeVariableMap["dst"], global.RuntimeVariableMap["dst_pkg_dir"], p.Name, p.Version)
	}
	if f.PathMode == "" {
		f.PathMode = global.RuntimeVariableMap["glo_dir_per"]
	}
	err = MakeDirAll(dstPath, f.PathMode)
	if err != nil {
		log.Println(err)
		return err
	}
	kindSlice := strings.Split(f.Name, ".")
	kind := kindSlice[len(kindSlice)-1]
	// log.Println(kind)
	switch kind {
	case "gz":
		err = CompressTarGz(srcFile, dstPath)
		if err != nil {
			log.Println(err)
			return err
		}
	case "xz":
		err = CompressTarXz(srcFile, dstPath)
		if err != nil {
			log.Println(err)
			return err
		}
	case "tar":
		err = CompressTar(srcFile, dstPath)
		if err != nil {
			log.Println(err)
			return err
		}
	default:
		return errors.New("the packge type is not support")
	}
	if f.SubDir != "" {
		err = Rename(fmt.Sprintf("%s/%s", dstPath, f.SubDir), fmt.Sprintf("%s/%s", dstPath, p.Name))
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}
func Rename(src, dst string) (err error) {
	return os.Rename(src, dst)

}

func HandleEnvironment(p model.Package, f model.Document) (err error) {
	srcPath := global.RuntimeVariableMap["src"]
	srcFile := filepath.Join(srcPath, global.RuntimeVariableMap["src_pkg_dir"], f.Name)
	dstPath := f.Path
	if dstPath == "" {
		dstPath = filepath.Join(global.RuntimeVariableMap["dst"], global.RuntimeVariableMap["dst_pkg_dir"], p.Name, p.Version)
	}
	dstFile := filepath.Join(dstPath, f.Name)
	err = MakeDirAll(dstPath, global.RuntimeVariableMap["glo_dir_per"])
	if err != nil {
		log.Println(err)
		return err
	}
	err = CopyFile(srcFile, dstFile)
	if err != nil {
		log.Println(err)
		return err
	}
	if f.PathMode == "" {
		f.PathMode = global.RuntimeVariableMap["glo_dir_per"]
	}
	err = Chmod(dstPath, f.PathMode)
	if err != nil {
		log.Println(err)
		return err
	}
	if f.FileMode == "" {
		f.FileMode = global.RuntimeVariableMap["glo_doc_per"]
	}
	err = Chmod(fmt.Sprintf("%s/%s", dstPath, f.Name), f.FileMode)
	if err != nil {
		log.Println(err)
		return err
	}
	// dstLinkPath := filepath.Join(global.RuntimeVariableMap["dst"], global.RuntimeVariableMap["dst_env_dir"])
	// err = MakeDirAll(dstLinkPath, global.RuntimeVariableMap["glo_dir_per"])
	// if err != nil {
	// 	log.Println(err)
	// 	return err
	// }
	// dstLinkFile := filepath.Join(dstLinkPath, f.Name)
	// err = SymlinkCover(dstFile, dstLinkFile)
	// if err != nil {
	// 	log.Println(err)
	// 	return err
	// }
	return nil
}
func CopyFile(src, dst string) (err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}
	return nil
}
func Chmod(name string, mode string) error {
	return os.Chmod(name, fs.FileMode(cast.ToUint32(mode)))
}
func Symlink(oldname string, newname string) error {

	return os.Symlink(oldname, newname)
}
func SymlinkCover(oldname string, newname string) error {
	if _, err := os.Lstat(newname); err == nil {
		if err := os.Remove(newname); err != nil {
			return err
		}
	}
	return os.Symlink(oldname, newname)
}
