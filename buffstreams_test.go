package buffstreams

import "testing"

func TestNewUsesDefaultMessageSize(t *testing.T) {
	cfg := BuffManagerConfig{}
	buffM := New(cfg)
	if buffM.maxMessageSize != DefaultMaxMessageSize {
		t.Errorf("Expected Max Message Size to be %d, actually got %d", DefaultMaxMessageSize, buffM.maxMessageSize)
	}
}

func TestNewUsesSpecifiedMaxMessageSize(t *testing.T) {
	cfg := BuffManagerConfig{
		MaxMessageSize: 8196,
	}
	buffM := New(cfg)
	if buffM.maxMessageSize != cfg.MaxMessageSize {
		t.Errorf("Expected Max Message Size to be %d, actually got %d", cfg.MaxMessageSize, buffM.maxMessageSize)
	}
}
