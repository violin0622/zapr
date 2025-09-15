package vencoder

import (
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

var defLevelEncoder = &vLevelEncoder{
	pool:    buffer.NewPool(),
	prefix:  `v`,
	left:    `(`,
	right:   `)`,
	err:     `error`,
	leveler: func(l zapcore.Level) int64 { return 0 - int64(l) },
}
var LevelEncoder = defLevelEncoder.encode

type vLevelEncoder struct {
	pool                     buffer.Pool
	prefix, left, right, err string
	leveler                  func(zapcore.Level) int64
}

func (e *vLevelEncoder) encode(l zapcore.Level, pae zapcore.PrimitiveArrayEncoder) {
	if l > 0 {
		pae.AppendString(e.err)
		return
	}

	b := e.pool.Get()
	b.Reset()
	b.AppendString(e.prefix)
	b.AppendString(e.left)
	b.AppendInt(e.leveler(l))
	b.AppendString(e.right)
	pae.AppendString(b.String())
}

type EncoderOption func(*vLevelEncoder)

func WithPrefix(p string) EncoderOption {
	return func(vle *vLevelEncoder) {
		vle.prefix = p
	}
}

func WithLeftParenthesis(l string) EncoderOption {
	return func(vle *vLevelEncoder) {
		vle.left = l
	}
}

func WithRightParenthesis(r string) EncoderOption {
	return func(vle *vLevelEncoder) {
		vle.right = r
	}
}

func WithNegativeLevel() EncoderOption {
	return func(vle *vLevelEncoder) {
		vle.leveler = func(l zapcore.Level) int64 { return int64(l) }
	}
}

func WithPositiveLevel() EncoderOption {
	return func(vle *vLevelEncoder) {
		vle.leveler = func(l zapcore.Level) int64 { return 0 - int64(l) }
	}
}

func WithErrorLevel(err string) EncoderOption {
	return func(vle *vLevelEncoder) {
		vle.err = err
	}
}

func NewLevelEncoder(opts ...EncoderOption) zapcore.LevelEncoder {
	le := &vLevelEncoder{
		pool:    buffer.NewPool(),
		prefix:  `v`,
		left:    `(`,
		right:   `)`,
		err:     `error`,
		leveler: func(l zapcore.Level) int64 { return 0 - int64(l) },
	}
	for _, o := range opts {
		o(le)
	}
	return le.encode
}
