package template

import "fmt"

func Get(tmpl string) (string, error) {
	switch tmpl {
	case "StandardController":
		return StandardController, nil
	case "StandardControllerWithSubreconcilers":
		return StandardControllerWithSubreconcilers, nil
	default:
		return StandardController, fmt.Errorf(
			"Yo! The template \"%s\" is not one I recognize.",
			tmpl,
		)
	}
}
