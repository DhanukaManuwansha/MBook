package rpc

import (
	"MedbookServer/internal/adapters/framework/drive/grpc/pb"
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (grpca GRPCAdapter) GetPatients(ctx context.Context, req *pb.GetPatientsParameters) (*pb.PatientListResult, error) {
	var err error
	patientList := &pb.PatientListResult{}

	//Get patients from API layer					-Dhanuka Manuwansha						-27/05/2022
	patients, err := grpca.patientsapi.GetPatientsApi()

	if err != nil {
		fmt.Println("GetLatestePrescriptions rpc error : ", err)
		return patientList, status.Error(codes.Internal, "unexpected Error")
	}

	var p []*pb.PatientResult

	for i, v := range patients {
		var temPatient = &pb.PatientResult{
			PatientId:      v.PatientID,
			Nic:            v.Nic,
			FirstName:      v.FirstName,
			MiddleName:     v.MiddleName,
			LastName:       v.LastName,
			Dob:            v.Dob.String(),
			AddressLineOne: v.AddressLineOne.String,
			AddressLineTwo: v.AddressLineTwo.String,
			City:           v.City.String,
			State:          v.State.String,
			Country:        v.Country,
			TellNo:         v.TellNo,
			Email:          v.Email,
			Password:       v.Password,
			Sex:            v.Sex.String,
			Height:         float32(v.Height.Float64),
			Weight:         float32(v.Weight.Float64),
			BloodGroup:     v.BloodGroup.String,
			MaritalState:   v.MaritalState.String,
			Occupation:     v.Occupation.String,
			Nationality:    v.Nationality.String,
			FamilyId:       v.FamilyID.Int64,
		}

		fmt.Println(i)

		p = append(p, temPatient)
	}

	patientList = &pb.PatientListResult{
		Patients: p,
	}

	return patientList, nil

}

func (grpca GRPCAdapter) GetPatientsByNIC(ctx context.Context, req *pb.GetPatientsByNICParameters) (*pb.PatientListResult, error) {
	var err error
	patientList := &pb.PatientListResult{}

	//Get patients from API layer					-Dhanuka Manuwansha						-27/05/2022
	patients, err := grpca.patientsapi.GetPatientBYNICApi(req.Nic)

	if err != nil {
		fmt.Println("GetLatestePrescriptions rpc error : ", err)
		return patientList, status.Error(codes.Internal, "unexpected Error")
	}

	var p []*pb.PatientResult

	for i, v := range patients {
		var temPatient = &pb.PatientResult{
			PatientId:      v.PatientID,
			Nic:            v.Nic,
			FirstName:      v.FirstName,
			MiddleName:     v.MiddleName,
			LastName:       v.LastName,
			Dob:            v.Dob.String(),
			AddressLineOne: v.AddressLineOne.String,
			AddressLineTwo: v.AddressLineTwo.String,
			City:           v.City.String,
			State:          v.State.String,
			Country:        v.Country,
			TellNo:         v.TellNo,
			Email:          v.Email,
			Password:       v.Password,
			Sex:            v.Sex.String,
			Height:         float32(v.Height.Float64),
			Weight:         float32(v.Weight.Float64),
			BloodGroup:     v.BloodGroup.String,
			MaritalState:   v.MaritalState.String,
			Occupation:     v.Occupation.String,
			Nationality:    v.Nationality.String,
			FamilyId:       v.FamilyID.Int64,
		}

		fmt.Println(i)

		p = append(p, temPatient)
	}

	patientList = &pb.PatientListResult{
		Patients: p,
	}

	return patientList, nil

}

func (grpca GRPCAdapter) GetPatientsByTellNo(ctx context.Context, req *pb.GetPatientsByTellNoParameters) (*pb.PatientListResult, error) {
	var err error
	patientList := &pb.PatientListResult{}

	//Get patients from API layer					-Dhanuka Manuwansha						-27/05/2022
	patients, err := grpca.patientsapi.GetPatientBYTellNoApi(req.TellNo)

	if err != nil {
		fmt.Println("GetLatestePrescriptions rpc error : ", err)
		return patientList, status.Error(codes.Internal, "unexpected Error")
	}

	var p []*pb.PatientResult

	for i, v := range patients {
		var temPatient = &pb.PatientResult{
			PatientId:      v.PatientID,
			Nic:            v.Nic,
			FirstName:      v.FirstName,
			MiddleName:     v.MiddleName,
			LastName:       v.LastName,
			Dob:            v.Dob.String(),
			AddressLineOne: v.AddressLineOne.String,
			AddressLineTwo: v.AddressLineTwo.String,
			City:           v.City.String,
			State:          v.State.String,
			Country:        v.Country,
			TellNo:         v.TellNo,
			Email:          v.Email,
			Password:       v.Password,
			Sex:            v.Sex.String,
			Height:         float32(v.Height.Float64),
			Weight:         float32(v.Weight.Float64),
			BloodGroup:     v.BloodGroup.String,
			MaritalState:   v.MaritalState.String,
			Occupation:     v.Occupation.String,
			Nationality:    v.Nationality.String,
			FamilyId:       v.FamilyID.Int64,
		}

		fmt.Println(i)

		p = append(p, temPatient)
	}

	patientList = &pb.PatientListResult{
		Patients: p,
	}

	return patientList, nil

}

func (grpca GRPCAdapter) GetPatientsByName(ctx context.Context, req *pb.GetPatientsByNameParameters) (*pb.PatientListResult, error) {
	var err error
	patientList := &pb.PatientListResult{}

	//Get patients from API layer					-Dhanuka Manuwansha						-27/05/2022
	patients, err := grpca.patientsapi.GetPatientBYNameApi(req.Name)

	if err != nil {
		fmt.Println("GetLatestePrescriptions rpc error : ", err)
		return patientList, status.Error(codes.Internal, "unexpected Error")
	}

	var p []*pb.PatientResult

	for i, v := range patients {
		var temPatient = &pb.PatientResult{
			PatientId:      v.PatientID,
			Nic:            v.Nic,
			FirstName:      v.FirstName,
			MiddleName:     v.MiddleName,
			LastName:       v.LastName,
			Dob:            v.Dob.String(),
			AddressLineOne: v.AddressLineOne.String,
			AddressLineTwo: v.AddressLineTwo.String,
			City:           v.City.String,
			State:          v.State.String,
			Country:        v.Country,
			TellNo:         v.TellNo,
			Email:          v.Email,
			Password:       v.Password,
			Sex:            v.Sex.String,
			Height:         float32(v.Height.Float64),
			Weight:         float32(v.Weight.Float64),
			BloodGroup:     v.BloodGroup.String,
			MaritalState:   v.MaritalState.String,
			Occupation:     v.Occupation.String,
			Nationality:    v.Nationality.String,
			FamilyId:       v.FamilyID.Int64,
		}

		fmt.Println(i)

		p = append(p, temPatient)
	}

	patientList = &pb.PatientListResult{
		Patients: p,
	}

	return patientList, nil

}

func (grpca GRPCAdapter) GetPatientById(ctx context.Context, req *pb.GetPatientsByIDParameters) (*pb.PatientResult, error) {
	var err error
	patientList := &pb.PatientResult{}

	//Get patients from API layer					-Dhanuka Manuwansha						-27/05/2022
	patients, err := grpca.patientsapi.GetPatientByIdApi(req.PatientId)

	if err != nil {
		fmt.Println("GetLatestePrescriptions rpc error : ", err)
		return patientList, status.Error(codes.Internal, "unexpected Error")
	}

	var patient = &pb.PatientResult{
		PatientId:      patients.PatientID,
		Nic:            patients.Nic,
		FirstName:      patients.FirstName,
		MiddleName:     patients.MiddleName,
		LastName:       patients.LastName,
		Dob:            patients.Dob.String(),
		AddressLineOne: patients.AddressLineOne.String,
		AddressLineTwo: patients.AddressLineTwo.String,
		City:           patients.City.String,
		State:          patients.State.String,
		Country:        patients.Country,
		TellNo:         patients.TellNo,
		Email:          patients.Email,
		Password:       patients.Password,
		Sex:            patients.Sex.String,
		Height:         float32(patients.Height.Float64),
		Weight:         float32(patients.Weight.Float64),
		BloodGroup:     patients.BloodGroup.String,
		MaritalState:   patients.MaritalState.String,
		Occupation:     patients.Occupation.String,
		Nationality:    patients.Nationality.String,
		FamilyId:       patients.FamilyID.Int64,
	}

	return patient, nil

}
