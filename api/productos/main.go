package main

import (
    "encoding/json"
    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "strconv"
)

type Producto struct {
    Id           string `json:"_id"`
    Nombre       string `json:"nombre"`
    Valor        string `json:"valor"`
    Distribuidor string `json:"distribuidor"`
    Disponible   string `json:"disponible"`
}

var productos []Producto

var jsonData string = `[{
	"_id": 1,
	"nombre": "Melones",
	"valor": "2500",
	"distribuidor": "Melones S.A",
	"disponible": "si"
},
{
	"_id": 2,
	"nombre": "Naranjas",
	"valor": "1500",
	"distribuidor": "Naranjas S.A",
	"disponible": "si"
},
{
	"_id": 3,
	"nombre": "Manzanas",
	"valor": "1000",
	"distribuidor": "Manzanas S.A",
	"disponible": "si"
}
]`

func FindProducto(id int) *Producto {
    for _, producto := range productos {
        if producto.Id == id {
            return &producto
        }
    }
    return nil
}

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
    id := req.QueryStringParameters["id"]
    var data []byte
    if id == "" {
        data, _ = json.Marshal(productos)
    } else {
        param, err := strconv.Atoi(id)
        if err == nil {
            producto := FindProducto(param)
            if producto != nil {
                data, _ = json.Marshal(*producto)
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
    _ = json.Unmarshal([]byte(jsonData), &productos)
    lambda.Start(handler)
}