package v_utils

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"strings"
)


/* ============================================
	Created by andy pangaribuan on 2020/11/23
	Copyright BoltIdea. All rights reserved.
   ============================================ */
func (*filesStruct) GetFileExtension(filename string) string {
	arr := strings.Split(filename, ".")
	if len(arr) < 2 {
		return ""
	}
	return strings.ToLower(arr[len(arr)-1])
}


func (*filesStruct) GetFileName(filePath string, includeExt bool) string {
	arr := strings.Split(filePath, string(os.PathSeparator))
	if includeExt {
		return arr[len(arr)-1]
	}
	return strings.Split(arr[len(arr)-1], ".")[0]
}


func (slf *filesStruct) GetFilesRecursively(dirPath string, extensions []string, limit int, condition func(fileName, fullPath string) bool) (files []FileScanResult, err error) {
	files = make([]FileScanResult, 0)
	for i:=0; i<len(extensions); i++ {
		extensions[i] = strings.ToLower(extensions[i])
	}

	if len(dirPath) > 0 {
		if dirPath[len(dirPath)-1:] == string(os.PathSeparator) {
			dirPath = dirPath[:len(dirPath)-1]
		}
	}

	var readDir func(dirPath string) (err error)
	readDir = func(dirPath string) (err error) {
		fi, _err := ioutil.ReadDir(dirPath)
		if _err != nil {
			err = errors.WithStack(_err)
			return
		}

		dirs := make([]string, 0)
		for _, file := range fi {
			if len(files) >= limit {
				break
			}

			if file.IsDir() {
				path := dirPath + string(os.PathSeparator) + file.Name()
				dirs = append(dirs, path)
			}

			if !file.IsDir() && file.Size() > 0 {
				add := len(extensions) == 0
				fileName := file.Name()
				fullPath := dirPath + string(os.PathSeparator) + fileName

				for _, ext := range extensions {
					if ext == slf.GetFileExtension(fileName) {
						add = true
						break
					}
				}

				if add && condition != nil {
					add = condition(fileName, fullPath)
				}
				if add {
					files = append(files, FileScanResult{
						FileName: fileName,
						DirPath:  dirPath,
						FilePath: fullPath,
					})
				}
			}
		}

		if len(files) < limit {
			for _, dir := range dirs {
				_err := readDir(dir)
				if _err != nil {
					err = _err
					break
				}

				if len(files) >= limit {
					break
				}
			}
		}

		return
	}


	err = readDir(dirPath)






	//fi, _err := ioutil.ReadDir(dirPath)
	//if _err != nil {
	//	err = errors.WithStack(_err)
	//	return
	//}
	//
	//for _, file := range fi {
	//	if len(files) >= limit {
	//		break
	//	}
	//
	//	if !file.IsDir() && file.Size() > 0 {
	//		add := extension == ""
	//		fileName := file.Name()
	//		fullPath := dirPath + string(os.PathSeparator) + fileName
	//
	//		if extension != "" {
	//			if extension == slf.GetFileExtension(fileName) {
	//				add = true
	//			}
	//		}
	//
	//		if add && condition != nil {
	//			add = condition(fileName, fullPath)
	//		}
	//		if add {
	//			files = append(files, fullPath)
	//		}
	//	}
	//
	//	if file.IsDir() {
	//
	//	}
	//}

	return
}
