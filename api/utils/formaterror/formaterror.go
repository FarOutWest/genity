package formaterror

import (
	"errors"
	"strings"
)

func FormatError(err string) error {

	if strings.Contains(err, "title") {
		return errors.New("Title Already Taken")
	}

	return errors.New("Incorrect Details")
}
