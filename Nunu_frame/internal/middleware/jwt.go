package middleware

import (
	v1 "Nunu_frame/api/v1"
	"Nunu_frame/pkg/jwt"
	"Nunu_frame/pkg/log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	ContextClaimsKey = "claims"
)

func getToken(ctx *gin.Context) string {
	token := ctx.GetHeader("Authorization")
	if token != "" {
		// 兼容 Bearer xxx
		token = strings.TrimSpace(token)
		token = strings.TrimPrefix(token, "Bearer ")
		token = strings.TrimPrefix(token, "bearer ")
		return token
	}

	token, _ = ctx.Cookie("accessToken")
	if token != "" {
		return token
	}

	return ctx.Query("accessToken")
}

// 安全写日志（避免 panic）
func recoveryLoggerFunc(ctx *gin.Context, logger *log.Logger) {
	claims, ok := ctx.Get(ContextClaimsKey)
	if !ok {
		return
	}
	if userInfo, ok := claims.(*jwt.MyCustomClaims); ok {
		logger.WithValue(ctx, zap.Any("UserId", userInfo.UserId))
	}
}

func OptionalAuth(j *jwt.JWT, logger *log.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := getToken(ctx)
		if token == "" {
			ctx.Next()
			return
		}

		claims, err := j.ParseToken(token)
		if err != nil {
			ctx.Next()
			return
		}

		ctx.Set(ContextClaimsKey, claims)
		recoveryLoggerFunc(ctx, logger)
		ctx.Next()
	}
}

func LoginRequired(j *jwt.JWT, logger *log.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := getToken(ctx)
		if token == "" {
			v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, "未登录")
			ctx.Abort()
			return
		}

		claims, err := j.ParseToken(token)
		if err != nil {
			v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, "登录失效")
			ctx.Abort()
			return
		}

		ctx.Set(ContextClaimsKey, claims)
		recoveryLoggerFunc(ctx, logger)
		ctx.Next()
	}
}

func TypeRequired(j *jwt.JWT, logger *log.Logger, types ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := getToken(ctx)
		if token == "" {
			v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, "未登录")
			ctx.Abort()
			return
		}

		claims, err := j.ParseToken(token)
		if err != nil {
			v1.HandleError(ctx, http.StatusUnauthorized, v1.ErrUnauthorized, "登录失效")
			ctx.Abort()
			return
		}

		userClaims := claims

		allowed := false
		for _, userType := range types {
			if userClaims.UserType == userType {
				allowed = true
				break
			}
		}

		if !allowed {
			v1.HandleError(ctx, http.StatusForbidden, v1.ErrUnauthorized, "权限不足")
			ctx.Abort()
			return
		}

		ctx.Set(ContextClaimsKey, userClaims)
		recoveryLoggerFunc(ctx, logger)
		ctx.Next()
	}
}
