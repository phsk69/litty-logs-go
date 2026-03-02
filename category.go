package litty

import "strings"

// ShortenCategory yeets the namespace bloat from category names.
// "github.com/user/pkg.Service" becomes just "Service" fr fr.
// handles both dot and slash separators because Go uses slashes for package paths bestie
func ShortenCategory(category string) string {
	// check for the last separator of either type — whichever comes last wins
	dotIdx := strings.LastIndexByte(category, '.')
	slashIdx := strings.LastIndexByte(category, '/')

	// pick whichever separator is furthest right — thats the most specific segment
	idx := dotIdx
	if slashIdx > idx {
		idx = slashIdx
	}

	if idx >= 0 {
		return category[idx+1:]
	}
	return category
}
