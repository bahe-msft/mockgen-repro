package mockgenrepro

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v4"
)

//go:generate mockgen -destination mock_main/foo.go -package mock_main . Interface

type Interface interface {
	BeginAbortLatestOperation(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string, options *armcontainerservice.AgentPoolsClientBeginAbortLatestOperationOptions) (*runtime.Poller[armcontainerservice.AgentPoolsClientAbortLatestOperationResponse], error)
}

func main() {
	fmt.Println("hello, world!")
}
