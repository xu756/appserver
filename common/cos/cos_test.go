package cos

import (
	"context"
	"server/common/config"
	"testing"
)

func TestCosClient_UploadFile(t *testing.T) {
	config.Init("../../configs/dev.yaml")
	client, err := NewClient()
	if err != nil {
		t.Error(err)
	}
	err = client.UploadFile(context.Background())
	if err != nil {
		t.Error(err)
		return
	}

}
