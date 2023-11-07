
功能

	a simple ini file parser
	一个简单的ini文件的解析器


ini文件(ini file)

```

[global]

pid = run/php-fpm.pid
error_log = /data1/logs/php-fpm_error.log
log_level = notice
emergency_restart_threshold = 10
emergency_restart_interval = 1m
process_control_timeout = 5s
daemonize = yes
```


简单的使用-simple use

```
package main

import (
	"fmt"
	"github.com/kjx98/golib/ini"
)

func main() {
	t_iniconfig()
}

func t_iniconfig() {
	cf, err := goini.ParserConfig("./conf/simple.conf", false)
	if err != nil {
		fmt.Println("get config faild")
		return
	}
	//获取配置
	fmt.Println(cf.GetConfig("global", "error_log",""))
	fmt.Println(cf.GetConfig("global", "log_level",""))
	fmt.Println(cf.GetConfig("global", "emergency_restart_threshold",""))
	fmt.Println(cf.GetConfig("global", "emergency_restart_interval",""))
	fmt.Println(cf.GetConfig("global", "process_control_timeout",""))
	fmt.Println(cf.GetConfig("global", "daemonize",""))

	//将配置写入到新的文件
	goini.SaveConfigToFile(cf, "./conf/simple_bak.conf")

	cf, err = goini.ParserConfig("./conf/simple_bak.conf", false)
	if err != nil {
		fmt.Println("get config faild")
		return
	}

	fmt.Println(cf.GetConfig("global", "error_log",""))
	fmt.Println(cf.GetConfig("global", "log_level",""))
	fmt.Println(cf.GetConfig("global", "emergency_restart_threshold",""))
	fmt.Println(cf.GetConfig("global", "emergency_restart_interval",""))
	fmt.Println(cf.GetConfig("global", "process_control_timeout",""))
	fmt.Println(cf.GetConfig("global", "daemonize",""))

}

```

代码输出 output
```
/data1/log/php-fpm_error.log
notice
10
1m
5s
yes
/data1/log/php-fpm_error.log
notice
10
1m
5s
yes
```





单元测试/性能测试 unittest/benchmarktest

```
BenchmarkParserConfig-4           	  100000	     12054 ns/op
BenchmarkParserConfig2-4          	  100000	     11961 ns/op
BenchmarkParserConfigParallel-4   	  200000	      7165 ns/op
PASS
ok  	config	4.174s
```
