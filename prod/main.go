package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	rsv "gitswarm.f5net.com/salton/reservations/pkg/reservation"

	"github.com/rs/cors"
	kafka "github.com/segmentio/kafka-go"
)

var (
	index = 1
)

func producerHandler(kafkaWriter *kafka.Writer) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(wrt http.ResponseWriter, req *http.Request) {
		log.Printf("Request accepted : %v %v", req.Method, req.URL)

		var r rsv.Reservation
		body, err := ioutil.ReadAll(req.Body)

		if err := json.Unmarshal(body, &r); err != nil {
			log.Printf("Error: request body incorrect.\nreason: %v", err)
			wrt.WriteHeader(http.StatusBadRequest)
			return
		}

		r.ID = index
		index++
		log.Printf("Reservation Details : %v", r)

		b, err := r.Bytes()
		if err != nil {
			fmt.Printf("Error: unexpected encoding issue.\nreason: %v", err)
			wrt.WriteHeader(http.StatusBadRequest)
			return
		}

		msg := kafka.Message{
			Key:   []byte(rsv.OPNEW),
			Value: b,
		}
		log.Printf("Submitting request to Kafka. Operation: %s,", string(rsv.OPNEW))
		err = kafkaWriter.WriteMessages(req.Context(), msg)

		if err != nil {
			log.Printf("Error: Kafka.WriteMessage unexpected error.\nreason :%v", err)
			wrt.WriteHeader(http.StatusInternalServerError)
			wrt.Write([]byte(err.Error()))
			log.Fatalln(err)
		}
	})
}

func updateHandler(kafkaWriter *kafka.Writer) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(wrt http.ResponseWriter, req *http.Request) {
		log.Printf("Request accepted : %v %v", req.Method, req.URL)

		var r rsv.Reservation
		body, err := ioutil.ReadAll(req.Body)

		if err := json.Unmarshal(body, &r); err != nil {
			log.Printf("Error: request body incorrect.\nreason: %v", err)
			wrt.WriteHeader(http.StatusBadRequest)
			return
		}

		params := mux.Vars(req)
		id, _ := strconv.Atoi(params["id"])
		r.ID = id

		log.Printf("Update Reservation Details : %v, %v", id, r)

		b, err := r.Bytes()
		if err != nil {
			fmt.Printf("Error: unexpected encoding issue.\nreason: %v", err)
			wrt.WriteHeader(http.StatusBadRequest)
			return
		}

		msg := kafka.Message{
			Key:   []byte(rsv.OPCHG),
			Value: b,
		}
		log.Printf("Submitting request to Kafka. Operation: %s,", string(rsv.OPCHG))
		err = kafkaWriter.WriteMessages(req.Context(), msg)

		if err != nil {
			log.Printf("Error: Kafka.WriteMessage unexpected error.\nreason :%v", err)
			wrt.WriteHeader(http.StatusInternalServerError)
			wrt.Write([]byte(err.Error()))
			log.Fatalln(err)
		}
	})
}

func deleteHandler(kafkaWriter *kafka.Writer) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(wrt http.ResponseWriter, req *http.Request) {

		log.Printf("DELETE Request accepted : %v %v", req.Method, req.URL)

		params := mux.Vars(req)
		id := params["id"]

		log.Printf("Delete Reservation Details : %v", id)

		//  b, err := r.Bytes()
		// if err != nil {
		// 	fmt.Printf("Error: unexpected encoding issue.\nreason: %v", err)
		// 	wrt.WriteHeader(http.StatusBadRequest)
		// 	return
		// }

		msg := kafka.Message{
			Key:   []byte(rsv.OPREM),
			Value: []byte(id),
		}
		log.Printf("Submitting request to Kafka. Operation: %s,", string(rsv.OPREM))
		err := kafkaWriter.WriteMessages(req.Context(), msg)

		if err != nil {
			log.Printf("Error: Kafka.WriteMessage unexpected error.\nreason :%v", err)
			wrt.WriteHeader(http.StatusInternalServerError)
			wrt.Write([]byte(err.Error()))
			log.Fatalln(err)
		}
	})
}

func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func main() {
	// get kafka writer using environment variables.
	server := os.Getenv("kafka")
	kafkaURL := server + ":9092"
	topic := "topic1"
	kafkaWriter := getKafkaWriter(kafkaURL, topic)
	defer kafkaWriter.Close()

	// Add handle func for producer.
	router := mux.NewRouter()

	router.HandleFunc("/reservations", producerHandler(kafkaWriter)).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/reservations/{id:[0-9]+}", deleteHandler(kafkaWriter)).Methods(http.MethodDelete, http.MethodOptions)
	router.HandleFunc("/reservations/{id:[0-9]+}", updateHandler(kafkaWriter)).Methods(http.MethodPut, http.MethodOptions)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://10.240.106.202:8084", "http://10.240.106.202:8080", "http://10.240.106.202:8081"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})

	// Insert the middleware
	handler := c.Handler(router)

	fmt.Println("starting producer-api...")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
