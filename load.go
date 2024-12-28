package dotenv

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
	d.mu.Lock()
	d.opts.lookupGit = true
	d.mu.Unlock()
}

func OptLookupMod() {
	d.mu.Lock()
	d.opts.lookupMod = true
	d.mu.Unlock()
}

func OptDisableFileExpand() {
	d.mu.Lock()
	d.opts.disableFileExpand = true
	d.mu.Unlock()
}

func OptDisablePathExpand() {
	d.mu.Lock()
	d.opts.disablePathExpand = true
	d.mu.Unlock()
}

func OptDebug() {
	d.mu.Lock()
	d.opts.debug = true
	d.mu.Unlock()
}

func OptLookupFile(file string) {
	d.mu.Lock()
	d.opts.lookupFile = append(d.opts.lookupFile, file)
	d.mu.Unlock()
}
