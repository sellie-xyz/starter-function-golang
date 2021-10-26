package __cookiecutter_project_name__

import (
	"context"
	"create-payment-method/utils"
	"fmt"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
	"time"
)

const (
	StatusReady     = "ready"
	StatusUsed      = "used"
	StatusExpired   = "expired"
	StatusCancelled = "cancelled"
	StatusInUse     = "in_use"
)

var (
	validate *validator.Validate
)

type repository interface {
	SaveObject(ctx context.Context, object *Object) error
	SaveStateTransition(ctx context.Context, statusTransitions *StatusTransitions) error
	SaveCurrentState(ctx context.Context, currentStatus *CurrentStatus) error
}

type Usecase struct {
	Repository repository
}

func (u *Usecase) SaveObject(ctx context.Context, createObject *CreateObject) (Object, error) {

	validate = validator.New()
	if err := validate.Struct(*createObject); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return Object{}, validationErrors
	}


	object := Object{
		ID:           u.newID(),
		Status:   StatusReady,
		Metadata: createObject.Metadata,
	}

	err := u.Repository.SaveObject(ctx, &object)
	if err != nil {
		return Object{}, errors.Wrap(err, "error creating new method")
	}

	return object, nil
}

func (u *Usecase) newID() string {
	id, err := utils.NewID()
	if err != nil {
		log.Fatalln(err)
	}
	return 	fmt.Sprintf("_live_%s", id)
}


func (u *Usecase) setCurrentStatus(ctx context.Context, status *string) error {

	currentStatus := CurrentStatus{
		Status: "status",
		Date:   time.Time{},
	}

	err := u.Repository.SaveCurrentState(ctx, &currentStatus)
	if err != nil {
		return errors.Wrap(err, "error saving status")
	}

	return nil
}
