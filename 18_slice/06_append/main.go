package main

func main() {
	greetings := make([]string, 3, 5)

	greetings[0] = "Hello"
	greetings[1] = "Hi"
	greetings[2] = "Xin chao"
	//greetings[3] = "bonjour"

	append(greetings, "bonjour")
}
