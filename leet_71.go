package main

import "strings"

func simplifyPath(path string) string {
	res := []string{}
	res = append(res, "/")
	ps := strings.Split(path, "/")

	for i := 0; i < len(ps); i++ {
		l := res[len(res)-1]
		switch ps[i] {
		case "":
			if l != "/" {
				res = append(res, "/")
			}
		case ".":
		case "..":
			if len(res) == 1 {
				continue
			}
			if l != "/" {
				res = res[:len(res)-1]
			} else {
				res = res[:len(res)-2]
			}

		default:
			if l != "/" {
				res = append(res, "/")
			}
			res = append(res, ps[i])
		}
	}

	if len(res) > 1 && res[len(res)-1] == "/" {
		res = res[:len(res)-1]
	}

	return strings.Join(res, "")
}
