package schema

import (
	"os"

	"github.com/followthepattern/adapticc/config"
)

func GetSchema(cfg config.Server) (string, error) {
	content, err := os.ReadFile(cfg.GraphqlSchemaFilepath)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
