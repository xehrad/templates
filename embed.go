package templates

import (
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/xehrad/git"
)

//go:embed starter
var starterFS embed.FS
var ErrNotValidStarter = errors.New("Err Not Valid Starter")

const (
	_ROOT  = "starter"
	Nextjs = "nextjs"
)

// GetStarterFiles reads all files from the embedded FS and prepares them for Git upload
func GetStarterFiles(name string) ([]git.FileNode, error) {
	if !isValidStarter(name) {
		return nil, ErrNotValidStarter
	}

	var files []git.FileNode
	starterPath := fmt.Sprintf("%s/%s", _ROOT, name)
	err := fs.WalkDir(starterFS, starterPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		content, err := starterFS.ReadFile(path)
		if err != nil {
			return err
		}
		strContent := string(content)

		// Remove the root prefix ("starter/nextjs") from the path so it sits at root of repo
		relPath, err := filepath.Rel(starterPath, path)
		if err != nil {
			return err
		}

		files = append(files, git.FileNode{
			Path:    relPath,
			Content: &strContent,
			Type:    git.FileTypeFile,
		})
		return nil
	})

	return files, err
}

func isValidStarter(name string) bool {
	switch name {
	case Nextjs:
		return true
	default:
		return false
	}
}
