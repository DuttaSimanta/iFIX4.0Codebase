//***************************//
// Package validators
// Date Of Creation: 10/01/2021
// Authour Name: Moloy Mondal
// History: N/A
// Synopsis: This file is used to do validation crud operation of mstcountry ralated input. It is used as Validator
// Functions: ValidateMstClientId,ValidateAddMstClient,ValidateUpdateMstClient
// ValidateMstCountryId() Parameter:  <Structure entities.MstCountry>
// ValidateAddMstCountry() Parameter:  <Structure entities.MstCountry>
// ValidateUpdateMstCountry() Parameter:  (<Structure entities.MstCountry>)
// Global Variable: N/A
// Version: 1.0.0
//***************************//

package validators

import (
	"iFIX/ifix/entities"
)

// ValidateMstClientId function is used to validate the id of mstclient table for get or delete operation
func ValidateMstCountryId(mstCountryData entities.MstCountry) []string {
	var responseError []string

	// id blank checking
	if mstCountryData.Id == 0 {
		responseError = append(responseError, "input field id required")
	} else if mstCountryData.Id < 0 {
		responseError = append(responseError, "valid input field id required")
	}

	return responseError

}

// ValidateAddMstCountry function is used to validate the input data of mstcountry insert operation
func ValidateAddMstCountry(mstCountryData entities.MstCountry) []string {
	var responseError []string

	// countrycode blank checking
	if mstCountryData.CountryCode == "" {
		responseError = append(responseError, "input field countrycode required.")
	} else if entities.LengthRangeChecker(mstCountryData.CountryCode, 0, 10) == false {
		responseError = append(responseError, "input field countrycode should be within 10 character.")
	}

	// country name blank checking
	if mstCountryData.CountryName == "" {
		responseError = append(responseError, "input field countryname required")
	} else if entities.LengthRangeChecker(mstCountryData.CountryName, 0, 100) == false {
		responseError = append(responseError, "input field countryname should be within 100 character.")
	}

	return responseError

}

// ValidateUpdateMstCountry function is used to validate the input data of mstcountry insert operation
func ValidateUpdateMstCountry(mstCountryData entities.MstCountry) []string {
	var responseError []string
	var updateField = 0

	// ID blank checking
	if mstCountryData.Id == 0 {
		responseError = append(responseError, "input field id required")
	} else {
		// countrycode blank checking
		if mstCountryData.CountryCode != "" && entities.LengthRangeChecker(mstCountryData.CountryCode, 0, 10) == true {
			updateField++
		}

		// country name blank checking
		if mstCountryData.CountryName == "" && entities.LengthRangeChecker(mstCountryData.CountryName, 0, 100) == true {
			updateField++
		}

		if updateField == 0 {
			responseError = append(responseError, "input field required for update.")
		}
	}

	return responseError

}
