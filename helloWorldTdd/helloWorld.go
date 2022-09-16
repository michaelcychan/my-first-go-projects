package main

func HelloWorld(name string) string {
	return "Hello " + correctInput(name) + "!"
}

func correctInput(name string) string {
	if name == "" {
		return "World"
	} else {
		return name
	}
}

func main() {

}
