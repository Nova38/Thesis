package rbac

import (
	_ "strings"
	// "github.com/rs/zerolog/log"
	_ "github.com/samber/lo"
	_ "github.com/samber/oops"

	_ "google.golang.org/protobuf/types/known/timestamppb"

	// "github.com/hyperledger-labs/cckit/identity"

	pb "github.com/nova38/thesis/lib/go/gen/rbac"
)

type AuthContractEval interface {
	GetCurrentUser() (*pb.User, error)
	GetCurrentUserId() (*pb.User_Id, error)
	GetUser(pb.GetUserRequest) (*pb.User, error)
	GetUserList() (*pb.User_List, error)

	GetCollection() (*pb.Collection, error)
	GetCollectionList() (*pb.Collection_List, error)
}

type AuthContractSubmit interface {
	UserRegister(*pb.UserRegisterRequest) (*pb.User, error)
	UserUpdateMembership(*pb.UpdateMembershipRequest) (*pb.User, error)
	CollectionCreate(*pb.CollectionCreateRequest) (*pb.Collection, error)
	CollectionUpdateRequest(*pb.CollectionUpdateRequest) (*pb.Collection, error)
}

type AuthService interface {
	AuthContractEval
	AuthContractSubmit
}

type AuthContractImpl struct{}
