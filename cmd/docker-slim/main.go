package main

import (
	"os"

	"github.com/docker-slim/docker-slim/pkg/app/master"
)

// docker-slim ��ڣ�
func main() {
	if len(os.Args) > 1 && os.Args[1] == "slim" {
		//hack to handle plugin invocations
		os.Args = append([]string{os.Args[0]}, os.Args[2:]...)
	}

	app.Run()
}
