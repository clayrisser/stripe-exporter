package main

import (
	"encoding/json"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gopkg.in/gin-gonic/gin.v1"
	"io/ioutil"
)

var customersCount = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "stripe_customers_count",
	Help: "Number of stripe customers registered on stripe",
})

func webhook(c *gin.Context) {
	data := getBody(c)["data"].(map[string]interface{})["object"].(map[string]interface{})
	if data["customer"] != nil {
		customersCount.Inc()
	}
	dataRaw, _ := json.Marshal(data)
	fmt.Println("*********************")
	fmt.Println(string(dataRaw))
	fmt.Println("---------------------")
	c.JSON(200, data)
}

func init() {
	prometheus.MustRegister(customersCount)
}

func main() {
	r := gin.Default()
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	r.POST("/webhook", webhook)
	r.Run()
}

func getBody(c *gin.Context) map[string]interface{} {
	var bodyMap map[string]interface{}
	bodyRaw, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(bodyRaw, &bodyMap)
	return bodyMap
}
