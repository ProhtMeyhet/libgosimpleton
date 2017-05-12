package processes

import(
	"strconv"
)

// tries to parse a list of strings as uints. if successful uints will filled,
// if not strings will be filled. if an entry is empty, it will be ignored.
func ParseAsNamesOrIds(froms ...string) (names []string, ids []uint) {
	for _, from := range froms {
		if from == "" { continue }
		id64, e := strconv.ParseUint(from, 10, 0); if e == nil {
			ids = append(ids, uint(id64))
		} else {
			names = append(names, from)
		}
	}; return
}
