package main

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/RaihanSultana/go-restaurant-management/services/common/genproto/orders"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()

	conn := NewGRPCClient(":9000")
	defer conn.Close()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c := orders.NewOrderServiceClient(conn)
		ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
		defer cancel()

		_, err := c.CreateOrder(ctx, &orders.CreateOrderRequest{
			CustomerID: 24,
			ProductID:  3124,
			Quantity:   2,
		})

		if err != nil {
			log.Fatalf("client error %v", err)
		}

		o, err := c.GetOrders(ctx, &orders.GetOrderRequest{
			CustomerID: 42,
		})

		if err != nil {
			log.Fatalf("client error: %v", err)
		}

		t := template.Must(template.New("orders").Parse(ordersTemplate))
		if err := t.Execute(w, o.GetOrders()); err != nil {
			log.Fatalf("template error: %v", err)
		}
	})

	log.Println("Starting servcer on: ", s.addr)
	return http.ListenAndServe(s.addr, router)
}

var ordersTemplate = `
<!DOCTYPE html>
<html>
<title>Kitchen Order</title>
</html>
<body>
<h1>Orders List</h1>
<table border="1">
<tr>
<th>Order ID</th>
<th>Customer ID</th>
<th>ProductID ID</th>
<th>Quantity</th>
</tr>
{{range .}}
<tr>
<th>{{.OrderID}}</th>
<th>{{.CustomerID}}</th>
<th>{{.ProductID}}</th>
<th>{{.Quantity}}</th>
</tr>
{{end}}
</table>
</body>
`
