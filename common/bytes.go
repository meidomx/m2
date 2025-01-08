package edcode

import (
	"encoding/binary"
	"io"
)

func WriteInt64BE(w io.Writer, val int64) error {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(val))
	_, err := w.Write(buf)
	return err
}
