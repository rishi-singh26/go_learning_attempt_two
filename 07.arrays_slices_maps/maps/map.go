package maps

import "fmt"

func main() {
	websites := map[string]string{
		"Google":              "google.com",
		"Amazon Web Services": "aws.com",
	}

	fmt.Println(websites)

	fmt.Println(websites["Amazon Web Services"])
	fmt.Println(websites["Google"])

	websites["LinkedIn"] = "linkedin.com"
	fmt.Println(websites)

	delete(websites, "Google")
	fmt.Println(websites)
}
