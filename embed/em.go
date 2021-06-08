package em

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

const folderPerm os.FileMode = 0755

var ErrExist error = os.ErrExist

func Tree(fsys embed.FS, outputPath string) error {
	return Walk(fsys, ".", func(dirpath string, de fs.DirEntry) error {
		fullpath := filepath.Join(outputPath, dirpath, de.Name())
		if de.IsDir() {
			return os.MkdirAll(fullpath, folderPerm)
		}
		return nil
	})
}

func Touch(fsys embed.FS, outputPath string) error {
	return Walk(fsys, ".", func(dirpath string, de fs.DirEntry) error {
		fullpath := filepath.Join(outputPath, dirpath, de.Name())
		if de.IsDir() {
			return os.MkdirAll(fullpath, folderPerm)
		}
		// unsure how IsNotExist works. this could be improved
		_, err := os.Stat(fullpath)
		if os.IsNotExist(err) {
			_, err = os.Create(fullpath)
		}
		return err
	})
}

func Write(fsys embed.FS, outputPath string) error {
	return Walk(fsys, ".", func(dirpath string, de fs.DirEntry) error {
		embedPath := sanitize(filepath.Join(dirpath, de.Name()))
		fullpath := filepath.Join(outputPath, embedPath)
		if de.IsDir() {
			return os.MkdirAll(fullpath, folderPerm)
		}
		return embedCopyToFile(fsys, embedPath, fullpath)
	})
}

func Patch(fsys embed.FS, outputPath string) error {
	return Walk(fsys, ".", func(dirpath string, de fs.DirEntry) error {
		embedPath := sanitize(filepath.Join(dirpath, de.Name()))
		fullpath := filepath.Join(outputPath, embedPath)
		if de.IsDir() {
			return os.MkdirAll(fullpath, folderPerm)
		}
		_, err := os.Stat(fullpath)
		if os.IsNotExist(err) {
			err = embedCopyToFile(fsys, embedPath, fullpath)
		}
		return err
	})
}

func Create(fsys embed.FS, outputPath string) error {
	err := Walk(fsys, ".", func(dirpath string, de fs.DirEntry) error {
		embedPath := filepath.Join(dirpath, de.Name())
		fullpath := filepath.Join(outputPath, embedPath)
		if de.IsDir() {
			return nil
		}
		_, err := os.Stat(fullpath)
		if os.IsNotExist(err) {
			return nil
		}
		if err != nil {
			return err
		}
		return ErrExist
	})
	if err != nil {
		return err
	}
	return Patch(fsys, outputPath)
}

func Walk(fsys embed.FS, startPath string, f func(path string, de fs.DirEntry) error) error {
	folders := make([]string, 0) // buffer of folders to process
	err := WalkDir(fsys, startPath, func(dirpath string, de fs.DirEntry) error {
		if de.IsDir() {
			folders = append(folders, filepath.Join(dirpath, de.Name()))
		}
		return f(dirpath, de)
	})
	if err != nil {
		if len(folders) == 0 {
			return fmt.Errorf("no folder found: %v", err)
		}
		return err
	}
	n := len(folders)
	for n != 0 {
		for i := 0; i < n; i++ {
			err = WalkDir(fsys, folders[i], func(dirpath string, de fs.DirEntry) error {
				if de.IsDir() {
					folders = append(folders, filepath.Join(dirpath, de.Name()))
				}
				return f(dirpath, de)
			})
			if err != nil {
				return err
			}
		}
		// we process n folders at a time, add new folders while
		//processing n folders, then discard those n folders once finished
		// and resume with a new n list of folders
		var newFolders int = len(folders) - n
		folders = folders[n : n+newFolders] // if found 0 new folders, end
		n = len(folders)
	}
	return nil
}

func WalkDir(fsys embed.FS, startPath string, f func(path string, de fs.DirEntry) error) error {
	startPath = sanitize(startPath)
	items, err := fsys.ReadDir(startPath)
	if err != nil {
		return err
	}
	for _, item := range items {
		if err := f(startPath, item); err != nil {
			return err
		}
	}
	return nil
}

func embedCopyToFile(fsys embed.FS, embedPath, path string) error {
	embedPath = sanitize(embedPath)
	fi, err := fsys.Open(embedPath)
	if err != nil {
		return fmt.Errorf("opening embedded file %v: %v", embedPath, err)
	}
	fo, err := os.Create(path)
	if err != nil {
		return err
	}
	_, err = io.Copy(fo, fi)
	return err
}

func sanitize(embedPath string) string {
	return strings.ReplaceAll(embedPath, "\\", "/")
}
