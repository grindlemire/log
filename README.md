# log [![](https://godoc.org/github.com/grindlemire/log?status.svg)](https://godoc.org/github.com/grindlemire/log)

a simple wrapper for [uber-go/zap](https://github.com/uber-go/zap) that simplifies configuration and handles log rotation

# Why?

Because I often need a logger that is extremely simple to hook up and just use. However I also want callers, logfile configuration, and structured logs. Zap is a good choice except I also want to have log rotation which zap doesn't support natively.

This package is intended to be a very lightweight wrapper on top of zap while greatly simplifying the configuration, providing capbaility to easily produce colored logs or json logs, and seamlessly integrate with rotating log files.

# Usage:

```Golang
func main() {
	// configuration is compatible with env loading or json if you want to store it in a file
	// I also provide log.Default for convenience

	// Simply call Init with the configuration and all the files you want to log to (if any)
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
```

# Configuration

This is the options struct you can use

```Golang
type Opts struct {
    Level              Level `json:"log_level"            env:"log_level"             default:"INFO"  description:"log level to log at (possible values are debug, info, warn, error, fatal, panic)"`
    MaxLogSize         int   `json:"log_max_size"         env:"log_max_size"          default:"10"    description:"Max size of a log in mb before rolling over"`
    MaxLogBackups      int   `json:"log_max_backups"      env:"log_max_backups"       default:"5"     description:"Max number of backups to keep"`
    CompressBackupLogs bool  `json:"log_compress_backups" env:"log_compress_backups"  default:"false" description:"Whether to compress backups or not"`
    Console            bool  `json:"log_console"          env:"log_console"           default:"true"  description:"Whether to log to the console or not (through stdout)"`
    CallerSkip         int   `json:"log_caller_skip"      env:"log_caller_skip"       default:"1"     description:"How many levels of stack to skip before logging in your application (defaults to 1 for this library)"`
}
```
