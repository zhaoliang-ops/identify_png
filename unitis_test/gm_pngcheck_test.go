package unitis_test

import (
	"fmt"
	"strings"
	"testing"
)

func TestGmPngCheck(t *testing.T) {
	// filePath := "/home/max/max/workplace/repaire_png/errpng/02.jpg"
	// filePath := "/mnt/workplace/local_test/02.jpg"
	// filePath := "/mnt/workplace/local_test/de9868ae7f99422ab2d283af1df88c87.png"
	names := filenames
	for _, fn := range names {
		filePath := "/home/max/max/workplace/repaire_png/errpng/" + fn
		identifyCmdStr := fmt.Sprintf("identify +ping %s", filePath)
		result, err := Cmd(identifyCmdStr, true)
		rsp := string(result)
		if strings.Contains(rsp, "warning") || strings.Contains(rsp, "error") {
			fmt.Println(fn, "identify image faild1", "waring")
			continue
		}
		if err != nil {
			fmt.Println(fn, "identify image faild2", "err")
			continue
		}
		pngCheckCmdStr := fmt.Sprintf("pngcheck -v %s", filePath)
		result, err = Cmd(pngCheckCmdStr, true)
		rsp = string(result)
		if strings.Contains(rsp, "ERRORS") {
			fmt.Println(fn, "pngCheck image faild", "ERRORS")
			continue
		}
		if err != nil {
			fmt.Println(fn, "pngCheck image faild", "ERRORS")
			continue
		}
		fmt.Println(fn, "corrected png image")
	}

}
