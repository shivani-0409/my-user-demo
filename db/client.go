package db
import(
	"github.com/go-swagger/go-swagger/examples/ShivaniCustomServerExample1/domain"
)

type DataStore interface{
	AddUser(user *domain.User) (string, error)
	ListUsers(filters map[string]interface{},limit int64) ([]*domain.User,error)
	DeleteUser(id string) error
	ViewUser(id string) (*domain.User,error)
}
type DatastoreFactory func() (DataStore, error)
var factories map[string]DatastoreFactory
func RegisterDataStore(key string,value DatastoreFactory){
	if factories == nil{
		factories=make(map[string]DatastoreFactory)
	}
	factories[key]=value
}
func NewDataStore(dbType string) (DataStore,error){
	return factories[dbType]()
}