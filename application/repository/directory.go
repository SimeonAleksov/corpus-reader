package repository

import (
	"errors"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"nu/corpus-reader/application/domain"
)

type DirectoryRepository interface {
	ListFiles(rootDirectory string) ([]string, error)
	GetFileContent(filePath string) (string, error)
}

type DirectoryRepositoryImplementation struct{}

func NewDirectoryRepository() *DirectoryRepositoryImplementation {
	return &DirectoryRepositoryImplementation{}
}

func (f *DirectoryRepositoryImplementation) ListFiles(rootDirectory string, exts []string) (*domain.RootDirectory, error) {
	if _, err := os.Stat(rootDirectory); errors.Is(err, os.ErrNotExist) {
		return nil, errors.New("Directory does not exists.")
	}

	var files []string
	err := filepath.WalkDir(rootDirectory, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		for _, s := range exts {
			if strings.HasSuffix(path, "."+s) {
				files = append(files, path)
				return nil
			}
		}
		return nil
	})
	return &domain.RootDirectory{
		Files: files,
	}, err
}

func (f *DirectoryRepositoryImplementation) GetFileContent(filePath string) (string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
