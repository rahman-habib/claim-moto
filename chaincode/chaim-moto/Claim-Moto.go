/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing a workflow
type SmartContract struct {
	contractapi.Contract
}

// QueryResult structure used for handling result of query
type QueryResult struct {
	Key    string `json:"Key"`
	Record *ClaimBlock
}

// QueryHistory structure used for handling result of query
type QueryHistory struct {
	Record string
}

type DeliveryReadyBlock struct {
	Url string `json:"Url"`
}
type DocumentBlock struct {
	Url string `json:"Url"`
}
type PartBlock struct {
	Amount    string `json:"Amount"`
	Brand     string `json:"Brand"`
	Id        string `json:"Id"`
	Discount  string `json:"Discount"`
	Name      string `json:"Name"`
	Quantity  string `json:"Quantity"`
	Articalem string `json:"Articalem"`
	SubTotal  string `json:"SubTotal"`
}
type ServiceBlock struct {
	Discount      string `json:"Discount"`
	Id            string `json:"Id"`
	Price         string `json:"Price"`
	ServiceType   string `json:"ServiceType"`
	ServiceName   string `json:"ServiceName"`
	ServiceDetail string `json:"ServiceDetail"`
	SubTotal      string `json:"SubTotal"`
}
type AgentBlock struct {
	AgencyGarageName            string               `json:"AgencyGarageName"`
	ApproveTime                 string               `json:"ApproveTime"`
	Area                        string               `json:"Area"`
	AssignedDate                string               `json:"AssignedDate"`
	BranchName                  string               `json:"BranchName"`
	AfterRepairDocuments        []DocumentBlock      `json:"AfterRepairDocuments"`
	ProofOfRepairRemarks        []DocumentBlock      `json:"ProofOfRepairRemarks"`
	CarNo                       string               `json:"CarNo"`
	CivilId                     string               `json:"CivilId"`
	ClaimId                     string               `json:"ClaimId"`
	ClaimType                   string               `json:"ClaimType"`
	Comments                    string               `json:"Comments"`
	DeliveryRemarks             string               `json:"DeliveryRemarks"`
	RepairRemarks               string               `json:"RepairRemarks"`
	DeliveredRemarks            string               `json:"DeliveredRemarks"`
	GarageRemarks               string               `json:"GarageRemarks"`
	Email                       string               `json:"Email"`
	EstimateDayNo               string               `json:"EstimateDayNo"`
	ClaimOfficeExternalComments string               `json:"ClaimOfficeExternalComments"`
	ClaimOfficeInternalComments string               `json:"ClaimOfficeInternalComments"`
	GarageExternalComments      string               `json:"GarageExternalComments"`
	GarageInternalComments      string               `json:"GarageInternalComments"`
	PaidExternalComments        string               `json:"PaidExternalComments"`
	PaidInternalComments        string               `json:"PaidInternalComments"`
	IncidentDate                string               `json:"IncidentDate"`
	KoPay                       string               `json:"KoPay"`
	Make                        string               `json:"Make"`
	MobileNo                    string               `json:"MobileNo"`
	Model                       string               `json:"Model"`
	Name                        string               `json:"Name"`
	NetTotal                    string               `json:"NetTotal"`
	PolicyNo                    string               `json:"PolicyNo"`
	PolicyType                  string               `json:"PolicyType"`
	PolicyValidity              string               `json:"PolicyValidity"`
	Region                      string               `json:"Region"`
	RepairOption                string               `json:"RepairOption"`
	Status                      string               `json:"Status"`
	Time                        string               `json:"Time"`
	TotalEstimate               string               `json:"TotalEstimate"`
	VehiclePartsAmount          string               `json:"VehiclePartsAmount"`
	VehicleServicesAmount       string               `json:"VehicleServicesAmount"`
	DeliveryReadyFor            []DeliveryReadyBlock `json:"DeliveryReadyFor"`
	Documents                   []DocumentBlock      `json:"Documents"`
	Parts                       []PartBlock          `json:"Parts"`
	Services                    []ServiceBlock       `json:"Services"`
	Status_1                    string               `json:"Status_1"`
	ClaimOfficeRejectionReasion string               `json:"ClaimOfficeRejectionReasion"`
	ClaimOfficeRejectionTime    string               `json:"ClaimOfficeRejectionTime"`
	GarageRejectionReasion      string               `json:"GarageRejectionReasion"`
	GarageRejectionTime         string               `json:"GarageRejectionTime"`
	EstimateApprovedDate        string               `json:"EstimateApprovedDate"`
}

type HolderBlock struct {
	AgencyGarageName            string               `json:"AgencyGarageName"`
	ApproveTime                 string               `json:"ApproveTime"`
	Area                        string               `json:"Area"`
	AssignedDate                string               `json:"AssignedDate"`
	BranchName                  string               `json:"BranchName"`
	AfterRepairDocuments        []DocumentBlock      `json:"AfterRepairDocuments"`
	ProofOfRepairRemarks        []DocumentBlock      `json:"ProofOfRepairRemarks"`
	CarNo                       string               `json:"CarNo"`
	CivilId                     string               `json:"CivilId"`
	ClaimId                     string               `json:"ClaimId"`
	ClaimType                   string               `json:"ClaimType"`
	Comments                    string               `json:"Comments"`
	GarageRemarks               string               `json:"GarageRemarks"`
	Email                       string               `json:"Email"`
	EstimateDayNo               string               `json:"EstimateDayNo"`
	ClaimOfficeExternalComments string               `json:"ClaimOfficeExternalComments"`
	GarageExternalComments      string               `json:"GarageExternalComments"`
	PaidExternalComments        string               `json:"PaidExternalComments"`
	IncidentDate                string               `json:"IncidentDate"`
	KoPay                       string               `json:"KoPay"`
	Make                        string               `json:"Make"`
	MobileNo                    string               `json:"MobileNo"`
	Model                       string               `json:"Model"`
	Name                        string               `json:"Name"`
	NetTotal                    string               `json:"NetTotal"`
	PolicyNo                    string               `json:"PolicyNo"`
	PolicyType                  string               `json:"PolicyType"`
	PolicyValidity              string               `json:"PolicyValidity"`
	Region                      string               `json:"Region"`
	RepairOption                string               `json:"RepairOption"`
	Status                      string               `json:"Status"`
	Time                        string               `json:"Time"`
	DeliveryRemarks             string               `json:"DeliveryRemarks"`
	RepairRemarks               string               `json:"RepairRemarks"`
	DeliveredRemarks            string               `json:"DeliveredRemarks"`
	TotalEstimate               string               `json:"TotalEstimate"`
	VehiclePartsAmount          string               `json:"VehiclePartsAmount"`
	VehicleServicesAmount       string               `json:"VehicleServicesAmount"`
	DeliveryReadyFor            []DeliveryReadyBlock `json:"DeliveryReadyFor"`
	Documents                   []DocumentBlock      `json:"Documents"`
	Parts                       []PartBlock          `json:"Parts"`
	Services                    []ServiceBlock       `json:"Services"`
	Status_1                    string               `json:"Status_1"`
	ClaimOfficeRejectionReasion string               `json:"ClaimOfficeRejectionReasion"`
	ClaimOfficeRejectionTime    string               `json:"ClaimOfficeRejectionTime"`
}

type GarageBlock struct {
	AgencyGarageName            string               `json:"AgencyGarageName"`
	ApproveTime                 string               `json:"ApproveTime"`
	Area                        string               `json:"Area"`
	AssignedDate                string               `json:"AssignedDate"`
	BranchName                  string               `json:"BranchName"`
	CarNo                       string               `json:"CarNo"`
	CivilId                     string               `json:"CivilId"`
	ClaimId                     string               `json:"ClaimId"`
	ClaimType                   string               `json:"ClaimType"`
	Comments                    string               `json:"Comments"`
	DeliveryRemarks             string               `json:"DeliveryRemarks"`
	RepairRemarks               string               `json:"RepairRemarks"`
	DeliveredRemarks            string               `json:"DeliveredRemarks"`
	GarageRemarks               string               `json:"GarageRemarks"`
	Email                       string               `json:"Email"`
	EstimateDayNo               string               `json:"EstimateDayNo"`
	ClaimOfficeExternalComments string               `json:"ClaimOfficeExternalComments"`
	ClaimOfficeInternalComments string               `json:"ClaimOfficeInternalComments"`
	GarageExternalComments      string               `json:"GarageExternalComments"`
	GarageInternalComments      string               `json:"GarageInternalComments"`
	IncidentDate                string               `json:"IncidentDate"`
	KoPay                       string               `json:"KoPay"`
	Make                        string               `json:"Make"`
	MobileNo                    string               `json:"MobileNo"`
	Model                       string               `json:"Model"`
	Name                        string               `json:"Name"`
	NetTotal                    string               `json:"NetTotal"`
	PolicyNo                    string               `json:"PolicyNo"`
	PolicyType                  string               `json:"PolicyType"`
	PolicyValidity              string               `json:"PolicyValidity"`
	Region                      string               `json:"Region"`
	RepairOption                string               `json:"RepairOption"`
	Status                      string               `json:"Status"`
	Time                        string               `json:"Time"`
	TotalEstimate               string               `json:"TotalEstimate"`
	VehiclePartsAmount          string               `json:"VehiclePartsAmount"`
	VehicleServicesAmount       string               `json:"VehicleServicesAmount"`
	DeliveryReadyFor            []DeliveryReadyBlock `json:"DeliveryReadyFor"`
	Documents                   []DocumentBlock      `json:"Documents"`
	Parts                       []PartBlock          `json:"Parts"`
	Services                    []ServiceBlock       `json:"Services"`
	Status_1                    string               `json:"Status_1"`
	ClaimOfficeRejectionReasion string               `json:"ClaimOfficeRejectionReasion"`
	ClaimOfficeRejectionTime    string               `json:"ClaimOfficeRejectionTime"`
	GarageRejectionReasion      string               `json:"GarageRejectionReasion"`
	GarageRejectionTime         string               `json:"GarageRejectionTime"`
	EstimateApprovedDate        string               `json:"EstimateApprovedDate"`
}
type ClaimBlock struct {
	ClaimId        string      `json:"ClaimId"`
	InsuranceAgent AgentBlock  `json:"InsuranceAgent"`
	PolicyHolder   HolderBlock `json:"PolicyHolder"`
	AgencyGarage   GarageBlock `json:"AgencyGarage"`
}

// InitLedger adds a base set of invoices to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {

	Claims := []ClaimBlock{}

	for i, claim := range Claims {
		claimAsBytes, _ := json.Marshal(claim)
		fmt.Println(i)
		err := ctx.GetStub().PutState(claim.ClaimId, claimAsBytes)

		if err != nil {
			return fmt.Errorf("Failed to put to world state. %s", err.Error())
		}
	}

	return nil
}

// CreateOrder adds a new order to the world state with given details
func (s *SmartContract) CreateClaim(ctx contractapi.TransactionContextInterface, ClaimId string,
	ClaimType string,
	RepairOption string,
	IncidentDate string,
	Region string,
	Area string,
	Comments string,
	Time string,
	Documents []DocumentBlock,
	PolicyNo string,
	PolicyType string,
	PolicyValidity string,
	CarNo string,
	Model string,
	Make string,
	CivilId string,
	Name string,
	Email string,
	MobileNo string,
	Status string,
	Status_1 string,
) error {

	claim := ClaimBlock{
		ClaimId: ClaimId,
		InsuranceAgent: AgentBlock{
			ClaimId:        ClaimId,
			ClaimType:      ClaimType,
			RepairOption:   RepairOption,
			IncidentDate:   IncidentDate,
			Region:         Region,
			Area:           Area,
			Comments:       Comments,
			Time:           Time,
			Documents:      Documents,
			PolicyNo:       PolicyNo,
			PolicyType:     PolicyType,
			PolicyValidity: PolicyValidity,
			CarNo:          CarNo,
			Model:          Model,
			Make:           Make,
			CivilId:        CivilId,
			Name:           Name,
			Email:          Email,
			MobileNo:       MobileNo,
			Status:         Status,
			Status_1:       Status_1,
		},
		PolicyHolder: HolderBlock{
			ClaimId:        ClaimId,
			ClaimType:      ClaimType,
			RepairOption:   RepairOption,
			IncidentDate:   IncidentDate,
			Region:         Region,
			Area:           Area,
			Comments:       Comments,
			Time:           Time,
			Documents:      Documents,
			PolicyNo:       PolicyNo,
			PolicyType:     PolicyType,
			PolicyValidity: PolicyValidity,
			CarNo:          CarNo,
			Model:          Model,
			Make:           Make,
			CivilId:        CivilId,
			Name:           Name,
			Email:          Email,
			MobileNo:       MobileNo,
			Status:         Status,
			Status_1:       Status_1,
		},
		AgencyGarage: GarageBlock{},
	}

	claimAsBytes, _ := json.Marshal(claim)

	return ctx.GetStub().PutState(ClaimId, claimAsBytes)
}

// UpdateOrder in world state
func (s *SmartContract) setClaimStatusApprove(ctx contractapi.TransactionContextInterface,
	ClaimId string, Status string, ExternalComments string,
	InternalComments string, ApproveTime string) error {

	claim, err := s.QueryClaim(ctx, ClaimId)

	if err != nil {
		return err
	}
	claim.InsuranceAgent.ApproveTime = ApproveTime
	claim.InsuranceAgent.ClaimOfficeInternalComments = InternalComments
	claim.InsuranceAgent.ClaimOfficeExternalComments = ExternalComments
	claim.InsuranceAgent.Status = Status
	claim.PolicyHolder.ApproveTime = ApproveTime
	claim.PolicyHolder.ClaimOfficeExternalComments = ExternalComments
	claim.PolicyHolder.Status = Status

	claimAsBytes, _ := json.Marshal(claim)

	return ctx.GetStub().PutState(ClaimId, claimAsBytes)
}

// UpdateOrder in world state
func (s *SmartContract) setClaimStatusReject(ctx contractapi.TransactionContextInterface,
	ClaimId string, Status string, ExternalComments string,
	InternalComments string, RejectionTime string, RejectionReasion string) error {

	claim, err := s.QueryClaim(ctx, ClaimId)

	if err != nil {
		return err
	}
	claim.InsuranceAgent.ClaimOfficeRejectionTime = RejectionTime
	claim.InsuranceAgent.ClaimOfficeInternalComments = InternalComments
	claim.InsuranceAgent.ClaimOfficeExternalComments = ExternalComments
	claim.InsuranceAgent.ClaimOfficeRejectionReasion = RejectionReasion
	claim.InsuranceAgent.Status = Status
	claim.PolicyHolder.ClaimOfficeRejectionTime = RejectionTime
	claim.PolicyHolder.ClaimOfficeExternalComments = ExternalComments
	claim.PolicyHolder.Status = Status
	claim.PolicyHolder.ClaimOfficeRejectionReasion = RejectionReasion

	claimAsBytes, _ := json.Marshal(claim)

	return ctx.GetStub().PutState(ClaimId, claimAsBytes)
}

// UpdateOrder in world state
func (s *SmartContract) AssignToGarage(ctx contractapi.TransactionContextInterface,
	ClaimId string, Status string, BranchName string,
	AssignedDate string, AgencyGarageName string) error {

	claim, err := s.QueryClaim(ctx, ClaimId)

	if err != nil {
		return err
	}
	claim.InsuranceAgent.BranchName = BranchName
	claim.InsuranceAgent.AssignedDate = AssignedDate
	claim.InsuranceAgent.AgencyGarageName = AgencyGarageName
	claim.InsuranceAgent.Status = Status
	claim.AgencyGarage = GarageBlock{
		ClaimId:                     claim.InsuranceAgent.ClaimId,
		ClaimType:                   claim.InsuranceAgent.ClaimType,
		RepairOption:                claim.InsuranceAgent.RepairOption,
		IncidentDate:                claim.InsuranceAgent.IncidentDate,
		Region:                      claim.InsuranceAgent.Region,
		Area:                        claim.InsuranceAgent.Area,
		Comments:                    claim.InsuranceAgent.Comments,
		Time:                        claim.InsuranceAgent.Time,
		Documents:                   claim.InsuranceAgent.Documents,
		PolicyNo:                    claim.InsuranceAgent.PolicyNo,
		PolicyType:                  claim.InsuranceAgent.PolicyType,
		PolicyValidity:              claim.InsuranceAgent.PolicyValidity,
		CarNo:                       claim.InsuranceAgent.CarNo,
		Model:                       claim.InsuranceAgent.Model,
		Make:                        claim.InsuranceAgent.Make,
		CivilId:                     claim.InsuranceAgent.CivilId,
		Name:                        claim.InsuranceAgent.Name,
		Email:                       claim.InsuranceAgent.Email,
		MobileNo:                    claim.InsuranceAgent.MobileNo,
		Status:                      claim.InsuranceAgent.Status,
		Status_1:                    claim.InsuranceAgent.Status_1,
		ApproveTime:                 claim.InsuranceAgent.ApproveTime,
		ClaimOfficeExternalComments: claim.AgencyGarage.ClaimOfficeExternalComments,
	}

	claimAsBytes, _ := json.Marshal(claim)

	return ctx.GetStub().PutState(ClaimId, claimAsBytes)
}

// UpdateOrder in world state
func (s *SmartContract) setGarageClaimReject(ctx contractapi.TransactionContextInterface,
	ClaimId string, Status string, Remarks string,
	RejectionTime string, RejectionReasion string) error {

	claim, err := s.QueryClaim(ctx, ClaimId)

	if err != nil {
		return err
	}
	claim.InsuranceAgent.GarageRemarks = Remarks
	claim.InsuranceAgent.GarageRejectionReasion = RejectionReasion
	claim.InsuranceAgent.GarageRejectionTime = RejectionTime
	claim.InsuranceAgent.Status = Status
	claim.AgencyGarage.GarageRemarks = Remarks
	claim.AgencyGarage.GarageRejectionReasion = RejectionReasion
	claim.AgencyGarage.GarageRejectionTime = RejectionTime
	claim.AgencyGarage.Status = Status

	claimAsBytes, _ := json.Marshal(claim)

	return ctx.GetStub().PutState(ClaimId, claimAsBytes)
}

// UpdateOrder in world state
func (s *SmartContract) setGarageInitialEstimate(ctx contractapi.TransactionContextInterface,
	ClaimId string, Status string, Parts []PartBlock,
	Services []ServiceBlock, EstimateDayno string, TotalEstimate string,
	VehiclePartsAmount string, VehicleServicesAmount string) error {

	claim, err := s.QueryClaim(ctx, ClaimId)

	if err != nil {
		return err
	}
	claim.InsuranceAgent.Parts = Parts
	claim.InsuranceAgent.Services = Services
	claim.InsuranceAgent.EstimateDayNo = EstimateDayno
	claim.InsuranceAgent.TotalEstimate = TotalEstimate
	claim.InsuranceAgent.VehiclePartsAmount = VehiclePartsAmount
	claim.InsuranceAgent.VehicleServicesAmount = VehicleServicesAmount
	claim.InsuranceAgent.Status = Status
	claim.AgencyGarage.Parts = Parts
	claim.AgencyGarage.Services = Services
	claim.AgencyGarage.EstimateDayNo = EstimateDayno
	claim.AgencyGarage.TotalEstimate = TotalEstimate
	claim.AgencyGarage.VehiclePartsAmount = VehiclePartsAmount
	claim.AgencyGarage.VehicleServicesAmount = VehicleServicesAmount
	claim.AgencyGarage.Status = Status

	claimAsBytes, _ := json.Marshal(claim)

	return ctx.GetStub().PutState(ClaimId, claimAsBytes)
}

// UpdateOrder in world state
func (s *SmartContract) ClaimEstimateStatusApprove(ctx contractapi.TransactionContextInterface,
	ClaimId string, Status string, EstimateDate string) error {

	claim, err := s.QueryClaim(ctx, ClaimId)

	if err != nil {
		return err
	}
	claim.InsuranceAgent.EstimateApprovedDate = EstimateDate
	claim.InsuranceAgent.Status = Status
	claim.AgencyGarage.EstimateApprovedDate = EstimateDate
	claim.AgencyGarage.Status = Status

	claimAsBytes, _ := json.Marshal(claim)

	return ctx.GetStub().PutState(ClaimId, claimAsBytes)
}

// UpdateOrder in world state
func (s *SmartContract) SendEstimateToHolder(ctx contractapi.TransactionContextInterface,
	ClaimId string, Status string, NetTotal string, KoPay string) error {

	claim, err := s.QueryClaim(ctx, ClaimId)

	if err != nil {
		return err
	}
	claim.InsuranceAgent.NetTotal = NetTotal
	claim.InsuranceAgent.KoPay = KoPay
	claim.InsuranceAgent.Status = Status
	claim.PolicyHolder = HolderBlock{
		AgencyGarageName:            claim.InsuranceAgent.AgencyGarageName,
		ApproveTime:                 claim.InsuranceAgent.ApproveTime,
		Area:                        claim.InsuranceAgent.Area,
		AssignedDate:                claim.InsuranceAgent.AssignedDate,
		BranchName:                  claim.InsuranceAgent.BranchName,
		CarNo:                       claim.InsuranceAgent.CarNo,
		CivilId:                     claim.InsuranceAgent.CivilId,
		ClaimId:                     claim.InsuranceAgent.ClaimId,
		ClaimType:                   claim.InsuranceAgent.ClaimType,
		Comments:                    claim.InsuranceAgent.Comments,
		GarageRemarks:               claim.InsuranceAgent.GarageRemarks,
		Email:                       claim.InsuranceAgent.Email,
		EstimateDayNo:               claim.InsuranceAgent.EstimateDayNo,
		ClaimOfficeExternalComments: claim.InsuranceAgent.ClaimOfficeExternalComments,
		GarageExternalComments:      claim.InsuranceAgent.GarageExternalComments,
		IncidentDate:                claim.InsuranceAgent.IncidentDate,
		KoPay:                       claim.InsuranceAgent.KoPay,
		Make:                        claim.InsuranceAgent.Make,
		MobileNo:                    claim.InsuranceAgent.MobileNo,
		Model:                       claim.InsuranceAgent.Model,
		Name:                        claim.InsuranceAgent.Name,
		NetTotal:                    claim.InsuranceAgent.NetTotal,
		PolicyNo:                    claim.InsuranceAgent.PolicyNo,
		PolicyType:                  claim.InsuranceAgent.PolicyType,
		PolicyValidity:              claim.InsuranceAgent.PolicyValidity,
		Region:                      claim.InsuranceAgent.Region,
		RepairOption:                claim.InsuranceAgent.RepairOption,
		Status:                      claim.InsuranceAgent.Status,
		Time:                        claim.InsuranceAgent.Time,
		TotalEstimate:               claim.InsuranceAgent.TotalEstimate,
		VehiclePartsAmount:          claim.InsuranceAgent.VehiclePartsAmount,
		VehicleServicesAmount:       claim.InsuranceAgent.VehicleServicesAmount,
		DeliveryReadyFor:            claim.InsuranceAgent.DeliveryReadyFor,
		Documents:                   claim.InsuranceAgent.Documents,
		Parts:                       claim.InsuranceAgent.Parts,
		Services:                    claim.InsuranceAgent.Services,
		Status_1:                    claim.InsuranceAgent.Status_1,
		ClaimOfficeRejectionReasion: claim.InsuranceAgent.ClaimOfficeRejectionReasion,
		ClaimOfficeRejectionTime:    claim.InsuranceAgent.ClaimOfficeRejectionTime,
	}

	claimAsBytes, _ := json.Marshal(claim)

	return ctx.GetStub().PutState(ClaimId, claimAsBytes)
}

func (s *SmartContract) ApprovalFromHolder(ctx contractapi.TransactionContextInterface,
	ClaimId string, Status string) error {

	claim, err := s.QueryClaim(ctx, ClaimId)

	if err != nil {
		return err
	}
	claim.InsuranceAgent.Status = Status
	claim.PolicyHolder.Status = Status
	claim.AgencyGarage.Status = Status

	claimAsBytes, _ := json.Marshal(claim)

	return ctx.GetStub().PutState(ClaimId, claimAsBytes)
}

func (s *SmartContract) ReadyforDelivery(ctx contractapi.TransactionContextInterface,
	ClaimId string, Status string, DeliverRemarks string, DeliveryReadyFor []DeliveryReadyBlock) error {

	claim, err := s.QueryClaim(ctx, ClaimId)

	if err != nil {
		return err
	}
	claim.InsuranceAgent.Status = Status
	claim.InsuranceAgent.DeliveryRemarks = DeliverRemarks
	claim.InsuranceAgent.DeliveryReadyFor = DeliveryReadyFor
	claim.PolicyHolder.Status = Status
	claim.PolicyHolder.DeliveryReadyFor = DeliveryReadyFor
	claim.PolicyHolder.DeliveryRemarks = DeliverRemarks
	claim.AgencyGarage.Status = Status
	claim.AgencyGarage.DeliveryRemarks = DeliverRemarks
	claim.AgencyGarage.DeliveryReadyFor = DeliveryReadyFor

	claimAsBytes, _ := json.Marshal(claim)

	return ctx.GetStub().PutState(ClaimId, claimAsBytes)
}

func (s *SmartContract) DeliverVehical(ctx contractapi.TransactionContextInterface,
	ClaimId string, Status string, DeliveredRemarks string) error {

	claim, err := s.QueryClaim(ctx, ClaimId)

	if err != nil {
		return err
	}
	claim.InsuranceAgent.Status = Status
	claim.InsuranceAgent.DeliveredRemarks = DeliveredRemarks

	claim.PolicyHolder.Status = Status

	claim.PolicyHolder.DeliveredRemarks = DeliveredRemarks
	claim.AgencyGarage.Status = Status
	claim.AgencyGarage.DeliveredRemarks = DeliveredRemarks

	claimAsBytes, _ := json.Marshal(claim)

	return ctx.GetStub().PutState(ClaimId, claimAsBytes)
}

func (s *SmartContract) UnderRepair(ctx contractapi.TransactionContextInterface,
	ClaimId string, Status string, RepairRemarks string) error {

	claim, err := s.QueryClaim(ctx, ClaimId)

	if err != nil {
		return err
	}
	claim.InsuranceAgent.Status = Status
	claim.InsuranceAgent.RepairRemarks = RepairRemarks
	claim.PolicyHolder.Status = Status
	claim.PolicyHolder.RepairRemarks = RepairRemarks
	claim.AgencyGarage.Status = Status
	claim.AgencyGarage.RepairRemarks = RepairRemarks

	claimAsBytes, _ := json.Marshal(claim)

	return ctx.GetStub().PutState(ClaimId, claimAsBytes)
}

func (s *SmartContract) SettelByCash(ctx contractapi.TransactionContextInterface,
	ClaimId string, Status string) error {

	claim, err := s.QueryClaim(ctx, ClaimId)

	if err != nil {
		return err
	}
	claim.InsuranceAgent.Status = Status
	claim.PolicyHolder.Status = Status

	claimAsBytes, _ := json.Marshal(claim)

	return ctx.GetStub().PutState(ClaimId, claimAsBytes)
}

func (s *SmartContract) AfterRepair(ctx contractapi.TransactionContextInterface,
	ClaimId string, Status string, ProofOfRepairRemarks []DocumentBlock, AfterRepairDocuments []DocumentBlock) error {

	claim, err := s.QueryClaim(ctx, ClaimId)

	if err != nil {
		return err
	}
	claim.InsuranceAgent.Status = Status
	claim.InsuranceAgent.AfterRepairDocuments = AfterRepairDocuments
	claim.InsuranceAgent.ProofOfRepairRemarks = ProofOfRepairRemarks
	claim.PolicyHolder.Status = Status
	claim.PolicyHolder.ProofOfRepairRemarks = ProofOfRepairRemarks
	claim.PolicyHolder.AfterRepairDocuments = AfterRepairDocuments

	claimAsBytes, _ := json.Marshal(claim)

	return ctx.GetStub().PutState(ClaimId, claimAsBytes)
}

func (s *SmartContract) MarkAsPaid(ctx contractapi.TransactionContextInterface,
	ClaimId string, Status string, PaidInternalComments string, PaidExternalComments string) error {

	claim, err := s.QueryClaim(ctx, ClaimId)

	if err != nil {
		return err
	}
	claim.InsuranceAgent.Status = Status
	claim.InsuranceAgent.PaidExternalComments = PaidExternalComments
	claim.InsuranceAgent.PaidInternalComments = PaidInternalComments
	claim.PolicyHolder.Status = Status
	claim.PolicyHolder.PaidExternalComments = PaidExternalComments

	claimAsBytes, _ := json.Marshal(claim)

	return ctx.GetStub().PutState(ClaimId, claimAsBytes)
}

//

// QueryAllClaims returns all Claims found in world state
func (s *SmartContract) QueryAllClaims(ctx contractapi.TransactionContextInterface) ([]QueryResult, error) {

	startKey := ""
	endKey := ""

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []QueryResult{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}

		claim := new(ClaimBlock)
		_ = json.Unmarshal(queryResponse.Value, claim)

		queryResult := QueryResult{Key: queryResponse.Key, Record: claim}
		fmt.Println("queryResult", string(queryResponse.Value))
		results = append(results, queryResult)
	}

	return results, nil
}

// GetHistoryClaims returns the claim stored in the world state with given id
func (s *SmartContract) GetHistoryClaim(ctx contractapi.TransactionContextInterface, ClaimId string) ([]QueryHistory, error) {

	history, err := ctx.GetStub().GetHistoryForKey(ClaimId)

	if err != nil {
		return nil, fmt.Errorf("Failed to read History from world state. %s", err.Error())
	}

	if history == nil {
		return nil, fmt.Errorf("%s does not exist", ClaimId)
	}

	results := []QueryHistory{}

	for history.HasNext() {
		modification, err := history.Next()
		if err != nil {
			fmt.Println(err.Error())
			return nil, fmt.Errorf("Failed to read History from world state. %s", err.Error())
		}
		queryResult := QueryHistory{Record: string(modification.Value)}
		results = append(results, queryResult)
		fmt.Println("Returning information about", string(modification.Value))
	}

	return results, nil
}

// QueryOrder returns the claim stored in the world state with given id
func (s *SmartContract) QueryClaim(ctx contractapi.TransactionContextInterface, ClaimId string) (*ClaimBlock, error) {

	claimAsBytes, err := ctx.GetStub().GetState(ClaimId)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if claimAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", ClaimId)
	}

	claim := new(ClaimBlock)
	_ = json.Unmarshal(claimAsBytes, claim)

	return claim, nil
}

func main() {

	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		fmt.Printf("Error creating chain-apparek chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting chain-apparek chaincode: %s", err.Error())
	}
}
