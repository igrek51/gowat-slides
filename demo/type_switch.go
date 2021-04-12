func UnknownToInt64(value interface{}) (int64, error) {
	switch value := value.(type) {
	case uint64:
		return int64(value), nil
	case uint32:
		return int64(value), nil
	case int64:
		return value, nil
	case int32:
		return int64(value), nil
	case float64:
		return int64(value), nil
	case float32:
		return int64(value), nil
	case bool:
		if value {
			return 1, nil
		} else {
			return 0, nil
		}
	case string:
		vali, err := strconv.Atoi(value)
		if err != nil {
			return 0, errors.Wrapf(err, "converting %s to int64 failed", value)
		}
		return int64(vali), nil
	case json.Number:
		vali, err := value.Int64()
		if err != nil {
			return 0, errors.Wrapf(err, "converting json.Number %s to int64 failed", value)
		}
		return vali, nil
	default:
		// use "%#v" in order to ignore implemented String() representation
		stringVal := fmt.Sprintf("%#v", value)
		if strings.HasPrefix(stringVal, "0x") {
			stringVal := strings.Replace(stringVal, "0x", "", -1)
			vali, _ := strconv.ParseInt(stringVal, 16, 64)
			return vali, nil
		}
		vali, err := strconv.Atoi(stringVal)
		if err != nil {
			return 0, errors.Wrap(err, "converting to int64 failed")
		}
		return int64(vali), nil
	}
}
