package routes

import (
	"github.com/GeldNetworkMVP/GeldMVPBackend/apiModel"
	"github.com/GeldNetworkMVP/GeldMVPBackend/model"
)

var AppUserRoutes = model.Routers{
	//Route will be used to add a new User to DB
	model.Router{
		Name:    "Create AppUser",
		Method:  "POST",
		Path:    "/appuser/save",
		Handler: apiModel.CreateUser,
	},
	//Will return Users based on Users ID provided
	model.Router{
		Name:    "Get User by UserID",
		Method:  "GET",
		Path:    "/appuser/{_id}",
		Handler: apiModel.GetUsersByID,
	},
	//Will update Users based on Users ID provided
	model.Router{
		Name:    "Update Users",
		Method:  "PUT",
		Path:    "/updateuser",
		Handler: apiModel.UpdateUser,
	},
	//Will delete User based on User ID provided
	model.Router{
		Name:    "Delete User by UserID",
		Method:  "DELETE",
		Path:    "/appuser/remove/{_id}",
		Handler: apiModel.DeleteUserByID,
	},
	//Will return users based adminuserid paginated response
	model.Router{
		Name:    "Get user data pagination",
		Method:  "Get",
		Path:    "/appuser/admin/{userid}",
		Handler: apiModel.GetPaginatedUserData,
	},
}
