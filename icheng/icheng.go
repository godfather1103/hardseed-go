package icheng

import (
	"fmt"
	"github.com/godfather1103/hardseed-go/vo"
	"github.com/godfather1103/hardseed-go/utils"
	"path/filepath"
	"strconv"
	"os"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"io"
	"bytes"
	"mime/multipart"
)

var params *vo.DownloadInfo

func DownloadBt(info *vo.DownloadInfo) {
	//fmt.Println(utils.GetHttpInfo("https://getman.cn/echo","POST",strings.NewReader("a=1&b=222")))
	params = info
	var htmlFiles []string = []string{}
	for i := info.PageStart; i <= info.PageEnd; i++ {
		fileName := downloadHtml(info.Url, info.Path, i)
		if len(fileName) > 0 {
			htmlFiles = append(htmlFiles, fileName)
		}
	}
}

func downloadHtml(action string, path string, pageIndex int64) string {

	err := utils.PathMkdir(path + string(filepath.Separator) + utils.HtmlPrefix + string(filepath.Separator))
	if err != nil {
		fmt.Printf("check dir error![%v]\n", err)
		return ""
	}

	fileName := path + string(filepath.Separator) + utils.HtmlPrefix + string(filepath.Separator) + "page-" + strconv.FormatInt(pageIndex, 10) + ".html"
	action += "&page=" + strconv.FormatInt(pageIndex, 10)

	exists, err := utils.PathExists(fileName)

	var f *os.File

	if err != nil {
		fmt.Printf("get file error![%v]\n", err)
		return ""
	}
	if exists {
		f, _ = os.OpenFile(fileName, os.O_RDWR, 0666)
	} else {
		f, _ = os.Create(fileName)
	}

	httpClientParam := utils.GetHttpClientParam(params.Proxy)

	resp, code := utils.GetHttpInfo(action, "GET", nil, httpClientParam)
	if code == 200 && len(resp) > 0 {
		f.WriteString(resp)
	} else {
		return ""
	}
	f.Close()
	parseHtml(fileName)
	return fileName
}

func parseHtml(fileName string) {
	f, _ := os.OpenFile(fileName, os.O_RDONLY, 0666)
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		fmt.Printf("get file error![%v]\n", err)
	}
	doc.Find(".tr3.t_one h3 a").Each(func(i int, selection *goquery.Selection) {
			title := selection.Text()
			title = utils.GetUTF8StrFromGB2312(title)
			if !utils.CheckKeywordContainTitle(title,params.HateList) {
				href,exists := selection.Attr("href")
				if exists {
					downloadPicAndBt(params.Host+href,params.Path,title)
					fmt.Printf("Node is === %v\t%v\n",title,href)
				}
			}
		})

}

func downloadPicAndBt(action string, path string,title string) {
	path += string(filepath.Separator) + utils.BtProfix + string(filepath.Separator)
	err := utils.PathMkdir(path)
	if err != nil {
		fmt.Printf("create dir error![%v]\n", err)
	}else {
		resp,code := utils.GetHttpInfo(action,"GET",nil,utils.GetHttpClientParam(params.Proxy))
		resp = utils.GetUTF8StrFromGB2312(resp)
		if code == 200 {
			doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp))
			if err!=nil {
				fmt.Printf("get doc error! %v\n",err)
				return
			}
			imgsrc,_ := doc.Find("#read_tpc img").Attr("src")
			if len(imgsrc)>0 {
				saveFile(imgsrc,"GET",nil,path,title+".jpg",nil,false)
			}

			ahref,_ := doc.Find("#read_tpc a").Attr("href")

			if len(ahref)>0 {
				code := ""
				if strings.Contains(ahref,"?") {
					ahref = ahref[strings.Index(ahref,"?")+1:]
				}
				urlParams := strings.Split(ahref,"&")

				for _,value := range urlParams {
					if strings.Contains(value,"ref=") {
						code = value[strings.Index(value,"ref=")+4:]
					}
				}
				if len(code)>0 {
					//writer := multipart.NewWriter(bytes.NewBufferString("code="+code))
					header := map[string]string{"code": code}
					saveFile("http://www.jandown.com/fetch.php","POST",nil,path,title+".torrent",header,true)
				}
			}
		}
	}
}

func saveFile(action string,method string,param io.Reader,path string,fileName string,header map[string]string,isMutil bool)  {
	var resp string
	var code int
	if isMutil {
		writer := multipart.NewWriter(&bytes.Buffer{})
		resp,code = utils.GetMultipart(action, header,writer.FormDataContentType(),utils.GetHttpClientParam(params.Proxy))
	}else {
		if header!=nil {
			resp,code = utils.GetHttpInfoHeader(action,method,param,header,utils.GetHttpClientParam(params.Proxy))
		}else {
			resp,code = utils.GetHttpInfo(action,method,param,utils.GetHttpClientParam(params.Proxy))
		}
	}
	if code==200 {
		exists,err := utils.PathExists(path+fileName)
		if err!=nil {
			fmt.Printf("get file error! %v\n",err)
			return
		}
		var f *os.File
		if exists {
			f, _ = os.OpenFile(path+fileName, os.O_RDWR, 0666)
		} else {
			f, _ = os.Create(path+fileName)
		}
		if f != nil && len(resp)>0{
			io.Copy(f,bytes.NewReader([]byte(resp)))
		}
	}
}
