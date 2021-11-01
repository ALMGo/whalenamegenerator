package whalenamegenerator

import "strings"

func GenerateWhaleName() string {
	adjective := strings.Title(RandomAdjective())
	color := strings.Title(RandomColor())
	whale := strings.Title(RandomWhale())
	return adjective + " " + color + " " + whale
}
