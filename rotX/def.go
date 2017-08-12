package rotX

import (
	"io"
)

type RotX struct {
	reader io.Reader
	shift byte
}

func Make(reader io.Reader, shift byte) RotX {
	return RotX{reader,shift}
}

func (rotXreader RotX) Read(b []byte) (count int, result error) {
	count, result = rotXreader.reader.Read(b)
	for i:=0;i<count;i++ {
		if b[i] == '\n' {
		} else if b[i] >= 'a' && b[i] <= 'z' {
			b[i] = 'z' + 'a' - b[i]
		} else if b[i] >= 'A' && b[i] <= 'Z' {
			b[i] = 'Z' + 'A' - b[i]
		}
	}
	return
}