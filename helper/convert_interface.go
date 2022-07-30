package helper

import "strconv"

func ConvertInterfaceToInt(data interface{}) int {
	var result int64

	_, status := data.(string)
	if status {
		price, err := strconv.ParseInt(data.(string), 10, 64)
		FatalIfError(err)
		result = price
	} else {
		result = int64(data.(float64))
	}

	return int(result)
}
