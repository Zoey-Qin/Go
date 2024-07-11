
package calculator

import (
	"testing"
)

func TestAdd(t *testing.T){
	r := add(2,4)
	if r!=6{
		t.Fatalf("expect:%d, actul:%d",6,r)
	}
	t.Logf("test add succ")
}
