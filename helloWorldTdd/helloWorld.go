package main

import "fmt"

func HelloWorld(name, lang string) string {
	wholeSentence := generateReply(name, lang)
	return wholeSentence
}

func correctInput(name string) string {
	if name == "" {
		return "World"
	} else {
		return name
	}
}

func generateReply(name, lang string) string {
	var sentence string
	switch lang {
	case "jp":
		sentence = fmt.Sprintf("こにちは、%v!", correctInput(name))
	case "zh":
		sentence = fmt.Sprintf("你好，%v!", correctInput(name))
	default:
		sentence = fmt.Sprintf("Hello %v!", correctInput(name))
	}
	return sentence
}

func main() {

}
