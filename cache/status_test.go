package cache

import (
	"errors"
	"testing"
)

func TestStatusCache_Parse(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name        string
		s           *StatusCache
		args        args
		wantErr     bool
		expectedErr error
	}{
		{
			name: "Success message with connections available",
			s: &StatusCache{
				ResponseChans: map[int]chan bool{
					1: make(chan bool, 1),
				},
			},
			args:    args{message: `{"op":"status","id":1,"statusCode":"SUCCESS","connectionClosed":false,"connectionsAvailable":9}`},
			wantErr: false,
		},
		{
			name: "Success message without connections available",
			s: &StatusCache{
				ResponseChans: map[int]chan bool{
					2: make(chan bool, 1),
				},
			},
			args:    args{message: `{"op":"status","id":2,"statusCode":"SUCCESS","connectionClosed":false}`},
			wantErr: false,
		},
		{
			name: "Failure message with error message",
			s: &StatusCache{
				ResponseChans: map[int]chan bool{
					3: make(chan bool, 1),
				},
			},
			args:        args{message: `{"op":"status","id":3,"statusCode":"FAILURE","connectionClosed":false,"errorMessage":"some error"}`},
			wantErr:     true,
			expectedErr: errors.New("some error"),
		},
		{
			name: "Failure message without error message",
			s: &StatusCache{
				ResponseChans: map[int]chan bool{
					4: make(chan bool, 1),
				},
			},
			args:        args{message: `{"op":"status","id":4,"statusCode":"FAILURE","connectionClosed":false}`},
			wantErr:     true,
			expectedErr: errors.New("betfair status message indicates failure"),
		},
		{
			name: "Unknown status code",
			s: &StatusCache{
				ResponseChans: map[int]chan bool{
					5: make(chan bool, 1),
				},
			},
			args:        args{message: `{"op":"status","id":5,"statusCode":"UNKNOWN","connectionClosed":false}`},
			wantErr:     true,
			expectedErr: errors.New("betfair status message indicates an unknown status code"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.s.Parse(tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("StatusCache.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && tt.expectedErr != nil && !errors.Is(err, tt.expectedErr) && err.Error() != tt.expectedErr.Error() {
				t.Errorf("StatusCache.Parse() error = %v, expectedErr %v", err, tt.expectedErr)
				return
			}
		})
	}
}
