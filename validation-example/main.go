package main

import (
	"fmt"
	"strings"

	"github.com/lf-edge/eve-api/go/register"
)

// TestCase represents a test case for validating register messages.
type TestCase struct {
	Name        string
	Message     *register.ZRegisterMsg
	Description string
}

func main() {
	testCases := []TestCase{
		{
			Name: "Valid message",
			Message: &register.ZRegisterMsg{
				OnBoardKey: "deprecated-key",
				PemCert:    []byte(strings.Repeat("A", 200)), // Valid length (200 bytes)
				Serial:     "DEVICE-123_ABC",                 // Valid serial
				SoftSerial: "SOFT-456_DEF",                   // Valid soft serial
			},
			Description: "All fields are valid",
		},
		{
			Name: "Invalid PemCert (too short)",
			Message: &register.ZRegisterMsg{
				PemCert:    []byte("short"), // Only 5 bytes (minimum is 100)
				Serial:     "DEVICE-123",
				SoftSerial: "SOFT-456",
			},
			Description: "PemCert is only 5 bytes, minimum is 100",
		},
		{
			Name: "Invalid Serial (empty)",
			Message: &register.ZRegisterMsg{
				PemCert:    []byte(strings.Repeat("B", 150)),
				Serial:     "", // Empty serial (minimum is 1)
				SoftSerial: "SOFT-456",
			},
			Description: "Serial is empty, minimum is 1 character",
		},
		{
			Name: "Invalid Serial (invalid characters)",
			Message: &register.ZRegisterMsg{
				PemCert:    []byte(strings.Repeat("C", 150)),
				Serial:     "DEVICE 123@#$", // Contains space and special characters
				SoftSerial: "SOFT-456",
			},
			Description: "Serial contains spaces and special characters",
		},
		{
			Name: "Invalid SoftSerial (invalid characters)",
			Message: &register.ZRegisterMsg{
				PemCert:    []byte(strings.Repeat("D", 150)),
				Serial:     "DEVICE-123",
				SoftSerial: "SOFT@456#", // Contains special characters
			},
			Description: "SoftSerial contains special characters",
		},
		{
			Name: "Multiple validation errors",
			Message: &register.ZRegisterMsg{
				PemCert:    []byte("short"), // Too short
				Serial:     "",              // Empty
				SoftSerial: "INVALID@CHARS", // Invalid characters
			},
			Description: "Multiple fields with validation errors",
		},
		{
			Name: "Valid message with empty SoftSerial",
			Message: &register.ZRegisterMsg{
				PemCert:    []byte(strings.Repeat("E", 150)),
				Serial:     "DEVICE-789",
				SoftSerial: "", // Empty is valid for SoftSerial
			},
			Description: "SoftSerial is optional and can be empty",
		},
		{
			Name: "Maximum allowed lengths",
			Message: &register.ZRegisterMsg{
				PemCert:    []byte(strings.Repeat("F", 10240)), // Maximum PemCert length
				Serial:     strings.Repeat("A", 256),           // Maximum Serial length
				SoftSerial: strings.Repeat("B", 256),           // Maximum SoftSerial length
			},
			Description: "All fields at maximum allowed lengths",
		},
	}

	// Iterate through test cases
	for i, testCase := range testCases {
		fmt.Printf("%d. %s:\n", i+1, testCase.Name)
		fmt.Printf("\tDescription: %s\n", testCase.Description)

		err := testCase.Message.Validate()
		if err != nil {
			fmt.Printf("\tValidation failed: %v\n", err)
			if validationErr, ok := err.(register.ZRegisterMsgValidationError); ok {
				fmt.Printf("\tError details:\n")
				fmt.Printf("\t\tField: %s\n", validationErr.Field())
				fmt.Printf("\t\tReason: %s\n", validationErr.Reason())
				fmt.Printf("\t\tError Name: %s\n", validationErr.ErrorName())
				fmt.Printf("\t\tFull Error: %s\n", validationErr.Error())
			}
		} else {
			fmt.Printf("\tValidation passed\n")
		}
		fmt.Println()
	}
}
