package simplemath;

import "testing";

func TestAdd(t *testing.T) {
	r := Add(1,2);
	if r != 3{
		t.Error("add error");

	}
}
