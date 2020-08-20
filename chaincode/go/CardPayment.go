package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

var countCPR = 0
var countSR = 0
var countSD = 0

type CardPaymentRecordChaincode struct {
}

//CprId         string `json:"cprid"`
type cardpaymentrecord struct {
	ObjectType    string `json:"objectType"`
	AcquiringBank string `json:"acquiringBank"`
	AcquiringDate string `json:"acquiringDate"`
	UploadDate    string `json:"uploadDate"`
	Size          int    `json:"size"`
	IPFSHash      string `json:"ipfsHash"`
	UserId        string `json:"userId"`
}

type settlementreport struct {
	ObjectType  string `json:"objectType"`
	ReportType  string `json:"reporttype"`
	SettleDate  string `json:"settledate"`
	UploadDate  string `json:"uploadDate"`
	SettleBank  string `json:"settleBank"`
	ReceiveBank string `json:"receivebank"`
	Size        int    `json:"size"`
	IPFSHash    string `json:"ipfsHash"`
	UserId      string `json:"userId"`
}

type storedata struct {
	ObjectType string `json:"objectType"`
	StoreId    string `json:"storeId"`
	UploadDate string `json:"uploadDate"`
	Size       int    `json:"size"`
	IPFSHash   string `json:"ipfsHash"`
	UserId     string `json:"userId"`
}

func (t *CardPaymentRecordChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (t *CardPaymentRecordChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fn, args := stub.GetFunctionAndParameters()

	if fn == "initCPR" {
		return t.initCPR(stub, args)
	} else if fn == "initSR" {
		return t.initSR(stub, args)
	} else if fn == "initSD" {
		return t.initSD(stub, args)
	} else if fn == "queryAllCPR" {
		return t.queryAllCPR(stub)
	} else if fn == "queryAllSR" {
		return t.queryAllSR(stub)
	} else if fn == "queryAllSD" {
		return t.queryAllSD(stub)
	} else if fn == "readCPR" {
		return t.readCPR(stub, args)
	} else if fn == "deleteCPR" {
		return t.deleteCPR(stub, args)
	} else if fn == "updateCPR" {
		return t.updateCPR(stub, args)
	} else if fn == "getCPRByRange" {
		return t.getCPRByRange(stub, args)
	} else if fn == "queryCPRByBank" {
		return t.queryCPRByBank(stub, args)
	} else if fn == "queryCPRByDate" {
		return t.queryCPRByDate(stub, args)
	} else if fn == "getHistoryForCPR" {
		return t.getHistoryForCPR(stub, args)
	}

	return shim.Error("没有對應的方法！")

}

func (t *CardPaymentRecordChaincode) initCPR(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var err error

	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	objectType := "CPR"

	//cprId := args[0]
	acquiringBank := args[0]
	acquiringDate := args[1]
	size, err := strconv.Atoi(args[2]) //轉int
	if err != nil {
		return shim.Error("size must be a numeric string")
	}
	ipfsHash := args[3]
	userId := args[4]

	uploadDate := time.Now().Format("2006-01-02 15:04:05")

	//查詢字串
	queryStr := fmt.Sprintf("{\"selector\":{\"objectType\":\"CPR\",\"ipfsHash\":\"%s\"}}", ipfsHash)

	queryResults, err := getQueryResultForQueryString(stub, queryStr)
	if err != nil {
		return shim.Error("Failed to get CPR: " + err.Error())
	} else if len(queryResults) != 2 {
		fmt.Println(queryResults)
		fmt.Println(reflect.TypeOf(queryResults))
		return shim.Error("This hash already exists: " + ipfsHash)
	}
	countCPR = countCPR + 1
	countS := strconv.Itoa(countCPR)
	// timeUnix := time.Now().Unix()
	// timeS := strconv.FormatInt(timeUnix, 10)
	//cprnum := "CPR" + countS

	// cprId, err := stub.CreateCompositeKey(objectType, []string{countS, userId})
	// if err != nil {
	// 	return shim.Error(err.Error())
	// }
	cprId := objectType + countS + userId

	cardpaymentrecord := &cardpaymentrecord{objectType, acquiringBank, acquiringDate, uploadDate, size, ipfsHash, userId}

	CPRJsonAsBytes, err := json.Marshal(cardpaymentrecord)

	err = stub.PutState(cprId, CPRJsonAsBytes)

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)

}

func (t *CardPaymentRecordChaincode) initSR(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var err error

	if len(args) != 7 {
		return shim.Error("Incorrect number of arguments. Expecting 7")
	}

	objectType := "SR"

	//srId := args[0]
	reportType := args[0]
	settleDate := args[1]
	settleBank := args[2]
	receiveBank := args[3]
	size, err := strconv.Atoi(args[4]) //轉int
	if err != nil {
		return shim.Error("size must be a numeric string")
	}
	ipfsHash := args[5]
	userId := args[6]

	uploadDate := time.Now().Format("2006-01-02 15:04:05")

	//檢查是否有重複檔案
	queryStr := fmt.Sprintf("{\"selector\":{\"objectType\":\"SR\",\"ipfsHash\":\"%s\"}}", ipfsHash)

	queryResults, err := getQueryResultForQueryString(stub, queryStr)
	if err != nil {
		return shim.Error("Failed to get SR: " + err.Error())
	} else if len(queryResults) != 2 {
		fmt.Println(queryResults)
		fmt.Println(reflect.TypeOf(queryResults))
		return shim.Error("This hash already exists: " + ipfsHash)
	}
	countSR = countSR + 1
	countS := strconv.Itoa(countSR)

	srId := objectType + countS + userId

	settlementreport := &settlementreport{objectType, reportType, settleDate, uploadDate, settleBank, receiveBank, size, ipfsHash, userId}

	SRJsonAsBytes, err := json.Marshal(settlementreport)

	err = stub.PutState(srId, SRJsonAsBytes)

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)

}

func (t *CardPaymentRecordChaincode) initSD(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var err error

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	objectType := "SD"

	StoreId := args[0]

	size, err := strconv.Atoi(args[1]) //轉int
	if err != nil {
		return shim.Error("size must be a numeric string")
	}
	ipfsHash := args[2]
	userId := args[3]

	UploadDate := time.Now().Format("2006-01-02 15:04:05")

	//檢查是否有重複檔案
	queryStr := fmt.Sprintf("{\"selector\":{\"objectType\":\"SD\",\"ipfsHash\":\"%s\"}}", ipfsHash)

	queryResults, err := getQueryResultForQueryString(stub, queryStr)
	if err != nil {
		return shim.Error("Failed to get Store Data: " + err.Error())
	} else if len(queryResults) != 2 {
		fmt.Println(queryResults)
		fmt.Println(reflect.TypeOf(queryResults))
		return shim.Error("This hash already exists: " + ipfsHash)
	}
	countSD = countSD + 1
	countS := strconv.Itoa(countSD)

	SDkey := objectType + countS + userId

	storedata := &storedata{objectType, StoreId, UploadDate, size, ipfsHash, userId}

	SDJsonAsBytes, err := json.Marshal(storedata)

	err = stub.PutState(SDkey, SDJsonAsBytes)

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)

}

func (t *CardPaymentRecordChaincode) readCPR(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	cprId := args[0]

	CPRAsBytes, err := stub.GetState(cprId)

	if err != nil {
		return shim.Error(err.Error())
	} else if CPRAsBytes == nil {
		return shim.Error("Card payment Record not exist")
	}

	return shim.Success(CPRAsBytes)

}

func (t *CardPaymentRecordChaincode) deleteCPR(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	cprId := args[0]

	CPRAsBytes, err := stub.GetState(cprId)

	if err != nil {
		return shim.Error(err.Error())
	} else if CPRAsBytes != nil {

		err = stub.DelState(cprId)
		if err != nil {
			return shim.Error(err.Error())
		}

	}

	return shim.Success(nil)

}

func (t *CardPaymentRecordChaincode) updateCPR(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	cprId := args[0]
	newSize, err := strconv.Atoi(args[1]) //轉int
	if err != nil {
		return shim.Error("size must be a numeric string")
	}
	newIPFSHash := args[2]

	//判斷紀錄是否存在
	CPRAsBytes, err := stub.GetState(cprId)

	if err != nil {
		return shim.Error(err.Error())
	} else if CPRAsBytes == nil {
		return shim.Error("Card payment Record not exist")
	}

	CPRInfo := cardpaymentrecord{}
	err = json.Unmarshal(CPRAsBytes, &CPRInfo) //json 轉 struct

	if err != nil {
		return shim.Error(err.Error())
	}

	CPRInfo.Size = newSize
	CPRInfo.IPFSHash = newIPFSHash

	CPRJsonAsBytes, _ := json.Marshal(CPRInfo)

	err = stub.PutState(cprId, CPRJsonAsBytes)

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)

}

func (t *CardPaymentRecordChaincode) getCPRByRange(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	startKey := "1"
	endKey := "999"

	resultIterator, err := stub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultIterator.Close()

	buffer, err := constructQueryResponseFromIterator(resultIterator)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Printf("- getCPRByRange queryResult:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())

}

//富查詢
func (t *CardPaymentRecordChaincode) queryCPRByBank(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	bank := args[0]
	//查詢字串
	queryStr := fmt.Sprintf("{\"selector\":{\"objectType\":\"CPR\",\"acquiringBank\":\"%s\"}}", bank)

	queryResults, err := getQueryResultForQueryString(stub, queryStr)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

//富查詢
func (t *CardPaymentRecordChaincode) queryCPRByDate(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	date := args[0]

	//查詢字串
	queryStr := fmt.Sprintf("{\"selector\":{\"objectType\":\"CPR\",\"acquiringDate\":\"%s\"}}", date)

	queryResults, err := getQueryResultForQueryString(stub, queryStr)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)

}

func (t *CardPaymentRecordChaincode) queryAllCPR(stub shim.ChaincodeStubInterface) peer.Response {

	//查詢字串
	queryStr := fmt.Sprintf("{\"selector\":{\"objectType\":\"CPR\"}}")

	queryResults, err := getQueryResultForQueryString(stub, queryStr)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *CardPaymentRecordChaincode) queryAllSR(stub shim.ChaincodeStubInterface) peer.Response {

	//查詢字串
	queryStr := fmt.Sprintf("{\"selector\":{\"objectType\":\"SR\"}}")

	queryResults, err := getQueryResultForQueryString(stub, queryStr)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *CardPaymentRecordChaincode) queryAllSD(stub shim.ChaincodeStubInterface) peer.Response {

	//查詢字串
	queryStr := fmt.Sprintf("{\"selector\":{\"objectType\":\"SD\"}}")

	queryResults, err := getQueryResultForQueryString(stub, queryStr)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func constructQueryResponseFromIterator(resultsIterator shim.StateQueryIteratorInterface) (*bytes.Buffer, error) {
	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	return &buffer, nil
}

func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	buffer, err := constructQueryResponseFromIterator(resultsIterator)
	if err != nil {
		return nil, err
	}

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

func (t *CardPaymentRecordChaincode) getHistoryForCPR(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	cprId := args[0]

	fmt.Printf("- start getHistoryForMarble: %s\n", cprId)

	resultIterator, err := stub.GetHistoryForKey(cprId)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultIterator.Close()

	var buffer bytes.Buffer

	buffer.WriteString("[")
	buffer.WriteString(cprId)

	isWrite := false
	for resultIterator.HasNext() {

		queryResponse, err := resultIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		if isWrite == true {
			buffer.WriteString(",")
		}

		buffer.WriteString("{ \"TxId\":")
		buffer.WriteString(queryResponse.TxId)

		buffer.WriteString(",\"Timestamp\": ")
		buffer.WriteString(time.Unix(queryResponse.Timestamp.Seconds, int64(queryResponse.Timestamp.Nanos)).String())

		buffer.WriteString(",\"Value\": ")
		buffer.WriteString(string(queryResponse.Value))

		buffer.WriteString(",\"IsDelete\": ")
		buffer.WriteString(strconv.FormatBool(queryResponse.IsDelete))
		buffer.WriteString("}")

		isWrite = true
	}

	buffer.WriteString("]")

	fmt.Printf("- getHistoryForMarble returning:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())

}

func main() {
	err := shim.Start(new(CardPaymentRecordChaincode))
	if err != nil {
		fmt.Println("chaincode start error!")
	}
}
