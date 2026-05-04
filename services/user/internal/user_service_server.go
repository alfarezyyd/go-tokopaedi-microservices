package internal

import (
	"context"
	"go-tokopaedi-microservices/pkg/exception"
	"go-tokopaedi-microservices/pkg/helper"
	"go-tokopaedi-microservices/pkg/validator"
	"go-tokopaedi-microservices/protobuf/genproto"

	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

type UserServiceServer struct {
	genproto.UnimplementedUserServiceServer
	userRepository    UserRepository
	validationService validator.Service
	gormDatabase      *gorm.DB
}

func NewUserServiceServer(userRepository UserRepository, gormDatabase *gorm.DB, validatorService validator.Service) UserServiceServer {
	return UserServiceServer{
		userRepository:    userRepository,
		gormDatabase:      gormDatabase,
		validationService: validatorService,
	}
}
func (userServiceServer *UserServiceServer) FindAll(
	ctx context.Context,
	_ *emptypb.Empty,
) (*genproto.QueryUserResponses, error) {

	var users []*genproto.QueryUserResponse

	err := userServiceServer.gormDatabase.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		userEntities, err := userServiceServer.userRepository.FindAll(tx)
		helper.CheckErrorOperation(err, exception.ToGrpcError(exception.ParseGormError(err)))

		for _, userEntity := range userEntities {
			userResponse := helper.MapEntityIntoResponse[User, *genproto.QueryUserResponse](userEntity)
			users = append(users, userResponse)
		}

		return nil
	})
	helper.CheckErrorOperation(err, exception.ToGrpcError(exception.ParseGormError(err)))
	return &genproto.QueryUserResponses{
		QueryUserResponses: users,
	}, nil
}

func (userServiceServer *UserServiceServer) HandleRegister(context context.Context, createUserRequest *genproto.CreateUserRequest) (*emptypb.Empty, error) {
	err := userServiceServer.validationService.ValidateStruct(createUserRequest)
	userServiceServer.validationService.ParseValidationError(err, createUserRequest)
	err = userServiceServer.gormDatabase.WithContext(context).Transaction(func(gormTransaction *gorm.DB) error {
		userEntity := helper.MapCreateRequestIntoEntity[genproto.CreateUserRequest, User](createUserRequest)
		err = userServiceServer.userRepository.Create(gormTransaction, userEntity)
		helper.CheckErrorOperation(err, exception.ToGrpcError(exception.ParseGormError(err)))
		return nil
	})
	helper.CheckErrorOperation(err, exception.ToGrpcError(exception.ParseGormError(err)))
	return nil, nil
}
