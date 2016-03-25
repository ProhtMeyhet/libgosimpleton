package simpleton

import(
	// empty
)

// give back a channel in which to push items of a list of []byte in a separate goroutine
func BytesListToChannel(list [][]byte) <-chan []byte {
	// be lenient on memory
	size := len(list); if size > 50 { size = 50 }
	channel := make(chan []byte, size)
	go func() {
		for _, item := range list {
			channel <-item
		}
		close(channel)
	}()
	return channel
}

// ----------delete everything below this line --------------- //

// give back a channel in which to push items of a list of string in a separate goroutine
func StringListToChannel(list []string) <-chan string {
	// be lenient on memory
	size := len(list); if size > 50 { size = 50 }
	channel := make(chan string, size)
	go func() {
		for _, item := range list {
			channel <-item
		}
		close(channel)
	}()
	return channel
}

// read from a channel, wait till its closed and return all read items
func StringChannelToList(channel <-chan string) (list []string) {
	for item := range channel {
		list = append(list, item)
	}
	return
}

// give back a channel in which to push items of a list of int in a separate goroutine
func IntListToChannel(list []int) <-chan int {
	// be lenient on memory
	size := len(list); if size > 50 { size = 50 }
	channel := make(chan int, size)
	go func() {
		for _, item := range list {
			channel <-item
		}
		close(channel)
	}()
	return channel
}

// read from a channel, wait till its closed and return all read items
func IntChannelToList(channel <-chan int) (list []int) {
	for item := range channel {
		list = append(list, item)
	}
	return
}

// give back a channel in which to push items of a list of uint in a separate goroutine
func UintListToChannel(list []uint) <-chan uint {
	// be lenient on memory
	size := len(list); if size > 50 { size = 50 }
	channel := make(chan uint, size)
	go func() {
		for _, item := range list {
			channel <-item
		}
		close(channel)
	}()
	return channel
}

// read from a channel, wait till its closed and return all read items
func UintChannelToList(channel <-chan uint) (list []uint) {
	for item := range channel {
		list = append(list, item)
	}
	return
}

// give back a channel in which to push items of a list of uint8 in a separate goroutine
func Uint8ListToChannel(list []uint8) <-chan uint8 {
	// be lenient on memory
	size := len(list); if size > 50 { size = 50 }
	channel := make(chan uint8, size)
	go func() {
		for _, item := range list {
			channel <-item
		}
		close(channel)
	}()
	return channel
}

// read from a channel, wait till its closed and return all read items
func Uint8ChannelToList(channel <-chan uint8) (list []uint8) {
	for item := range channel {
		list = append(list, item)
	}
	return
}

// give back a channel in which to push items of a list of uint16 in a separate goroutine
func Uint16ListToChannel(list []uint16) <-chan uint16 {
	// be lenient on memory
	size := len(list); if size > 50 { size = 50 }
	channel := make(chan uint16, size)
	go func() {
		for _, item := range list {
			channel <-item
		}
		close(channel)
	}()
	return channel
}

// read from a channel, wait till its closed and return all read items
func Uint16ChannelToList(channel <-chan uint16) (list []uint16) {
	for item := range channel {
		list = append(list, item)
	}
	return
}

// give back a channel in which to push items of a list of uint32 in a separate goroutine
func Uint32ListToChannel(list []uint32) <-chan uint32 {
	// be lenient on memory
	size := len(list); if size > 50 { size = 50 }
	channel := make(chan uint32, size)
	go func() {
		for _, item := range list {
			channel <-item
		}
		close(channel)
	}()
	return channel
}

// read from a channel, wait till its closed and return all read items
func Uint32ChannelToList(channel <-chan uint32) (list []uint32) {
	for item := range channel {
		list = append(list, item)
	}
	return
}

// give back a channel in which to push items of a list of uint64 in a separate goroutine
func Uint64ListToChannel(list []uint64) <-chan uint64 {
	// be lenient on memory
	size := len(list); if size > 50 { size = 50 }
	channel := make(chan uint64, size)
	go func() {
		for _, item := range list {
			channel <-item
		}
		close(channel)
	}()
	return channel
}

// read from a channel, wait till its closed and return all read items
func Uint64ChannelToList(channel <-chan uint64) (list []uint64) {
	for item := range channel {
		list = append(list, item)
	}
	return
}

// give back a channel in which to push items of a list of int8 in a separate goroutine
func Int8ListToChannel(list []int8) <-chan int8 {
	// be lenient on memory
	size := len(list); if size > 50 { size = 50 }
	channel := make(chan int8, size)
	go func() {
		for _, item := range list {
			channel <-item
		}
		close(channel)
	}()
	return channel
}

// read from a channel, wait till its closed and return all read items
func Int8ChannelToList(channel <-chan int8) (list []int8) {
	for item := range channel {
		list = append(list, item)
	}
	return
}

// give back a channel in which to push items of a list of int16 in a separate goroutine
func Int16ListToChannel(list []int16) <-chan int16 {
	// be lenient on memory
	size := len(list); if size > 50 { size = 50 }
	channel := make(chan int16, size)
	go func() {
		for _, item := range list {
			channel <-item
		}
		close(channel)
	}()
	return channel
}

// read from a channel, wait till its closed and return all read items
func Int16ChannelToList(channel <-chan int16) (list []int16) {
	for item := range channel {
		list = append(list, item)
	}
	return
}

// give back a channel in which to push items of a list of int32 in a separate goroutine
func Int32ListToChannel(list []int32) <-chan int32 {
	// be lenient on memory
	size := len(list); if size > 50 { size = 50 }
	channel := make(chan int32, size)
	go func() {
		for _, item := range list {
			channel <-item
		}
		close(channel)
	}()
	return channel
}

// read from a channel, wait till its closed and return all read items
func Int32ChannelToList(channel <-chan int32) (list []int32) {
	for item := range channel {
		list = append(list, item)
	}
	return
}

// give back a channel in which to push items of a list of int64 in a separate goroutine
func Int64ListToChannel(list []int64) <-chan int64 {
	// be lenient on memory
	size := len(list); if size > 50 { size = 50 }
	channel := make(chan int64, size)
	go func() {
		for _, item := range list {
			channel <-item
		}
		close(channel)
	}()
	return channel
}

// read from a channel, wait till its closed and return all read items
func Int64ChannelToList(channel <-chan int64) (list []int64) {
	for item := range channel {
		list = append(list, item)
	}
	return
}

// give back a channel in which to push items of a list of float32 in a separate goroutine
func Float32ListToChannel(list []float32) <-chan float32 {
	// be lenient on memory
	size := len(list); if size > 50 { size = 50 }
	channel := make(chan float32, size)
	go func() {
		for _, item := range list {
			channel <-item
		}
		close(channel)
	}()
	return channel
}

// read from a channel, wait till its closed and return all read items
func Float32ChannelToList(channel <-chan float32) (list []float32) {
	for item := range channel {
		list = append(list, item)
	}
	return
}

// give back a channel in which to push items of a list of float64 in a separate goroutine
func Float64ListToChannel(list []float64) <-chan float64 {
	// be lenient on memory
	size := len(list); if size > 50 { size = 50 }
	channel := make(chan float64, size)
	go func() {
		for _, item := range list {
			channel <-item
		}
		close(channel)
	}()
	return channel
}

// read from a channel, wait till its closed and return all read items
func Float64ChannelToList(channel <-chan float64) (list []float64) {
	for item := range channel {
		list = append(list, item)
	}
	return
}

// give back a channel in which to push items of a list of complex64 in a separate goroutine
func Complex64ListToChannel(list []complex64) <-chan complex64 {
	// be lenient on memory
	size := len(list); if size > 50 { size = 50 }
	channel := make(chan complex64, size)
	go func() {
		for _, item := range list {
			channel <-item
		}
		close(channel)
	}()
	return channel
}

// read from a channel, wait till its closed and return all read items
func Complex64ChannelToList(channel <-chan complex64) (list []complex64) {
	for item := range channel {
		list = append(list, item)
	}
	return
}

// give back a channel in which to push items of a list of complex128 in a separate goroutine
func Complex128ListToChannel(list []complex128) <-chan complex128 {
	// be lenient on memory
	size := len(list); if size > 50 { size = 50 }
	channel := make(chan complex128, size)
	go func() {
		for _, item := range list {
			channel <-item
		}
		close(channel)
	}()
	return channel
}

// read from a channel, wait till its closed and return all read items
func Complex128ChannelToList(channel <-chan complex128) (list []complex128) {
	for item := range channel {
		list = append(list, item)
	}
	return
}

// give back a channel in which to push items of a list of byte in a separate goroutine
func ByteListToChannel(list []byte) <-chan byte {
	// be lenient on memory
	size := len(list); if size > 50 { size = 50 }
	channel := make(chan byte, size)
	go func() {
		for _, item := range list {
			channel <-item
		}
		close(channel)
	}()
	return channel
}

// read from a channel, wait till its closed and return all read items
func ByteChannelToList(channel <-chan byte) (list []byte) {
	for item := range channel {
		list = append(list, item)
	}
	return
}

// give back a channel in which to push items of a list of rune in a separate goroutine
func RuneListToChannel(list []rune) <-chan rune {
	// be lenient on memory
	size := len(list); if size > 50 { size = 50 }
	channel := make(chan rune, size)
	go func() {
		for _, item := range list {
			channel <-item
		}
		close(channel)
	}()
	return channel
}

// read from a channel, wait till its closed and return all read items
func RuneChannelToList(channel <-chan rune) (list []rune) {
	for item := range channel {
		list = append(list, item)
	}
	return
}
