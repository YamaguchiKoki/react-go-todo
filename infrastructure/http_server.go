package infrastructure

import (
	"strconv"
	"time"

	"github.com/YamaguchiKoki/react-go-todo/adapter/logger"
	"github.com/YamaguchiKoki/react-go-todo/adapter/repository"
	"github.com/YamaguchiKoki/react-go-todo/adapter/validator"
	"github.com/YamaguchiKoki/react-go-todo/infrastructure/database"
	"github.com/YamaguchiKoki/react-go-todo/infrastructure/log"
	"github.com/YamaguchiKoki/react-go-todo/infrastructure/router"
	"github.com/YamaguchiKoki/react-go-todo/infrastructure/validation"
)

//infrastructure層の依存関係、アプリケーションの設定を抽象化した構造体
type config struct {
	appName string
	logger logger.Logger
	validator validator.Validator
	dbSQL repository.SQL
	dbNoSQL repository.NoSQL
	ctxTimeout time.Duration
	webServerPort router.Port
	webServer router.Server
}

func NewConfig() *config {
	return &config{}
}

func (c *config) ContextTimeout(t time.Duration) *config {
	c.ctxTimeout = t
	return c
}

func (c *config) Name(name string) *config {
	c.appName = name
	return c
}

func (c *config) Logger(instance int) *config {
	log, err := log.NewLoggerFactory(instance)
	if err != nil {
		log.Fatalln(err)
	}

	c.logger = log
	c.logger.Infof("Successfully configured log")

	return c
}

// func (c *config) DbSQL(instance int) *config {
// 	db, err := database.NewDatabaseSQLFactory()
// }

func (c *config) DbNoSQL(instance int) *config {
	db, err := database.NewDatabaseNoSQLFactory(instance)
	if err != nil {
		c.logger.Fatalln(err, "Could not make a connection to the database")
	}

	c.logger.Infof("Successfully connected to the NoSQL database")

	c.dbNoSQL = db
	return c
}

func (c *config) Validator(instance int) *config {
	v, err := validation.NewValidatorFactory(instance)
	if err != nil {
		c.logger.Fatalln(err)
	}

	c.logger.Infof("Successfully configured validator")

	c.validator = v
	return c
}

func (c *config) WebServer(instance int) *config {
	s, err := router.NewWebServerFactory(
		instance,
		c.logger,
		c.dbSQL,
		c.dbNoSQL,
		c.validator,
		c.webServerPort,
		c.ctxTimeout,
	)

	if err != nil {
		c.logger.Fatalln(err)
	}

	c.logger.Infof("Successfully configured router server")

	c.webServer = s
	return c
}

func (c *config) WebServerPort(port string) *config {
	p, err := strconv.ParseInt(port, 10, 64)
	if err != nil {
		c.logger.Fatalln(err)
	}

	c.webServerPort = router.Port(p)
	return c
}

func (c *config) Start() {
	c.webServer.Listen()
}