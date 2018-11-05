package main

import (
	"beego/model"
	"encoding/json"
	"fmt"
	"github.com/redpois0n/wallpaper"
	"gopkg.in/resty.v1"
	"io/ioutil"
	"math/rand"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {

	n:=strconv.Itoa(rand.Intn(100 - 1) + 1)

	resp, err := resty.R().Get("http://www.bing.com/HPImageArchive.aspx?format=js&idx="+n+"&n=1&mkt=zh-CN")
	if err!=nil {
	}

	wall := model.Wallpaper{}
	json.Unmarshal([]byte(resp.String()), &wall)

	url := "http://www.bing.com"+wall.Images[0].URL

	imgResp,err := resty.R().Get(url)
	path,err:=filepath.Abs("wallpaper.png")
	fmt.Println(path)
	ioutil.WriteFile(path,imgResp.Body(),0)
	path = strings.Replace(path,"\\","/",-1)
	wallpaper.SetWallpaper(path)

	//os.Remove(path)

}
