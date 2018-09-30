package ini

import (
	"errors"
	//"log"
	"strconv"
)

const (
	DEFAULT_NAME   = "default_key"
	SectionComment = "comment-section-%s"
)

var (
	FORMAT_ERROR = errors.New("ini format error")
	NOT_FOUND    = errors.New("not found config")
)

type IniConfig struct {
	Sections map[string]sectionConfigs
}

type sectionConfig struct {
	key string
	val string //需要的时候可以去转换
}
type sectionConfigs map[string]*sectionConfig

//创建一个configlist
var configList = make(map[string]*IniConfig)

//创建一个空的config
func NewIniConfig() (*IniConfig, error) {
	acf := &IniConfig{}
	acf.Sections = make(map[string]sectionConfigs)

	return acf, nil
}

//获取配置, defV 为默认
func (cf *IniConfig) GetConfig(secting, name string, defV string) string {
	sec, ok := cf.Sections[secting][name]
	if !ok {
		return defV
	}
	return sec.val
}

//获取配置的各种方式 int, defV 为默认
func (cf *IniConfig) GetConfigInt(secting, name string, defV int) int {
	ret, ok := cf.Sections[secting][name]
	if !ok {
		return defV
	}
	i, err := strconv.Atoi(ret.val)
	if err != nil {
		return defV
	}
	return i
}

//更新或者创建配置
func (cf *IniConfig) PutConfig(secting, name, val string) bool {
	sec, ok := cf.Sections[secting]
	if !ok {
		cf.Sections[secting] = make(map[string]*sectionConfig)
		cf.Sections[secting][name] = NewSection(name, val)
		return true
	}
	sec[name].key = name
	sec[name].val = val
	return true
}

//删除配置
func (cf *IniConfig) DelConfigData(secting, name string) bool {
	_, ok := cf.Sections[secting][name]
	if !ok {
		return true
	}
	delete(cf.Sections[secting], name)
	return true
}

//获取配置的各种方式 double, defV 为默认
func (cf *IniConfig) GetConfigDouble(secting, name string, defV float64) float64 {
	ret, ok := cf.Sections[secting][name]
	if !ok {
		return defV
	}
	f, err := strconv.ParseFloat(ret.val, 64)
	if err != nil {
		return defV
	}
	return f
}

//创建一个配置
func NewSection(k, v string) *sectionConfig {
	return &sectionConfig{k, v}
}
