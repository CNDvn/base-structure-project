package services

import (
	"errors"
	"gobase/pkg/errormsg"
	"gobase/pkg/helpers"
	"gobase/pkg/reqdto"
	"gobase/pkg/resdto"
	"gobase/pkg/schemas"
	"gobase/pkg/utils"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type TAuth struct{}
type tJWTPayload struct {
	UID      primitive.ObjectID `json:"uid"`
	Username string             `json:"username"`
	Email    string             `json:"email"`
	jwt.RegisteredClaims
}

func (t *TAuth) SignUp(dto *reqdto.TSignUpWithUsername) (*resdto.SignUpSuccess, *utils.CustomError) {
	if dto == nil {
		utils.PrintLog("func (t *TAuth) SignUp", "data dto sign up nil")
		return nil, &utils.CustomError{
			Status:  http.StatusBadRequest,
			Message: errormsg.CAN_NOT_SIGN_UP,
		}
	}

	hashedPassword, err := t.HashPassword(dto.Password)
	if err != nil {
		utils.PrintLog("func (t *TAuth) SignUp", err.Error())
		return nil, &utils.CustomError{
			Status:  http.StatusBadRequest,
			Message: errormsg.CAN_NOT_SIGN_UP,
		}
	}

	newUser := &schemas.TUser{
		Name:     dto.Name,
		UserName: dto.Username,
		Password: hashedPassword,
	}

	if err := User.Create(newUser); err != nil {
		utils.PrintLog("func (t *TAuth) SignUp", err.Message)
		return nil, err
	}

	claim := &tJWTPayload{
		UID:      newUser.ID,
		Email:    newUser.Email,
		Username: newUser.UserName,
	}
	token, err := t.CreateJwt(claim)
	if err != nil {
		utils.PrintLog("func (t *TAuth) SignUp", "generate jwt fail")
		return nil, &utils.CustomError{
			Status:  http.StatusBadRequest,
			Message: errormsg.CAN_NOT_SIGN_UP,
		}
	}

	refreshToken, err := t.CreateRefreshToken(claim)
	if err != nil {
		utils.PrintLog("func (t *TAuth) SignUp", "generate refresh token fail")
		return nil, &utils.CustomError{
			Status:  http.StatusBadRequest,
			Message: errormsg.CAN_NOT_SIGN_UP,
		}
	}

	return &resdto.SignUpSuccess{
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}

func (t *TAuth) SignIn(dto *reqdto.TSignInWithUsername) (*resdto.SignInSuccess, *utils.CustomError) {
	if dto == nil {
		utils.PrintLog("func (t *TAuth) SignUp", "data dto sign in nil")
		return nil, &utils.CustomError{
			Status:  http.StatusBadRequest,
			Message: errormsg.SIGN_IN_FAIL,
		}
	}

	filter := bson.M{
		"user_name": dto.Username,
	}

	user, err := User.FindOne(filter)
	if err != nil {
		utils.PrintLog("func (t *TAuth) SignUp", err.Error())
		return nil, &utils.CustomError{
			Status:  http.StatusBadRequest,
			Message: errormsg.SIGN_IN_FAIL,
		}
	}

	if !t.CheckPasswordHash(dto.Password, user.Password) {
		utils.PrintLog("func (t *TAuth) SignUp", "Password wrong")
		return nil, &utils.CustomError{
			Status:  http.StatusBadRequest,
			Message: errormsg.SIGN_IN_FAIL,
		}
	}

	claim := &tJWTPayload{
		UID:      user.ID,
		Email:    user.Email,
		Username: user.UserName,
	}
	token, err := t.CreateJwt(claim)
	if err != nil {
		utils.PrintLog("func (t *TAuth) SignUp", "generate jwt fail")
		return nil, &utils.CustomError{
			Status:  http.StatusBadRequest,
			Message: errormsg.CAN_NOT_SIGN_UP,
		}
	}

	refreshToken, err := t.CreateRefreshToken(claim)
	if err != nil {
		utils.PrintLog("func (t *TAuth) SignUp", "generate refresh token fail")
		return nil, &utils.CustomError{
			Status:  http.StatusBadRequest,
			Message: errormsg.SIGN_IN_FAIL,
		}
	}

	return &resdto.SignInSuccess{
		SignUpSuccess: resdto.SignUpSuccess{
			Token:        token,
			RefreshToken: refreshToken,
		},
	}, nil
}

func (t *TAuth) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (t *TAuth) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (t *TAuth) CreateJwt(claims *tJWTPayload) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(expirationTime)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(helpers.GetENV().JWT_SECRET_KEY))
}

func (t *TAuth) CreateRefreshToken(claims *tJWTPayload) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)

	claims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(expirationTime)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(helpers.GetENV().REFRESH_TOKEN_SECRET_KEY))
}

func (t *TAuth) ParseJwt(jwtStr string) (*tJWTPayload, error) {
	claims := &tJWTPayload{}

	tkn, err := jwt.ParseWithClaims(jwtStr, claims, func(token *jwt.Token) (any, error) {
		return []byte(helpers.GetENV().JWT_SECRET_KEY), nil
	})
	if err != nil {
		// if err == jwt.ErrSignatureInvalid {
		// 	// w.WriteHeader(http.StatusUnauthorized)
		// 	return nil, err
		// }
		// w.WriteHeader(http.StatusBadRequest)
		return nil, err
	}
	if !tkn.Valid {
		// w.WriteHeader(http.StatusUnauthorized)
		return nil, errors.New("token invalid")
	}
	return claims, nil
}

func (t *TAuth) ParseRefreshToken(refreshTokenStr string) (*tJWTPayload, error) {
	claims := &tJWTPayload{}

	tkn, err := jwt.ParseWithClaims(refreshTokenStr, claims, func(token *jwt.Token) (any, error) {
		return []byte(helpers.GetENV().REFRESH_TOKEN_SECRET_KEY), nil
	})
	if err != nil {
		return nil, err
	}
	if !tkn.Valid {
		return nil, errors.New("token invalid")
	}
	return claims, nil
}
