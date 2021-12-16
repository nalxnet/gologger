package zap

import (
	"bytes"
	"github.com/naucon/gologger"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"testing"
)

func TestStdLogAdapter_NewAdapter(t *testing.T) {
	t.Run("TestStdLogAdapter_NewAdapter_ShouldImplementLoggerInterface", func(t *testing.T) {
		adapter, _ := newAdapter()

		assert.Implements(t, (*logger.Logger)(nil), adapter)
	})
}

func TestStdLogAdapter_Error(t *testing.T) {
	t.Run("TestStdLogAdapter_Error_ShouldLog", func(t *testing.T) {
		expectedMsg := "error message"
		adapter, out := newAdapter()
		adapter.Error(expectedMsg)

		assert.Contains(t, out.String(), `"msg":"error message"`)
	})
}

func TestStdLogAdapter_ErrorWithFields(t *testing.T) {
	t.Run("TestStdLogAdapter_ErrorWithFields_ShouldLog", func(t *testing.T) {
		expectedMsg := "error message"
		expectedFields := map[string]interface{}{
			"key": "value",
		}
		adapter, out := newAdapter()
		adapter.ErrorWithFields(expectedFields, expectedMsg)

		assert.Contains(t, out.String(), `"msg":"error message"`)
		assert.Contains(t, out.String(), `"key":"value"`)
	})
}

func TestStdLogAdapter_Errorf(t *testing.T) {
	t.Run("TestStdLogAdapter_Errorf_ShouldLog", func(t *testing.T) {
		expectedMsg := "error message: %v"
		adapter, out := newAdapter()
		adapter.Errorf(expectedMsg, errors.New("some error"))

		assert.Contains(t, out.String(), `"msg":"error message: some error"`)
	})
}

func TestStdLogAdapter_ErrorfWithFields(t *testing.T) {
	t.Run("TestStdLogAdapter_ErrorfWithFields_ShouldLog", func(t *testing.T) {
		expectedMsg := "error message: %v"
		expectedFields := map[string]interface{}{
			"key": "value",
		}
		adapter, out := newAdapter()
		adapter.ErrorfWithFields(expectedFields, expectedMsg, errors.New("some error"))

		assert.Contains(t, out.String(), `"msg":"error message: some error"`)
		assert.Contains(t, out.String(), `"key":"value"`)
	})
}

func TestStdLogAdapter_ErrorErr(t *testing.T) {
	t.Run("TestStdLogAdapter_ErrorErr_ShouldLog", func(t *testing.T) {
		adapter, out := newAdapter()
		adapter.ErrorErr(errors.New("some error"))

		assert.Contains(t, out.String(), `"msg":"some error"`)
		assert.Contains(t, out.String(), `"error":"some error"`)
		assert.Contains(t, out.String(), `"errorVerbose":`)
	})

	t.Run("TestStdLogAdapter_ErrorErr_ShouldLogWrappedError", func(t *testing.T) {
		adapter, out := newAdapter()
		adapter.ErrorErr(errors.Wrap(errors.New("inner error"), "outer error"))

		assert.Contains(t, out.String(), `"msg":"outer error: inner error"`)
		assert.Contains(t, out.String(), `"error":"outer error: inner error"`)
		assert.Contains(t, out.String(), `"errorVerbose":`)
	})
}

func TestStdLogAdapter_Warn(t *testing.T) {
	t.Run("TestStdLogAdapter_Warn_ShouldLog", func(t *testing.T) {
		expectedMsg := "warning message"
		adapter, out := newAdapter()
		adapter.Warn(expectedMsg)

		assert.Contains(t, out.String(), `"msg":"warning message"`)
	})
}

func TestStdLogAdapter_WarnWithFields(t *testing.T) {
	t.Run("TestStdLogAdapter_WarnWithFields_ShouldLog", func(t *testing.T) {
		expectedMsg := "warning message"
		expectedFields := map[string]interface{}{
			"key": "value",
		}
		adapter, out := newAdapter()
		adapter.WarnWithFields(expectedFields, expectedMsg)

		assert.Contains(t, out.String(), `"msg":"warning message"`)
		assert.Contains(t, out.String(), `"key":"value"`)
	})
}

func TestStdLogAdapter_Warnf(t *testing.T) {
	t.Run("TestStdLogAdapter_Warnf_ShouldLog", func(t *testing.T) {
		expectedMsg := "warning message: %v"
		adapter, out := newAdapter()
		adapter.Warnf(expectedMsg, errors.New("some error"))

		assert.Contains(t, out.String(), `"msg":"warning message: some error"`)
	})
}

func TestStdLogAdapter_WarnfWithFields(t *testing.T) {
	t.Run("TestStdLogAdapter_WarnfWithFields_ShouldLog", func(t *testing.T) {
		expectedMsg := "warning message: %v"
		expectedFields := map[string]interface{}{
			"key": "value",
		}
		adapter, out := newAdapter()
		adapter.WarnfWithFields(expectedFields, expectedMsg, errors.New("some error"))

		assert.Contains(t, out.String(), `"msg":"warning message: some error"`)
		assert.Contains(t, out.String(), `"key":"value"`)
	})
}

func TestStdLogAdapter_WarnErr(t *testing.T) {
	t.Run("TestStdLogAdapter_WarnErr_ShouldLog", func(t *testing.T) {
		adapter, out := newAdapter()
		adapter.WarnErr(errors.New("some error"))

		assert.Contains(t, out.String(), `"msg":"some error"`)
		assert.Contains(t, out.String(), `"error":"some error"`)
		assert.Contains(t, out.String(), `"errorVerbose":`)
	})
}

func TestStdLogAdapter_Info(t *testing.T) {
	t.Run("TestStdLogAdapter_Info_ShouldLog", func(t *testing.T) {
		expectedMsg := "info message"
		adapter, out := newAdapter()
		adapter.Info(expectedMsg)

		assert.Contains(t, out.String(), `"msg":"info message"`)
	})
}

func TestStdLogAdapter_InfoWithFields(t *testing.T) {
	t.Run("TestStdLogAdapter_InfoWithFields_ShouldLog", func(t *testing.T) {
		expectedMsg := "info message"
		expectedFields := map[string]interface{}{
			"key": "value",
		}
		adapter, out := newAdapter()
		adapter.InfoWithFields(expectedFields, expectedMsg)

		assert.Contains(t, out.String(), `"msg":"info message"`)
		assert.Contains(t, out.String(), `"key":"value"`)
	})
}

func TestStdLogAdapter_Infof(t *testing.T) {
	t.Run("TestStdLogAdapter_Infof_ShouldLog", func(t *testing.T) {
		expectedMsg := "info message: %v"
		adapter, out := newAdapter()
		adapter.Infof(expectedMsg, errors.New("some error"))

		assert.Contains(t, out.String(), `"msg":"info message: some error"`)
	})
}

func TestStdLogAdapter_InfofWithFields(t *testing.T) {
	t.Run("TestStdLogAdapter_InfofWithFields_ShouldLog", func(t *testing.T) {
		expectedMsg := "info message: %v"
		expectedFields := map[string]interface{}{
			"key": "value",
		}
		adapter, out := newAdapter()
		adapter.InfofWithFields(expectedFields, expectedMsg, errors.New("some error"))

		assert.Contains(t, out.String(), `"msg":"info message: some error"`)
		assert.Contains(t, out.String(), `"key":"value"`)
	})
}

func TestStdLogAdapter_Debug(t *testing.T) {
	t.Run("TestStdLogAdapter_Debug_ShouldLog", func(t *testing.T) {
		expectedMsg := "debug message"
		adapter, out := newAdapter()
		adapter.Debug(expectedMsg)

		assert.Contains(t, out.String(), `"msg":"debug message"`)
	})
}

func TestStdLogAdapter_DebugWithFields(t *testing.T) {
	t.Run("TestStdLogAdapter_DebugWithFields_ShouldLog", func(t *testing.T) {
		expectedMsg := "debug message"
		expectedFields := map[string]interface{}{
			"key": "value",
		}
		adapter, out := newAdapter()
		adapter.DebugWithFields(expectedFields, expectedMsg)

		assert.Contains(t, out.String(), `"msg":"debug message"`)
		assert.Contains(t, out.String(), `"key":"value"`)
	})
}

func TestStdLogAdapter_Debugf(t *testing.T) {
	t.Run("TestStdLogAdapter_Debugf_ShouldLog", func(t *testing.T) {
		expectedMsg := "debug message: %v"
		adapter, out := newAdapter()
		adapter.Debugf(expectedMsg, errors.New("some error"))

		assert.Contains(t, out.String(), `"msg":"debug message: some error"`)
	})
}

func TestStdLogAdapter_DebugfWithFields(t *testing.T) {
	t.Run("TestStdLogAdapter_DebugfWithFields_ShouldLog", func(t *testing.T) {
		expectedMsg := "debug message: %v"
		expectedFields := map[string]interface{}{
			"key": "value",
		}
		adapter, out := newAdapter()
		adapter.DebugfWithFields(expectedFields, expectedMsg, errors.New("some error"))

		assert.Contains(t, out.String(), `"msg":"debug message: some error"`)
		assert.Contains(t, out.String(), `"key":"value"`)
	})
}

func newAdapter() (*zapAdapter, *bytes.Buffer) {
	out := new(bytes.Buffer)
	encoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		MessageKey: "msg",
	})
	lvl := zapcore.Level(zapcore.DebugLevel)
	ws := zapcore.AddSync(out)
	core := zapcore.NewCore(encoder, ws, lvl)
	zapLogger := zap.New(core, zap.WithCaller(true))
	return NewAdapter(zapLogger), out
}
