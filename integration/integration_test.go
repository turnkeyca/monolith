package integration

import (
	"fmt"
	"os"
	"testing"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/joho/godotenv"
	"github.com/turnkeyca/monolith/integration/client"
)

//RH - this function is too long on purpose.
func Test(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	transport := httptransport.New(fmt.Sprintf(`localhost:%s`, os.Getenv("PORT")), "", nil)
	// LOG IN
	cl := client.New(transport, strfmt.Default)
	err = loginBadSecretShouldFail(t, cl)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	err = loginNewUserNewUserFlagFalse(t, cl)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	userId, token, err := login(t, cl)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	err = loginReturningUserNewUserFlag(t, cl, userId)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}

	// DELETE USER
	err = deleteUser(t, cl, userId, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	err = deleteUserNotFound(t, cl, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	// UPDATE USER
	err = updateUser(t, cl, userId, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	err = updateUserValidationError(t, cl, userId, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	err = updateUserNotFound(t, cl, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	//GET USER
	err = getUser(t, cl, userId, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	err = getUserNotFound(t, cl, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
}

func assert(obj1 string, obj2 string, name string) error {
	if obj1 != obj2 {
		return fmt.Errorf(`assert failed: expected %s to be %s but was %s`, name, obj1, obj2)
	}
	return nil
}

func assertTrue(obj bool, name string) error {
	if !obj {
		return fmt.Errorf(`assert failed: expected %s to be %t but was %t`, name, obj, !obj)
	}
	return nil
}
