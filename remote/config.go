package remote

import (
	"fmt"
	"google.golang.org/grpc"
)

// RemotingOption configures how the remote infrastructure is started
type RemotingOption func(*remoteConfig)

func defaultRemoteConfig() *remoteConfig {
	return &remoteConfig{
		advertisedAddress:        "",
		dialOptions:              []grpc.DialOption{grpc.WithInsecure()},
		endpointWriterBatchSize:  1000,
		endpointManagerBatchSize: 1000,
		endpointWriterQueueSize:  1000000,
		endpointManagerQueueSize: 1000000,
	}
}

func WithEndpointWriterBatchSize(batchSize int) RemotingOption {
	return func(config *remoteConfig) {
		config.endpointWriterBatchSize = batchSize
	}
}

func WithEndpointWriterQueueSize(queueSize int) RemotingOption {
	return func(config *remoteConfig) {
		config.endpointWriterQueueSize = queueSize
	}
}

func WithEndpointManagerBatchSize(batchSize int) RemotingOption {
	return func(config *remoteConfig) {
		config.endpointManagerBatchSize = batchSize
	}
}

func WithEndpointManagerQueueSize(queueSize int) RemotingOption {
	return func(config *remoteConfig) {
		config.endpointManagerQueueSize = queueSize
	}
}

func WithDialOptions(options ...grpc.DialOption) RemotingOption {
	return func(config *remoteConfig) {
		config.dialOptions = options
	}
}

func WithServerOptions(options ...grpc.ServerOption) RemotingOption {
	return func(config *remoteConfig) {
		config.serverOptions = options
	}
}

func WithCallOptions(options ...grpc.CallOption) RemotingOption {
	return func(config *remoteConfig) {
		config.callOptions = options
	}
}

func WithAdvertisedAddress(address string) RemotingOption {
	return func(config *remoteConfig) {
		config.advertisedAddress = address
	}
}

func BindTo(host string, port int) *remoteConfig {
	c := defaultRemoteConfig()
	c.host = host
	c.port = port
	return c
}

func BindToLocalhost(port int) *remoteConfig {
	return BindTo("127.0.0.1", port)
}

func (rc *remoteConfig) Address() string {
	return fmt.Sprintf("%v:%v", rc.host, rc.port)
}

type remoteConfig struct {
	host                     string
	port                     int
	advertisedAddress        string
	serverOptions            []grpc.ServerOption
	callOptions              []grpc.CallOption
	dialOptions              []grpc.DialOption
	endpointWriterBatchSize  int
	endpointWriterQueueSize  int
	endpointManagerBatchSize int
	endpointManagerQueueSize int
}
