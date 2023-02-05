package token

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"time"
)

const TokenExpireDuration = time.Hour * 2

type MyClaims struct {
	UserID   uint64 `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

var MySecret = []byte("WaWa123!@#FGH")

// GenToken 生成JWT
func GenToken(userID uint64, username string) (string, error) {
	// 创建一个我们自己的声明的数据
	c := MyClaims{
		userID,
		username, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add( //这里注意INT不能直接与time.hour相乘,所以用time.duration转换一下
				time.Duration(viper.GetInt("auth.jet_expire")) * time.Hour).Unix(), // 过期时间
			Issuer: "bluebell", // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	// 令牌有效
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
