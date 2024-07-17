package parser

import (
	"regexp"
	"strings"

	"github.com/reonardoleis/cherry/pkg/component"
)

func HTML[T any](root component.Component[T], raw string) string {
	indexes := make(map[string]int)
	for k := range root.Children() {
		indexes[k] = 0
	}

	re := regexp.MustCompile(`<\.(\w+)>.*?<\/\.\w+>`)
	matches := re.FindAllStringSubmatch(raw, -1)

	for _, match := range matches {
		tag := match[0]
		componentName := strings.Split(tag, "</.")[0]
		componentName = strings.Replace(componentName, "<", "", 1)
		componentName = strings.Replace(componentName, ".", "", 1)
		componentName = strings.Replace(componentName, ">", "", 1)

		index := indexes[componentName]
		component := root.Children()[componentName][index]

		raw = strings.Replace(raw, tag, component.Render(), 1)
		indexes[componentName] += 1
	}

	return raw
}
