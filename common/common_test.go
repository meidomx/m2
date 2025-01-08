package edcode_test

import (
	"testing"

	edcode "github.com/meidomx/m2/common"
)

func TestReverseStr(t *testing.T) {
	t.Log(edcode.ReverseStrBytes("helloworld"))
	t.Log(edcode.ReverseStrBytes(""))
	t.Log(edcode.ReverseStrBytes("1"))
	t.Log(edcode.ReverseStrBytes("as"))
	t.Log(edcode.ReverseStrBytes("asd"))
}
