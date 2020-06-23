// Code generated by sdkgen. DO NOT EDIT.

//nolint
package clickhouse

import (
	"context"

	"google.golang.org/grpc"

	clickhouse "github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/clickhouse/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// MlModelServiceClient is a clickhouse.MlModelServiceClient with
// lazy GRPC connection initialization.
type MlModelServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Create implements clickhouse.MlModelServiceClient
func (c *MlModelServiceClient) Create(ctx context.Context, in *clickhouse.CreateMlModelRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return clickhouse.NewMlModelServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements clickhouse.MlModelServiceClient
func (c *MlModelServiceClient) Delete(ctx context.Context, in *clickhouse.DeleteMlModelRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return clickhouse.NewMlModelServiceClient(conn).Delete(ctx, in, opts...)
}

// Get implements clickhouse.MlModelServiceClient
func (c *MlModelServiceClient) Get(ctx context.Context, in *clickhouse.GetMlModelRequest, opts ...grpc.CallOption) (*clickhouse.MlModel, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return clickhouse.NewMlModelServiceClient(conn).Get(ctx, in, opts...)
}

// List implements clickhouse.MlModelServiceClient
func (c *MlModelServiceClient) List(ctx context.Context, in *clickhouse.ListMlModelsRequest, opts ...grpc.CallOption) (*clickhouse.ListMlModelsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return clickhouse.NewMlModelServiceClient(conn).List(ctx, in, opts...)
}

type MlModelIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err     error
	started bool

	client  *MlModelServiceClient
	request *clickhouse.ListMlModelsRequest

	items []*clickhouse.MlModel
}

func (c *MlModelServiceClient) MlModelIterator(ctx context.Context, clusterId string, opts ...grpc.CallOption) *MlModelIterator {
	return &MlModelIterator{
		ctx:    ctx,
		opts:   opts,
		client: c,
		request: &clickhouse.ListMlModelsRequest{
			ClusterId: clusterId,
			PageSize:  1000,
		},
	}
}

func (it *MlModelIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	response, err := it.client.List(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.MlModels
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *MlModelIterator) Value() *clickhouse.MlModel {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *MlModelIterator) Error() error {
	return it.err
}

// Update implements clickhouse.MlModelServiceClient
func (c *MlModelServiceClient) Update(ctx context.Context, in *clickhouse.UpdateMlModelRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return clickhouse.NewMlModelServiceClient(conn).Update(ctx, in, opts...)
}