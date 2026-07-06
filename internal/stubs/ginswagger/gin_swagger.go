package ginSwagger

import "github.com/gin-gonic/gin"

func WrapHandler(interface{}, ...interface{}) gin.HandlerFunc { return func(*gin.Context) {} }
