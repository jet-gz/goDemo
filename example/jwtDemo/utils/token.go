package utils

import (
	"errors"
	"fmt"

	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	JwtTokenOK            int    = 200100                      //token有效
	JwtTokenInvalid       int    = -400100                     //无效的token
	JwtTokenExpired       int    = -400101                     //过期的token
	JwtTokenFormatErrCode int    = -400102                     //提交的 token 格式错误
	JwtTokenFormatErrMsg  string = "提交的 token 格式错误"            //提交的 token 格式错误
	JwtTokenMustValid     string = "token为必填项,请在请求header部分提交!" //提交的 token 格式错误

)

type userToken struct {
	userJwt *JwtSign
}

// CreateUserFactory 创建 userToken 工厂
func CreateUserFactory() *userToken {
	return &userToken{
		userJwt: CreateMyJWT("22223"), //(variable.ConfigYml.GetString("Token.JwtTokenSignKey")),
	}
}

// GenerateToken 生成token   expireAt 单位秒
func (u *userToken) GenerateToken(userid int64, username string, phone string, expireAt int64) (tokens string, err error) {

	// 根据实际业务自定义token需要包含的参数，生成token，注意：用户密码请勿包含在token
	customClaims := CustomClaims{
		UserId: userid,
		Name:   username,
		Phone:  phone,
		// 特别注意，针对前文的匿名结构体，初始化的时候必须指定键名，并且不带 jwt. 否则报错：Mixture of field: value and value initializers
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 10,       // 生效开始时间
			ExpiresAt: time.Now().Unix() + expireAt, // 失效截止时间
		},
	}
	return u.userJwt.CreateToken(customClaims)
}

// ParseToken 将 token 解析为绑定时传递的参数
func (u *userToken) ParseToken(tokenStr string) (c CustomClaims, err error) {
	if customClaims, err := u.userJwt.ParseToken(tokenStr); err == nil {
		return *customClaims, nil
	} else {
		return CustomClaims{}, errors.New("token 无效")
	}
}

// 判断token本身是否未过期
// 参数解释：
// token： 待处理的token值
// expireAtSec： 过期时间延长的秒数，主要用于用户刷新token时，判断是否在延长的时间范围内，非刷新逻辑默认为0
func (u *userToken) isNotExpired(token string, expireAtSec int64) (*CustomClaims, int) {
	if customClaims, err := u.userJwt.ParseToken(token); err == nil {

		if time.Now().Unix()-(customClaims.ExpiresAt+expireAtSec) < 0 {
			// token有效
			return customClaims, JwtTokenOK
		} else {
			// 过期的token
			return customClaims, JwtTokenExpired
		}
	} else {
		// 无效的token
		return nil, JwtTokenInvalid
	}
}

func (u *userToken) IsEffective(token string) bool {
	customClaims, code := u.isNotExpired(token, 0)

	if JwtTokenOK == code {
		fmt.Println(customClaims)
		//1.首先在redis检测是否存在某个用户对应的有效token，如果存在就直接返回，不再继续查询mysql，否则最后查询mysql逻辑，确保万无一失
		// if variable.ConfigYml.GetInt("Token.IsCacheToRedis") == 1 {
		// 	tokenRedisFact := token_cache_redis.CreateUsersTokenCacheFactory(customClaims.UserId)
		// 	if tokenRedisFact != nil {
		// 		defer tokenRedisFact.ReleaseRedisConn()
		// 		if tokenRedisFact.TokenCacheIsExists(token) {
		// 			return true
		// 		}
		// 	}
		// }
		// //2.token符合token本身的规则以后，继续在数据库校验是不是符合本系统其他设置，例如：一个用户默认只允许10个账号同时在线（10个token同时有效）
		// if model.CreateUserFactory("").OauthCheckTokenIsOk(customClaims.UserId, token) {
		// 	return true
		// }
		return true
	}
	return false
}
