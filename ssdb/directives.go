package ssdb

import (
	"errors"
	"fmt"
	"strconv"
)

func (c *Client) Hsize(name string) (int, error) {
	result, err := c.Do("hsize", name)
	if err != nil {
		return 0, err
	}
	status := result[0]
	if status != "ok" {
		return 0, errors.New(status)
	}

	size, _ := strconv.Atoi(result[1])

	return size, nil
}

func (c *Client) Hget(name, key string) (string, error) {
	result, err := c.Do("hget", name, key)
	if err != nil {
		return "", err
	}
	status := result[0]
	if status != "ok" {
		return "", errors.New(status)
	}

	return result[1], nil
}

func (c *Client) Multi_hget(name string, keys []string) ([]string, error) {
	result, err := c.Do("multi_hget", name, keys)
	if err != nil {
		return nil, err
	}
	status := result[0]
	if status != "ok" {
		fmt.Println(result)
		return nil, errors.New(status)
	}

	return result[1:], nil
}

func (c *Client) Hset(name, key, value string) error {
	result, err := c.Do("hset", name, key, value)
	if err != nil {
		panic(err)
		return err
	}
	status := result[0]
	if status != "ok" {
		return errors.New(status)
	}

	return nil
}

func (c *Client) Hdel(name, key string) error {
	result, err := c.Do("hdel", name, key)
	if err != nil {
		return err
	}
	status := result[0]
	if status != "ok" {
		return errors.New(status)
	}

	return nil
}

func (c *Client) Hscan(name, key_start, key_end string, limit int) ([]string, error) {
	result, err := c.Do("hscan", name, key_start, key_end, limit)
	if err != nil {
		return nil, err
	}
	status := result[0]
	if status != "ok" {
		return nil, errors.New(status)
	}

	return result[1:], nil
}

func (c *Client) Hclear(name string) (int, error) {
	result, err := c.Do("hclear", name)
	if err != nil {
		return 0, err
	}
	status := result[0]
	if status != "ok" {
		return 0, errors.New(status)
	}

	amount, _ := strconv.Atoi(result[1])

	return amount, nil
}

func (c *Client) Zsize(name string) (int, error) {
	result, err := c.Do("zsize", name)
	if err != nil {
		return 0, err
	}
	status := result[0]
	if status != "ok" {
		return 0, errors.New(status)
	}

	size, _ := strconv.Atoi(result[1])

	return size, nil
}

func (c *Client) Zget(name, key string) (string, error) {
	result, err := c.Do("zget", name, key)
	if err != nil {
		return "", err
	}
	status := result[0]
	if status != "ok" {
		return "", errors.New(status)
	}

	return result[1], nil
}

func (c *Client) Zset(name, key string, score int64) error {
	result, err := c.Do("zset", name, key, score)
	if err != nil {
		return err
	}
	status := result[0]
	if status != "ok" {
		return errors.New(status)
	}

	return nil
}

func (c *Client) Zdel(name, key string) error {
	result, err := c.Do("zdel", name, key)
	if err != nil {
		return err
	}
	status := result[0]
	if status != "ok" {
		return errors.New(status)
	}

	return nil
}

func (c *Client) Zscan(name, key_start, score_start, score_end string, limit int) ([]string, error) {
	result, err := c.Do("zscan", name, key_start, score_start, score_end, limit)
	if err != nil {
		return nil, err
	}
	status := result[0]
	if status != "ok" {
		return nil, errors.New(status)
	}
	// fmt.Println(result)

	return result[1:], nil
}

func (c *Client) Zrscan(name, key_start, score_start, score_end string, limit int) ([]string, error) {
	result, err := c.Do("zrscan", name, key_start, score_start, score_end, limit)
	if err != nil {
		return nil, err
	}
	status := result[0]
	if status != "ok" {
		return nil, errors.New(status)
	}
	// fmt.Println(result)

	return result[1:], nil
}
