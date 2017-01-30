package goetcd

import (
"log"
"time"

"golang.org/x/net/context"
"github.com/coreos/etcd/client"
)

var kapi client.KeysAPI

func SetKey(route string, value string) (error) {
	resp, err := kapi.Set(context.Background(), route, value, nil)
	if err != nil {
		log.Fatal(err)
		return err
	} else {
		log.Printf("Set is done. Metadata is %q\n", resp)
		return nil
	}
}

func GetKey(route string) (string, error) {
	resp, err := kapi.Get(context.Background(), route, nil)
	if err != nil {
		log.Fatal(err)
		return "", err
	} else {
		log.Printf("Get is done. Metadata is %q\n", resp)
		log.Printf("%q key has %q value\n", resp.Node.Key, resp.Node.Value)
		return resp.Node.Value, nil
	}
}

func SetConf(url string) (client.KeysAPI, error) {
	log.Printf("etcd connecting.. %q", url)
	cfg := client.Config{
		Endpoints:               []string{url},
		Transport:               client.DefaultTransport,
		// set timeout per request to fail fast when the target endpoint is unavailable
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := client.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	kapi = client.NewKeysAPI(c)
	return kapi, err
}
