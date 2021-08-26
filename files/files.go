package files

import (
	"bufio"
	"errors"
	"io/fs"
	"organLib/paths"
	"os"
	"path/filepath"
	"strconv"
)

type File struct {
	Name  string
	Paths []string
	Size  []int64
}

func SearchFile(root *paths.RootDir, fileName string) (File, error) {
	var files File
	nonExisting := errors.New("error : the file you choose ( " + fileName + " ) doesn't exist")
	err := filepath.Walk(root.Root, func(path string, info fs.FileInfo, err error) error {
		if info.Name() == fileName {
			files.Paths = append(files.Paths, path)
			files.Size = append(files.Size, info.Size())
		}
		return nil
	},
	)

	files.Name = fileName

	if err != nil {
		return files, err
	}

	if len(files.Paths) <= 0 {
		panic(nonExisting)
	}

	return files, nil
}

func (file *File) CreateLogFile(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(f)
	defer f.Close()

	writer.WriteString("Paths in Root for " + file.Name + "\n\n")
	for i, str := range file.Paths {
		writer.WriteString("Size :> " + strconv.Itoa(int(file.Size[i])) + " | Path :> " + str + "\n")
	}

	writer.Flush()
	return nil
}

func DeleteFileInRoot(fileToDelete *File, root *paths.RootDir) error {
	err := os.Remove(fileToDelete.Name)
	if err != nil {
		return err
	}
	return nil
}
