package user

import (
	"context"
	"time"
	dto "user-service/app/api-user/dto"
	modelUser "user-service/models/user"
	cErr "user-service/pkg/err"
	"user-service/pkg/jwt"
	"user-service/pkg/password"
)

type UserService struct {
	storer UserStorer
	jwt    jwt.Jwt
	pass   password.Password
}

func NewUserService(store UserStorer, jwt jwt.Jwt, pass password.Password) *UserService {
	return &UserService{
		storer: store,
		jwt:    jwt,
		pass:   pass,
	}
}

func (u *UserService) CreateOne(ctx context.Context, user dto.UserDTO) (dto.UserResp, error) {
	var (
		userEntity = user.ToEntity("")
		hashedPass = u.pass.Hashed(user.Password)
	)

	userEntity.Password = hashedPass
	err := u.storer.CreateOne(ctx, userEntity)
	if err != nil {
		return dto.UserResp{}, err
	}

	return dto.UserResp{
		Id: userEntity.UserId,
	}, nil
}

func (u *UserService) Login(ctx context.Context, body dto.LoginDTO) (dto.LoginResp, error) {
	var (
		now    = time.Now()
		result = dto.LoginResp{
			ExpiredIn: now.Add(time.Hour * 2).Unix(),
		}
	)

	user, err := u.storer.GetByEmailOrPhoneNumber(ctx, body.Username)
	if err != nil {
		return result, err
	}

	status := u.pass.CompareHashedAndPassword(body.Password, user.Password)
	if !status {
		return result, cErr.UnAuthorizedError
	}

	claims := modelUser.Claims{
		Id:        user.UserId,
		ExpiredIn: result.ExpiredIn,
	}

	token, err := u.jwt.CreateToken(claims)
	if err != nil {
		return result, cErr.UnAuthorizedError
	}

	result.AccessToken = token.AccessToken
	result.RefreshToken = token.RefreshToken

	return result, nil
}

func (u *UserService) RefreshToken(ctx context.Context, body dto.RefreshTokenDTO) (dto.LoginResp, error) {
	var (
		now    = time.Now()
		result = dto.LoginResp{
			ExpiredIn: now.Add(time.Hour * 2).Unix(),
		}
	)

	token := modelUser.Token{
		AccessToken:  body.AccessToken,
		RefreshToken: body.RefreshToken,
	}

	user, err := u.jwt.ValidateRefreshToken(token)
	if err != nil {
		return result, cErr.InvalidToken
	}

	claims := modelUser.Claims{
		Id:        user.Id,
		ExpiredIn: result.ExpiredIn,
	}

	newToken, err := u.jwt.CreateToken(claims)
	if err != nil {
		return result, cErr.UnAuthorizedError
	}

	result.AccessToken = newToken.AccessToken
	result.RefreshToken = newToken.RefreshToken

	return result, nil
}
