package simpleton

import(
	"testing"

	"reflect"
)

func TestStringListToChannel(t *testing.T) {
	list := []string{ "acdc", "are", "the", "greatest" }
	key := 0

	channel := StringListToChannel(list...)
	for item := range channel {
		if !reflect.DeepEqual(list[key], item) {
			t.Errorf("expected %v, got %v\n", list[key], item)
		}
		key++
	}
}

func TestStringChannelToList(t *testing.T) {
	list := []string{ "acdc", "are", "the", "greatest" }
	channel := make(chan string, len(list))
	go func() {
		for _, item := range list {
			channel <-item
		}
		close(channel)
	}()

	listFromChannel := StringChannelToList(channel)
	if !reflect.DeepEqual(list, listFromChannel) {
		t.Fatalf("expected %v, got %v\n", list, listFromChannel)
	}
}

//TODO more testing
