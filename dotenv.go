package dotenv

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type dotenv struct {
	files []string
	opts  *options
}

var LoadFailedErr = errors.New("[dotenv] failed to load files")

func (d *dotenv) Load(files ...string) error {
	var loadErr error
	parsedFiles := d.opts.ParseFilePaths(files...)
	for _, parsedFile := range parsedFiles {
		if err := loadFile(parsedFile, false); err != nil && d.opts.debug {
			if loadErr != nil {
				loadErr = fmt.Errorf("%s\n%w", loadErr.Error(), err)
			} else {
				loadErr = err
			}
			log.Println(fmt.Sprintf("[dotenv] Loading parsedFile %s failed with error %s", parsedFile, err.Error()))
			continue
		}
		d.files = append(d.files, parsedFile)
	}

	return loadErr
}

func loadFile(file string, overload bool) error {
	fileEnv, err := godotenv.Read(file)
	if err != nil {
		return err
	}

	osEnv := map[string]bool{}
	for _, rawEnvLine := range os.Environ() {
		key := strings.Split(rawEnvLine, "=")[0]
		osEnv[key] = true
	}

	for key, value := range fileEnv {
		if !osEnv[key] || overload {
			_ = os.Setenv(key, value)
		}
	}

	return nil
}

func (d *dotenv) Overload(files ...string) error {
	parsedFiles := d.opts.ParseFilePaths(files...)
	for _, parsedFile := range parsedFiles {
		if err := loadFile(parsedFile, true); err != nil && d.opts.debug {
			log.Println(fmt.Sprintf("[dotenv] Overloading parsedFile %s failed with error %s", parsedFile, err.Error()))
			continue
		}
		d.files = append(d.files, parsedFile)
	}
	return nil
}

func new() *dotenv {
	return &dotenv{opts: newOpts()}
}
