package mysql

import (
	"log"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/OscarTango1995/AuthServiceMicroservice/db"
	"github.com/OscarTango1995/AuthServiceMicroservice/models"
)

func Test_client_CreateUser(t *testing.T) {
	envVariables()

	type args struct {
		user *models.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "success - add user in db",
			args:    args{user: &models.User{Name: "Ovais", Email: "a@a.com", Password: "password", CreatedAt: time.Now().UTC()}},
			wantErr: false,
		},
		{
			name:    "fail - add invalid user in db",
			args:    args{user: &models.User{ID: "1", Name: "Junaid", Email: "b@a.com", Password: "password", CreatedAt: time.Now().UTC()}},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := NewClient(db.Option{})
			if err != nil {
				log.Print(err, c)
				return
			}
			_, err2 := c.CreateUser(tt.args.user)
			if (err2 != nil) != tt.wantErr {
				t.Errorf("CreateUser error = %v, wantErr %v", err2, tt.wantErr)
				return
			}
		})
	}
}
func Test_client_DeleteUser(t *testing.T) {
	envVariables()
	c, _ := NewClient(db.Option{})
	user := &models.User{Name: "Shahzad", Email: "c@a.com", Password: "password123", CreatedAt: time.Now().UTC()}
	_, _ = c.CreateUser(user)
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "success - delete user from db",
			args:    args{id: user.ID},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := c.DeleteUser(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func Test_client_ListUser(t *testing.T) {
	envVariables()

	c, err := NewClient(db.Option{})
	if err != nil {
		t.Errorf("List() error = %v, wantErr %v", err, c)
		return
	}
	user := &models.User{Name: "Junaid", Email: "b@a.com", Password: "password", CreatedAt: time.Now().UTC().Truncate(time.Minute)}
	_, _ = c.CreateUser(user)

	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.User
		wantErr bool
	}{
		{
			name:    "success - get user from db",
			args:    args{id: user.ID},
			want:    user,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.ListUser(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_UpdateUser(t *testing.T) {
	envVariables()
	c, err := NewClient(db.Option{})
	if err != nil {
		log.Println(err, c)
	}
	type args struct {
		user *models.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "success - update user in db",
			args:    args{user: &models.User{ID: "78bc55bb-1c18-41a4-af0e-56f6acaff137", Name: "XYZ Updated", Password: "password", Email: "xyz@xyz.com", CreatedAt: time.Now().UTC().Truncate(time.Minute)}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := c.UpdateUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_CreateSession(t *testing.T) {
	envVariables()

	type args struct {
		session *models.Session
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "success - add user in db",
			args:    args{session: &models.Session{UserID: "1", Token: "abcdefghi", CreatedAt: time.Now().UTC(), ExpiresAt: time.Now().UTC()}},
			wantErr: false,
		},
		{
			name:    "fail - add invalid user in db",
			args:    args{session: &models.Session{ID: "2", UserID: "2", Token: "qwertyuiop", CreatedAt: time.Now().UTC(), ExpiresAt: time.Now().UTC()}},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := NewClient(db.Option{})
			if err != nil {
				log.Print(err, c)
				return
			}
			_, err2 := c.CreateSession(tt.args.session)
			if (err2 != nil) != tt.wantErr {
				t.Errorf("CreateSession() error = %v, wantErr %v", err2, tt.wantErr)
				return
			}
		})
	}
}

func Test_client_DeleteSession(t *testing.T) {
	envVariables()

	c, _ := NewClient(db.Option{})
	session := &models.Session{UserID: "4", Token: "asdfewqty", CreatedAt: time.Now().UTC()}
	_, _ = c.CreateSession(session)
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "success - delete user from db",
			args:    args{id: session.ID},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := c.DeleteSession(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteSession() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_ListSession(t *testing.T) {
	envVariables()

	c, _ := NewClient(db.Option{})
	session := &models.Session{UserID: "5", Token: "abcdefdfghi", CreatedAt: time.Now().UTC().Truncate(time.Minute), ExpiresAt: time.Now().UTC().Truncate(time.Minute)}
	_, _ = c.CreateSession(session)
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.Session
		wantErr bool
	}{
		{
			name:    "success - get user from db",
			args:    args{id: session.ID},
			want:    session,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.ListSession(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListSession() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListSession() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_UpdateSession(t *testing.T) {
	envVariables()

	c, err := NewClient(db.Option{})
	if err != nil {
		log.Println(err, c)
	}
	type args struct {
		session *models.Session
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "success - update user in db",
			args:    args{session: &models.Session{ID: "656950f6-bee1-42a6-8304-ae1aba1a1692", UserID: "2", Token: "qwertyuiop", CreatedAt: time.Now().UTC().Truncate(time.Minute), ExpiresAt: time.Now().UTC().Truncate(time.Minute)}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := c.UpdateSession(tt.args.session); (err != nil) != tt.wantErr {
				t.Errorf("UpdateSession() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func envVariables() {
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_HOST", "auth-microservice-mysql-db")
	os.Setenv("DB_USER", "root")
}
