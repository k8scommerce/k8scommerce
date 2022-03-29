package gcache

import (
	"bytes"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/mailgun/groupcache/v2"
	"github.com/zeromicro/go-zero/core/discov"
)

func Serve(pool *groupcache.HTTPPool, listenOn string) http.Server {
	server := http.Server{
		Addr:    getListenOn(listenOn),
		Handler: pool,
	}

	go func() {
		fmt.Printf("Starting gCache server at %s...\n", server.Addr)
		if err := server.ListenAndServe(); err != nil {
			panic(err)
		}
	}()
	return server
}

func PeerListener(pool *groupcache.HTTPPool, selfAddr string, c discov.EtcdConf) {
	sub, err := discov.NewSubscriber(c.Hosts, c.Key)
	if err != nil {
		fmt.Println("ERROR:", err)
	}

	// the pool ip address is always the main port + 1
	update := func() {
		var poolSubs []string
		for _, listenOn := range sub.Values() {
			if listenOn != selfAddr {
				listenOn = getListenOn(listenOn)
			}
			poolSubs = append(poolSubs, listenOn)
		}

		pool.Set(poolSubs...)
		fmt.Printf("gCache pool.Set: %#v\n", poolSubs)
	}
	sub.AddListener(update)
	update()
}

// the listenOn port for gcache is always the main port plus one
// so if the service rpc is listening to port 8080 gcache would listen on 8081
func getListenOn(listenOn string) string {
	in := []byte(listenOn)
	r := regexp.MustCompile(`:([0-9]{4,5})`)
	out := r.ReplaceAllFunc(in, func(port []byte) []byte {
		parts := bytes.Split(port, []byte(":"))
		if len(parts) > 0 {
			portNum, _ := strconv.Atoi(string(parts[1]))
			return []byte(fmt.Sprintf(":%d", portNum+1))
		}
		return port
	})
	if out != nil {
		return string(out)
	}
	return listenOn
}
