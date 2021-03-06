// +build go1.9

package qdb

/*
	#include <qdb/client.h>
*/
import "C"
import (
	"time"
	"unsafe"
)

func toQdbTimespec(time time.Time) C.qdb_timespec_t {
	return C.qdb_timespec_t{C.qdb_time_t(time.Unix()), C.qdb_time_t(time.Nanosecond())}
}

func toQdbTime(t time.Time) C.qdb_time_t {
	if t.Equal(PreserveExpiration()) {
		val := C.ulonglong(C.qdb_preserve_expiration)
		ptr := unsafe.Pointer(&val)
		return *(*C.qdb_time_t)(ptr)
	}
	return C.qdb_time_t(t.UnixNano() / int64(time.Millisecond))
}

func (cval C.qdb_timespec_t) toStructG() time.Time {
	return time.Unix(int64(cval.tv_sec), int64(cval.tv_nsec))
}
