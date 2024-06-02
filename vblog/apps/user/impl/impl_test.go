package impl_test

import (
	"context"
	"crypto/md5"
	"fmt"
	"testing"

	"github.com/qinchi-ops/govlog/vblog/apps/user"
	"github.com/qinchi-ops/govlog/vblog/apps/user/impl"
	"golang.org/x/crypto/bcrypt"
)

var (
	// 声明被测试的对象
	serviceImpl user.Service
	ctx         = context.Background()
)

// 招到对象
func init() {
	serviceImpl = impl.NewUserServiceImpl()
}

func TestCreateUser(t *testing.T) {
	req := user.NewCreateUserRequest()
	req.Username = "admin"
	req.Password = "123"
	// req.Role = 1
	// req.Label =
	ins, err := serviceImpl.CreateUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestQueryUser(t *testing.T) {
	req := user.NewQueryUserRequest()
	ins, err := serviceImpl.QueryUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestMd5(t *testing.T) {
	h := md5.New()
	h.Write([]byte("kasd182"))
	fmt.Printf("%x", h.Sum(nil))
}

// ab4441353cf1f3f57b91e868e8cc00b1--- PASS: TestMd5 (0.00s)
func TestPasswordHash(t *testing.T) {
	password := "secret"
	hash, _ := HashPassword(password) // ignore error for the sake of simplicity

	fmt.Println("Password:", password)
	fmt.Println("Hash:    ", hash)

	match := CheckPasswordHash(password, hash)
	fmt.Println("Match:   ", match)
}

// Hash:     $2a$12$qtw7.vjKoWsc8i8vU6ANcOkPy39YWMcU7txonZd3TyQX3iTPJy7f.
// Hash:     $2a$12$Y9GWAM9BbkDOK06Y.oqT5O7o0khvJZb8QiUoXxES2DfuN0xESkDci
// Hash:     $2a$12$1J/FDjEY3PKxD6.MmsBA4Ocp/7wVklNoTW/2MoCBPN5VqGeOFlnHe
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func TestCheckUser(t *testing.T) {
	req := user.NewCreateUserRequest()
	req.Username = "admin"
	req.Password = "123"

	u := user.NewUser(req)
	u.HashPassword()
	t.Log(u.CheckPassword("123456"))
}
