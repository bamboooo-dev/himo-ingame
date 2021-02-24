package env

import (
	"bytes"
	"os"
	"text/template"

	"golang.org/x/xerrors"
)

// generateYamlFromTemplate テンプレートのパスを指定して環境変数を元に Yaml を生成、file に書き込む
func generateYamlFromTemplate(yamlTemplatePath string, filePath string) (err error) {
	file, err := os.Create(filePath)
	if err != nil {
		return
	}
	defer func() {
		closeErr := file.Close()
		if closeErr != nil {
			err = xerrors.Errorf("%v: %w", closeErr, err)
		}
	}()

	var buf bytes.Buffer
	envMap, err := envToMap()
	if err != nil {
		return
	}
	t, err := template.ParseFiles(yamlTemplatePath)
	if err != nil {
		return
	}
	t = t.Option("missingkey=zero")
	if err = t.Execute(&buf, envMap); err != nil {
		return
	}
	_, err = file.Write(buf.Bytes())
	if err != nil {
		return
	}

	return
}
