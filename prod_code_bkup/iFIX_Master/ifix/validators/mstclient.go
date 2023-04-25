//***************************//
// Package validators
// Date Of Creation: 10/01/2021
// Authour Name: Moloy Mondal
// History: N/A
// Synopsis: This file is used to do validation crud operation of mstclient ralated input. It is used as Validator
// Functions: ValidateMstClientId,ValidateAddMstClient,ValidateUpdateMstClient
// ValidateMstClientId() Parameter:  <Structure entities.MstClient>
// ValidateAddMstClient() Parameter:  <Structure entities.MstClient>
// ValidateUpdateMstClient() Parameter:  (<Structure entities.MstClient>)
// Global Variable: N/A
// Version: 1.0.0
//***************************//

package validators

import (
	"iFIX/ifix/entities"
)

// ValidateMstClientId function is used to validate the id of mstclient table for get or delete operation
func ValidateMstClientId(mstClientData entities.MstClient) []string {
	var responseError []string

	// id blank checking
	if mstClientData.Id == 0 {
		responseError = append(responseError, "input field id required")
	} else if mstClientData.Id < 0 {
		responseError = append(responseError, "valid input field id required")
	}

	return responseError

}

// ValidateAddMstClient function is used to validate the input data of mstclient insert operation
func ValidateAddMstClient(mstClientData entities.MstClient) []string {
	var responseError []string

	// mstcode blank checking
	if mstClientData.Code == "" {
		responseError = append(responseError, "input field code required")
	}

	// mstname blank checking
	if mstClientData.Name == "" {
		responseError = append(responseError, "input field name required")
	}

	// description blank checking
	if mstClientData.Description == "" {
		responseError = append(responseError, "input field description required")
	}

	// clientauditflag blank checking
	if mstClientData.ClientAuditFlg == "" {
		responseError = append(responseError, "input field clientauditflag required")
	} else if entities.LengthChecker(mstClientData.ClientAuditFlg, 1) == false { // clientauditflag length checking
		responseError = append(responseError, "Client Audit Flag should be single character")
	}

	return responseError

}

// ValidateUpdateMstClient function is used to validate the input data of mstclient insert operation
func ValidateUpdateMstClient(mstClientData entities.MstClient) []string {
	var responseError []string
	var updateField = 0

	// clientauditflag blank checking
	if mstClientData.Id == 0 {
		responseError = append(responseError, "input field code required")
	} else {
		// mstcode blank checking
		if mstClientData.Code != "" {
			updateField++
		}

		// mstname blank checking
		if mstClientData.Name != "" {
			updateField++
		}

		// description blank checking
		if mstClientData.Description != "" {
			updateField++
		}

		// clientauditflag blank checking
		if mstClientData.ClientAuditFlg != "" {
			if entities.LengthChecker(mstClientData.ClientAuditFlg, 1) == false { // clientauditflag length checking
				responseError = append(responseError, "Client Audit Flag should be single character")
			} else {
				updateField++
			}
		}

		if updateField == 0 {
			responseError = append(responseError, "input field required for update.")
		}
	}

	return responseError

}
