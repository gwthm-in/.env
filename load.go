package dotenv

var d = new(newOpts())

func Load(files ...string) error {
	return d.Load(files...)
}

func Reset() {
	d = new(newOpts())
}

func Overload(files ...string) error {
	return d.Overload(files...)
}
