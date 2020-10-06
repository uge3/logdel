package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/astaxie/beego/config" //读取配置文件
)

var sema = make(chan struct{}, 20)

//读取JSON文件内的值
func filename() (names string) {
	//文件读取
	file, err := os.Open("./file_name.json")
	if err != nil {
		fmt.Println("读取错误", err)
		return
	}
	defer file.Close() //关闭文件

	readerfile := bufio.NewReader(file) //读取文件内容

	str, err := readerfile.ReadString('\n')
	if err == io.EOF { // io.EOF表示文件结尾

	}
	if err != nil {
		fmt.Printf("read filefailed,err:%v\n", readerfile)
	}

	names = str

	return names
}

//读取配置文件
func configfile() {
	conf, err := config.NewConfig("ini", "./config.conf") //读取一个新的配置文件

	if err != nil {
		fmt.Println("new config failed err:", err)
		return
	}
	fmt.Println("配置文件:", conf)
}

func main() {
	data := filename() //获取指定文件
	fmt.Println("json内文件名:", data)
	configfile()
	fi, err := os.Stat(data)
	if err == nil {
		fmt.Println("name:", fi.Name())
		fmt.Println("size:", fi.Size())
		fmt.Println("is dir:", fi.IsDir())
		fmt.Println("mode::", fi.Mode())
		fmt.Println("modTime:", fi.ModTime())
	} else if err != nil {
		fmt.Println("查询文件错误:", err)
		fmt.Println(fi)
	}
	if fi != nil {
		fisize := fi.Size() / 1024 * 1000
		fmt.Println("文件大小(k):", fisize)
		if int(fisize) > 1500 {
			err := os.Remove(data)

			if err != nil {
				fmt.Println("删除失败")
				// 删除失败

			} else {
				// 删除成功
				fmt.Println("删除成功")

			}
		}
	} else {
		fmt.Println("没有该文件!")
	}

}
