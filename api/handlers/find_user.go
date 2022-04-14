package handlers
import (
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1/gen/restapi/operations/users"
	"github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1/gen/models"
	"github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1"
	// "github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1/service"
)

// func NewFindUser(rt *runtime) users.FindUsersHandler{
func NewFindUser(rt *ShivaniCustomServerExample1.Runtime) users.FindUsersHandler{
	return &findUser{rt:rt}
}
type findUser struct{
	rt *ShivaniCustomServerExample1.Runtime
}
func (f *findUser) Handle(fup users.FindUsersParams) middleware.Responder{
	// fup.Name
	// v:=service.NewManager()
	// v.ListUsers()
	
	fmt.Println("Application Name is",f.rt.GetApplicationName())
	// n:="Shivani Sharma"
	// us:= []*models.User{{Address:"ABC",ID:2,Name:&n}}
	
	us,_:= f.rt.GetManager().ListUsers(fup.Name,fup.Limit)
	usResponse:=[]*models.User{}
	for _,usr:=range us{
		usResponse=append(usResponse,asUserResponse(usr))
	}
	res:=users.NewFindUsersOK().WithPayload(usResponse)
	return res
	// print the appication name
}

