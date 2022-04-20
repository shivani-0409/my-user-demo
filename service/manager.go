package service
import(
	"time"
	"github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1/domain"
	"github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1/db"
	"github.com/segmentio/ksuid"
)

type manager struct{
	ds db.DataStore
}

func (m *manager) CreateUser(input *domain.User) (string,error){
	input.CreatedAt=time.Now().UTC()
	input.ID=ksuid.New().String()
		if len(input.Name) <=2 {
			return "",&domain.Error{Code:400,Message:"Name should be atleast 3 chars long"}
		}
		return m.ds.AddUser(input)
}

func (m *manager) ListUsers(filters map[string]interface{}, limit int64) ([]*domain.User,error){
	return m.ds.ListUsers(filters,limit)
}

func (m *manager) DeleteUser(id string) error{
	if _,err:=m.ds.ViewUser(id); err != nil{
		return &domain.Error{Code:404,Message:"User doesn't exist"}
	}
	return  m.ds.DeleteUser(id)
  }

func (m *manager) ViewUser(id string) (*domain.User,error){
	return m.ds.ViewUser(id)
  }
   
   