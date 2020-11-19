package main

import (
    "encoding/json"
    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "strconv"
)

type ProductoRef struct {
    Producto_id int    `json:"producto_id"`
    Producto  string `json:"producto"`
}

type ClienteRef struct {
    Cliente_id int    `json:"cliente_id"`
    Cliente  string `json:"cliente"`
}

type Factura struct {
	Id         string         `json:"_id"`
    Fecha      string         `json:"fecha"`
    Hora       string         `json:"hora"`
    Cliente    string         `json:"cliente"`
    Cliente_id string         `json:"cliente_id"`
    Total      string         `json:"total"`
    Productos  []ProductoRef  `json:"productos"`
}

var items []Factura

var jsonData string = ` [{
	"_id": 1,
	"fecha": "2020-11-06",
	"hora": "16:00:00",
	"cliente": "Daniel",
	"cliente_id": 1,
	"total": "5.000",
	"productos": [{
			"producto": "Naranjas",
			"producto_id": 2
		},
		{
			"producto": "Melones",
			"producto_id": 1
		},
		{
			"producto": "Manzanas",
			"producto_id": 3
		}
	]
},
{
	"_id": 2,
	"fecha": "2020-11-06",
	"hora": "16:30:00",
	"cliente": "Cynthia",
	"cliente_id": 2,
	"total": "5.000",
	"productos": [{
			"producto": "Naranjas",
			"producto_id": 2
		},
		{
			"producto": "Melones",
			"producto_id": 1
		},
		{
			"producto": "Manzanas",
			"producto_id": 3
		}
	]
},
{
	"_id": 3,
	"fecha": "2020-11-06",
	"hora": "16:35:00",
	"cliente": "Bruno",
	"cliente_id": 3,
	"total": "5.000",
	"productos": [{
			"producto": "Naranjas",
			"producto_id": 2
		},
		{
			"producto": "Melones",
			"producto_id": 1
		},
		{
			"producto": "Manzanas",
			"producto_id": 3
		}
	]
},
{
	"_id": 4,
	"fecha": "2020-11-06",
	"hora": "16:40:00",
	"cliente": "Mali",
	"cliente_id": 4,
	"total": "5.000",
	"productos": [{
			"producto": "Naranjas",
			"producto_id": 2
		},
		{
			"producto": "Melones",
			"producto_id": 1
		},
		{
			"producto": "Manzanas",
			"producto_id": 3
		}
	]
}
]`

func FindItem(id int) *Factura {
    for _, item := range items {
        if item.Id == id {
            return &item
        }
    }
    return nil
}

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
    id := req.QueryStringParameters["id"]
    var data []byte
    if id == "" {
        data, _ = json.Marshal(items)
    } else {
        param, err := strconv.Atoi(id)
        if err == nil {
            item := FindItem(param)
            if item != nil {
                data, _ = json.Marshal(*item)
            } else {
                data = []byte("error\n")
            }
        }
    }
    return &events.APIGatewayProxyResponse{
        StatusCode:      200,
        Headers:         map[string]string{"Content-Type": "application/json"},
        Body:            string(data),
        IsBase64Encoded: false,
    }, nil
}

func main() {
    _ = json.Unmarshal([]byte(jsonData), &items)
    lambda.Start(handler)
}