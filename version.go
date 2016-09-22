package libgosimpleton

import(
	"fmt"
)

const(
	// only changes with major changes
	MAJOR_VERSION = "0"

	// new features, maybe incompatible changes
	MINOR_VERSION = "3"

	// no incompatible changes, just bug fixes
	BUGFIX_VERSION = "0"

	// munchies
	VERSION = MAJOR_VERSION + "." + MINOR_VERSION + "." + BUGFIX_VERSION

	// should one be using this
	FROM_MASTER = true
)

func Version() {
	fmt.Println("libgosimpleton version: " + VERSION)
}
