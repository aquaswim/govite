package govite

func appendIfNotExists(target []string, src string) []string {
	for _, s := range target {
		if s == src {
			return target
		}
	}
	return append(target, src)
}
