package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// type userService struct {
// 	userRepository repository.UserRepository
// }

func TestApiVersion(t *testing.T) {
	version := "1.0.0"
	assert.NotNil(t, version)
	assert.Equal(t, version, "1.0.0")
	fmt.Println("running test api version")
}

// func Test_IndexUsers(t *testing.T) {
// 	recorder := httptest.NewRecorder()
// 	_, err := http.NewRequest("GET", "/api/v1/users", nil)
// 	assert.Nil(t, err, nil)
// 	assert.Equal(t, recorder.Code, http.StatusOK)
// }

// func Test_CreateUsers(t *testing.T) {
// 	var input models.User
// 	input.ID = uuid.NewString()
// 	input.Name = "hasrul"
// 	input.Email = "hs.rhul@gmail.com"
// 	input.Password = "123"

// 	_, err := http.NewRequest("POST", "/api/v1/users", nil)
// 	// assert.Nil(t, err, nil)
// 	// assert.Equal(t, recorder.Code, http.StatusOK)
// }
