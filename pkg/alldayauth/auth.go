package alldayauth

import (
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/winai-pgm-itsystem/all-day-shop/config"
	"github.com/winai-pgm-itsystem/all-day-shop/modules/users"
)

type TonkenType string

const (
	Access  TonkenType = "access"
	Refresh TonkenType = "refresh"
	Admin   TonkenType = "admin"
	ApiKey  TonkenType = "apikey"
)

type alldayAuth struct {
	mapClaims *alldayMapClaims
	cfg       config.IJwtConfig
}

type alldayMapClaims struct {
	Claims *users.UserClaims `josn:"claims"`
	jwt.RegisteredClaims
}

type IAlldayAuth interface {
	SignToken() string
}

func jwtTimeDurationCal(t int) *jwt.NumericDate {
	return jwt.NewNumericDate(time.Now().Add(time.Duration(int64(t) * int64(math.Pow10(9)))))
}

func jwtTimeRepeatAdapter(t int64) *jwt.NumericDate {
	return jwt.NewNumericDate(time.Unix(t, 0))
}

func (a *alldayAuth) SignToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, a.mapClaims)
	ss, _ := token.SignedString(a.cfg.SecretKey())
	return ss
}

func ParseToken(cfg config.IJwtConfig, tokenString string) (*alldayMapClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &alldayMapClaims{}, func(t *jwt.Token) (interface{}, error) {
		method, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok || method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("signing method is invalid")
		}
		return cfg.SecretKey(), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return nil, fmt.Errorf("token format is invalid")
		} else if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, fmt.Errorf("token had expired")
		} else {
			return nil, fmt.Errorf("parse token failed: %v", err)
		}
	}
	if claims, ok := token.Claims.(*alldayMapClaims); ok {
		return claims, nil
	} else {
		return nil, fmt.Errorf("claims type is invalid")
	}

}

func RepeatToken(cfg config.IJwtConfig, claims *users.UserClaims, exp int64) string {
	obj := &alldayAuth{
		cfg: cfg,
		mapClaims: &alldayMapClaims{
			Claims: claims,
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    "allday-api",
				Subject:   "refresh-token",
				Audience:  []string{"customer", "admin"},
				ExpiresAt: jwtTimeRepeatAdapter(exp),
				NotBefore: jwt.NewNumericDate(time.Now()),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
			},
		},
	}
	return obj.SignToken()
}

func NewAlldayAuth(tokenType TonkenType, cfg config.IJwtConfig, claims *users.UserClaims) (IAlldayAuth, error) {
	switch tokenType {
	case Access:
		return newAccessToken(cfg, claims), nil
	case Refresh:
		return newRefreshToken(cfg, claims), nil
	default:
		return nil, fmt.Errorf("unkhnow token type")
	}
}

func newAccessToken(cfg config.IJwtConfig, claims *users.UserClaims) IAlldayAuth {
	return &alldayAuth{
		cfg: cfg,
		mapClaims: &alldayMapClaims{
			Claims: claims,
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    "alldayshop-api",
				Subject:   "access-token",
				Audience:  []string{"customer", "admin"},
				ExpiresAt: jwtTimeDurationCal(cfg.AccessExpiresAt()),
				NotBefore: jwt.NewNumericDate(time.Now()),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
			},
		},
	}
}

func newRefreshToken(cfg config.IJwtConfig, claims *users.UserClaims) IAlldayAuth {
	return &alldayAuth{
		cfg: cfg,
		mapClaims: &alldayMapClaims{
			Claims: claims,
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    "alldayshop-api",
				Subject:   "refresh-token",
				Audience:  []string{"customer", "admin"},
				ExpiresAt: jwtTimeDurationCal(cfg.RefreshExpiresAt()),
				NotBefore: jwt.NewNumericDate(time.Now()),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
			},
		},
	}
}
