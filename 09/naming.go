package naming

func Color(name string) string {
	switch name {
	case "red":
		return "#FF0000"
	case "green":
		return "#00FF00"
	case "blue":
		return "#0000FF"
	default:
		return "#000000"
	}
}
