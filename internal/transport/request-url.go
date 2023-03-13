package transport

import "fmt"

func BuildRequestURL(endpoint string, version string, resource string) string {
	return fmt.Sprintf("%s/%s/%s", endpoint, version, resource)
}
