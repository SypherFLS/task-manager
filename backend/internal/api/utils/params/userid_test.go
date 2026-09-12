package params

import (
	"net/http"
	"net/http/httptest"
    "context"
    "tm/internal/apperrors"
	"testing"
)

func TestGetUserID(t *testing.T) {
	tests := []struct {
		name    string
		data    any
		key     contextKey
        wantErr error
		want    bool
	}{
		{
			name: "usual",
			data: 10,
            key : UserIDKey,
            wantErr: nil,
            want : true,
		},
        {
			name: "wrong data type",
			data: "something",
            key : UserIDKey,
            wantErr : apperrors.WrongUserID, 
            want : false,
		},
        {
			name: "blank data",
			data: "",
            key : UserIDKey,
            wantErr : apperrors.BlankUserID,
            want : false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
            ctx := context.WithValue(
                context.Background(),
                UserIDKey,
                tt.data,
            )

            r := httptest.NewRequest(
                http.MethodPost,
                "/webhook",
                nil,
            )

            r = r.WithContext(ctx)

            _, err := GetUserID(r)

            got := err == nil 

            if got != tt.want{
                if err != tt.wantErr {
                    t.Errorf("test %v failed with wrong error %v\n", tt.name, err)
                } else {
                    t.Errorf("test %v failed with error %v\n", tt.name, err)
                }
            }

        })
    }
}
