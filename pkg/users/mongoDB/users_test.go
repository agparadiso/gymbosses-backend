package mongoDB

import "testing"

type userTestCase struct {
	name     string
	lastname string
	email    string
	out      string
}

func TestCreateUser(t *testing.T) {
	usrSrv := NewUsersSrv()
	testCases := []userTestCase{
		{
			name:     "Gabriel",
			lastname: "Paradiso",
			email:    "gaboparadiso@gmail.com",
			out:      "user Gabriel Paradiso with email gaboparadiso@gmail.com created in mongoDB",
		},
	}
	for _, tc := range testCases {
		actual := usrSrv.CreateUser(tc.name, tc.lastname, tc.email)
		if actual != tc.out {
			t.Fatalf("expected: %q, but got: %q", tc.out, actual)
		}
	}
}
