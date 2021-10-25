package trace

import (
	"bytes"
	"fmt"
	"testing"
)

func TestNew(t *testing.T){
	var buf bytes.Buffer
	tracer:= New(&buf)
	if tracer == nil{
		t.Error("Tracer is Nil")
	}else{
		msg:="Hello Trace Package."
		tracer.Trace(msg)
		if buf.String()!=msg{
			t.Errorf("Trace should not write '%s'.", buf.String())
			fmt.Println("Issue:",buf.String())

		}else{
			fmt.Println("Printed:",buf.String())
		}
	}
}
