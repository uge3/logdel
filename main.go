package main

import (
	"fmt"
	"os"
	"time"

	"github.com/astaxie/beego/config" //读取配置文件
)

//读取配置文件
func readconfig() (file_path string, maxsize int, unit int, ratio int, delay int) {
	conf, err := config.NewConfig("ini", "./config.conf") //读取一个新的配置文件
	if err != nil {
		fmt.Println("new config failed err:", err)
		return
	}
	fmt.Println("conf内容", conf)
	file_path = conf.String("server::file_path")
	fmt.Println("文件路径：", file_path)
	maxsize, err = conf.Int("server::maxsize")
	if err != nil {
		fmt.Println("获取最大文件出错!server::maxsize failed err:", err)
		return
	}
	unit, err = conf.Int("server::unit")
	if err != nil {
		fmt.Println("获取单位出错!server::unit failed err:", err)
		return
	}
	ratio, err = conf.Int("server::ratio")
	if err != nil {
		fmt.Println("获取倍数出错!server::ratio failed err:", err)
		return
	}
	delay, err = conf.Int("server::delay")
	if err != nil {
		fmt.Println("获取等待时间出错!server::delay failed err:", err)
		return
	}
	return string(file_path), int(maxsize), int(unit), int(ratio), int(delay)
}

func main() {
	var data = ""
	var maxsize = 0
	var unit = 0
	var ratio = 0
	var delay = 0
	data, maxsize, unit, ratio, delay = readconfig()
	fmt.Println("文件路径+文件名:", data)
	fmt.Println("最大值：", maxsize)
	fmt.Println("单位：", unit)
	fmt.Println("倍数：", ratio)
	fmt.Println("等待时间", delay)
	deltime := 0 //休眠时间
	for true {
		fmt.Println("停止时间为:", deltime, "秒！")
		time.Sleep(time.Duration(deltime) * time.Second) //进入休眠
		fi, err := os.Stat(data)                         //读取文件信息
		if err == nil {
			fmt.Println("name:", fi.Name())    //文件名
			fmt.Println("size:", fi.Size())    //文件大小
			fmt.Println("is dir:", fi.IsDir()) //文件路径
			fmt.Println("mode::", fi.Mode())
			fmt.Println("modTime:", fi.ModTime())
		} else if err != nil {
			fmt.Println("查询文件错误:", err)
			fmt.Println(fi)
		}
		if fi != nil {
			fisize := (fi.Size() / int64(unit)) * int64(ratio) //以k为单位
			fmt.Println("文件大小(k):", fisize)
			if int(fisize) > maxsize {
				err := os.Remove(data) //如果超过额定大小，进行删除
				if err != nil {
					fmt.Println("删除失败") // 删除失败
				} else {
					fmt.Println("删除成功") // 删除成功
				}
			} else {
				fmt.Println("文件不够大")
				deltime = int(float64(maxsize) - float64(fisize)) //对再次删除的时间进行计算，并退出当前循环
				continue
			}
		} else {
			fmt.Println("没有该文件!")
			deltime = delay //没有文件等待生成
			//deltime = 10
			continue
		}
	}
}
