package main

import (
	"context"
	"testing"

	"github.com/jeffOSRocko/rest_grpc/rpc_defs/src/keyvalue"
	"github.com/stretchr/testify/assert"
)

func TestAddVal(t *testing.T) {
	s := Server{}
	s.Init(false)

	kv := &keyvalue.KeyValue{Key: "foo", Val: "bar"}

	resp, err := s.AddVal(context.Background(), kv)
	assert.NoError(t, err)
	assert.Equal(t, keyvalue.KeyValue{Key: "foo", Val: "bar"}, *resp)

	// Attempting to add the same key should fail
	_, err = s.AddVal(context.Background(), kv)
	assert.Error(t, err)
}

func TestModifyVal(t *testing.T) {
	s := Server{}
	s.Init(false)

	s.stored_vals["foo"] = "testbar"

	kv := &keyvalue.KeyValue{Key: "foo", Val: "newbar"}

	resp, err := s.ModifyVal(context.Background(), kv)
	assert.NoError(t, err)
	assert.Equal(t, keyvalue.KeyValue{Key: "foo", Val: "newbar"}, *resp)
}

func TestDeleteVal(t *testing.T) {
	s := Server{}
	s.Init(false)

	s.stored_vals["foo"] = "bar"

	k := &keyvalue.Key{Key: "foo"}

	resp, err := s.DeleteVal(context.Background(), k)
	assert.NoError(t, err)
	assert.Equal(t, keyvalue.KeyValue{Key: "foo", Val: "bar"}, *resp)

	assert.Empty(t, s.stored_vals)
}

func TestGetValSuccess(t *testing.T) {
	s := Server{}
	s.Init(false)

	s.stored_vals["foo"] = "testbar"

	k := &keyvalue.Key{Key: "foo"}

	resp, err := s.GetVal(context.Background(), k)
	assert.NoError(t, err)
	assert.Equal(t, "testbar", resp.Val)

	s.stored_vals["foo"] = "testbar2"

	resp, err = s.GetVal(context.Background(), k)
	assert.NoError(t, err)
	assert.Equal(t, "testbar2", resp.Val)
}

func TestGetValFailure(t *testing.T) {
	s := Server{}
	s.Init(false)

	s.stored_vals["foo"] = "testbar"

	k := &keyvalue.Key{Key: "not here"}

	_, err := s.GetVal(context.Background(), k)
	assert.NotNil(t, err)
}

func TestGetAllSuccess(t *testing.T) {
	s := Server{}
	s.Init(false)

	s.stored_vals["foo"] = "testbar"
	s.stored_vals["notfoo"] = "nottestbar"

	getAllReq := &keyvalue.GetAllRequest{}

	resp, err := s.GetAll(context.Background(), getAllReq)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(resp.Keyvals))
	assert.Contains(t, resp.Keyvals, &keyvalue.KeyValue{Key: "foo", Val: "testbar"})
	assert.Contains(t, resp.Keyvals, &keyvalue.KeyValue{Key: "notfoo", Val: "nottestbar"})
}
