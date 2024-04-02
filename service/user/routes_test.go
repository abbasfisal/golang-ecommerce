package user

import (
	"github.com/abbasfisal/golang-ecommerce/types"
	"testing"
)

func TestUser(t *testing.T) {
	//userStore := &mockUserStore{}
	//Handler := NewHandler(userStore)
	//
	//t.Run("user payload is invalid", func(t *testing.T) {
	//	payload, _ := json.Marshal(types.RegisterUserPayload{
	//		FirstName: "ali",
	//		LastName:  "momeni",
	//		Email:     "",
	//		Password:  "password",
	//	})
	//	req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(payload))
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//	_ := req
	//})
}

type mockUserStore struct {
}

func (s *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, nil
}
func (s *mockUserStore) GetUserById(id int) (*types.User, error) {
	return nil, nil
}
func (s *mockUserStore) CreateUser(user types.User) error {
	return nil
}
