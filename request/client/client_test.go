package client

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCreateClientRequest_CreateRequest(t *testing.T) {
	expectedRequest := &ClientRequest{
		operation: "exec-operation", clientId: "client-id-1", requestNum: 1,
	}

	request, err := CreateRequest("REQUEST exec-operation client-id-1 1")
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(expectedRequest, request) {
		t.Error(fmt.Sprintf("Require: %#v Got: %#v", expectedRequest, request))
	}
}

func TestCreateClientRequest_InvalidRequest(t *testing.T) {
	_, err := CreateRequest("REQUEST exec-operation client-id-1")

	message := fmt.Sprint(err)

	if message != "wrong req: [REQUEST exec-operation client-id-1]" {
		t.Error(fmt.Sprintf("Require: %#v Got: %#v", "wrong req REQUEST exec-operation client-id-1", message))
	}
}
