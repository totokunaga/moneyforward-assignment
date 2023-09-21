package types

import (
	"encoding/json"
	"fmt"

	httpRequest "github.com/mfx/util/functions/httpRequest"
)

const endpointOrigin = "https://mfx-recruit-dev.herokuapp.com"

type User struct {
	Id       int
	Name     string
	Accounts []GetUserAccountType
}

// Fill required user information by calling the APIs
func (user *User) LoadData(userId int) error {
	user.Id = userId
	userCh := make(chan GetUserType)
	accountsCh := make(chan GetUserAccountType)
	errCh := make(chan error)

	go func() {
		user, err := GetUser(userId)
		if err != nil {
			errCh <- err
			return
		}

		userCh <- *user
		close(userCh)
	}()

	go func() {
		accounts, err := GetUserAccount(userId)
		if err != nil {
			errCh <- err
			return
		}

		for _, item := range accounts {
			accountsCh <- item
		}
		close(accountsCh)
	}()

	for userChOpen, accountsChOpen := true, true; userChOpen || accountsChOpen; {
		select {
		case fetchedUser, ok := <-userCh:
			if !ok {
				userChOpen = false
				continue
			}
			user.Name = fetchedUser.Name
		case account, ok := <-accountsCh:
			if !ok {
				accountsChOpen = false
				continue
			}
			user.Accounts = append(user.Accounts, account)
		case err := <-errCh:
			return err
		}
	}

	return nil
}

func GetUser(userId int) (*GetUserType, error) {
	getUserEndpoint := fmt.Sprintf("%s/users/%d", endpointOrigin, userId)

	body, err := httpRequest.Get(getUserEndpoint)
	if err != nil {
		return nil, err
	}

	var user GetUserType
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, err
	}

	return &user, err
}

func GetUserAccount(userId int) ([]GetUserAccountType, error) {
	getUserAccountEndpoint := fmt.Sprintf("%s/users/%d/accounts", endpointOrigin, userId)

	body, err := httpRequest.Get(getUserAccountEndpoint)
	if err != nil {
		return nil, err
	}

	var accounts []GetUserAccountType
	err = json.Unmarshal(body, &accounts)
	if err != nil {
		return nil, err
	}

	return accounts, err
}
