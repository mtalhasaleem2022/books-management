package mocks

import (
	"context"
	"time"
)

// MockRedisClient is a mock implementation of the RedisClient interface
type MockRedisClient struct {
	GetFunc             func(ctx context.Context, key string) (string, error)
	SetFunc             func(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	DeleteFunc          func(ctx context.Context, key string) error
	DeleteByPatternFunc func(ctx context.Context, pattern string) error
	CloseFunc           func()
}

func (m *MockRedisClient) Get(ctx context.Context, key string) (string, error) {
	return m.GetFunc(ctx, key)
}

func (m *MockRedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return m.SetFunc(ctx, key, value, expiration)
}

func (m *MockRedisClient) Delete(ctx context.Context, key string) error {
	return m.DeleteFunc(ctx, key)
}

func (m *MockRedisClient) DeleteByPattern(ctx context.Context, pattern string) error {
	return m.DeleteByPatternFunc(ctx, pattern)
}

func (m *MockRedisClient) Close() {
	m.CloseFunc()
}
