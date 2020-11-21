package gbk

import (
    "fmt"
    "testing"
)

//测试获取数据
func TestParserConfig(t *testing.T) {
	//测试获取数据
    s := "GBK 与 UTF-8 编码转换测试"
    gbk, err := Utf8ToGbk([]byte(s))
    if err != nil {
        t.Error(err)
    } else {
        fmt.Println(string(gbk))
    }

    utf8, err := GbkToUtf8(gbk)
    if err != nil {
        t.Error(err)
    } else {
        fmt.Println(string(utf8))
    }
}

func BenchmarkGbk2Utf(b *testing.B) {
    s := "GBK 与 UTF-8 编码转换测试"
	gbk, _ := Utf8ToGbk([]byte(s))
	for i := 0; i < b.N; i++ {
		if _, err := GbkToUtf8(gbk); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkUtf2Gbk(b *testing.B) {
    s := "GBK 与 UTF-8 编码转换测试"
	for i := 0; i < b.N; i++ {
		_, _ = Utf8ToGbk([]byte(s))
	}
}
