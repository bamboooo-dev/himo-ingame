package env

import (
	"os"
	"strings"

	"golang.org/x/xerrors"
)

func envToMap() (map[string]string, error) {
	envMap := make(map[string]string)
	for _, env := range os.Environ() {
		vals := strings.SplitN(env, "=", 2)
		if len(vals) != 2 {
			return nil, xerrors.New("envToMap failed during SplitN")
		}
		envMap[vals[0]] = vals[1]
	}
	return envMap, nil
}
