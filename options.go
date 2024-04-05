package dotenv

import (
	"os"
	"os/exec"
	"path/filepath"
)

const defaultConfigFile = ".env"

type options struct {
	lookupMod bool // look up for go.mod file, by default false
	lookupGit bool // look up for .git directory, by default false

	lookupFile  []string // file type of .env file, by default .env, ex: .env.test
	lookupPaths []string // look up for .env file in these paths, by default the current directory

	disableFileExpand bool // disable expanding lookupFile to find .env.${ENVIRONMENT} files, by default false
	disablePathExpand bool // disable expanding lookupPaths to find .env file, by default false
	debug             bool // enable debug mode, by default false
}

func (o *options) FilesOrDefault() []string {
	if len(o.lookupFile) == 0 {
		return []string{defaultConfigFile}
	}
	return o.lookupFile
}

// ParseFilePaths parses the given files and returns the absolute path of the files
func (o *options) ParseFilePaths() []string {
	var parsedFiles []string
	files := o.FilesOrDefault()
	d.logf("[dotenv] Files to parse: %+v", files)

	for _, file := range files {
		for _, path := range o.lookupPaths {
			if filepath.IsAbs(file) {
				parsedFiles = append(parsedFiles, file)
				continue
			}

			fullPath := filepath.Join(path, file)
			if _, err := os.Stat(fullPath); os.IsNotExist(err) {
				d.logf("[dotenv] File does not exist: %s", fullPath)
				continue
			}

			if o.lookupGit {
				repoPath := gitRepoPath()
				if repoPath != "" {
					fullPath = filepath.Join(repoPath, file)
				}
			}

			if o.lookupMod {
				modPath := modPath()
				if modPath != "" {
					fullPath = filepath.Join(modPath, file)
				}
			}

			parsedFiles = append(parsedFiles, fullPath)
		}
	}

	d.logf("[dotenv] Parsed files: %s", parsedFiles)
	return filterValidFiles(parsedFiles)
}

func filterValidFiles(files []string) []string {
	var validFilePaths []string
	for _, file := range files {
		if _, err := os.Stat(file); err == nil {
			validFilePaths = append(validFilePaths, file)
		}
	}
	return validFilePaths
}

func gitRepoPath() string {
	bytes, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err == nil {
		return string(bytes)
	}
	return ""
}

func modPath() string {
	bytes, err := exec.Command("go", "env", "GOMOD").Output()
	if err == nil {
		return filepath.Dir(string(bytes))
	}
	return ""
}

func newOpts() *options {
	return &options{
		lookupFile:  []string{},
		lookupPaths: []string{"./"},
	}
}
