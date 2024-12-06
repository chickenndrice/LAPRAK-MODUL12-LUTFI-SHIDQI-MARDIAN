package main

import "fmt"

func main() {
	var usn, pw string
	var kond int
	fmt.Scan(&usn, &pw)
	kond = 0
	
	for usn != "Admin" && pw != "Admin" {
		fmt.Scan(&usn, &pw)
		kond = kond + 1
	}
	fmt.Println(kond, "Percobaan gagal login")
}
