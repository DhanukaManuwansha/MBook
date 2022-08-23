package rpc

import (
	"MedbookServer/internal/adapters/framework/drive/grpc/pb"
	db "MedbookServer/internal/adapters/framework/driven/db/sqlc"
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (grpca GRPCAdapter) GetDrugChart(ctx context.Context, req *pb.GetDrugChartParameters) (*pb.DrugChartResult, error) {
	var err error
	drugchart := &pb.DrugChartResult{}
	var drugorders []db.DrugOrder
	var drugs []db.Drug

	//Check is patient Id null or not					-Dhanuka Manuwansha						-27/05/2022
	if req.PatientId == "" {
		fmt.Println("Get Drug Chart rpc error : patient ID cannot be null")
		return drugchart, status.Error(codes.InvalidArgument, "missing required params")
	}

	//Get all DrugChart from API layer					-Dhanuka Manuwansha						-27/05/2022
	prescriptions, err := grpca.prescriptionapi.GetActivePrescriptionsApi(req.PatientId)
	fmt.Println("llllllllllllllllllll")
	fmt.Println(prescriptions)
	fmt.Println(prescriptions[0].PrescriptionID)
	fmt.Println(prescriptions[0].Sketch)
	fmt.Println(len(prescriptions))
	if err != nil {
		fmt.Println("GetDrugChart rpc error : ", err)
		return drugchart, status.Error(codes.Internal, "unexpected Error")
	}

	for i := 0; i < len(prescriptions); i++ {
		drugorder, err := grpca.drugorderapi.GetDrugOrdersByPrescriptionIDApi(1)
		drugorders = append(drugorders, drugorder...)
		fmt.Println("aaaaaaaaa")
		fmt.Println(drugorder)
		if err != nil {
			fmt.Println("GetDrugChart rpc error : ", err)
			return drugchart, status.Error(codes.Internal, "unexpected Error")
		}
	}

	for i := 0; i < len(drugorders); i++ {
		drug, err := grpca.drugapi.GetDrugByIdApi(int64(drugorders[i].DrugID.Int64))
		fmt.Println("wwwwwwwwwwwwwwwwwwww")
		fmt.Println(drugorders[i].DrugID)
		fmt.Println(drug.DrugName)
		fmt.Println(drug.DrugID)
		drugs = append(drugs, drug)

		if err != nil {
			fmt.Println("GetDrugChart rpc error : ", err)
			return drugchart, status.Error(codes.Internal, "unexpected Error")
		}
	}

	// var tempDrugs []db.Drug
	// tempDrugs = append(tempDrugs, drugs[0])
	// for i := 0; i < len(drugs); i++ {
	// 	for j := 0; j < len(tempDrugs); j++ {
	// 		if tempDrugs[j].DrugID != drugs[i].DrugID {
	// 			tempDrugs = append(tempDrugs, drugs[i])
	// 		}
	// 	}
	// }

	keys := make(map[db.Drug]bool)
	list := []db.Drug{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range drugs {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	drugs = list
	var drugList []*pb.DrugResult
	for i, v := range drugs {
		var tempDrug = &pb.DrugResult{
			DrugId:             v.DrugID,
			DrugName:           v.DrugName,
			ScientficName:      v.ScientificName.String,
			DrugCategory:       v.DrugCategory,
			StorageTemperature: v.StorageTemperature.String,
			DangerousLevel:     v.DangerousLevel,
			Manufacture:        v.Manufacture,
			NoOfUnits:          int64(v.NoOfUnits),
			Note:               v.Notes.String,
		}

		fmt.Println(i)

		drugList = append(drugList, tempDrug)
	}

	drugchart = &pb.DrugChartResult{
		Drugs: drugList,
	}

	return drugchart, nil

}

func (grpca GRPCAdapter) GetDrugChartList(ctx context.Context, req *pb.GetDrugChartParameters) (*pb.DrugChartResultForDashBoard, error) {
	var err error
	drugchart := &pb.DrugChartResultForDashBoard{}

	//Check is patient Id null or not					-Dhanuka Manuwansha						-27/05/2022
	if req.PatientId == "" {
		fmt.Println("Get Drug Chart rpc error : patient ID cannot be null")
		return drugchart, status.Error(codes.InvalidArgument, "missing required params")
	}

	//Get all nursenotes from API layer					-Dhanuka Manuwansha						-27/05/2022
	drugs, err := grpca.drugchartapi.GetDrugChartApi(req.PatientId)

	if err != nil {
		fmt.Println("GetDrugChart rpc error : ", err)
		return drugchart, status.Error(codes.Internal, "unexpected Error")
	}

	var drugList []*pb.DrugResultForDashBoard
	for i, v := range drugs {
		var tempDrug = &pb.DrugResultForDashBoard{
			DrugName:   v.DrugName,
			Dose:       v.Dose,
			Dosage:     v.Dosage,
			Frequency:  v.Frequency,
			GivenDate:  v.Givendate.String(),
			GiveUntill: v.Giveuntil.String(),
		}

		fmt.Println(i)

		drugList = append(drugList, tempDrug)
	}

	drugchart = &pb.DrugChartResultForDashBoard{
		Drugs: drugList,
	}

	return drugchart, nil

}

func (grpca GRPCAdapter) GetDrugChartListForNurseDesk(ctx context.Context, req *pb.DrugChartForNurseDesktParam) (*pb.DrugChartForNurseDeskResultList, error) {
	var err error
	drugchart := &pb.DrugChartForNurseDeskResultList{}

	if req.PatientId == "" {
		fmt.Println("Get Drug Chart rpc error : patient ID cannot be null")
		return drugchart, status.Error(codes.InvalidArgument, "missing required params")
	}

	if req.Date == "" {
		fmt.Println("Get Drug Chart rpc error : Date cannot be null")
		return drugchart, status.Error(codes.InvalidArgument, "missing required params")
	}

	var param db.GetDrugchartForNurseDeskParams
	param.PatientID = req.PatientId
	date, err := time.Parse("2006-01-02", req.Date)
	param.SummeryDate = date
	//Get all nursenotes from API layer					-Dhanuka Manuwansha						-27/05/2022
	drugs, err := grpca.drugchartapi.GetDrugchartForNurseDeskApi(param)

	if err != nil {
		fmt.Println("GetDrugChart rpc error : ", err)
		return drugchart, status.Error(codes.Internal, "unexpected Error")
	}

	var morningDrugList []*pb.DrugChartForNurseDeskResult
	var noonDrugList []*pb.DrugChartForNurseDeskResult
	var eveDrugList []*pb.DrugChartForNurseDeskResult
	var nightDrugList []*pb.DrugChartForNurseDeskResult
	for i, v := range drugs {

		var tempDrug = &pb.DrugChartForNurseDeskResult{
			DrugName:         v.DrugName,
			Dose:             v.Dose,
			Dosage:           v.Dosage,
			Frequency:        v.Frequency,
			GivenDate:        v.Givendate.String(),
			GiveUntill:       v.Giveuntil.String(),
			DrugsummeryId:    v.DrugsummeryID,
			SummeryDate:      v.SummeryDate.String(),
			SummeryStatus:    v.SummeryStatus,
			FirstDoseStatus:  v.FirstDoseIsGiven,
			SecondDoseStatus: v.SecondDoseIsGiven,
			ThirdDoseStatus:  v.ThirdDoseIsGiven,
			FourthDoseStatus: v.FourthDoseIsGiven,
			DrugorderId:      int64(v.DrugorderID.Int64),
			PatientId:        v.PatientID,
		}

		if v.Frequency == 4 {
			var eveDrug = &pb.DrugChartForNurseDeskResult{
				DrugName:         v.DrugName,
				Dose:             v.Dose,
				Dosage:           v.Dosage,
				Frequency:        v.Frequency,
				GivenDate:        v.Givendate.String(),
				GiveUntill:       v.Giveuntil.String(),
				DrugsummeryId:    v.DrugsummeryID,
				SummeryDate:      v.SummeryDate.String(),
				SummeryStatus:    v.SummeryStatus,
				FirstDoseStatus:  v.FirstDoseIsGiven,
				SecondDoseStatus: v.SecondDoseIsGiven,
				ThirdDoseStatus:  v.ThirdDoseIsGiven,
				FourthDoseStatus: v.FourthDoseIsGiven,
				DrugorderId:      int64(v.DrugorderID.Int64),
				PatientId:        v.PatientID,
			}
			eveDrugList = append(eveDrugList, eveDrug)
		}

		fmt.Println(i)

		morningDrugList = append(morningDrugList, tempDrug)
		noonDrugList = append(noonDrugList, tempDrug)
		nightDrugList = append(nightDrugList, tempDrug)
	}

	drugchart = &pb.DrugChartForNurseDeskResultList{
		MorningList: morningDrugList,
		NoonList:    noonDrugList,
		EveList:     eveDrugList,
		NightList:   nightDrugList,
	}

	return drugchart, nil

}
