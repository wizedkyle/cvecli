package build

var (
	Build   = "Dev"
	Date    = "Dev"
	Version = "Dev"
)

func GetVersion() string {
	return Version + " " + Build + " " + Date
}
