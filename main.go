package main

import (
	"fmt"
	"flag"
	"strings"
	"os"
	"os/user"
	"path/filepath"
	"time"
	"strconv"
	"github.com/godfather1103/hardseed-go/utils"
	"github.com/godfather1103/hardseed-go/vo"
	"github.com/godfather1103/hardseed-go/icheng"
	"github.com/godfather1103/hardseed-go/caoliu"
)

var g_version string  = "1.0"
var g_softname string  = "hardseed-go"
var g_myemail string = "chuchuanbao@gmail.com"
var g_mywebspace string = "https://github.com/godfather1103/"+g_softname
var g_myemail_color string = "\x1b[1m"+"\x1b[32m"+g_myemail+"\x1b[0m"
var g_mywebspace_color string = "\x1b[1m"+"\x1b[32m"+g_mywebspace+"\x1b[0m"


func contains(args [] string,str string) (bool) {
	if args == nil || len(str)<1 {
		return false
	}
	for i:=0; i< len(args);i++  {
		if args[i] == str {
			return true
		}
	}
	return false
}

func showHelpInfo()  {
	fmt.Print(g_softname + " is a batch seeds and pictures download utiltiy from CaoLiu and AiCheng forum. ")
	fmt.Print("It's easy and simple to use. Usually, you could issue it as follow: " + "\n")
	fmt.Print("  $ hardseed" + "\n")
	fmt.Print("or" + "\n")
	fmt.Print("  $ hardseed --saveas-path ~/downloads --concurrent-tasks 4 --topics-range 8 64")
	fmt.Print(" --av-class aicheng_west --timeout-download-picture 32 --hate X-Art --proxy http://127.0.0.1:8087" + "\n")

	fmt.Print("\n")
	fmt.Print("  --help" + "\n")
	fmt.Print("  Show this help infomation what you are seeing. " + "\n")

	fmt.Print("\n")
	fmt.Print("  --version" + "\n")
	fmt.Print("  Show current version. " + "\n")

	fmt.Print("\n")
	fmt.Print("  --av-class" + "\n")
	fmt.Print("  There are 13 av classes: caoliu_west_reposted, caoliu_cartoon_reposted, ")
	fmt.Print("caoliu_asia_mosaicked_reposted, caoliu_asia_non_mosaicked_reposted, caoliu_west_original, ")
	fmt.Print("caoliu_cartoon_original, caoliu_asia_mosaicked_original, caoliu_asia_non_mosaicked_original, ")
	fmt.Print("caoliu_selfie, aicheng_west, aicheng_cartoon, aicheng_asia_mosaicked and aicheng_asia_non_mosaicked. " + "\n")
	fmt.Print("  As the name implies, \"caoliu\" stands for CaoLiu forum, \"aicheng\" for AiCheng forum, ")
	fmt.Print("\"reposted\" and \"original\" are clearity, and the \"selfie\" is photos by oneself, you konw ")
	fmt.Print("which one is your best lover (yes, only one). " + "\n")
	fmt.Print("  The default is aicheng_asia_mosaicked. " + "\n")

	fmt.Print("\n")
	fmt.Print("  --concurrent-tasks" + "\n")
	fmt.Print("  You can set more than one proxy, each proxy could more than one concurrent tasks. This option ")
	fmt.Print("set the number of concurrent tasks of each prox. " + "\n")
	fmt.Print("  The default number is 8. " + "\n")

	fmt.Print("\n")
	fmt.Print("  --timeout-download-picture" + "\n")
	fmt.Print("  Some pictures too big to download in few seconds. So, you should set the download picture ")
	fmt.Print("timeout seconds. " + "\n")
	fmt.Print("  The default timeout is 16 seconds." + "\n")

	fmt.Print("\n")
	fmt.Print("  --topics-range" + "\n")
	fmt.Print("  Set the range of to download topics. E.G.: " + "\n")
	fmt.Print("    --topics-range 2 16" + "\n")
	fmt.Print("    --topics-range 8 (I.E., --topics-range 1 8)" + "\n")
	fmt.Print("    --topics-range -1 (I.E., all topics of this av class)" + "\n")
	fmt.Print("  The default topics range is 1 to 64. " + "\n")

	fmt.Print("\n")
	fmt.Print("  --saveas-path" + "\n")
	fmt.Print("  Set the path to save seeds and pictures. The rule of dir: [avclass][range]@hhmmss. E.G., ")
	fmt.Print("[aicheng_west][2~32]@124908/. " + "\n")
	fmt.Print("  The default directory is home directory (or windows is C:\\). " + "\n")

	fmt.Print("\n")
	fmt.Print("  --hate" + "\n")
	fmt.Print("  If you hate some subject topics, you can ignore them by setting this option with keywords ")
	fmt.Print("in topic title, split by space-char ' ', and case sensitive. E.G., --hate 孕妇 重口味. ")
	fmt.Print("When --hate keywords list conflict with --like, --hate first. " + "\n")

	fmt.Print("\n")
	fmt.Print("  --like" + "\n")
	fmt.Print("  If you like some subject topics, you can grab them by setting this option with keywords ")
	fmt.Print("in topic title, split by space-char ' ', and case sensitive. E.G., --like 苍井空 小泽玛利亚. ")
	fmt.Print("When --like keywords list conflict with --hate, --hate first. " + "\n")

	fmt.Print("\n")
	fmt.Print("  --proxy" + "\n")
	fmt.Print("  As you know, the government likes blocking adult websites, so, I do suggest you to set ")
	fmt.Print("--proxy option. Hardseed supports more proxys: " + "\n")
	fmt.Print("    a) GoAgent (STRONGLY recommended), --proxy http://127.0.0.1:8087" + "\n")
	fmt.Print("    b) shadowsocks, --proxy socks5://127.0.0.1:1080, or socks5h://127.0.0.1:1080" + "\n")
	fmt.Print("    c) SSH, --proxy socks4://127.0.0.1:7070" + "\n")
	fmt.Print("    d) VPN (PPTP and openVPN), --proxy \"\"" + "\n")
	fmt.Print("  It is important that you should know, you can set more proxys at the same time, split by ")
	fmt.Print("space-char ' '. As the --concurrent-tasks option says, each proxy could more than one concurrent ")
	fmt.Print("tasks, now, what about more proxys? Yes, yes, the speed of downloading seed and pictures is ")
	fmt.Print("very very fast. E.G., --concurrent-tasks 8 --proxy http://127.0.0.1:8087 socks5://127.0.0.1:1080 ")
	fmt.Print("socks4://127.0.0.1:7070, the number of concurrent tasks is 8*3. " + "\n")
	fmt.Print("  If you wanna how to install and configure various kinds of proxy, please access my homepage ")
	fmt.Print("\"3.2 搭梯翻墙\" https://github.com/yangyangwithgnu/the_new_world_linux#3.2 " + "\n")
	fmt.Print("  If you are not good at computer, there is a newest goagent for floks who are not good at computer ")
	fmt.Print("by me, yes, out of box. see https://github.com/yangyangwithgnu/goagent_out_of_box_yang ")
	fmt.Print("  The default http://127.0.0.1:8087. " + "\n")

	fmt.Print("\n")
	fmt.Print("  That's all. Any suggestions let me know by ")
	fmt.Print(g_myemail_color)
	fmt.Print(" or ")
	fmt.Print(g_mywebspace_color)
	fmt.Print(", big thanks to you. Kiddo, take care of your body. :-)" + "\n\n")
}

func showVersionInfo()  {
	fmt.Print(g_softname+" version " + g_version + "\n")
	fmt.Print("email " + g_myemail_color + "\n")
	fmt.Print("webspace " + g_mywebspace_color + "\n\n")
}


type Cmds struct {
	help bool
	version bool
	av_class string
	concurrent_tasks int64
	timeout_download_picture int64
	topics_range int64
	saveas_path string
	hate string
	like string
	proxy string
}

func main() {
	loginUser, _ := user.Current()
	args := os.Args
	if contains(args,"--help") || contains(args,"-h") {
		showHelpInfo()
		return
	}else if contains(args,"--version") {
		showVersionInfo()
		return
	}
	cmds := new(Cmds)

	cmds.saveas_path = utils.GetOrDefault("saveas-path",loginUser.HomeDir+string(filepath.Separator)+g_softname)
	cmds.concurrent_tasks,_ = strconv.ParseInt(utils.GetOrDefault("concurrent-tasks","8"), 10, 64)
	cmds.timeout_download_picture,_ = strconv.ParseInt(utils.GetOrDefault("timeout-download-picture","16"),10,64)
	cmds.topics_range,_ = strconv.ParseInt(utils.GetOrDefault("topics-range","64"), 10, 64)


	flag.StringVar(&cmds.av_class,"av-class",utils.GetOrDefault("av-class","aicheng_asia_mosaicked"),"视频类型")
	flag.Int64Var(&cmds.concurrent_tasks,"concurrent-tasks",cmds.concurrent_tasks,"并发数")
	flag.Int64Var(&cmds.timeout_download_picture,"timeout-download-picture",cmds.timeout_download_picture,"图片下载超时时间")
	flag.Int64Var(&cmds.topics_range,"topics-range",cmds.topics_range,"下载范围")
	flag.StringVar(&cmds.saveas_path,"saveas-path", cmds.saveas_path,"文件存放地方")
	flag.StringVar(&cmds.hate,"hate",utils.GetOrDefault("hate",""),"需要排除的关键词")
	flag.StringVar(&cmds.like,"like",utils.GetOrDefault("like",""),"重点取用的关键词")
	flag.StringVar(&cmds.proxy,"proxy",utils.GetOrDefault("proxy",""),"代理地址")
	flag.Parse()

	av_class_name := cmds.av_class

	urlprefix,err := utils.GetUrlPrfix(av_class_name)
	if err != nil {
		fmt.Printf("error![%v]\n", err)
		return
	}

	var b_aicheng bool = true

	path := cmds.saveas_path

	err = utils.PathMkdir(path)
	if err != nil {
		fmt.Printf("check failed![%v]\n", err)
		return
	}


	if ("caoliu_west_original" == av_class_name) {
		b_aicheng = false
	} else if ("caoliu_cartoon_original" == av_class_name) {
		b_aicheng = false
	} else if ("caoliu_asia_mosaicked_original" == av_class_name) {
		b_aicheng = false
	} else if ("caoliu_asia_non_mosaicked_original" == av_class_name) {
		b_aicheng = false
	} else if ("caoliu_west_reposted" == av_class_name) {
		b_aicheng = false
	} else if ("caoliu_cartoon_reposted" == av_class_name) {
		b_aicheng = false
	} else if ("caoliu_asia_mosaicked_reposted" == av_class_name) {
		b_aicheng = false
	} else if ("caoliu_asia_non_mosaicked_reposted" == av_class_name) {
		b_aicheng = false
	} else if ("caoliu_selfie" == av_class_name) {
		b_aicheng = false
	}  else if ("aicheng_west" == av_class_name) {
		b_aicheng = true
	} else if ("aicheng_asia_mosaicked" == av_class_name) {
		b_aicheng = true
	} else if ("aicheng_cartoon" == av_class_name) {
		b_aicheng = true
	} else if ("aicheng_asia_non_mosaicked" == av_class_name) {
		b_aicheng = true
	} else {
		fmt.Println("ERROR! the --av-class argument invalid. More info please see --help. \n")
		return
	}

	fmt.Print("  the av class \""+av_class_name+"\"; \n")

	fmt.Print("  the download picture timeout seconds ")
	fmt.Print("\"" + strconv.FormatInt(cmds.timeout_download_picture,10) + "\"; \n")

	fmt.Print("  the range of parsing topics ")
	fmt.Print("[1~" +strconv.FormatInt(cmds.topics_range,10)+ "]; \n")

	path += string(filepath.Separator)+"["
	path += av_class_name
	path += "]"
	path += "[1~" + strconv.FormatInt(cmds.topics_range,10)  + "]@"
	path += time.Now().Format("150405")

	err = utils.PathMkdir(path)
	if err != nil {
		fmt.Printf("check dir error![%v]\n", err)
		return
	}

	fmt.Print("  the path to save seeds and pictures \""+path+"\"; \n")
	fmt.Print("  the number of concurrent tasks \"" + strconv.FormatInt(cmds.concurrent_tasks,10) + "\"; \n")

	hate_keywords_list := [] string { "连发", "連发", "连發", "連發",
		"连弹", "★㊣", "合辑", "合集",
		"合輯", "nike", "最新の美女骑兵㊣",
		"精選", "精选","版规" }

	var splits [] string
	if len(cmds.hate)>1{
		splits = strings.Split(cmds.hate,",")
	}
	if splits != nil{
		for i := 0; i< len(splits);i++  {
			hate_keywords_list = append(hate_keywords_list, splits[i])
		}
	}
	fmt.Printf("  ignore some topics which include the keywords as follow \"%v\";\n" , hate_keywords_list)


	like_keywords_list := strings.Split(cmds.like,",")
	fmt.Printf("  just parse topics which include the kewords as follow \"%v\";\n",like_keywords_list)

	fmt.Printf("  the proxy \"%s\";\n",cmds.proxy)

	caoliu_portal_url := utils.GetOrDefault("caoliu","http://t66y.com/")
	icheng_portal_url := utils.GetOrDefault("aicheng","http://www.ac168.info/bt/")

	info := new(vo.DownloadInfo)
	info.HateList = hate_keywords_list
	info.LikeList = like_keywords_list
	info.PageStart = 1
	info.PageEnd = cmds.topics_range
	info.Path = path
	info.Proxy = cmds.proxy
	info.TimeoutDownloadPic = cmds.timeout_download_picture
	info.Urlprefix = urlprefix

	if b_aicheng {
		info.Host = icheng_portal_url
		icheng_portal_url += urlprefix
		info.Url = icheng_portal_url
		icheng.DownloadBt(info)
	}else {
		info.Host = caoliu_portal_url
		caoliu_portal_url += urlprefix
		info.Url = caoliu_portal_url
		caoliu.DownloadHtml(info)
	}
	return
}
