package edcode_test

import (
	"testing"
)

import (
	"gorepo.moetang.info/prod/m2/common/edcode"
)

func TestReverseStr(t *testing.T) {
	t.Log(edcode.ReverseStrBytes("helloworld"))
	t.Log(edcode.ReverseStrBytes(""))
	t.Log(edcode.ReverseStrBytes("1"))
	t.Log(edcode.ReverseStrBytes("as"))
	t.Log(edcode.ReverseStrBytes("asd"))
}
