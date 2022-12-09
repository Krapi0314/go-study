package main

import "golang.org/x/tour/reader"

type MyReader struct{}
type MyReaderError int

func (e MyReaderError) Error() string {
	return "capacity max error, cap: " + string(e)
}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (myReader MyReader) Read(b []byte) (int, error) {
	if cap(b) < 1 {
		return 0, MyReaderError(cap(b))
	}
	
	b[0] = 'A'
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}
