package entrance

import (
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"os"
	"time"
)

type httpServer struct {
	httpAddr   string
	message      message
	configData ConfigData
}

func NewHttpServer(configData ConfigData, message message) *httpServer {
	return &httpServer{
		configData: configData,
		message:      message,
		httpAddr:   configData.HttpAddr,
	}
}

func (ht *httpServer) Start() error {

	gin.SetMode(gin.ReleaseMode)
	fileWriterAccsse, err := os.OpenFile("http_accsse.log", os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer fileWriterAccsse.Close()
	gin.DefaultWriter = fileWriterAccsse

	fileWriterError , err := os.OpenFile("http_error.log", os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer fileWriterError.Close()
	gin.DefaultErrorWriter = fileWriterError

	router := gin.Default()

	router.GET("/product_detail/:siteid/userid/:uid/detailid/:detailid", func(c *gin.Context) {
		siteId := c.Param("siteid")
		detailId := c.Param("detailid")
		ht.message.AddMessage("siteid:" + siteId + "detailId:" + detailId)
		c.String(http.StatusOK, "ok")
	})

	s := &http.Server{
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	httpListener, err := net.Listen("tcp", ht.httpAddr)
	if err != nil {
		return err
	}

	return s.Serve(httpListener)
}
