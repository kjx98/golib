package ini

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"strings"
)

// 将数据写入到文件, 无法保证配置的顺序
func SaveConfigToFile(cf *IniConfig, filename string) bool {
	fp, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Println("open failed")
		return false
	}
	defer fp.Close()

	for secName, sections := range cf.Sections {
		fp.WriteString("[" + secName + "]\n")
		for _, val := range sections {
			fp.WriteString(val.key + "=" + val.val + "\n")
		}
	}

	return true
}

// 解析config
func ParserConfig(filename string, reload bool) (*IniConfig, error) {
	if !reload {
		if cf, ok := configList[filename]; ok {
			return cf, nil
		}
	}
	if acf, err := parserFile(filename); err == nil {
		configList[filename] = acf
		return acf, err
	} else {
		return nil, err
	}
}

// 真正的解析config的文件
func parserFile(filename string) (*IniConfig, error) {

	acf, err := NewIniConfig()
	if err != nil {
		return nil, err
	}

	fp, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	bufRead := bufio.NewReader(fp)

	var curSectionName string
	curSectionName = ""
	for {

		line, err := bufRead.ReadBytes('\n')
		if len(line) == 0 {
			break
		}
		//去除空格
		line = bytes.TrimSpace(line)
		//去除空行
		if len(line) == 0 {
			continue
		}
		if line[0] == '[' {
			line = bytes.Trim(line, "[")
			line = bytes.Trim(line, "]")
			curSectionName = string(line)
		} else {
			if line[0] == '#' || line[0] == ';' {
				continue
			}
			if len(curSectionName) == 0 {
				return nil, FORMAT_ERROR
			}
			//处理内容
			lineString := string(line)
			lineArr := strings.Split(lineString, "=")

			if len(lineArr) == 2 {
				dataKey := strings.TrimSpace(lineArr[0])
				dataVal := strings.TrimSpace(lineArr[1])
				//获取内容
				if _, ok := acf.Sections[curSectionName]; !ok {
					acf.Sections[curSectionName] = make(map[string]*sectionConfig)
				}
				acf.Sections[curSectionName][dataKey] = NewSection(dataKey, dataVal)
				//log.Println(acf.Sections[curSectionName][dataKey].key)
			}
		}
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}

	}
	return acf, nil
}
