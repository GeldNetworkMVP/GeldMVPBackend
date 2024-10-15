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
		Name:      "Get User by ID",
		Method:    "GET",
		Path:      "/appuser/{_id}",
		Handler:   apiModel.GetUsersByID,
		Protected: true,
	},
	//Will update Users based on Users ID provided
	model.Router{
		Name:      "Update Users",
		Method:    "PUT",
		Path:      "/updateuser",
		Handler:   apiModel.UpdateUser,
		Protected: true,
	},
	//Will delete User based on User ID provided
	model.Router{
		Name:      "Delete User by ID",
		Method:    "DELETE",
		Path:      "/appuser/remove/{_id}",
		Handler:   apiModel.DeleteUserByID,
		Protected: true,
	},
	//Will return users based adminuserid paginated response
	// model.Router{
	// 	Name:    "Get user data pagination",
	// 	Method:  "Get",
	// 	Path:    "/appuser/admin/{userid}",
	// 	Handler: apiModel.GetPaginatedUserData,
	// },

	//Will return all the users
	model.Router{
		Name:      "Test normal Get All users",
		Method:    "GET",
		Path:      "/users",
		Handler:   apiModel.TestGetAllUsers,
		Protected: true,
	},
	//Will return Users based on status provided
	model.Router{
		Name:      "Get User by Status",
		Method:    "GET",
		Path:      "/appuser/status/{status}",
		Handler:   apiModel.GetUsersByStatus,
		Protected: true,
	},
	//Will update Users status based on Users ID provided
	model.Router{
		Name:      "Update Users Status",
		Method:    "PUT",
		Path:      "/updateuserstatus",
		Handler:   apiModel.UpdateUserStatus,
		Protected: true,
	},
	//Login
	model.Router{
		Name:    "User Sign In",
		Method:  "GET",
		Path:    "/usersignin",
		Handler: apiModel.UserSignIn,
	},
	//Check User Existence
	model.Router{
		Name:    "User Existence Check",
		Method:  "GET",
		Path:    "/userexists/{email}",
		Handler: apiModel.UserExistence,
	},
}
