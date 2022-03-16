package virtualfs

import (
	"io"
	"net/http"
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()

		router := gin.New()

		info := map[string][]byte{
			"a": []byte("a"),
		}
		fs := NewFS(info)

		router.GET("/:fileName", func(c *gin.Context) {
			fileName := c.Param("fileName")
			c.FileFromFS(fileName, http.FS(fs))
		})

		router.Run(":9383")
	}()
	wg.Done()
}

func TestNewFS(t *testing.T) {
	res, err := http.Get("http://localhost:9383/a")
	if !assert.Nil(t, err) {
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	assert.Nil(t, err)
	assert.Equal(t, "a", string(body))
}
