package main

import (
	"fmt"
	"github.com/mmcdole/gofeed"
	"os"
	"strings"
)

func main() {
	file, err := os.OpenFile("README.md", os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://holicc.github.io/index.xml")
	var builder strings.Builder
	//write header
	builder.WriteString("### （def Hi 👋 I'm Joe")
	builder.WriteString("\n")
	builder.WriteString("\n")
	//write github start
	builder.WriteString("<img align=\"right\" src=\"https://github-readme-stats.vercel.app/api?username=holicc&show_icons=true&icon_color=805AD5&text_color=718096&bg_color=ffffff&hide_title=true\" />\n")
	builder.WriteString("\n")
	//write content
	for i := 0; i < 7; i++ {
		item := feed.Items[i]
		content := fmt.Sprintf("* [%s](%s) - %s", item.Title, item.Link, item.Published[:17])
		builder.WriteString(content)
		builder.WriteString("\n")
	}
	//write footer
	builder.WriteString("More on [Job's Blog](https://holicc.github.io/)")
	//
	_, err = file.WriteString(builder.String())
	if err != nil {
		panic(err)
	} else {
		fmt.Println(builder.String())
	}
}
