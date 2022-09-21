package main

import (
	"context"
	"testing"

	"test/mock"

	"github.com/golang/mock/gomock"
)

func TestDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockdb := mock.NewMockDB(ctrl)
	// ctx, _ := context.WithCancel(context.Background())
	// mockdb.EXPECT().Ping(ctx).Return(nil)
	mockdb.EXPECT().Ping(context.Background()).Return(nil)
	usecase := NewUsecase()
	// DBに繋いでいない！！！
	if err := usecase(mockdb, context.Background()); err != nil {
		t.Errorf("unexpected error err = %s", err.Error())
		return
	}
}
