package main

import (
	"bytes"
	"context"
	"grpc-auth-app/auth-server/api"
	"grpc-auth-app/auth-server/internal/handler"
	"grpc-auth-app/auth-server/pkg/config"
	"grpc-auth-app/auth-server/pkg/di"
	"io"
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
)

func TestSignupUserGrpc(t *testing.T) {
	config.InitConfig()
	svc := di.BuildContainer()

	go handler.StartGrpcServer(svc)
	port := config.Config.GRPCPort
	// lis,err:= ne
	cc, err := grpc.NewClient(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("error while connceting to the client", err)
	}
	assert.NoError(t, err)
	client := api.NewUserServiceClient(cc)
	payload := `{
		"id": "66ade7f0463df84987f80c51",
		"index": 0,
		"guid": "19db8fff-a996-44fe-95d0-dbaadfdbcd53",
		"isActive": false,
		"balance": "$2,314.75",
		"picture": "http://placehold.it/32x32",
		"age": 32,
		"eyeColor": "brown",
		"name": "Chrystal Freeman",
		"gender": "female",
		"company": "MELBACOR",
		"email": "chrystalfreeman@melbacor.com",
		"phone": "+1 (889) 429-3149",
		"address": "705 Gerritsen Avenue, Ribera, Maine, 5918",
		"about": "Veniam irure et excepteur cupidatat laborum ullamco laborum labore esse eu anim dolor magna. Commodo consectetur incididunt elit laborum consectetur aliqua anim. Excepteur laboris tempor quis sit ullamco qui ex labore. Elit magna irure excepteur excepteur velit sint reprehenderit ut.\r\n",
		"registered": "2020-06-08T07:06:33 -06:-30",
		"latitude": 66.641033,
		"longitude": -27.260329,
		"tags": [
		  "occaecat",
		  "ad",
		  "ut",
		  "nisi",
		  "occaecat",
		  "sunt",
		  "commodo"
		],
		"friends": [
		  {
			"id": 0,
			"name": "Concepcion Bauer"
		  },
		  {
			"id": 1,
			"name": "Leona Juarez"
		  },
		  {
			"id": 2,
			"name": "Hayes Golden"
		  }
		],
		"greeting": "Hello, Chrystal Freeman! You have 6 unread messages.",
		"favoriteFruit": "strawberry"
	  }`
	// bytes, err := json.Marshal(payload)
	// if err != nil {
	// 	t.Log("error occured while marshaling payload")
	// }
	assert.NoError(t, err)
	user := new(api.User)
	err = protojson.Unmarshal([]byte(payload), user)
	if err != nil {
		t.Log("error occured while unmarshaling payload", err)
		// return
	}
	// t.Log("user proto", user.Email)
	assert.NoError(t, err)
	res, err := client.CreateUser(context.Background(), user)
	if err != nil {
		t.Log("error occured while caling method")
	}
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Token)
	t.Log("test success", res.Token.Token)
}

func TestSignupServiceHttp(t *testing.T) {
	config.InitConfig()
	svc := di.BuildContainer()

	go handler.StartHttpServer(svc)
	payload := `{
		"id": "66ade7f0463df84987f80c51",
		"index": 0,
		"guid": "19db8fff-a996-44fe-95d0-dbaadfdbcd53",
		"isActive": false,
		"balance": "$2,314.75",
		"picture": "http://placehold.it/32x32",
		"age": 32,
		"eyeColor": "brown",
		"name": "Chrystal Freeman",
		"gender": "female",
		"company": "MELBACOR",
		"email": "chrystalfreeman@melbacor.com",
		"phone": "+1 (889) 429-3149",
		"address": "705 Gerritsen Avenue, Ribera, Maine, 5918",
		"about": "Veniam irure et excepteur cupidatat laborum ullamco laborum labore esse eu anim dolor magna. Commodo consectetur incididunt elit laborum consectetur aliqua anim. Excepteur laboris tempor quis sit ullamco qui ex labore. Elit magna irure excepteur excepteur velit sint reprehenderit ut.\r\n",
		"registered": "2020-06-08T07:06:33 -06:-30",
		"latitude": 66.641033,
		"longitude": -27.260329,
		"tags": [
		  "occaecat",
		  "ad",
		  "ut",
		  "nisi",
		  "occaecat",
		  "sunt",
		  "commodo"
		],
		"friends": [
		  {
			"id": 0,
			"name": "Concepcion Bauer"
		  },
		  {
			"id": 1,
			"name": "Leona Juarez"
		  },
		  {
			"id": 2,
			"name": "Hayes Golden"
		  }
		],
		"greeting": "Hello, Chrystal Freeman! You have 6 unread messages.",
		"favoriteFruit": "strawberry"
	  }`
	req, _ := http.NewRequest("POST", "http://localhost:35580/signup", bytes.NewReader([]byte(payload)))

	client := &http.Client{}
	// http.DefaultServeMux.ServeHTTP(w, req)
	res, err := client.Do(req)
	if err != nil {
		t.Logf("unexpected error %v", res.Status)
	}
	assert.NoError(t, err)

	// result := res.
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status 200 OK, got %v", res.Status)
	}
	body, _ := io.ReadAll(res.Body)
	// body := res.Body.String()
	// if body != "Hello, world!\n" {
	// 	t.Errorf("expected body 'Hello, world!', got '%s'", body)
	// }
	t.Log("token", string(body))
}
