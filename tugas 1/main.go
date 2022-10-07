package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	println("Massukan huruf (konsonan/vocal) atau kata  (polindrone)")
	text, _ := reader.ReadString('\n')

	if len(text)-1 != 1 {
		polindrone(text)
	} else {
		vocalOrKonsonan(text)

		println("================================================polindrone===========================================")
	}
}
func vocalOrKonsonan(huruf string) {
	println("================================================ vocal or consonant ===========================================")
	println("vocal or konsonan")
	print(huruf, " adalah huruf ")
	huruf = string(huruf[0])
	if huruf == "a" || huruf == "i" || huruf == "u" || huruf == "e" || huruf == "o" {
		print(" vocal\n")
	} else {
		println("konsonan\n")
	}

}
func polindrone(input string) {
	input = strings.ToLower(input)
	for i := 0; i < len(input)/2; i++ {
		awal := string(input[i])
		akhir := string(input[(len(input) - 2 - i)])

		if awal != akhir {

			print("bukan palindrone")
			return
		}

	}
	print(" merupakan palindrone \n")
}
