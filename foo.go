package main

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v4"
)

type Interface interface {
	BeginAbortLatestOperation(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string, options *armcontainerservice.AgentPoolsClientBeginAbortLatestOperationOptions) (*runtime.Poller[armcontainerservice.AgentPoolsClientAbortLatestOperationResponse], error)
}

func main() {
	fmt.Println("hello, world!")
}
