package _go

import "testing"

func TestLock(t *testing.T) {
	t.Run("test_run", func(t *testing.T) {
		info, err := Lock("key_1", 2000, 5000, false)
		if err != nil {
			t.Errorf("错误: %v", err)
			return
		}

		t.Logf("信息： %v", info)
	})

	t.Run("test_timeout", func(t *testing.T) {
		info, err := Lock("key_1", 2000, 1000, false)
		if err != nil {
			t.Errorf("错误: %v", err)
			return
		}

		t.Logf("信息： %v", info)
	})

	t.Run("test_err", func(t *testing.T) {
		info, err := Lock("key_1", 2000, 5000, true)
		if err != nil {
			t.Errorf("错误: %v", err)
			return
		}

		t.Logf("信息： %v", info)
	})
}
