package dotenv

var d = new()

func Load(files ...string) error {
	return d.Load(files...)
}

func Reset() {
	d = new()
}

func Overload(files ...string) error {
	return d.Overload(files...)
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
	d.opts.lookupFile = file
}
