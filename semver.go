// Semantic version comparing library.
package gosemver

import (
	"strings"
)

var replaceMap = map[string]string{"v": "", "-": "."}

// Compare 2 version formats for which is greater.
// For more information see http://semver.org
// Returns: 0 if equals, 1 if the first version arg is greater,
// 2 if the second, -1 if problem occured.
func Compare(v1, v2 string) int {
	//keywords := {"alpha,beta,rc,p"}
	for k, v := range replaceMap {
		if strings.Contains(v1, k) {
			strings.Replace(v1, k, v, -1)
		}
		if strings.Contains(v2, k) {
			strings.Replace(v2, k, v, -1)
		}
	}
	ver1 := strings.Split(v1, ".")
	ver2 := strings.Split(v2, ".")

	var shorter int
	if len(ver1) > len(ver2) {
		shorter = len(ver2)
	} else {
		shorter = len(ver1)
	}

	for i := 0; i < shorter; i++ {
		if ver1[i] == ver2[i] {
			if shorter-1 == i {
				if len(ver1) == len(ver2) {
					return 0
				} else {
					// todo check for keywords
					if len(ver1) > len(ver2) {
						return 1
					} else {
						return 2
					}
				}
			}
		} else if ver1[i] > ver2[i] {
			return 1
		} else {
			return 2
		}
	}
	return -1
}
