package main

type Data struct {
	Name string
}

func dataMap() []*Data {
	return nil
}

func main() {
	result := dataMap()
	dataMapStr := make(map[string]*Data, len(result))

	for i, value := range result {
		dataMapStr[value.Name] = result[i]
	}

	println(len(dataMapStr))
	println(dataMapStr["name"] == nil)
}
