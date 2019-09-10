package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/unidoc/unioffice/document"
)

const (
	A = "少年张小凡"
	B = "少年林惊羽"
	C = "小孩A"
	D = "小孩B"
	E = "小孩C"
	F = "普智"
	G = "黑衣人"
	H = "群杂"

	W = "场景"
	X = "旁白"
	Y = "<"
	Z = ">"
)

var names = []string{
	A, B, C, D, E, F, G, H, W, X, Y, Z,
}

func main() {
	doc, err := document.Open("zhuxian.docx")
	if nil != err {
		log.Fatalf("error opening document:%s", err)
	}

	key := ""
	info := ""
	dialogs := make(map[string][]string)
	for _, p := range doc.Paragraphs() {
		for _, run := range p.Runs() {
			isNew := false
			line := run.Text()
			for _, name := range names {
				if 0 == strings.Index(line, string(name)) {
					if len(key) > 0 {

						dialogs[key] = append(dialogs[key], info)
					}
					key = name
					info = ""
					isNew = true
					break
				}
			}

			if isNew {
				continue
			}

			if len(line) >= 1 {
				info = info + line
			}
		}
	}

	for name, ps := range dialogs {
		fmt.Println(name, ":")
		for index, p := range ps {
			fmt.Printf("【%d】:%s\n", index, p)
		}

		fmt.Println()
		fmt.Println()
	}
}
