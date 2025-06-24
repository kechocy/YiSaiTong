package main

import (
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"

	// "github.com/duke-git/lancet/v2/strutil"
	"github.com/gen2brain/dlgs"
)

var count int = 0
var root = flag.String("d", ".", "待解密文件所在根目录")

func main() {
	// path, _ := os.Executable()
	// _, selfName := filepath.Split(path)
	flag.Parse()

	// 校验目录是否存在
	info, err := os.Stat(*root)
	if err != nil {
		dlgs.Error("错误", fmt.Sprintf("目录错误：%v", *root))
		os.Exit(1)
	}
	if !info.IsDir() {
		dlgs.Error("错误", fmt.Sprintf("指定路径不是目录：%v", *root))
		os.Exit(1)
	}

	// str, _ := os.Getwd()
	allFile, _ := getAllFileIncludeSubFolder(*root)
	for _, path := range allFile {
		// if strutil.AfterLast(path, "Unlock.exe") == "" {
		// 	continue
		// }
		// if strutil.AfterLast(path, selfName) == "" {
		// 	continue
		// }
		ext := strings.ToLower(filepath.Ext(path))

		if ext == ".pdf" || ext == ".doc" || ext == ".docx" || ext == ".xls" || ext == ".xlsx" || ext == ".ppt" || ext == ".pptx" {
			dstFilePath := path + ".temp"
			copyFile(path, dstFilePath)
			err := os.Remove(path)
			if err != nil {
				dlgs.Error("错误", fmt.Sprintf("文件 %v 未执行成功", path))
				continue
			}
			renameFile(dstFilePath, path)
		}
	}
	dlgs.Info("提示", fmt.Sprintf("自动解密完成 %v 项！", count))
}

func renameFile(sourcePath, dstFilePath string) {
	str, _ := os.Getwd()
	unlockPath := filepath.Join(str, "Unlock.exe")
	arg := fmt.Sprintf(` -source="%v" -dest="%v"`, sourcePath, dstFilePath)
	cmd := exec.Command(unlockPath)
	cmd.SysProcAttr = &syscall.SysProcAttr{CmdLine: "/c" + arg}
	output, err := cmd.Output()
	if err != nil {
		dlgs.Error("错误", fmt.Sprintf("%v \n文件重命名失败，可能是 Unlock.exe 不存在", err))
		dlgs.Warning("警告", fmt.Sprintf("请手动将 %v 重命名为 %v", sourcePath, dstFilePath))
	} else {
		info := string(output)
		if info != "" {
			dlgs.Info("提示", string(output))
		}
		count++
	}
}

func copyFile(sourcePath, dstFilePath string) (err error) {
	source, _ := os.Open(sourcePath)
	destination, _ := os.OpenFile(dstFilePath, os.O_CREATE|os.O_WRONLY, 0666)
	defer source.Close()
	defer destination.Close()
	buf := make([]byte, 1024)
	for {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		if _, err := destination.Write(buf[:n]); err != nil {
			return err
		}
	}
	return nil
}

// getAllFileIncludeSubFolder
//
//	@Description: 获取目录下所有文件（包含子目录）
//	@param folder
//	@return []string
//	@return error
func getAllFileIncludeSubFolder(folder string) ([]string, error) {
	var result []string
	filepath.WalkDir(folder, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			dlgs.Error("错误", err.Error())
			return err
		}
		if !d.IsDir() {
			result = append(result, path)
		}
		return nil
	})
	return result, nil
}
