package cmd

import (
	"errors"
	"os"
	"runtime/debug"
	"strings"

	"github.com/ProtonMail/gopenpgp/v2/constants"
)

const VERSION = "0.0.0"

// Version prints version information about gosop, and/or the
// underlying OpenPGP library/libraries.
func Version() error {
	if !backend || extended {
		_, err := os.Stdout.WriteString("gosop " + VERSION + "\n")
		if err != nil {
			return versionErr(err)
		}
	}
	if backend || extended {
		_, err := os.Stdout.WriteString("GopenPGP " + constants.Version + "\n")
		if err != nil {
			return versionErr(err)
		}
	}
	if extended {
		info, ok := debug.ReadBuildInfo()
		if !ok {
			return versionErr(errors.New("couldn't read debug information"))
		}
		for i, module := range info.Deps {
			if module.Path == "github.com/ProtonMail/go-crypto" ||
				module.Path == "golang.org/x/crypto" {
				for module.Replace != nil {
					module = module.Replace
				}
				version := module.Version
				versionParts := strings.Split(version, "-")
				_, err := os.Stdout.WriteString(info.Deps[i].Path[11:] + " " + versionParts[len(versionParts)-1] + "\n")
				if err != nil {
					return versionErr(err)
				}
			}
		}
		_, err := os.Stdout.WriteString("Compiled using " + info.GoVersion + "\n")
		if err != nil {
			return versionErr(err)
		}
	}
	return nil
}

func versionErr(err error) error {
	return Err99("version", err)
}
