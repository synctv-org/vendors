package bootstrap

import (
	"context"
	"io"
	"log"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/synctv-org/vendors/cmd/flags"
	"github.com/synctv-org/vendors/utils"
)

func setLog(l *logrus.Logger) {
	if flags.Dev {
		l.SetLevel(logrus.DebugLevel)
		l.SetReportCaller(true)
	} else {
		l.SetLevel(logrus.InfoLevel)
		l.SetReportCaller(false)
	}
}

func InitLog() (err error) {
	setLog(logrus.StandardLogger())
	forceColor := utils.ForceColor()
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:      forceColor,
		DisableColors:    !forceColor,
		ForceQuote:       flags.Dev,
		DisableQuote:     !flags.Dev,
		DisableSorting:   true,
		FullTimestamp:    true,
		TimestampFormat:  time.DateTime,
		QuoteEmptyFields: true,
	})
	log.SetOutput(logrus.StandardLogger().Writer())
	return nil
}

func InitDiscardLog(ctx context.Context) error {
	setLog(logrus.StandardLogger())
	logrus.StandardLogger().SetOutput(io.Discard)
	log.SetOutput(logrus.StandardLogger().Writer())
	return nil
}
