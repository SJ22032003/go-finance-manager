package conversion

import "strconv"

func StringToFloat64(s string) (float64, error) {
	float, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.0, err
	}
	return float, nil
}
