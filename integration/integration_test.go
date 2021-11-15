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
	err := godotenv.Load("../.env")
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

	//CREATE EMPLOYMENT
	err = createEmployment(t, cl, userId, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	err = createEmploymentIncorrectUserId(t, cl, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	err = createEmploymentValidationError(t, cl, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	//GET EMPLOYMENT
	employmentId, err := getEmploymentByUserId(t, cl, userId, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	err = getEmployment(t, cl, employmentId, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	err = getEmploymentNotFound(t, cl, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	err = getEmploymentByUserIdNotFound(t, cl, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	//UPDATE EMPLOYMENT
	err = updateEmployment(t, cl, userId, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	err = updateEmploymentNotFound(t, cl, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	err = updateEmploymentValidationError(t, cl, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	//DELETE EMPLOYMENT
	err = deleteEmployment(t, cl, employmentId, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	err = deleteEmploymentNotFound(t, cl, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}

	//CREATE PERMISSION
	//ACCEPT PERMISSION
	//GET PERMISSION
	//DELETE PERMISSION

	//CREATE PET
	err = createPet(t, cl, userId, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	err = createPetIncorrectUserId(t, cl, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	err = createPetValidationError(t, cl, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	//GET PET
	petId, err := getPetByUserId(t, cl, userId, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	err = getPet(t, cl, petId, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	err = getPetNotFound(t, cl, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	err = getPetByUserIdNotFound(t, cl, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	//UPDATE PET
	err = updatePet(t, cl, userId, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	err = updatePetNotFound(t, cl, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	err = updatePetValidationError(t, cl, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	//DELETE PET
	err = deletePet(t, cl, petId, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	err = deletePetNotFound(t, cl, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}

	//CREATE REFERENCE
	err = createReference(t, cl, userId, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	err = createReferenceIncorrectUserId(t, cl, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	err = createReferenceValidationError(t, cl, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	//GET REFERENCE
	referenceId, err := getReferenceByUserId(t, cl, userId, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	err = getReference(t, cl, petId, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	err = getReferenceNotFound(t, cl, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	err = getReferenceByUserIdNotFound(t, cl, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	//UPDATE REFERENCE
	err = updateReference(t, cl, userId, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	err = updateReferenceNotFound(t, cl, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	err = updateReferenceValidationError(t, cl, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	//DELETE REFERENCE
	err = deleteReference(t, cl, referenceId, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	err = deleteReferenceNotFound(t, cl, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}

	//CREATE ROOMMATE
	err = createRoommate(t, cl, userId, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	err = createRoommateIncorrectUserId(t, cl, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	err = createRoommateValidationError(t, cl, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	//GET ROOMMATE
	roommateId, err := getRoommateByUserId(t, cl, userId, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	err = getRoommate(t, cl, petId, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	err = getRoommateNotFound(t, cl, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	err = getRoommateByUserIdNotFound(t, cl, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	//UPDATE ROOMMATE
	err = updateRoommate(t, cl, userId, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	err = updateRoommateNotFound(t, cl, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	err = updateRoommateValidationError(t, cl, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	//DELETE ROOMMATE
	err = deleteRoommate(t, cl, roommateId, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	err = deleteRoommateNotFound(t, cl, token)
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
