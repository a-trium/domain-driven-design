package repository_test

import (
	"fmt"
	"github.com/a-trium/domain-driven-design/implementation-duk/service-gateway/internal/domain/user"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User.Repository", func() {

	When("When find a user", func() { // 조건
		It("should return dummy user", func() {
			// given
			mockUser := createMockUser()

			mockCtrl := gomock.NewController(T)
			repository := user.NewMockRepository(mockCtrl)
			defer mockCtrl.Finish()

			// when
			repository.EXPECT().FindOne(gomock.Any()).Return(mockUser)

			// then
			result := repository.FindOne(1)
			fmt.Println(result)
			Expect(result.Name).To(Equal(mockUser.Name))
		})
	})
})

func createMockUser() *user.User {
	return &user.User{Name:"hello"}
}
