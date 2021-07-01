package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/aureuneun/bitcoin/explorer"
	"github.com/aureuneun/bitcoin/rest"
)

func usage() {
	fmt.Printf("Welcome to bitcoin\n\n")
	fmt.Printf("Please use the following flags:\n\n")
	fmt.Printf("-mode		Choose 'html', 'rest' or 'both'\n")
	fmt.Printf("-port		Set port of the server\n")

	os.Exit(0)
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}

	// rest := flag.NewFlagSet("rest", flag.ExitOnError)
	// port := rest.Int("port", 4000, "Set port")
	// rest.Parse(os.Args[2:])

	// switch os.Args[1] {
	// case "rest":
	// 	fmt.Printf("rest %d\n", *port)
	// case "html":
	// 	fmt.Printf("html %d\n", *port)
	// default:
	// 	usage()
	// }

	mode := flag.String("mode", "rest", "Choose 'html', 'rest' or 'both'")
	port := flag.Int("port", 4000, "Set port of the server")
	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	case "both":
		go rest.Start(*port)
		explorer.Start(*port + 1)
	default:
		usage()
	}
}
