package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

func main3() {
	c1 := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "test",
		Help: "test",
	})

	c1.Set(66666)
	c2 := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "test33",
		Help: "test",
	})

	c2.Set(66666)
	p:= push.
		New("http://localhost:9091", "db_backup").
		Grouping("service", "cps").
		Grouping("instance","tt").
		Collector(c2).
		Collector(c1)
	 err := p.Push();
	 if err != nil {
		fmt.Println("Could not push completion time to Pushgateway:", err)
	}
}
