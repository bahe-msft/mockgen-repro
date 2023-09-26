package mockgenrepro

import (
	"mockgenrepro/mock_main"
	"testing"

	"go.uber.org/mock/gomock"
)

func Test_Foo(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	mock := mock_main.NewMockInterface(mockCtrl)
	mock.EXPECT().BeginAbortLatestOperation(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, nil).AnyTimes()
}
