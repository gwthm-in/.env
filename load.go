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
