package main

import ( 
	"context"
//	"github.com/gin-gonic/gin"
//	"database/sql" 
	"encoding/json"
	"fmt"
//	"net/http"
	"time"
	"os"
	"io/ioutil"
	"github.com/jackc/pgx/v5"
	//"github.com/nats-io/nats.go"
	//stan "github.com/nats-io/stan.go" 
	//"log"
)
type JsStruct struct {
	OrderUid    string `json:"order_uid"`
	TrackNumber string `json:"track_number"`
	Entry       string `json:"entry"`
	Delivery    struct {
		Name    string `json:"name"`
		Phone   string `json:"phone"`
		Zip     string `json:"zip"`
		City    string `json:"city"`
		Address string `json:"address"`
		Region  string `json:"region"`
		Email   string `json:"email"`
	} `json:"delivery"`
	Payment struct {
		Transaction  string `json:"transaction"`
		RequestId    string `json:"request_id"`
		Currency     string `json:"currency"`
		Provider     string `json:"provider"`
		Amount       int    `json:"amount"`
		PaymentDt    int    `json:"payment_dt"`
		Bank         string `json:"bank"`
		DeliveryCost int    `json:"delivery_cost"`
		GoodsTotal   int    `json:"goods_total"`
		CustomFee    int    `json:"custom_fee"`
	} `json:"payment"`
	Items []struct {
		ChrtId      int    `json:"chrt_id"`
		TrackNumber string `json:"track_number"`
		Price       int    `json:"price"`
		Rid         string `json:"rid"`
		Name        string `json:"name"`
		Sale        int    `json:"sale"`
		Size        string `json:"size"`
		TotalPrice  int    `json:"total_price"`
		NmId        int    `json:"nm_id"`
		Brand       string `json:"brand"`
		Status      int    `json:"status"`
	} `json:"items"`
	Locale            string    `json:"locale"`
	InternalSignature string    `json:"internal_signature"`
	CustomerId        string    `json:"customer_id"`
	DeliveryService   string    `json:"delivery_service"`
	Shardkey          string    `json:"shardkey"`
	SmId              int       `json:"sm_id"`
	DateCreated       time.Time `json:"date_created"`
	OofShard          string    `json:"oof_shard"`
}

var dataCache = make(map[string]JsStruct, 10) 

func main() { 
	// читаем .json и парсим его
	jsonFile, err := os.Open("model.json")
	if err != nil {
  		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	defer jsonFile.Close()
	var JsElement JsStruct
	json.Unmarshal(byteValue, &JsElement)
	fmt.Print("%v\n"  , JsElement)

	fmt.Print("---------------------------------------------------\n")
	fmt.Print(string(byteValue))
	
	// открываем базу 
	DB_url := "postgres://adam:@localhost:5432/adam?sslmode=disable"
        db_, err := pgx.Connect(context.Background(), os.Getenv(DB_url))
        defer db_.Close(context.Background())
        if err != nil {
                fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
                os.Exit(1)
        }
        var dur_req JsStruct
        if err != nil {
                fmt.Print(err)
        }
        a := db_.QueryRow(context.Background(), "SELECT data from orders;").Scan(&dur_req)
        fmt.Print(a)

	// записываем данные в базу сразу
	
	fmt.Print("--------------------------------------\n") 
	//fmt.Print(jsonFileToInsert)
	ret := db_.QueryRow(context.Background(), "insert into orders (id, data) values ($1, $2);", JsElement.OrderUid, string(byteValue))
	fmt.Print(ret) 
	
	fmt.Print("---------------------final----------------------------\n")
	a = db_.QueryRow(context.Background(), "SELECT data from orders;").Scan(&dur_req)
        fmt.Print(dur_req)
	


}	
