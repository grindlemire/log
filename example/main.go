package main

import (
	"github.com/grindlemire/log"
)

func main() {
	err := log.Init(log.Default, "./test.log")
	if err != nil {
		log.Fatalf("unable to initialize logger: %v", err)
	}

	log.Debug("no formatting")
	log.Debugf("formatting [%s]", "here")
	log.Debugw("extra fields", log.Fields{"like": "this", "or_numbers": 1})

	log.Info("no formatting")
	log.Infof("formatting [%s]", "here")
	log.Infow("extra fields", log.Fields{"like": "this", "or_numbers": 1})

	log.Warn("no formatting")
	log.Warnf("formatting [%s]", "here")
	log.Warnw("extra fields", log.Fields{"like": "this", "or_numbers": 1})

	log.Error("no formatting")
	log.Errorf("formatting [%s]", "here")
	log.Errorw("extra fields", log.Fields{"like": "this", "or_numbers": 1})

	log.Fatal("All done")
}
