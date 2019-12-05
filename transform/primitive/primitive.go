package primitive

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
)

// Transform an image into a primitive image
func Transform(in, ext string, mode int, nbShapes int) (io.Reader, error) {
	out, err := ioutil.TempFile("", fmt.Sprintf("out_*%s", ext))
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(out.Name())
	cmd := exec.Command("primitive", "-v", "-i", in, "-o", out.Name(), "-m", strconv.Itoa(mode), "-n", strconv.Itoa(nbShapes))
	result, err := cmd.CombinedOutput()
	log.Println(string(result))
	if err != nil {
		return nil, err
	}
	b := bytes.NewBuffer(nil)
	_, err = io.Copy(b, out)
	if err != nil {
		return nil, err
	}
	return b, nil
}
