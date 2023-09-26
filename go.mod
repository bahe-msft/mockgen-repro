module mockgenrepro

go 1.21.0

require (
	github.com/Azure/azure-sdk-for-go/sdk/azcore v1.7.1
	github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v4 v4.3.0
	go.uber.org/mock v0.3.0
)

require (
	github.com/Azure/azure-sdk-for-go/sdk/internal v1.3.0 // indirect
	golang.org/x/net v0.14.0 // indirect
	golang.org/x/text v0.12.0 // indirect
)

replace go.uber.org/mock => github.com/bcho/mock v0.0.0-20230926075755-6ed4df535fbf
