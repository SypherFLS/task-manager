package auth 

import (
	"testing"
)

func TestSolPassword(t *testing.T) {
	passwordA, passwordB := "password", "password"

	hashA, errA := HashPassword(passwordA)
	hashB, errB := HashPassword(passwordB)

	if errA != nil {
		t.Fatalf("failed hashing passwordA with error %v \n", errA)
	}

	if errB != nil {
		t.Fatalf("failed hashing passwordB with error %v \n", errB)
	}

	if hashA == hashB {
		t.Fatal("password must be different")
	}
}

func TestCheckPassword(t *testing.T) {
	tests := []struct {
		Name string
		Password string
		Want bool
	}{
		{
			Name : "usual password",
			Password : "UsualPassword",
			Want : true,
		},
		{
			Name : "empty password",
			Password : "",
			Want : true,
		},
		{
			Name : "password with special symbols",
			Password : "пароль密码🔥",
			Want : true,
		},
		{
			Name : "password with spaces",
			Password : "this password have spaces",
			Want : true,
		},
	}

	for _, tt := range tests{
		t.Run(tt.Name, func(t *testing.T){
			hash, err := HashPassword(tt.Password)

			if err != nil {
				t.Fatalf("failed hashin password %v with error %v", tt.Name, err)
			}

			if hash == "" {
				t.Fatalf("empty hash in test %v", tt.Name)
			}

			erro := CheckPassword(tt.Password, hash)

			got := erro == nil

			if got != tt.Want {
				t.Errorf("mismatching passwords in test %v", tt.Name)
			}
		})
	}
}