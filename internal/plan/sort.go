package plan

import "sort"

func Normalize(value Plan) Plan {
	sort.Slice(value.Actions, func(i, j int) bool {
		if value.Actions[i].Type != value.Actions[j].Type {
			return value.Actions[i].Type < value.Actions[j].Type
		}

		return value.Actions[i].Identity < value.Actions[j].Identity
	})

	return value
}
