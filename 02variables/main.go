package main

import "fmt"

const LoginToken string = "hgjfkdld"

func main() {
	var username string ="aditya"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type: %T \n", isLoggedIn)

	var smallVal int = 256
	fmt.Println(smallVal)
	fmt.Printf("variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.94736378
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type: %T \n", smallFloat)

	var anotherVal int 
	fmt.Println(anotherVal)
	fmt.Printf("Variable is of type: %T \n", anotherVal)

	var website = "learncodeonline.in"
	fmt.Println(website)

    // no var style
	numberOfUser := 3000000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)



}
