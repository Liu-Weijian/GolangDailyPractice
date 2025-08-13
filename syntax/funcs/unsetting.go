package main

func YouName(name string, aliases ...string) {

}

func main() {

	YouName("11")
	YouName("11", "aa")
	YouName("11", "aa", "bb")

	aliases := []string{"aa", "bb"}
	YouName("11", aliases...)
}
