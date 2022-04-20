package handlers
import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1/gen/restapi/operations/users"
	"github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1/gen/models"
	"github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1"
	"github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1/domain"
	// "log"
	"fmt"
)

func NewFindUser(rt *ShivaniCustomServerExample1.Runtime) users.FindUsersHandler{
	return &findUser{rt:rt}
}
type findUser struct{
	rt *ShivaniCustomServerExample1.Runtime
}
func (f *findUser) Handle(fup users.FindUsersParams) middleware.Responder{
	us,err := f.rt.GetManager().ListUsers(extractFilters(fup), *fup.Limit)
	if err != nil{
		fmt.Println(err)
		return users.NewFindUsersDefault(404).WithPayload(asErrorResponse(&domain.Error{Message:"No user found"}))
	}
	usResponse:=[]*models.User{}
	for _,usr:=range us{
		usResponse=append(usResponse,asUserResponse(usr))
	}
	res:=users.NewFindUsersOK().WithPayload(usResponse)
	return res
}
func extractFilters(fup users.FindUsersParams) map[string]interface{}{
	filterMap := make(map[string]interface{})
	if fup.Name != nil {
		filterMap["name"] = *fup.Name
	}
	return filterMap
}
