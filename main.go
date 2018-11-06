package main

import (
	"beego/model"
	"encoding/json"
	"github.com/redpois0n/wallpaper"
	"gopkg.in/resty.v1"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	idx :=  strconv.Itoa(random(1, 7))
	//mkt := "zh-CN"
	mkt :=""
	resp, err := resty.R().Get("http://www.bing.com/HPImageArchive.aspx?format=js&idx="+idx+"&n=8&mkt="+mkt)
	if err!=nil {
	}

	wall := model.Wallpaper{}
	json.Unmarshal([]byte(resp.String()), &wall)

	length := len(wall.Images)
	rand.Seed(time.Now().UnixNano())
	n :=  random(0, length-1)
	url := "http://www.bing.com"+wall.Images[n].URL
	println("设置的壁纸："+url)

	imgResp,err := resty.R().Get(url)
	path,err:=filepath.Abs("wallpaper.png")

	ioutil.WriteFile(path,imgResp.Body(),0)
	path = strings.Replace(path,"\\","/",-1)
	wallpaper.SetWallpaper(path)
	time.Sleep(1 * time.Second)
	os.Remove(path)

}

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}
