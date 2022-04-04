package firestore

import (
	logger "github.com/sirupsen/logrus"
)

func log() *logger.Entry {
	logger.SetFormatter(&logger.TextFormatter{
		FullTimestamp: true,
	})

	return logger.WithFields(logger.Fields{
		"package": "firestore",
	})
}
