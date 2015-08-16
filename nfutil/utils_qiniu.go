package nfutil

import (
	"bytes"

	"golang.org/x/net/context"
	"qiniupkg.com/api.v7/kodo"
	"qiniupkg.com/api.v7/kodocli"
)

const (
	AK         string = "oqWdc7mQYLMbciPiosBR9BU2McR-EAkgP7oWFWRn"
	SK         string = "qCP3Gue30GXYXgMdSKFMcPTiL6wraEpotIG-KFF-"
	BUCKETNAME string = "j2car"
	DOMAIN     string = "7xl4c2.com1.z0.glb.clouddn.com"
)

var (
	client *kodo.Client
	bucket kodo.Bucket
)

func getClient() {
	if client != nil {
		return
	}
	kodo.SetMac(AK, SK)
	zone := 0                    // 您空间(Bucket)所在的区域
	client = kodo.New(zone, nil) // 用默认配置创建 Client
	bucket = client.Bucket(BUCKETNAME)
	return
}

func PutLocalToCloud(b []byte, size int64, remoteKey string) (url string, err error) {
	reader := bytes.NewReader(b)
	if client == nil {
		getClient()
	}
	ctx := context.Background()
	_err := bucket.Put(ctx, nil, remoteKey, reader, size, nil)
	if err != nil {
		return "", _err
	} else {
		url = GetCloudFileUrl(remoteKey)
		return url, nil
	}
}

func PutLocalFileToCloud(localFile string, remoteKey string) (url string, err error) {
	if client == nil {
		getClient()
	}
	ctx := context.Background()
	_err := bucket.PutFile(ctx, nil, remoteKey, localFile, nil)
	if err != nil {
		return "", _err
	} else {
		url = GetCloudFileUrl(remoteKey)
		return url, nil
	}
}

func PutLocalFileToCloudWithoutKey(localFile string) (url string, err error) {
	if client == nil {
		getClient()
	}
	ctx := context.Background()
	var ret kodocli.PutRet
	_err := bucket.PutFileWithoutKey(ctx, &ret, localFile, nil)
	if _err != nil {
		return "", _err
	} else {
		url = GetCloudFileUrl(ret.Key)
		return url, nil
	}
}

func GetCloudFileUrl(remoteKey string) (url string) {
	baseUrl := kodo.MakeBaseUrl(DOMAIN, remoteKey) // 得到下载 url
	return baseUrl
}
