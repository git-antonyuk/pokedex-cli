package string_utils_print_list

import (
	"fmt"
	"strings"
)

func PrintList(list []string) {
	var sb strings.Builder
	for _, str := range list {
		fmt.Fprintf(&sb, "%s\n", str)
	}
	fmt.Print(sb.String())
}
