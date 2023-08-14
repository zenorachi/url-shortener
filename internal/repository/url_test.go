package repository

import (
	"context"
	"errors"
	"github.com/go-redis/redismock/v9"
	"github.com/stretchr/testify/assert"
	"github.com/zenorachi/url-shortener/model"
	"testing"
)

func TestUrls_Create(t *testing.T) {
	client, mock := redismock.NewClientMock()
	urlRepo := NewUrls(client)

	type args struct {
		url model.URL
	}
	type mockBehaviour func(args args)

	tests := []struct {
		name          string
		mockBehaviour mockBehaviour
		args          args
		wantErr       bool
	}{
		{
			name: "OK",
			args: args{
				url: model.NewUrl("google.com", "abc123"),
			},
			mockBehaviour: func(args args) {
				mock.ExpectSet(args.url.Shortened, args.url.Original, 0).SetVal("OK")
			},
		},
		{
			name: "ERROR",
			args: args{},
			mockBehaviour: func(args args) {
				mock.ExpectSet(args.url.Shortened, args.url.Original, 0).SetErr(errors.New("test error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehaviour(tt.args)
			err := urlRepo.Create(context.Background(), tt.args.url)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				err = mock.ExpectationsWereMet()
				assert.NoError(t, err)
			}
		})
	}
}

func TestUrls_GetByShorted(t *testing.T) {
	client, mock := redismock.NewClientMock()
	urlRepo := NewUrls(client)

	type args struct {
		originalUrl  string
		shortenedUrl string
	}
	type mockBehaviour func(args args)

	tests := []struct {
		name          string
		mockBehaviour mockBehaviour
		args          args
		wantErr       bool
	}{
		{
			name: "OK",
			args: args{
				originalUrl:  "google.com",
				shortenedUrl: "abc123",
			},
			mockBehaviour: func(args args) {
				mock.ExpectGet(args.shortenedUrl).SetVal(args.originalUrl)
			},
		},
		{
			name: "ERROR",
			args: args{},
			mockBehaviour: func(args args) {
				mock.ExpectGet(args.shortenedUrl).SetErr(errors.New("test error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehaviour(tt.args)
			result, err := urlRepo.GetByShorted(context.Background(), tt.args.shortenedUrl)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.args.originalUrl, result)
				err = mock.ExpectationsWereMet()
				assert.NoError(t, err)
			}
		})
	}
}
