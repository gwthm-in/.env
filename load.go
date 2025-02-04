package dotenv

import "github.com/fsnotify/fsnotify"

var d = new()

func Load() error {
	return d.Load()
}

func Reset() {
	d = new()
}

func Overload() error {
	return d.Overload()
}

func OptLookupGit() {
	d.opts.lookupGit = true
}

func OptLookupMod() {
	d.opts.lookupMod = true
}

func OptDisableFileExpand() {
	d.opts.disableFileExpand = true
}

func OptDisablePathExpand() {
	d.opts.disablePathExpand = true
}

func OptDebug() {
	d.opts.debug = true
}

func OptLookupFile(file string) {
	d.opts.lookupFile = append(d.opts.lookupFile, file)
}

func OptDynamicLookupFile(file string) {
	d.opts.dynamicLookupFile = append(d.opts.dynamicLookupFile, file)
}

func WatchConfig() {
	d.opts.watchConfig = true
	files := d.opts.ParseFilePaths()
	for _, file := range files {
		go d.WatchConfig(file)
	}
}

func OnConfigChange(fn func(fsnotify.Event)) {
	d.opts.onConfigChange = fn
}
