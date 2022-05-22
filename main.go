package main

import (
	"fmt"
	"houseduino-be/config"
	"houseduino-be/controllers"
	"net/http"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	loc, errT := time.LoadLocation("Europe/Rome")
	if errT != nil {
		config.PrintErrorLog("Main", errT)
	}
	time.Local = loc

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
	})
	dbsqlx := config.ConnectDBSqlx()
	hsqlx := controllers.NewBaseHandlerSqlx(dbsqlx)

	r.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// documentation for developers
	opts := middleware.SwaggerUIOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.SwaggerUI(opts, nil)
	r.Handle("/docs", sh)

	// documentation for share
	opts1 := middleware.RedocOpts{SpecURL: "/swagger.yaml", Path: "docs"}
	sh1 := middleware.Redoc(opts1, nil)
	r.Handle("/docs", sh1)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"hello\": \"world\"}"))
	})

	houseduino_be := r.PathPrefix("/").Subrouter()

	houseduino_be.HandleFunc("/altitude/insert", hsqlx.PostAltitudeSqlx).Methods("POST")
	houseduino_be.HandleFunc("/altitude/all", hsqlx.GetAltitudesSqlx).Methods("GET")
	houseduino_be.HandleFunc("/altitude/lasthour", hsqlx.GetAltitudesLastHourSqlx).Methods("GET")
	houseduino_be.HandleFunc("/altitude/showdata/{recordNumber}", hsqlx.GetShowDataSqlx).Methods("GET")

	houseduino_be.HandleFunc("/activity/insert", hsqlx.PostActivitySqlx).Methods("POST")
	houseduino_be.HandleFunc("/activity/all", hsqlx.GetActivitiesSqlx).Methods("GET")
	houseduino_be.HandleFunc("/activity/lasthour", hsqlx.GetActivitiesLastHourSqlx).Methods("GET")
	houseduino_be.HandleFunc("/activity/isactive", hsqlx.GetIsActiveSqlx).Methods("GET")

	houseduino_be.HandleFunc("/message/insert", hsqlx.PostMessageSqlx).Methods("POST")
	houseduino_be.HandleFunc("/message/lasthour", hsqlx.GetMessagesLastHourSqlx).Methods("GET")

	houseduino_be.HandleFunc("/rain/insert", hsqlx.PostRainSqlx).Methods("POST")
	houseduino_be.HandleFunc("/rain/all", hsqlx.GetRainsSqlx).Methods("GET")
	houseduino_be.HandleFunc("/rain/lasthour", hsqlx.GetRainsLastHourSqlx).Methods("GET")
	houseduino_be.HandleFunc("/rain/showdata/{recordNumber}", hsqlx.GetRainShowDataSqlx).Methods("GET")

	houseduino_be.HandleFunc("/pressure/insert", hsqlx.PostPressureSqlx).Methods("POST")
	houseduino_be.HandleFunc("/pressure/all", hsqlx.GetPressuresSqlx).Methods("GET")
	houseduino_be.HandleFunc("/pressure/lasthour", hsqlx.GetPressuresLastHourSqlx).Methods("GET")
	houseduino_be.HandleFunc("/pressure/showdata/{recordNumber}", hsqlx.GetPressureShowDataSqlx).Methods("GET")

	houseduino_be.HandleFunc("/temperature/insert", hsqlx.PostTemperatureSqlx).Methods("POST")
	houseduino_be.HandleFunc("/temperature/all", hsqlx.GetTemperaturesSqlx).Methods("GET")
	houseduino_be.HandleFunc("/temperature/lasthour", hsqlx.GetTemperaturesLastHourSqlx).Methods("GET")
	houseduino_be.HandleFunc("/temperature/showdata/{recordNumber}", hsqlx.GetTemperatureShowDataSqlx).Methods("GET")

	houseduino_be.HandleFunc("/humidity/insert", hsqlx.PostHumiditySqlx).Methods("POST")
	houseduino_be.HandleFunc("/humidity/all", hsqlx.GetHumiditiesSqlx).Methods("GET")
	houseduino_be.HandleFunc("/humidity/lasthour", hsqlx.GetHumiditiesLastHourSqlx).Methods("GET")
	houseduino_be.HandleFunc("/humidity/showdata/{recordNumber}", hsqlx.GetHumidityShowDataSqlx).Methods("GET")

	houseduino_be.HandleFunc("/weather/now", hsqlx.GetWeatherSqlx).Methods("GET")

	houseduino_be.HandleFunc("/plant/all", hsqlx.GetPlantAllSqlx).Methods("GET")
	houseduino_be.HandleFunc("/plant/status", hsqlx.GetPlantsStatusSqlx).Methods("GET")
	houseduino_be.HandleFunc("/plant/humidity/insert", hsqlx.PostPlantHumiditySqlx).Methods("POST")
	houseduino_be.HandleFunc("/plant/humidity/all/{id_plant}", hsqlx.GetPlantHumiditiesSqlx).Methods("GET")
	houseduino_be.HandleFunc("/plant/humidity/last/{id_plant}", hsqlx.GetPlantLastSqlx).Methods("GET")
	houseduino_be.HandleFunc("/plant/humidity/lasthour/{id_plant}", hsqlx.GetPlantHumiditiesLastHourSqlx).Methods("GET")
	houseduino_be.HandleFunc("/plant/humidity/showdata/{id_plant}/{recordNumber}", hsqlx.GetPlantHumidityShowDataSqlx).Methods("GET")

	now := time.Now()
	fmt.Println("The current datetime is:", now)

	hsqlx.GetTestSqlx()
	http.Handle("/", r)
	/* 	s := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", "0.0.0.0", "5559"),
		Handler: cors.Default().Handler(r),
	} */
	s := &http.Server{
		Handler:      cors.Default().Handler(r),
		Addr:         ":5557",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err := s.ListenAndServe()
	if err != nil {
		config.PrintErrorLog("Main", err)
	}

}
