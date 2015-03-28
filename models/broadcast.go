package models

func SetBroadcast(content string) error {
	_, err := db.Set(broadcast, content)
	if err != nil {
		return err
	}
	return nil
}

func GetBroadcast() (string, error) {
	raw, err := db.Get(broadcast)
	if err != nil {
		return "", err
	}
	if raw == nil {
		return "", nil
	}

	return raw.(string), nil
}
