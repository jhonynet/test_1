package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

//meta: map[string]interface{} {
//// withdraw-related attributes
//"withdraw.enabled": true,
//"withdraw.min_deposit_amount": 100,
//"withdraw.max_deposit_amount": 100,
//// deposit-related attributes
//"deposit.enabled": true,
//"deposit.min_deposit_amount": 100,
//// real world assets
//"stock.enabled": true,
//// display rules
//'visibility': true

type AssetRepository struct {
	Store map[int64]*Asset
}

func newAssetRepository() *AssetRepository {
	return &AssetRepository{
		Store: make(map[int64]*Asset),
	}
}

func (r *AssetRepository) List() map[int64]*Asset {
	c := r.Store
	return c
}

func (r *AssetRepository) Add(assets ...*Asset) {
	for idx := range assets {
		r.Store[assets[idx].ID] = assets[idx]
	}
}

func (r *AssetRepository) Update(a *Asset) {
	r.Store[a.ID] = a
}

func (r *AssetRepository) Delete(a *Asset) {
	delete(r.Store, a.ID)
}

type Asset struct {
	ID int64

	Name   string
	Ticker string

	Supply int
	Price  float64

	Meta map[string]interface{}
}

func main() {
	repository := newAssetRepository()

	asset1 := &Asset{
		ID:     1,
		Name:   "BTC",
		Ticker: "BTC",
		Supply: 1000,
		Price:  100000.02,
		Meta: map[string]interface{}{
			"withdraw.enabled": true,
			"deposit.enabled":  true,
		},
	}

	asset2 := &Asset{
		ID:     2,
		Name:   "DAI",
		Ticker: "DAI",
		Supply: 10,
		Price:  1000.03,
		Meta: map[string]interface{}{
			"withdraw.enabled": true,
			"deposit.enabled":  false,
		},
	}

	repository.Add(asset1, asset2)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		fmt.Fprint(w, "Hello, World!")
	})

	http.HandleFunc("/assets", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.NotFound(w, r)
			return
		}

		assets := repository.List()
		// todo: add correct headers (like content-type)
		err := json.NewEncoder(w).Encode(assets)
		if err != nil {
			_, err = fmt.Fprintln(w, `{"status": "success"}`)
			if err != nil {
				log.Fatal(err)
			}
		}
	})

	http.HandleFunc("/assets/add", func(writer http.ResponseWriter, request *http.Request) {
		// todo: add
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}

}
