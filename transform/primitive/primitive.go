package primitive

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

// Transform an image into a primitive image
func Transform(in, dst string, mode Mode, nbShapes int) error {
	fileName := filepath.Base(in)
	outFilePath := fmt.Sprintf("%s/%s_%s", dst, strings.ToLower(mode.String()), fileName)
	out, err := os.Create(outFilePath)
	if err != nil {
		return err
	}
	cmd := exec.Command("primitive", "-v", "-i", in, "-o", out.Name(), "-m", strconv.Itoa(int(mode)), "-n", strconv.Itoa(nbShapes))
	result, err := cmd.CombinedOutput()
	log.Printf("[INFO] %s", string(result))
	if err != nil {
		return err
	}
	return nil
}
