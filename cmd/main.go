package main

import (
	"MedbookServer/internal/adapters/app/api"
	gRPC "MedbookServer/internal/adapters/framework/drive/grpc"
	db "MedbookServer/internal/adapters/framework/driven/db/sqlc"
	patientdb "MedbookServer/internal/adapters/framework/driven/db/sqlc/patientDB_sqlc"
	"MedbookServer/internal/ports"
	"MedbookServer/utils"
	"fmt"

	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {

	var err error

	//ports
	var observationAppAdapter ports.ObservationAPIPort
	var userAppAdapter ports.UserAPIPorts
	var wardAppAdapter ports.WardAPIPort
	var bedAppAdapter ports.BedAPIPort
	var superAdminAppAdapter ports.SuperAdminAPIPorts
	var centerAdminAppAdapter ports.CenterAdminAPIPorts
	var nurseAppAdapter ports.NurseAPIPorts
	var doctorAppAdapter ports.DoctorAPIPorts
	var nusrNoteAppAdapter ports.NurseNoteAPIPort
	var prescriptionAppAdapter ports.PrescriptionAPIPort
	var drugOrderAppAdapter ports.DrugOrderAPIPort
	var drugAppAdapter ports.DrugAPIPort
	var drugChartAppAdapter ports.DrugChartAPIPort
	var drugSummeryAppAdapter ports.DrugSummeryAPIPort
	var gRPCAdapter ports.GRPCPort
	var patientAppAdapter ports.PatientAPIPort
	var admitPatientAppAdapter ports.AdmitPatientAPIPort
	var centerAppAdapter ports.CenterAPIPort

	//database connection attributes
	config, err := utils.LoadConfig("../")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	patientConfig, err := utils.LoadPatientConfig("../")
	if err != nil {
		log.Fatal("Cannot load patient config:", err)
	}

	dbConnection, err := sql.Open(config.DBDriver, config.DBSource)

	patientDBConnection, err := sql.Open(patientConfig.PatientDBDriver, patientConfig.PatientDBSource)
	if err != nil {
		log.Fatalf("failed to initiate dbase connection: %v", err)
		print("failed to initiate dbconnection")
	}

	query := db.New(dbConnection)
	patientQuery := patientdb.New(patientDBConnection)
	observationAppAdapter = api.NewObservationAPIAdapter(query)
	userAppAdapter = api.NewUserAPIAdapter(query)
	wardAppAdapter = api.NewWardAPIAdapter(query)
	bedAppAdapter = api.NewBedAPIAdapter(query)
	superAdminAppAdapter = api.NewSuperAdminAPIAdapter(query)
	centerAdminAppAdapter = api.NewCenterAdminAPIAdapter(query)
	nurseAppAdapter = api.NewNurseAPIAdapter(query)
	doctorAppAdapter = api.NewDoctorAPIAdapter(query)
	observationAppAdapter = api.NewObservationAPIAdapter(query)
	nusrNoteAppAdapter = api.NewNurseNoteAPIAdapter(query)
	prescriptionAppAdapter = api.NewPrescriptionAPIAdapter(query)
	drugOrderAppAdapter = api.NewDrugOrderAPIAdapter(query)
	drugAppAdapter = api.NewDrugAPIAdapter(query)
	drugChartAppAdapter = api.NewDrugChartAPIAdapter(query)
	drugSummeryAppAdapter = api.NewDrugSummeryAPIAdapter(query)
	patientAppAdapter = api.NewPatientAPIAdapter(patientQuery)
	admitPatientAppAdapter = api.NewAdmitPatientAPIAdapter(query)
	centerAppAdapter = api.NewCenterAPIAdapter(query)

	gRPCAdapter = gRPC.NewGRPCAdapter(observationAppAdapter, userAppAdapter, bedAppAdapter, wardAppAdapter, superAdminAppAdapter, centerAdminAppAdapter, nurseAppAdapter, doctorAppAdapter, nusrNoteAppAdapter, prescriptionAppAdapter, drugOrderAppAdapter, drugAppAdapter, drugChartAppAdapter, drugSummeryAppAdapter, patientAppAdapter, admitPatientAppAdapter, centerAppAdapter)
	fmt.Println(patientQuery)

	gRPCAdapter.Run()

}
