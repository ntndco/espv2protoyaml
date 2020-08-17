package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type backends []string

type backendRules []backendRule
type backendRule struct {
	Address  string
	Selector string
}

type usageRules []usageRule
type usageRule struct {
	Selector          string
	AllowUnregistered bool
}

func (b *backends) String() string {
	l, _ := json.MarshalIndent(b, "", " ")
	return fmt.Sprintf("%v\n", string(l))
}

// Set ...
func (b *backends) Set(value string) error {
	var err error
	for _, s := range strings.Split(value, ",") {
		*b = append(*b, s)
	}
	return err
}

// Set ...
func (br *backendRules) Set(value string) error {
	res := strings.Split(value, ",")
	if len(res) != 2 {
		return errors.New("Must provide two values")
	}
	r := backendRule{
		Address:  res[1],
		Selector: res[0],
	}
	*br = append(*br, r)
	return nil
}

// String ...
func (br *backendRules) String() string {
	l, _ := json.MarshalIndent(br, "", " ")
	return fmt.Sprintf("%v\n", string(l))
}

// Set ...
func (ur *usageRules) Set(value string) error {
	auc := false
	res := strings.Split(value, ",")
	if len(res) != 2 {
		return errors.New("Must provide two values")
	}
	if res[1] == "true" {
		auc = true
	}
	if res[1] != "true" || res[1] == "false" {
		return errors.New("Must provide bool value for 'allow unregistered calls' value")
	}
	r := usageRule{
		Selector:          res[0],
		AllowUnregistered: auc,
	}
	*ur = append(*ur, r)
	return nil
}

// String ...
func (ur *usageRules) String() string {
	l, _ := json.MarshalIndent(ur, "", " ")
	return fmt.Sprintf("%v\n", string(l))
}
