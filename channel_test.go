package libgosimpleton

import(
	"testing"
)

func TestStringListToChannel(t *testing.T) {
	list := []string{ "acdc", "are", "the", "greatest" }
	key := 0

	channel := StringListToChannel(list)
	for item := range channel {
		Assert(t, list[key], item)
		key++
	}
}

//TODO more testing
