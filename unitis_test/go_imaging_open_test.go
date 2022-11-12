package unitis_test

import (
	"fmt"
	"os/exec"
	"testing"

	"github.com/disintegration/imaging"
)

var filenames = []string{
	"01.png",
	"02.png",
	"03.png",
	"04.png",
	"05.png",
	"06.png",
	"07.png",
	"08.png",
	"09.png",
	"10.png",
	"00.jpg",
	"01.jpg",
	"02.jpg",
	"03.jpg",
	"04.jpg",
	"05.jpg",
	"06.jpg",
	"11.jpg",
	"12.jpg",
	"13.jpg",
	"14.jpg",
	"15.jpg",
	"16.jpg",
	"17.jpg",
	"18.jpg",
	"19.jpg",
	"20.jpg",
	"b1343e43c0f14b03881af41f460943fd.png",
	"de9868ae7f99422ab2d283af1df88c87.png",
	"25.jpg",
}

func Cmd(cmd string, shell bool) ([]byte, error) {
	if shell {
		return exec.Command("bash", "-c", cmd).CombinedOutput()
	} else {
		return exec.Command(cmd).Output()
	}
}
func TestGoImageOpen(t *testing.T) {
	// filePath := "/home/max/max/workplace/repaire_png/errpng/02.jpg"
	names := filenames
	for _, fn := range names {
		_, err := imaging.Open("/home/max/max/workplace/repaire_png/errpng/" + fn)
		if err != nil {
			fmt.Println(err, "failed1")
		} else {
			fmt.Println(err, "successed1")
		}
	}

}
func BenchmarkGoImageOpen(b *testing.B) {
	filename := "/home/max/max/workplace/repaire_png/errpng/02.jpg"
	image, err := imaging.Open(filename)
	if err != nil {
		b.Fatal(err)
	}
	b.SetBytes(int64(image.Bounds().Dx() * image.Bounds().Dy()))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		imaging.Open(filename)
	}
}
