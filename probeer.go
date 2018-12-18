package main

import (
	"fmt"
	"log"
	"os/exec"
	"reflect"
	"strings"
)

func main() {
	cmd := exec.Command("lvs")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))

	fmt.Printf("%v\n", reflect.TypeOf(out))

	for ii, ll := range strings.Split(string(out), "\n") {
		fmt.Println(ii, ll)
	}

}
