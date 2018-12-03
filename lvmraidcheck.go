package main

import (
	"fmt"
	"log"
	"os/exec"
	"reflect"
)

func main() {
	cmd := exec.Command("lvs")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))

	fmt.Println(reflect.TypeOf(out))

	for ii, rr := range out {
		fmt.Println(ii, rr)
	}
}
