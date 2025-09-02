package zapr

import (
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

var lvlEnc = vLevelEncoder{pool: buffer.NewPool()}

type vLevelEncoder struct {
	pool buffer.Pool
}

func (e *vLevelEncoder) Encode(l zapcore.Level, pae zapcore.PrimitiveArrayEncoder) {
	if l > 0 {
		pae.AppendString(l.CapitalString()[:4])
		return
	}

	b := e.pool.Get()
	b.Reset()
	b.AppendString("V(")
	b.AppendInt(0 - int64(l))
	b.AppendString(`)`)
	pae.AppendString(b.String())
}

type EncoderOption func(*vLevelEncoder)

func NewLevelEncoder(opts ...EncoderOption) vLevelEncoder {
	return vLevelEncoder{pool: buffer.NewPool()}
}
