package plan

import (
	"fmt"
	"io"
)

func WriteText(
	writer io.Writer,
	value Plan,
) error {
	for _, action := range Normalize(value).Actions {
		if _, err := fmt.Fprintf(
			writer,
			"%s %s\n",
			action.Type,
			action.Identity,
		); err != nil {
			return err
		}
	}

	return nil
}
