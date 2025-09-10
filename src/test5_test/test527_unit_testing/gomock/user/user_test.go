package user

// file: user_test.go

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 假设我们有个业务逻辑函数
func GetUserName(repo UserRepository, id int) (string, error) {
	u, err := repo.GetUser(id)
	if err != nil {
		return "", err
	}
	return u.Name, nil
}

func TestGetUserName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // 确保所有期望都被验证

	mockRepo := NewMockUserRepository(ctrl)

	// 定义 mock 行为
	mockRepo.EXPECT().
		GetUser(1).
		Return(&User{ID: 1, Name: "Alice"}, nil)

	// 调用业务逻辑
	name, err := GetUserName(mockRepo, 1)

	// 验证结果
	assert.NoError(t, err)
	assert.Equal(t, "Alice", name)
}

func TestGetUserName_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockUserRepository(ctrl)

	// 模拟错误返回
	mockRepo.EXPECT().
		GetUser(2).
		Return(nil, errors.New("user not found"))

	name, err := GetUserName(mockRepo, 2)

	assert.Error(t, err)
	assert.Empty(t, name)
}

// mockgen -source=user.go -destination=mock_user_test.go -package=user
// 会生成一个 mock_user_test.go 文件，里面包含 MockUserRepository。
