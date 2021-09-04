package main

import (
	"errors"
	"github.com/naucon/gologger"
	zapAdapter "github.com/naucon/gologger/adapter/zap"
	"go.uber.org/zap"
)

func main() {
	zapLogger, _ := zap.NewProduction()
	defer zapLogger.Sync()

	l := zapAdapter.NewAdapter(zapLogger)
	doSomething(l)
}

func doSomething(l logger.Logger) {
	err := errors.New("some error")

	l.Error("a meaningful std message")
	l.Errorf("a meaningful std message: %v", err)
	l.ErrorErr(err)
	l.Warn("a meaningful std message")
	l.Warnf("a meaningful std message: %v", err)
	l.WarnErr(err)
	l.Info("a meaningful std message")
	l.Infof("a meaningful std message: %v", err)
	l.Debug("a meaningful std message")
	l.Debugf("a meaningful std message: %v", err)
}