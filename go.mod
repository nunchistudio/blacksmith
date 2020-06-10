module github.com/nunchistudio/blacksmith

go 1.14

require (
	github.com/hashicorp/go-getter v1.4.1
	github.com/mitchellh/go-homedir v1.1.0
	github.com/segmentio/ksuid v1.0.2
	github.com/sirupsen/logrus v1.6.0
)

replace cloud.google.com/go => cloud.google.com/go v0.58.0

replace cloud.google.com/go/pubsub => cloud.google.com/go/pubsub v1.3.1

replace cloud.google.com/go/storage => cloud.google.com/go/storage v1.9.0

replace golang.org/x/net => golang.org/x/net v0.0.0-20200602114024-627f9648deb9

replace golang.org/x/sys => golang.org/x/sys v0.0.0-20200323222414-85ca7c5b95cd

replace go.opencensus.io => go.opencensus.io v0.22.3

replace github.com/aws/aws-sdk-go => github.com/aws/aws-sdk-go v1.31.13
