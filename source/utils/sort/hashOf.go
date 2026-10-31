package sort

func hashOf(value any) string {

	if hashable, ok := value.(interface{ Hash() string }); ok {
		return hashable.Hash()
	}

	return ""

}
