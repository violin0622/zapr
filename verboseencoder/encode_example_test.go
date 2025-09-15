package vencoder_test

import (
	"errors"
	"os"

	"github.com/go-logr/zapr"
	ve "github.com/violin0622/zapr/verboseencoder"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func ExampleLevelEncoder() {
	encoderCfg := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		NameKey:        "logger",
		EncodeLevel:    ve.LevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
	}
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderCfg), os.Stdout, zapcore.Level(-99))
	zlog := zap.New(core)
	defer zlog.Sync()
	logr := zapr.NewLogger(zlog)

	logr.Error(errors.New(`an error`), "Bad thing happened.")
	logr.Info(`Hello world!`)
	logr.V(1).Info(`Good luck!`)
	logr.V(2).Info(`Blah blah blah.`)

	//Output:
	//{"level":"error","msg":"Bad thing happened.","error":"an error"}
	//{"level":"v(0)","msg":"Hello world!"}
	//{"level":"v(1)","msg":"Good luck!"}
	//{"level":"v(2)","msg":"Blah blah blah."}
}

func Example_customLevelEncoder() {
	lvlEnc := ve.NewLevelEncoder(
		ve.WithErrorLevel(`CustomErrorLevelString`),
		ve.WithLeftParenthesis(`{`),
		ve.WithRightParenthesis(`]`),
		ve.WithPrefix(``), // custom level string prefix like "v", empty in this case.
		ve.WithNegativeLevel())

	encoderCfg := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		NameKey:        "logger",
		EncodeLevel:    lvlEnc,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
	}
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderCfg), os.Stdout, zapcore.Level(-99))
	zlog := zap.New(core)
	defer zlog.Sync()
	logr := zapr.NewLogger(zlog)

	logr.Error(errors.New(`an error`), "Bad thing happened.")
	logr.Info(`Hello world!`)
	logr.V(1).Info(`Good luck!`)
	logr.V(2).Info(`Blah blah blah.`)

	//Output:
	//{"level":"CustomErrorLevelString","msg":"Bad thing happened.","error":"an error"}
	//{"level":"{0]","msg":"Hello world!"}
	//{"level":"{-1]","msg":"Good luck!"}
	//{"level":"{-2]","msg":"Blah blah blah."}
}
