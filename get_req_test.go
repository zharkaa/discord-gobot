package request_test

import (
	"fmt"
	"testing"
	"main"
)

func TestRequest(t *testing.T) {
	result := get_Quotes()
	fmt.Println(result)
}