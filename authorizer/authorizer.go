package authorizer

import (
	"fmt"
	"log"

	"github.com/turnkeyca/monolith/db"
)

type Authorizer struct {
	logger *log.Logger
	db     *db.Database
}

func New(logger *log.Logger, db *db.Database) *Authorizer {
	return &Authorizer{
		logger: logger,
		db:     db,
	}
}

func (a *Authorizer) CheckUserIdAndToken(userId string, loggedInUserId string, perm PermissionType) error {
	var count []int
	err := a.db.Select(&count, `select count(*) from "permission" where user_id=$1 and on_user_id=$2 and "permission"=$3`, loggedInUserId, userId, perm)
	if err != nil {
		return fmt.Errorf("user [%s] does not have permission for [%s]: %s", loggedInUserId, userId, err)
	}
	if count[0] <= 0 {
		return fmt.Errorf("user [%s] does not have permission for [%s]", loggedInUserId, userId)
	}
	return a.CheckUserActive(userId)
}

func (a *Authorizer) CheckUserIdAndTokenNoActiveCheck(userId string, loggedInUserId string, perm PermissionType) error {
	var count []int
	err := a.db.Select(&count, `select count(*) from "permission" where user_id=$1 and on_user_id=$2 and "permission"=$3`, loggedInUserId, userId, perm)
	if err != nil {
		return fmt.Errorf("user [%s] does not have permission for [%s]: %s", loggedInUserId, userId, err)
	}
	if count[0] <= 0 {
		return fmt.Errorf("user [%s] does not have permission for [%s]", loggedInUserId, userId)
	}
	return nil
}

func (a *Authorizer) CheckUserIdsAndTokenAny(userIds []string, loggedInUserId string, perm PermissionType) error {
	var err error
	for _, userId := range userIds {
		err = a.CheckUserIdAndToken(userId, loggedInUserId, perm)
		if err == nil {
			return nil
		}
	}
	return err
}

func (a *Authorizer) CheckUserActive(userId string) error {
	var count []int
	err := a.db.Select(&count, `select count(*) from users where id = $1 and user_status = 'inactive';`, userId)
	if err != nil {
		return err
	}
	if count[0] > 0 {
		return fmt.Errorf("user [%s] is inactive", userId)
	}
	return nil
}
