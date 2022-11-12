package unitis_test

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"
	"testing"
)

func TestGmDebugCoder(t *testing.T) {
	// filePath := "/home/max/max/workplace/repaire_png/errpng/01.jpg"
	// filePath := "/mnt/workplace/local_test/02.jpg"
	// filePath := "/home/max/max/workplace/repaire_png/errpng/de9868ae7f99422ab2d283af1df88c87.png"
	names := filenames
	for _, fn := range names {
		filePath := "/home/max/max/workplace/repaire_png/errpng/" + fn
		identifyDebugCoderCmdStr := fmt.Sprintf("identify +ping  -debug coder %s", filePath)
		result, err := Cmd(identifyDebugCoderCmdStr, true)
		rsp := string(result)
		if strings.Contains(rsp, "warning") || strings.Contains(rsp, "error") {
			fmt.Println(fn, "image have warning or error")
			continue
		}
		if err != nil {
			// t.Errorf(rsp, rsp)
			fmt.Println(fn, "identify image faild2")
			continue
		}
		fmt.Println(fn, "corrected png image")
	}

}
func BenchmarkGmDebugCoder(b *testing.B) {
	filename := "/home/max/max/workplace/repaire_png/errpng/02.jpg"
	identifyDebugCoderCmdStr := fmt.Sprintf("identify +ping %s", filename)
	result, err := Cmd(identifyDebugCoderCmdStr, true)
	rsp := string(result)
	if strings.Contains(rsp, "warning") || strings.Contains(rsp, "error") {
		b.Fatal("have warning or error")
	}
	if err != nil {
		b.Fatal(err)
	}
	buf := bytes.NewBuffer(result)
	var ret uint64
	binary.Read(buf, binary.BigEndian, &ret)
	b.SetBytes(int64(ret))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		identifyDebugCoderCmdStr := fmt.Sprintf("identify +ping  -debug coder %s", filename)
		Cmd(identifyDebugCoderCmdStr, true)
	}
}
