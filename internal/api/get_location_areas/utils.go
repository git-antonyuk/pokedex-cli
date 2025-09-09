package api_get_location_areas

func ConvertLocationToNameList(items []LocationItem) []string {
	list := []string{}
	for _, item := range items {
		list = append(list, item.Name)
	}
	return list
}
