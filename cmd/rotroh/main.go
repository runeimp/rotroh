package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/runeimp/rotroh"
)

const (
	appName    = "RotRoh"
	appVersion = "0.1.0"
	usage      = `Usage: %s [OPTIONS] STRINGS

OPTIONS:
`
)

var appLabel = fmt.Sprintf("%s v%s", appName, appVersion)

func main() {
	base64Ptr := flag.Bool("base64", false, "Use Base64 codex")
	helpPtr := flag.Bool("help", false, "Display this help info")
	pipeModePtr := flag.Bool("pipe", false, "Do not output a newline at the end")
	rot13Ptr := flag.Bool("rot13", false, "Use ROT-13 transform")
	rot47Ptr := flag.Bool("rot47", false, "Use ROT-47 transform")
	rotCustomPtr := flag.String("rot-custom", "", "Use a custom ROT transform set")
	rotRohPtr := flag.Bool("rotroh", true, "Use RotRoh codex")
	versionPtr := flag.Bool("version", false, "Display version info")

	noRotRoh := false

	flag.Usage = func() {
		fmt.Println(appLabel)
		fmt.Fprintf(flag.CommandLine.Output(), usage, filepath.Base(os.Args[0]))

		flag.VisitAll(func(f *flag.Flag) {
			optionName := fmt.Sprintf("-%s", f.Name)
			if len(f.DefValue) > 0 {
				fmt.Fprintf(flag.CommandLine.Output(), "  %-15s  %s (default: %v)\n", optionName, f.Usage, f.DefValue)
			} else {
				fmt.Fprintf(flag.CommandLine.Output(), "  %-15s  %s\n", optionName+" SET", f.Usage)
			}
		})
		fmt.Println()
	}

	flag.Parse()

	if *helpPtr {
		flag.Usage()
		os.Exit(0)
	}

	if *versionPtr {
		fmt.Println(appLabel)
		os.Exit(0)
	}

	args := flag.Args()

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		b, _ := ioutil.ReadAll(os.Stdin)
		args = []string{string(b)}
	}

	for _, arg := range args {
		if *rot13Ptr && *rot47Ptr {
			fmt.Fprintf(os.Stderr, "You can only use one transform of either ROT-13, ROT-47, Base64, or RotRoh\n")
			os.Exit(1)
		}
		// fmt.Printf("arg: %q\n", arg)

		var err error
		result := arg

		if *rot47Ptr {
			// fmt.Println("ROT-47")
			result = rotroh.Rot47String(result)
			noRotRoh = true
		} else if *rot13Ptr {
			// fmt.Println("ROT-13")
			result = rotroh.Rot13String(result)
			noRotRoh = true
		} else if *rotCustomPtr != "" {
			// fmt.Println("RotCustom")
			result, err = rotroh.RotCustomString(result, *rotCustomPtr)
			noRotRoh = true
		} else if *base64Ptr {
			// fmt.Println("Base64")
			result, err = rotroh.Base64String(result)
			noRotRoh = true
		}

		if *rotRohPtr && noRotRoh == false {
			// fmt.Println("RotRoh")
			result, err = rotroh.RotRoh47String(result)
		}

		if err != nil {
			fmt.Println(err.Error())
		}
		if *pipeModePtr {
			fmt.Print(result)
		} else {
			fmt.Println(result)
		}
	}
}

func test() {
	result := rotroh.Rot13String("Hi Ma!")
	log.Printf("main() | result: %q\n", result)

	result = rotroh.Rot13String(result)
	log.Printf("main() | result: %q\n", result)

	result = rotroh.Rot47String("Ahegao!")
	log.Printf("main() | result: %q\n", result)

	result = rotroh.Rot47String(result)
	log.Printf("main() | result: %q\n", result)

	result, err := rotroh.RotRoh47String("Ahegao!")
	if err != nil {
		log.Printf("main() | error: %q\n", err.Error())
	} else {
		log.Printf("main() | result: %q\n", result)
	}

	result, err = rotroh.RotRoh47String(result)
	if err != nil {
		log.Printf("main() | error: %q\n", err.Error())
	} else {
		log.Printf("main() | result: %q\n", result)
	}
}
