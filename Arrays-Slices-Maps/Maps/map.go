package main

import "fmt"

func main(){
	websites := map[string]string {
		"Google": "https://www.google.com",
		"Facebook": "https://www.facebook.com",
		"Amazon Web Service": "https://aws.amazon.com",
	}

	fmt.Println(websites)
	fmt.Println(websites["Amazon Web Service"])

	websites["Microsoft"] = "https://www.microsoft.com"
	fmt.Println(websites)

	delete(websites, "Facebook")
	fmt.Println(websites)
}