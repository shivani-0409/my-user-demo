package handlers
import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1/gen/restapi/operations/users"
	"github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1"
	"github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1/domain"
)
type deleteUser struct{
	rt *ShivaniCustomServerExample1.Runtime
}
func NewDeleteUser(rt *ShivaniCustomServerExample1.Runtime) users.DeleteUserHandler{
	return &deleteUser{rt:rt}
}
func (d *deleteUser) Handle(del users.DeleteUserParams) middleware.Responder{
	if err := d.rt.GetManager().DeleteUser(del.ID); err!=nil{
		derr,ok:=err.(domain.Err)
		if ok {
			switch derr.StatusCode(){
				case 404:
					return users.NewDeleteUserNotFound().WithPayload(asErrorResponse(err.(*domain.Error)))
			}
		} else{
			return users.NewDeleteUserDefault(500).WithPayload(asErrorResponse(&domain.Error{Message:"Internal Server Error"}))
		}
	}
	return users.NewDeleteUserNoContent()
} 