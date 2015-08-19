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
	DOMAIN     string = "img.papakaka.com"
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

func CopyCloudFile(srcKey string, destKey string) (url string, err error) {
	if client == nil {
		getClient()
	}
	ctx := context.Background()
	err = bucket.Copy(ctx, srcKey, destKey)
	url = GetCloudFileUrl(destKey)
	return url, err
}

func ListCloudFile(prefix string) (ret []kodo.ListItem) {
	if client == nil {
		getClient()
	}
	ctx := context.Background()
	var allentries []kodo.ListItem = make([]kodo.ListItem, 0, 10240)
	var marker string = ""
	for {
		entries, markerOut, err := bucket.List(ctx, prefix, marker, 2048)
		if len(entries) > 0 {
			allentries = append(allentries, entries...)
		}
		if err != nil {
			break
		}
		if markerOut != "" {
			marker = markerOut
		} else {
			break
		}
	}
	return allentries
}

func DellCloudFileByPathPrefix(prefix string) error {
	allentries := ListCloudFile(prefix)
	keys := make([]string, 0, len(allentries))
	for _, v := range allentries {
		keys = append(keys, v.Key)
	}
	if client == nil {
		getClient()
	}
	ctx := context.Background()
	_, err := bucket.BatchDelete(ctx, keys...)
	return err
}
