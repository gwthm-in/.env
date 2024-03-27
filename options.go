package dotenv

import (
	"os"
	"os/exec"
	"path/filepath"
)

type options struct {
	lookupMod bool // look up for go.mod file, by default false
	lookupGit bool // look up for .git directory, by default false

	lookupFile  string   // file type of .env file, by default .env, ex: .env.test
	lookupPaths []string // look up for .env file in these paths, by default the current directory

	disableFileExpand bool // disable expanding lookupFile to find .env.${ENVIRONMENT} files, by default false
	disablePathExpand bool // disable expanding lookupPaths to find .env file, by default false
	debug             bool // enable debug mode, by default false
}

func (o *options) FilesOrDefault(files ...string) []string {
	if len(files) == 0 {
		files = append(files, o.lookupFile)
	}
	return files
}

func (o *options) ParseFilePaths(files ...string) []string {
	var parsedFiles []string
	files = o.FilesOrDefault(files...)
	d.logf("[dotenv] Default files to parse: %+s", files)
	for _, file := range files {
		parsedFiles = append(parsedFiles, o.ParseFilePath(file))
	}
	d.logf("[dotenv] Parsed files: %s", parsedFiles)

	if !filepath.IsAbs(o.lookupFile) {
		if o.lookupGit {
			if repoPath := gitRepoPath(); repoPath != "" {
				parsedFiles = append(parsedFiles, o.ParseFilePath(repoPath))
			}
			d.logf("[dotenv] Parsed files after git lookup: %s", parsedFiles)
		}

		if o.lookupMod {
			if modPath := modPath(); modPath != "" {
				parsedFiles = append(parsedFiles, o.ParseFilePath(modPath))
			}
			d.logf("[dotenv] Parsed files after mod lookup: %s", parsedFiles)
		}
	}

	return filterValidFiles(parsedFiles)
}

func filterValidFiles(files []string) []string {
	validFilePaths := []string{}
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
		return filepath.Dir(string(bytes))
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

func (o *options) ParseFilePath(file string) string {
	if isDirectory(file) {
		file = filepath.Join(file, o.lookupFile)
	}

	if o.disableFileExpand {
		return file
	}

	return os.Expand(file, os.Getenv)
}

func isDirectory(file string) bool {
	fInfo, err := os.Stat(file)
	if err != nil {
		return false
	}

	return fInfo.IsDir()
}

func newOpts() *options {
	return &options{
		lookupFile:  ".env",
		lookupPaths: []string{"./"},
	}
}
