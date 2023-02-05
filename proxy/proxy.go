package proxy

// proxy handler get video baru tiap 5 detik

import (
	"fmt"
	"time"
)

type ThirdPartyYoutubeLib interface {
	ListVideo() []string
}

type ThirdPartyYoutubeImpl struct {
}

func (t ThirdPartyYoutubeImpl) ListVideo() []string {
	<-time.After(200 * time.Millisecond)
	return []string{
		"youtube.com/watch/youtube1",
		"youtube.com/watch/youtube2",
	}
}

type CacheYoutube struct {
	Service ThirdPartyYoutubeLib
	cache   []string
}

func (c *CacheYoutube) run() {
	for range time.Tick(5 * time.Second) {
		c.cache = []string{}
	}
}

func (c *CacheYoutube) ListVideo() []string {
	if len(c.cache) == 0 {
		c.cache = c.Service.ListVideo()
		fmt.Println("call real api get video")
	} else {
		fmt.Println("return list video video")
	}

	return c.cache
}

func init() {
	impl := ThirdPartyYoutubeImpl{}
	cache := CacheYoutube{
		Service: &impl,
	}
	go cache.run()
	fmt.Println(cache.ListVideo())
	fmt.Println(cache.ListVideo())

	<-time.After(7 * time.Second)

	fmt.Println(cache.ListVideo())

}
