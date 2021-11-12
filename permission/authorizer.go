package permission

import "fmt"

func (a *Authorizer) CheckUserIdAndToken(userId string, loggedInUserId string, perm PermissionType) error {
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
