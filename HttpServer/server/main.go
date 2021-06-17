package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
)

type AdjustWebhookData struct {
	ActivityKind string  `query:"activity_kind"`
	Adid         string  `query:"adid"`
	AppName      string  `query:"app_name"`
	CreatedAt    int64   `query:"created_at"`
	Event        string  `query:"event"`
	EventName    string  `query:"event_name"`
	GpsAdid      string  `query:"gps_adid"`
	Idfa         string  `query:"idfa"`
	Idfv         string  `query:"idfv"`
	RevenueUsd   float64 `query:"revenue_usd"`
	Tracker      string  `query:"tracker"`
	TrackerName  string  `query:"tracker_name"`
}

func main() {
	//server1()
	data := getAdjustWebhookData(url.Values{
		"tracker":     []string{"123"},
		"revenue_usd": []string{"1.11345"},
		"created_at":[]string{"123"},
	})
	log.Println(data)
}

func server2() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		m := map[string]string{}
		if err := c.BindQuery(&m); err != nil {
			c.String(200, err.Error())
			return
		}
		c.JSON(200, m)
	})
	if err := r.Run(); err != nil {
		panic(err)
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getAdjustWebhookData(values url.Values) *AdjustWebhookData {
	data := &AdjustWebhookData{}
	value := reflect.ValueOf(data).Elem()
	tp := value.Type()
	for i := 0; i < value.NumField(); i++ {
		fi := tp.Field(i)
		vi := value.Field(i)
		queryTag := fi.Tag.Get("query")
		if len(values[queryTag]) == 0 {
			continue
		}
		queryValue := values[queryTag][0]
		switch fi.Type.Kind() {
		case reflect.Float64:
			queryFloat, err := strconv.ParseFloat(queryValue, 64)
			if err != nil {
				panic(err)
			}
			vi.SetFloat(queryFloat)
		case reflect.Int64:
			queryInt , err := strconv.ParseInt(queryValue, 10, 64)
			if err != nil {
				panic(err)
			}
			vi.SetInt(queryInt)
		default:
			vi.SetString(values[queryTag][0])
		}
	}
	return data
}

func server1() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		getAdjustWebhookData(request.URL.Query())
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
