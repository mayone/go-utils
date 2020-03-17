package build

import (
	"fmt"
)

// These variables are set using -ldflags
var (
	name           string
	version        string
	gitBranch      string
	lastCommitSHA  string
	lastCommitTime string
	goVersion      string
)

// Info returns building setup info in string
func Info() string {
	return fmt.Sprintf(`
Name             : %v
Release          : %v
Commit SHA1      : %v
Commit Timestamp : %v
Branch           : %v
Go Version       : %v
`,
		name, version, lastCommitSHA, lastCommitTime, gitBranch, goVersion)
}
