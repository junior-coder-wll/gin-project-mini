//@Author: wulinlin
//@Description:
//@File:  main
//@Version: 1.0.0
//@Date: 2023/03/10 03:31

package main

import (
	"gin-project-mini/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "第一个gin框架")
	})
	logger := middlewares.NewLogger("runtime/logs", 100, 1)
	r.Use(middlewares.LoggerMiddleware(logger))
	r.Run()
}

/*
- cmd/
  - main.go
- configs/
  - config.yaml
- controllers/
  - auth_controller.go
  - user_controller.go
- middlewares/
  - auth_middleware.go
  - error_middleware.go
- models/
  - user_model.go
- repositories/
  - user_repository.go
- routes/
  - router.go
- utils/
  - utils.go
*/
