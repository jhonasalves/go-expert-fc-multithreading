package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jhonasalves/go-expert-fc-multithreading/models"
)

const (
	viaCEPURL    = "http://viacep.com.br/ws/"
	brasilAPIURL = "https://brasilapi.com.br/api/cep/v1/"
)

// GetCep godoc
// @Summary      Get an address by CEP
// @Description  Get an address information based on CEP (postal code)
// @Tags         cep
// @Accept       json
// @Produce      json
// @Param        cep   path   string   true   "CEP (postal code)"
// @Success      200   {}   "Successful response"
// @Failure      404   {}   "CEP not found"
// @Router       /cep/{cep} [get]
func GetCep(w http.ResponseWriter, r *http.Request) {
	cep := chi.URLParam(r, "cep")

	if cep == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	c1 := make(chan *models.ViaCEP)
	c2 := make(chan *models.BrasilAPI)

	go func() {
		data, err := GetViaCEP(cep)
		if err != nil {
			panic(err)
		}

		c1 <- data
	}()

	go func() {
		data, err := GetBrasilAPI(cep)
		if err != nil {
			panic(err)
		}

		c2 <- data
	}()

	select {
	case result := <-c1:
		data, err := json.Marshal(result)
		if err != nil {
			panic(err)
		}

		fmt.Println("Received from ViaCEP")
		fmt.Printf("Data: %s\n", string(data))
		json.NewEncoder(w).Encode(result)

	case result := <-c2:
		data, err := json.Marshal(result)
		if err != nil {
			panic(err)
		}

		fmt.Println("Received from BrasilAPI")
		fmt.Printf("Data: %s\n", string(data))
		json.NewEncoder(w).Encode(result)

	case <-time.After(time.Second):
		fmt.Println("timeout")
		w.WriteHeader(http.StatusRequestTimeout)
	}
}

func GetViaCEP(cep string) (*models.ViaCEP, error) {
	url := viaCEPURL + cep + "/json/"
	req, error := http.NewRequest("GET", url, nil)
	if error != nil {
		return nil, error
	}

	resp, error := http.DefaultClient.Do(req)
	if error != nil {
		return nil, error
	}
	defer resp.Body.Close()

	body, error := io.ReadAll(resp.Body)
	if error != nil {
		return nil, error
	}

	var c models.ViaCEP
	error = json.Unmarshal(body, &c)
	if error != nil {
		return nil, error
	}

	return &c, nil
}

func GetBrasilAPI(cep string) (*models.BrasilAPI, error) {
	url := brasilAPIURL + cep
	req, error := http.NewRequest("GET", url, nil)
	if error != nil {
		return nil, error
	}

	resp, error := http.DefaultClient.Do(req)
	if error != nil {
		return nil, error
	}
	defer resp.Body.Close()

	body, error := io.ReadAll(resp.Body)
	if error != nil {
		return nil, error
	}

	var c models.BrasilAPI
	error = json.Unmarshal(body, &c)
	if error != nil {
		return nil, error
	}

	return &c, nil
}
