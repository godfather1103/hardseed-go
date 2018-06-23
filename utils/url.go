package utils

import (
	"errors"
	"strings"
)

var url_map map[string]string

func initParam() {
	url_map = make(map[string]string)
	url_map["caoliu_west_reposted"] = "thread0806.php?fid=19"
	url_map["caoliu_cartoon_reposted"] = "thread0806.php?fid=24"
	url_map["caoliu_asia_mosaicked_reposted"] = "thread0806.php?fid=18"
	url_map["caoliu_asia_non_mosaicked_reposted"] = "thread0806.php?fid=17"
	url_map["caoliu_west_original"] = "thread0806.php?fid=4"
	url_map["caoliu_cartoon_original"] = "thread0806.php?fid=5"
	url_map["caoliu_asia_mosaicked_original"] = "thread0806.php?fid=15"
	url_map["caoliu_asia_non_mosaicked_original"] = "thread0806.php?fid=2"
	url_map["caoliu_selfie"] = "thread0806.php?fid=16"
	url_map["aicheng_west"] = "thread.php?fid=5"
	url_map["aicheng_cartoon"] = "thread.php?fid=6"
	url_map["aicheng_asia_mosaicked"] = "thread.php?fid=4"
	url_map["aicheng_asia_non_mosaicked"] = "thread.php?fid=16"
}

func GetUrlPrfix(key string) (string,error){
	if len(key)<1 {
		return "",errors.New("字段不能为空！")
	}else{
		if url_map==nil {
			initParam()
		}
		return url_map[key],nil
	}
}

func CheckKeywordContainTitle(title string,keywords []string) bool {
	if len(title)<1 || len(keywords)<1 {
		return false
	}else {
		for _,key := range keywords{
			if strings.Index(title,key) >= 0 {
				return true
			}
		}
		return false
	}
}