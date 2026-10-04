package persistence

func Changed(
	oldHash string,
	newHash string,
) bool {

	if oldHash == "" {
		return true
	}

	return oldHash != newHash
}
