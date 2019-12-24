package main

import (
	"fmt"
	"github.com/huichen/sego"
)



func main() {
	// 载入词典
	var segmenter sego.Segmenter
	segmenter.LoadDictionary(`test\分词\dictionary.txt`)

	// 分词
	text := []byte("1922	www.Torrenting.org - Catch 22 S01E06 720p HDTV x264-MORiTZ")
	segments := segmenter.Segment(text)
  
	// 处理分词结果
	// 支持普通模式和搜索模式两种分词，见代码中SegmentsToString函数的注释。
	fmt.Println(sego.SegmentsToSlice(segments,true))
for _, v := range sego.SegmentsToSlice(segments,true) {
	fmt.Println(v)
}
	
}