package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pooya-dehghan/repository/mysql"
	"github.com/pooya-dehghan/service/userservice"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/users/register", userRegisterHandler)
	http.ListenAndServe(":8787", mux)
}

func userRegisterHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		fmt.Print(writer, "invalid method")
	}

	data, err := io.ReadAll(req.Body)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`error in read body %s`, err.Error())))
		return
	}
	var reqData userservice.RegisterRequest
	err = json.Unmarshal(data, &reqData)
	fmt.Println(reqData)

	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`error in read body %s`, err.Error())))
		return
	}
	mysqlRep := mysql.New()

	userSvc := userservice.New(mysqlRep)

	userCreated, err := userSvc.Register(reqData)

	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`error in user creation %s`, err.Error())))
		return
	}

	fmt.Println(userCreated)
}
