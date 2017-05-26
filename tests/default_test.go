package test

import (
	_ "github.com/riveryang/sysgo/routers"
	"path/filepath"
	"runtime"
	"testing"

	"encoding/base64"
	"fmt"
	"github.com/astaxie/beego"

	"github.com/docker/go/canonical/json"
	"github.com/riveryang/sysgo/aes"
	"github.com/riveryang/sysgo/models"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

// TestGet is a sample to run an endpoint test
func TestAes(t *testing.T) {
	if pcStat, err := models.NewPcStat(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(pcStat)

		val, _ := pcStat.Encrypt()
		fmt.Println(val)

		decrypt, _ := base64.StdEncoding.DecodeString(val)
		bytes, _ := aes.AesDecrypt(decrypt, aes.DefaultKey())
		fmt.Println(string(bytes))

		var stat = models.PCStat{}
		json.Unmarshal(bytes, &stat)
		fmt.Println(stat)
	}
}
