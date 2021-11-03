package permission

import "fmt"

func (a *Authorizer) CheckUserIdAndToken(userId string, loggedInUserId string, perm PermissionType) error {
	if userId == loggedInUserId {
		return nil
	}
	var count int
	err := a.db.Select(&count, `select count(*) from permission where user_id=$1 and on_user_id=$2 and permission=$3`, loggedInUserId, userId, perm)
	if err != nil && count > 0 {
		return nil
	}
	return fmt.Errorf("user [%s] does not have permission for [%s]", loggedInUserId, userId)
}
