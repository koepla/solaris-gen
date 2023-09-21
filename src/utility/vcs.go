package utility

import "runtime/debug"

// CommitVersion tries to obtain the version control system revision of the Go binary
func CommitVersion() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, settingsEntry := range info.Settings {
			if settingsEntry.Key == "vcs.revision" {
				return settingsEntry.Value
			}
		}
	}
	return "unknown"
}
