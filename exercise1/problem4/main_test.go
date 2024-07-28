package main

import (
	"testing"
)

func TestDetectWord(t *testing.T) {
	table := []struct {
		crowd string
		exp   string
	}{
		{"UcUNFYGaFYFYGtNUH", "cat"},
		{"bEEFGBuFBRrHgUHlNFYaYr", "burglar"},
		{"YFemHUFBbezFBYzFBYLleGBYEFGBMENTment", "embezzlement"},
		{"cLXSNVVJVOJBIQRVKIZWKJOIVHXELVReLXSNVVJVOJBIQRVKIZWKJOIVHXELVRrLXSNVVJVOJBIQRVKIZWKJOIVHXELVRtLXSNVVJVOJBIQRVKIZWKJOIVHXELVRaLXSNVVJVOJBIQRVKIZWKJOIVHXELVRiLXSNVVJVOJBIQRVKIZWKJOIVHXELVRn", "certain"},
		{"cUEOYCSUXVOaUEOYCSUXVOt", "cat"},
		{"vJAQSZNYRQTFUHDHSDMBDPUNFQJXSXeJAQSZNYRQTFUHDHSDMBDPUNFQJXSXgJAQSZNYRQTFUHDHSDMBDPUNFQJXSXeJAQSZNYRQTFUHDHSDMBDPUNFQJXSXtJAQSZNYRQTFUHDHSDMBDPUNFQJXSXaJAQSZNYRQTFUHDHSDMBDPUNFQJXSXbJAQSZNYRQTFUHDHSDMBDPUNFQJXSXlJAQSZNYRQTFUHDHSDMBDPUNFQJXSXe", "vegetable"},
		{"dATIQTJLBZFHSRXWOZQMOKZPANOUGMeATIQTJLBZFHSRXWOZQMOKZPANOUGMlATIQTJLBZFHSRXWOZQMOKZPANOUGMiATIQTJLBZFHSRXWOZQMOKZPANOUGMgATIQTJLBZFHSRXWOZQMOKZPANOUGMhATIQTJLBZFHSRXWOZQMOKZPANOUGMt", "delight"},
		{"pUBOKJGODIJBSXPMTODCGHATrUBOKJGODIJBSXPMTODCGHATiUBOKJGODIJBSXPMTODCGHATcUBOKJGODIJBSXPMTODCGHATeUBOKJGODIJBSXPMTODCGHATy", "pricey"},
		{"sWRRKMVJVHHZTKAQTJUQDPKHSHPOYCnWRRKMVJVHHZTKAQTJUQDPKHSHPOYCaWRRKMVJVHHZTKAQTJUQDPKHSHPOYCkWRRKMVJVHHZTKAQTJUQDPKHSHPOYCe", "snake"},
		{"aJULRJHMOVLEFVJZnJULRJHMOVLEFVJZgJULRJHMOVLEFVJZlJULRJHMOVLEFVJZe", "angle"},
		{"aJWCHXONGQCXGPXLZQBKEIHZWwJWCHXONGQCXGPXLZQBKEIHZWaJWCHXONGQCXGPXLZQBKEIHZWrJWCHXONGQCXGPXLZQBKEIHZWe", "aware"},
		{"nNUZKGKNEVZBPQZQQLHZZPaNUZKGKNEVZBPQZQQLHZZPmNUZKGKNEVZBPQZQQLHZZPe", "name"},
		{"cLBFKXYQFLLElLBFKXYQFLLEeLBFKXYQFLLEvLBFKXYQFLLEeLBFKXYQFLLEr", "clever"},
		{"bDUNEPWILKUFNTRGMBRSVGAABBFCCXErDUNEPWILKUFNTRGMBRSVGAABBFCCXEaDUNEPWILKUFNTRGMBRSVGAABBFCCXEsDUNEPWILKUFNTRGMBRSVGAABBFCCXEh", "brash"},
		{"fXTTJVWFCHYZMaXTTJVWFCHYZMsXTTJVWFCHYZMt", "fast"},
		{"dJYPHZIRXYOLDGAQUPHIZTXJOKNoJYPHZIRXYOLDGAQUPHIZTXJOKNwJYPHZIRXYOLDGAQUPHIZTXJOKNnJYPHZIRXYOLDGAQUPHIZTXJOKNtJYPHZIRXYOLDGAQUPHIZTXJOKNoJYPHZIRXYOLDGAQUPHIZTXJOKNwJYPHZIRXYOLDGAQUPHIZTXJOKNn", "downtown"},
		{"pKICNUFWFNFORlKICNUFWFNFORaKICNUFWFNFORnKICNUFWFNFORtKICNUFWFNFORs", "plants"},
		{"wLAXIBDWXVPRQOOQRRTOYRODLAQHiLAXIBDWXVPRQOOQRRTOYRODLAQHnLAXIBDWXVPRQOOQRRTOYRODLAQHdLAXIBDWXVPRQOOQRRTOYRODLAQHy", "windy"},
		{"sELJQETMYLTDKXYNSSOISZFPMAtELJQETMYLTDKXYNSSOISZFPMAaELJQETMYLTDKXYNSSOISZFPMArELJQETMYLTDKXYNSSOISZFPMAt", "start"},
		{"wQYKDHGMNYMKUHKDeQYKDHGMNYMKUHKDt", "wet"},
		{"kVOJQJIFILEHVnVOJQJIFILEHViVOJQJIFILEHVfVOJQJIFILEHVe", "knife"},
		{"nBKCXNIJYJSVDoBKCXNIJYJSVDtBKCXNIJYJSVDe", "note"},
		{"bOEYZAJVFYUGXQWZXrOEYZAJVFYUGXQWZXuOEYZAJVFYUGXQWZXsOEYZAJVFYUGXQWZXh", "brush"},
		{"tEMVSYRPYHSZRLJNOMTYRPREIHoEMVSYRPYHSZRLJNOMTYRPREIHoEMVSYRPYHSZRLJNOMTYRPREIHtEMVSYRPYHSZRLJNOMTYRPREIHhEMVSYRPYHSZRLJNOMTYRPREIHbEMVSYRPYHSZRLJNOMTYRPREIHrEMVSYRPYHSZRLJNOMTYRPREIHuEMVSYRPYHSZRLJNOMTYRPREIHsEMVSYRPYHSZRLJNOMTYRPREIHh", "toothbrush"},
		{"sWRIQGRPNHQQPSIPRoWRIQGRPNHQQPSIPRgWRIQGRPNHQQPSIPRgWRIQGRPNHQQPSIPRy", "soggy"},
		{"fRBODZACXIIXHZRGKJQMDLOONTlRBODZACXIIXHZRGKJQMDLOONToRBODZACXIIXHZRGKJQMDLOONTwRBODZACXIIXHZRGKJQMDLOONTeRBODZACXIIXHZRGKJQMDLOONTrRBODZACXIIXHZRGKJQMDLOONTs", "flowers"},
		{"dPWUSQZDQIHANDHEQUZBLAULSoPWUSQZDQIHANDHEQUZBLAULSlPWUSQZDQIHANDHEQUZBLAULSl", "doll"},
		{"aOGSREBZUHUEJYSSBUlOGSREBZUHUEJYSSBUoOGSREBZUHUEJYSSBUoOGSREBZUHUEJYSSBUf", "aloof"},
		{"aGQEAESDQIBWRUTuGQEAESDQIBWRUTsGQEAESDQIBWRUTpGQEAESDQIBWRUTiGQEAESDQIBWRUTcGQEAESDQIBWRUTiGQEAESDQIBWRUToGQEAESDQIBWRUTuGQEAESDQIBWRUTs", "auspicious"},
		{"mPVIWSNGHMXHaPVIWSNGHMXHrPVIWSNGHMXHkPVIWSNGHMXHePVIWSNGHMXHt", "market"},
		{"dUMIHKRZLPJFAGUKPGXHiUMIHKRZLPJFAGUKPGXHlUMIHKRZLPJFAGUKPGXHiUMIHKRZLPJFAGUKPGXHgUMIHKRZLPJFAGUKPGXHeUMIHKRZLPJFAGUKPGXHnUMIHKRZLPJFAGUKPGXHt", "diligent"},
		{"sPRTRRRETBCDTtPRTRRRETBCDTrPRTRRRETBCDTiPRTRRRETBCDTpPRTRRRETBCDTePRTRRRETBCDTd", "striped"},
		{"mDKXCLZDVPRNMGGFGEOZoDKXCLZDVPRNMGGFGEOZoDKXCLZDVPRNMGGFGEOZn", "moon"},
		{"aQQWPQYQEEDILHYDSGQAINQZWCABYcQQWPQYQEEDILHYDSGQAINQZWCABYcQQWPQYQEEDILHYDSGQAINQZWCABYoQQWPQYQEEDILHYDSGQAINQZWCABYuQQWPQYQEEDILHYDSGQAINQZWCABYnQQWPQYQEEDILHYDSGQAINQZWCABYt", "account"},
		{"sJPUQNBZOQYREGGQSYPmJPUQNBZOQYREGGQSYPeJPUQNBZOQYREGGQSYPlJPUQNBZOQYREGGQSYPlJPUQNBZOQYREGGQSYPy", "smelly"},
		{"iWHDAZIAOYUDTHYYCUNBXQnWHDAZIAOYUDTHYYCUNBXQk", "ink"},
		{"mOMTJYOJTLFBKGMYISFQHiOMTJYOJTLFBKGMYISFQHsOMTJYOJTLFBKGMYISFQHcOMTJYOJTLFBKGMYISFQHrOMTJYOJTLFBKGMYISFQHeOMTJYOJTLFBKGMYISFQHaOMTJYOJTLFBKGMYISFQHnOMTJYOJTLFBKGMYISFQHt", "miscreant"},
		{"qFEUYWIKGXCZVXOPZKOBCKHEBuFEUYWIKGXCZVXOPZKOBCKHEBiFEUYWIKGXCZVXOPZKOBCKHEBxFEUYWIKGXCZVXOPZKOBCKHEBoFEUYWIKGXCZVXOPZKOBCKHEBtFEUYWIKGXCZVXOPZKOBCKHEBiFEUYWIKGXCZVXOPZKOBCKHEBc", "quixotic"},
		{"dXKIIKPMULMUIDCSOFTJrXKIIKPMULMUIDCSOFTJaXKIIKPMULMUIDCSOFTJcXKIIKPMULMUIDCSOFTJoXKIIKPMULMUIDCSOFTJnXKIIKPMULMUIDCSOFTJiXKIIKPMULMUIDCSOFTJaXKIIKPMULMUIDCSOFTJn", "draconian"},
		{"cVBMNIAWBKZCBuVBMNIAWBKZCBrVBMNIAWBKZCBiVBMNIAWBKZCBoVBMNIAWBKZCBuVBMNIAWBKZCBs", "curious"},
		{"dWMZKRYZEXCEVEiWMZKRYZEXCEVEsWMZKRYZEXCEVEtWMZKRYZEXCEVEuWMZKRYZEXCEVErWMZKRYZEXCEVEbWMZKRYZEXCEVEeWMZKRYZEXCEVEd", "disturbed"},
		{"lJMDJPLYPPJTAPOSeJMDJPLYPPJTAPOSaJMDJPLYPPJTAPOSn", "lean"},
		{"gDTWSJJAFFHHMNMPXTAWKQOVrDTWSJJAFFHHMNMPXTAWKQOVoDTWSJJAFFHHMNMPXTAWKQOVuDTWSJJAFFHHMNMPXTAWKQOVcDTWSJJAFFHHMNMPXTAWKQOVhDTWSJJAFFHHMNMPXTAWKQOVy", "grouchy"},
		{"aNHHJIPROAMxNHHJIPROAMiNHHJIPROAMoNHHJIPROAMmNHHJIPROAMaNHHJIPROAMtNHHJIPROAMiNHHJIPROAMc", "axiomatic"},
		{"tXBGCUQSBNTSGZMAVNNIYOVVVAZOQKeXBGCUQSBNTSGZMAVNNIYOVVVAZOQKnXBGCUQSBNTSGZMAVNNIYOVVVAZOQKuXBGCUQSBNTSGZMAVNNIYOVVVAZOQKoXBGCUQSBNTSGZMAVNNIYOVVVAZOQKuXBGCUQSBNTSGZMAVNNIYOVVVAZOQKs", "tenuous"},
		{"yVBNHOPAMPHUKGZJFATSHCZAeVBNHOPAMPHUKGZJFATSHCZAaVBNHOPAMPHUKGZJFATSHCZAr", "year"},
		{"gNWUOMXIDOFQLKrNWUOMXIDOFQLKaNWUOMXIDOFQLKb", "grab"},
		{"bTVORYGRQELJJOQKZWIENrTVORYGRQELJJOQKZWIENoTVORYGRQELJJOQKZWIENtTVORYGRQELJJOQKZWIENhTVORYGRQELJJOQKZWIENeTVORYGRQELJJOQKZWIENr", "brother"},
		{"sYZYERJOTTELSPOSAMmYZYERJOTTELSPOSAMeYZYERJOTTELSPOSAMlYZYERJOTTELSPOSAMl", "smell"},
		{"cCPHANPWHKQWLRFDBJOCKTBNUCFXeCPHANPWHKQWLRFDBJOCKTBNUCFXnCPHANPWHKQWLRFDBJOCKTBNUCFXt", "cent"},
		{"rBZFHMFKHMKXEDMSuBZFHMFKHMKXEDMStBZFHMFKHMKXEDMShBZFHMFKHMKXEDMSlBZFHMFKHMKXEDMSeBZFHMFKHMKXEDMSsBZFHMFKHMKXEDMSs", "ruthless"},
		{"pFGEMWBMWIHLPLVJFaFGEMWBMWIHLPLVJFnFGEMWBMWIHLPLVJFiFGEMWBMWIHLPLVJFcFGEMWBMWIHLPLVJFkFGEMWBMWIHLPLVJFy", "panicky"},
		{"tIBIEPZZNNVJWMJNTUKRADYXWXZAeIBIEPZZNNVJWMJNTUKRADYXWXZAdIBIEPZZNNVJWMJNTUKRADYXWXZAiIBIEPZZNNVJWMJNTUKRADYXWXZAoIBIEPZZNNVJWMJNTUKRADYXWXZAuIBIEPZZNNVJWMJNTUKRADYXWXZAs", "tedious"},
	}

	for _, r := range table {
		out := detectWord(r.crowd)
		if out != r.exp {
			t.Errorf("detectWord(%s) was incorrect, expected: %s, got: %s.", r.crowd, r.exp, out)
		}
	}
}
