package dotenv

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/joho/godotenv"
)

var NoFailsToLoadErr = errors.New("[dotenv] No files to load")

type dotenv struct {
	mu sync.Mutex

	files []string
	opts  *options
}

func (d *dotenv) Load() error {
	var loadErr error
	parsedFiles := d.opts.ParseFilePaths()
	if len(parsedFiles) == 0 {
		d.logf("[dotenv] No files found to load")
		return NoFailsToLoadErr
	}
	d.logf("[dotenv] Loading parsedFiles %s", strings.Join(parsedFiles, ", "))

	for _, parsedFile := range parsedFiles {
		if err := loadFile(parsedFile, false); err != nil {
			loadErr = wrapError(loadErr, err)
			d.logf("[dotenv] Loading parsedFile %s failed with error %s", parsedFile, err.Error())
			continue
		}
		d.logf("[dotenv] Loaded parsedFile %s", parsedFile)
		d.files = append(d.files, parsedFile)
	}
	d.logf("[dotenv] Loaded parsedFiles %s", strings.Join(d.files, ", "))
	return loadErr
}

func (d *dotenv) logf(args ...interface{}) {
	if d.opts.debug {
		log.Println(fmt.Sprintf(args[0].(string), args[1:]...))
	}
}

func wrapError(loadErr error, err error) error {
	if loadErr != nil {
		loadErr = fmt.Errorf("%s\n%w", loadErr.Error(), err)
	} else {
		loadErr = err
	}
	return loadErr
}

func loadFile(file string, overload bool) error {
	fileEnv, err := godotenv.Read(file)
	if err != nil {
		d.logf("[dotenv] Loading file %s failed with error %s", file, err.Error())
		return err
	}
	d.logf("[dotenv] Loading file %s", file)

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

func (d *dotenv) Overload() error {
	parsedFiles := d.opts.ParseFilePaths()
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
