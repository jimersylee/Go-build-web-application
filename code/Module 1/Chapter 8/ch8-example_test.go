package example_test

func TestSquare(t *Testing.T) {
	if Square(4) != 16 {
		t.Error("expected", 16, "got", 4)
	}
}
