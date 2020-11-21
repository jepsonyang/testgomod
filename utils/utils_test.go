package utils

import (
	"github.com/jepsonyang/testmod1/jep2"
	"testing"
)

func Test_SayHello(t *testing.T) {
	SayHello()
	t.Logf("test is ok")

	jep2.Jep2()
}
