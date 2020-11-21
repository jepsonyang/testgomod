package utils

import (
	"github.com/jepsonyang/testmod1/jep1"
	"testing"
)

func Test_SayHello(t *testing.T) {
	SayHello()
	t.Logf("test is ok")

	jep1.Jep1()
}
