package entity

import (
	"fmt"

	"github.com/cloudstateio/go-support/cloudstate/eventsourced"
	"github.com/golang/protobuf/proto"
	domain "github.wdf.sap.corp/ICN-Nanjing-Projects/Collaboration-board/credit/persistence"
)

type Credit struct {
	credit *domain.Credit
}

func NewCredit(eventsourced.EntityID) eventsourced.EntityHandler {
	return &Credit{
		credit: &domain.Credit{
			Quantity: 0,
		},
	}
}

func (r *Credit) HandleCommand(ctx *eventsourced.Context, name string, cmd proto.Message) (proto.Message, error) {
	switch c := cmd.(type) {
	case *GetUserCredit:
		return r.GetCredit(ctx, c)
	case *AddUserCredit:
		return r.AddCredit(ctx, c)
	default:
		return nil, nil
	}
}

func (r *Credit) GetCredit(*eventsourced.Context, *GetUserCredit) (*UserCredit, error) {
	return &UserCredit{
		Quantity: r.credit.Quantity,
	}, nil
}

func (r *Credit) AddCredit(ctx *eventsourced.Context, userCredit *AddUserCredit) (*UserCredit, error) {
	quantity := userCredit.GetQuantity()
	if quantity <= 0 {
		return nil, fmt.Errorf("cannot add non-positive quantity %d to user credit", quantity)
	}
	ctx.Emit(&domain.CreditAdded{
		Quantity: quantity,
	})
	return &UserCredit{
		Quantity: r.credit.Quantity,
	}, nil
}

func (r *Credit) HandleEvent(ctx *eventsourced.Context, event interface{}) error {
	switch e := event.(type) {
	case *domain.CreditAdded:
		return r.CreditAdded(e)
	default:
		return nil
	}
}

func (r *Credit) CreditAdded(added *domain.CreditAdded) error {
	r.credit.Quantity += added.GetQuantity()
	return nil
}
