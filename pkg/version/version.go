package version

import "runtime/debug"

var Version = "dev"

func String() string {
	if Version != "dev" {
		return Version
	}
	info, ok := debug.ReadBuildInfo()
	if ok && info.Main.Version != "(devel)" {
		return Version
	}
	return Version
}
