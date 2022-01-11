package domain

import (
	"net/http"
	"testing"
)

var (
	userDaoMock usersDaoInterface
)

func TestGetUserNoUserFound(t *testing.T) {
	// Initialization

	user, err := userDaoMock.GetUser(0)
	if user != nil {
		t.Error("We are not expecting a user with UserID 0")
	}

	if err == nil {
		t.Error("We are expecting a user with UserID 0")
	}

	if err.StatusCode != http.StatusNotFound {
		t.Error("We are expecting 404 for a user with UserID 0")
	}
}
