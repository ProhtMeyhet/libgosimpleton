package libgocredentials

import(

)

func IsTransactionIsolationLevel(level string) bool {
	switch(level) {
	case TRANSACTION_READ_UNCOMMITTED:
		fallthrough
	case TRANSACTION_READ_COMMITTED:
		fallthrough
	case TRANSACTION_REPEATABLE_READ:
		fallthrough
	case TRANSACTION_SERIALIZE:
		return true
	}

	return false
}
