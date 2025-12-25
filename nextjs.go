package templates

import (
	"embed"
	"io/fs"
	"path/filepath"

	"github.com/xehrad/git"
)

//go:embed starter/*/*
var starterFS embed.FS

type ProjectStarter string

const (
	Nextjs ProjectStarter = "nextjs-starter"
)

// GetNextjsStarterFiles reads all files from the embedded FS and prepares them for Git upload
func GetStarterFiles(root ProjectStarter) ([]git.FileNode, error) {
	var files []git.FileNode

	err := fs.WalkDir(starterFS, string(root), func(path string, d fs.DirEntry, err error) error {
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

		// Remove the root prefix ("nextjs-starter/") from the path so it sits at root of repo
		relPath, _ := filepath.Rel(string(root), path)
		// Convert Windows paths to Git standard (forward slash)
		relPath = filepath.ToSlash(relPath)

		files = append(files, git.FileNode{
			Path:    relPath,
			Content: &strContent,
			Type:    git.FileTypeFile,
		})
		return nil
	})

	return files, err
}
