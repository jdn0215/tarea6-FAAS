package main

import (
    "encoding/json"
    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "strconv"
)


type Cliente struct {
    Id       string    `json:"_id"`
    Nombre   string    `json:"nombre"`
    Apellido string    `json:"apellido"`
    Correo   string    `json:"correo"`
    Telefono string    `json:"telefono"`
}

var items []Cliente

var jsonData string = `[{
	"_id": 1,
	"nombre": "Daniel",
	"apellido": "Salas",
	"correo": "js@mail.com",
	"telefono": "888888888"
},
{
	"_id": 2,
	"nombre": "Cynthia",
	"apellido": "Madrigal",
	"correo": "cm@mail.com",
	"telefono": "888888888"
},
{
	"_id": 3,
	"nombre": "Bruno",
	"apellido": "Salas",
	"correo": "bs@mail.com",
	"telefono": "888888888"
},
{
	"_id": 4,
	"nombre": "Mali",
	"apellido": "Madrigal",
	"correo": "mm@mail.com",
	"telefono": "888888888"
}
]`

func FindItem(id int) *Cliente {
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