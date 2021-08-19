package asset

import (
	"io"
	"io/ioutil"
)

func Asset(name string) (io.ReadCloser, error) {
	return assets.Open(name)
}

func AssetString(name string) (string, error) {
	asset, err := Asset(name)
	if err != nil {
		return "", err
	}
	defer asset.Close()

	b, err := ioutil.ReadAll(asset)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func MustAsset(name string) io.ReadCloser {
	r, err := Asset(name)
	if err != nil {
		panic(err)
	}

	return r
}

func MustAssetString(name string) string {
	data, err := AssetString(name)
	if err != nil {
		panic(err)
	}

	return data
}
