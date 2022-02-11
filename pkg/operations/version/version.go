package version

import "fmt"

const version string = "1.0.0"

func FetchVersion() string {
	return version
}

func DisplayVersion(ver string) {
	fmt.Printf("%s\n", ver)
}
