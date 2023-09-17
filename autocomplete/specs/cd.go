// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["cd"] = model.Subcommand{
		Name:        []string{"cd"},
		Description: `Change the shell working directory`,
		Args: []model.Arg{{
			Suggestions: []model.Suggestion{{
				Name:        []string{`-`},
				Description: `Switch to the last used folder`,
			}, {
				Name:        []string{`~`},
				Description: `Switch to the home directory`,
			}},
			FilterStrategy: model.FilterStrategyFuzzy,
			Generator:      nil, // TODO: port over generator
		}},
	}
}