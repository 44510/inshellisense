// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["nano"] = model.Subcommand{
		Name:        []string{"nano"},
		Description: `Nano's ANOther editor, an enhanced free Pico clone`,
		Args: []model.Arg{{
			Templates: []model.Template{model.TemplateFilepaths},
		}},
	}
}