package rpc

import (
	"MedbookServer/internal/adapters/framework/drive/grpc/pb"
	ports "MedbookServer/internal/ports"
	"log"
	"net"

	"google.golang.org/grpc"
)

type GRPCAdapter struct {
	api             ports.ObservationAPIPort
	userapi         ports.UserAPIPorts
	bedapi          ports.BedAPIPort
	wardapi         ports.WardAPIPort
	superadminapi   ports.SuperAdminAPIPorts
	centeradminapi  ports.CenterAdminAPIPorts
	nurseapi        ports.NurseAPIPorts
	doctorapi       ports.DoctorAPIPorts
	nursenoteapi    ports.NurseNoteAPIPort
	prescriptionapi ports.PrescriptionAPIPort
	drugorderapi    ports.DrugOrderAPIPort
	drugapi         ports.DrugAPIPort
	drugchartapi    ports.DrugChartAPIPort
	drugsummeryapi  ports.DrugSummeryAPIPort
	patientsapi     ports.PatientAPIPort
	admitpatientapi ports.AdmitPatientAPIPort
	centerapi       ports.CenterAPIPort
}

func NewGRPCAdapter(api ports.ObservationAPIPort, userapi ports.UserAPIPorts, bedapi ports.BedAPIPort, wardapi ports.WardAPIPort, superadminapi ports.SuperAdminAPIPorts, centeradminapi ports.CenterAdminAPIPorts, nurseapi ports.NurseAPIPorts, doctorapi ports.DoctorAPIPorts, nursenoteapi ports.NurseNoteAPIPort, prescriptionapi ports.PrescriptionAPIPort, drugorderapi ports.DrugOrderAPIPort, drugapi ports.DrugAPIPort, drugchartapi ports.DrugChartAPIPort, drugsummeryapi ports.DrugSummeryAPIPort, patientsapi ports.PatientAPIPort, admitpatientapi ports.AdmitPatientAPIPort, centerapi ports.CenterAPIPort) *GRPCAdapter {
	return &GRPCAdapter{
		api:             api,
		userapi:         userapi,
		bedapi:          bedapi,
		wardapi:         wardapi,
		superadminapi:   superadminapi,
		centeradminapi:  centeradminapi,
		nurseapi:        nurseapi,
		doctorapi:       doctorapi,
		nursenoteapi:    nursenoteapi,
		prescriptionapi: prescriptionapi,
		drugorderapi:    drugorderapi,
		drugapi:         drugapi,
		drugchartapi:    drugchartapi,
		drugsummeryapi:  drugsummeryapi,
		patientsapi:     patientsapi,
		admitpatientapi: admitpatientapi,
		centerapi:       centerapi,
	}
}

func (grpca GRPCAdapter) Run() {

	var err error

	listen, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatalf("failed to listen on port 9000: %v", err)
	}

	grpcServer := grpc.NewServer()

	//observation server services
	observationServiceServer := grpca
	nursenoteServiceServer := grpca
	drugchartServiceServer := grpca
	drugSummeryServiceServer := grpca
	prescriptionServiceServer := grpca
	patientServiceServer := grpca
	admitPatientServiceServer := grpca

	// user server services 	- Thaveesha - 30/05/2022
	userServiceServer := grpca

	// bed server services 		- Thaveesha - 30/05/2022
	bedServiceServer := grpca

	// ward server services  	- Thaveesha - 30/05/2022
	wardServiceServer := grpca

	// super admin service server 	- Thaveesha - 7/06/2022
	superAdminServiceServer := grpca

	// center admin service server
	centerAdminServiceServer := grpca

	nurseServicesServer := grpca

	doctorServiceServer := grpca
	centerServiceServer := grpca

	pb.RegisterObservationServiceServer(grpcServer, observationServiceServer)
	pb.RegisterNurseNoteServiceServer(grpcServer, nursenoteServiceServer)
	pb.RegisterDrugChartServiceServer(grpcServer, drugchartServiceServer)
	pb.RegisterDrugSummeryServiceServer(grpcServer, drugSummeryServiceServer)
	pb.RegisterPrescriptionServiceServer(grpcServer, prescriptionServiceServer)
	pb.RegisterPatientServiceServer(grpcServer, patientServiceServer)
	pb.RegisterAdmitPatientServiceServer(grpcServer, admitPatientServiceServer)

	pb.RegisterUserServiceServer(grpcServer, userServiceServer)

	pb.RegisterBedServiceServer(grpcServer, bedServiceServer)

	pb.RegisterWardServiceServer(grpcServer, wardServiceServer)

	pb.RegisterSuperAdminServicesServer(grpcServer, superAdminServiceServer)

	pb.RegisterCenterAdminServicesServer(grpcServer, centerAdminServiceServer)

	pb.RegisterNurseServicesServer(grpcServer, nurseServicesServer)

	pb.RegisterDoctorServicesServer(grpcServer, doctorServiceServer)

	pb.RegisterCenterServicesServer(grpcServer, centerServiceServer)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}

}
