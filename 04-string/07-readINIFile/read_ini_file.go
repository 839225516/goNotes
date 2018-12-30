package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

//从ini文件中取值的函数
//	filename: ini 文件的文件名
//	expectSection: 期望读取的段
//	expectKey: 期望读取段中的键
func getValue(filename, expectSection, expectKey string) string {

	file, err := os.Open(filename)
	if err != nil {
		log.Println("打不开文件：", file)
	}

	//函数结束时，关闭文件
	defer file.Close()

	//使用读取器读取文件
	reader := bufio.NewReader(file)

	//当前读取的段名
	var sectionName string

	for {

		//读取文件的一行
		linestr, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		//切掉左右两边的空格符
		linestr = strings.TrimSpace(linestr)

		//忽略空行
		if linestr == "" {
			continue
		}

		//忽略注释
		if linestr[0] == ';' {
			continue
		}

		//读取段和键值

		//读了段
		//行首和行尾分别是方括号的，说明是段
		if linestr[0] == '[' && linestr[len(linestr)-1] == ']' {
			//取出段名
			sectionName = linestr[1 : len(linestr)-1]

			//如果段名为期望的段名，则查找key
		} else if sectionName == expectSection {
			pair := strings.Split(linestr, "=")

			//保证切开后只有一个等号分割的键值情况
			if len(pair) == 2 {

				//去掉键多余的空白符
				key := strings.TrimSpace(pair[0])

				if key == expectKey {
					//返回去掉空白符的值
					return strings.TrimSpace(pair[1])
				}
			}
		}

	}

	//如果没有找到则返回空
	return ""
}

func main() {
	iniFile := "conf.ini"
	section := "remote \"origin\""
	key := "url"

	remoteurl := getValue(iniFile, section, key)
	fmt.Println(remoteurl)

}
