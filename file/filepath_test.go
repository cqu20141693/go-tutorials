package file

import (
	"fmt"
	"github.com/cqu20141693/go-service-common/file"
	"path/filepath"
	"testing"
)

func TestFilePathJoin(t *testing.T) {
	service := "sip-server"
	join := filepath.Join(file.GetCurrentPath(), "log", service+".log-%Y%m%d%H%M")
	fmt.Println(join)
}
