package model

type HomeeMessage struct {
	Nodes 		*[]Node 	`json:"nodes"`
	Node		*Node		`json:"node"`
	Attribute 	*Attribute 	`json:"attribute"`
	Code 		int		`json:"code"`
	Error		*ErrorMessage	`json:"error"`
	Warning		*WarningMessage `json:"warning"`
}

type ErrorMessage struct {
	Code 		int		`json:"code"`
	Message 	string 		`json:"message"`
	Description	string		`json:"description"`
	Data 		*interface{}	`json:"data"`
}


type WarningMessage struct {
	Code 		int		`json:"code"`
	Message		string		`json:"message"`
	Description 	string		`json:"description"`
	Data 		*interface{}	`json:"data"`
}

const (
	 ERROR_CODE_MISSING_PARAMETER		 = 1
	 ERROR_CODE_INVALID_FORMAT		 = 2
	 ERROR_CODE_INVALID_PATH		 = 3
	 ERROR_CODE_INVALID_VERSION		 = 4
	 ERROR_CODE_OBJECT_NOT_FOUND		 = 5
	 ERROR_CODE_INVALID_RESOURCE_ID		 = 6
	 ERROR_CODE_INTERNAL_SERVICE_EXCEPTION	 = 7
	 ERROR_CODE_TOO_MANY_PARAMETERS		 = 8
	 ERROR_CODE_INVALID_PARAMETERS		 = 9
	 ERROR_CODE_CONFLICTIING_PARAMETERS	 = 10


	//Cube Warnings 
	 WARNING_CODE_CubeAdded				= 100
	 WARNING_CODE_CubeRemoved 			= 101
	 WARNING_CODE_CubeIsMissing 			= 102
	 WARNING_CODE_CubeInLearnMode 			= 103
	 WARNING_CODE_CubeLearnModeStarted 		= 104
	 WARNING_CODE_CubeLearnModeSuccessful 		= 105
	 WARNING_CODE_CubeLearnModeTimeout 		= 106
	 WARNING_CODE_CubeLearnModeNodeAlreadyAdded 	= 107
	 WARNING_CODE_CubeLearnModeFailed 		= 108
	 WARNING_CODE_CubeInRemoveMode 			= 109
	 WARNING_CODE_CubeRemoveModeStarted 		= 110
	 WARNING_CODE_CubeRemoveModeSuccessful 		= 111
	 WARNING_CODE_CubeRemoveModeTimeout 		= 112
	 WARNING_CODE_CubeRemoveModeNodeAlreadyDeleted 	= 113
	 WARNING_CODE_CubeRemoveModeFailed 		= 114
	 WARNING_CODE_CubeScannedNodes 			= 115
	 WARNING_CODE_CubeUpdateInProgess 		= 116
	 WARNING_CODE_CubeUpdateStarted 		= 117
	 WARNING_CODE_CubeUpdateEnded 			= 118
	 WARNING_CODE_CubeUpdateFailed 			= 119
	 WARNING_CODE_CubeUserInteractionRequired 	= 120
	 WARNING_CODE_CubeRemoveModeCanceled 		= 121
	 WARNING_CODE_CubeLearnModeCanceled 		= 122
	 WARNING_CODE_CubeAuthorizationRequired 	= 123
	 WARNING_CODE_CubeLearnModeInitializing 	= 124
	 WARNING_CODE_CubeRemoveModeForbidden 		= 125
	
	// Node
	 WARNING_CODE_NodeBadLinkQuality                 = 200 
	 WARNING_CODE_NodeIsMissing                      = 201 
	 WARNING_CODE_NodeWaterDetected                  = 202 
	 WARNING_CODE_NodeSmokeDetected                  = 203 
	 WARNING_CODE_NodeBatteryLow                     = 204 
	 WARNING_CODE_NodeLocked                         = 205 
	 WARNING_CODE_NodeNotCompatible                  = 206 
	 WARNING_CODE_NodeResetSuccessful                = 207 
	 WARNING_CODE_NodeResetStarted                   = 208 
	 WARNING_CODE_NodeResetFailed                    = 209 
	 WARNING_CODE_NodeResetTimeout                   = 210 
	 WARNING_CODE_NodeWrongHVACMode                  = 211 
	 WARNING_CODE_NodeRangeError                     = 212 
	 WARNING_CODE_NodeBlocked                        = 213 
	 WARNING_CODE_NodeWrongAwayMode                  = 214 
	 WARNING_CODE_NodeResetCanceled                  = 215 
	
	// Settings
	 WARNING_CODE_SettingRemoteAccessActivated       = 300 
	 WARNING_CODE_SettingRemoteAccessDeactivated     = 301 
	 WARNING_CODE_SettingOnline                      = 302 
	 WARNING_CODE_SettingOffline                     = 303 
	 WARNING_CODE_SettingNetworkUninitialized        = 304 
	 WARNING_CODE_SettingNetworkInitializing         = 305 
	 WARNING_CODE_SettingNetworkInitialized          = 306 
	
	// Update
	 WARNING_CODE_UpdateAvailable                    = 400 
	 WARNING_CODE_UpdateDownloading                  = 401 
	 WARNING_CODE_UpdateInstalling                   = 402 
	 WARNING_CODE_UpdateInProgress                   = 403 
	 WARNING_CODE_UpdateSuccessful                   = 404 
	 WARNING_CODE_UpdateConnectionFailed             = 405 
	 WARNING_CODE_UpdateDownloadFailed               = 406 
	 WARNING_CODE_UpdateInstallationFailed           = 407 
	 WARNING_CODE_UpdatePreparing                    = 408 
	 WARNING_CODE_NoUpdateAvailable                  = 409 
	 WARNING_CODE_USBUpdateAvailable                 = 410 
	 WARNING_CODE_NoUSBUpdateAvailable               = 411 
	 WARNING_CODE_USBUpdateCanceled                  = 412 
	
	 WARNING_CODE_BackupAvailable                    = 450 
	 WARNING_CODE_BackupCreationFailed               = 451 
	 WARNING_CODE_RestoreAvailable                   = 452 
	 WARNING_CODE_NoRestoreAvailable                 = 453 
	 WARNING_CODE_RestoreFailed                      = 454 
	 WARNING_CODE_RestoreRestarting                  = 455 
	 WARNING_CODE_RestoreCanceled                    = 456 
	
	// Access
	 WARNING_CODE_PermissionDenied                   = 500 
	 WARNING_CODE_TeachInForbidden                   = 501 
	 WARNING_CODE_PermissionGranted                  = 502 
	 WARNING_CODE_VideocodeWrong                     = 503 
	 WARNING_CODE_PasscodeWrong                      = 504 
	 WARNING_CODE_ChangeVideocodeSuccessful          = 505 
	 WARNING_CODE_ChangePasscodeSuccessful           = 506 
	 WARNING_CODE_ChangeVideocodeFailed              = 507 
	 WARNING_CODE_ChangePasscodeFailed               = 508 
	
	// Device
	 WARNING_CODE_DeviceRemoved                      = 600 
	 WARNING_CODE_DeviceAdded                        = 601 
	
	// User
	 WARNING_CODE_UserRemoved                        = 700 
	 WARNING_CODE_AllUsersRemoved                    = 701 
	 WARNING_CODE_UserPasswordChangeRequired         = 702 
	 WARNING_CODE_UserPasswordChangeFailed           = 703 
	 WARNING_CODE_UserPasswordChangeSuccessful       = 704 
	 WARNING_CODE_UserAlreadyExists                  = 705
	
	// Inova
	 WARNING_CODE_InovaIntrusionDetected             = 800 
	 WARNING_CODE_InovaError                         = 801 
	 WARNING_CODE_InovaDetectorEjected               = 802 
	 WARNING_CODE_InovaDoorOpen                      = 803 
	 WARNING_CODE_InovaWrongOperationMode            = 804 
	 WARNING_CODE_InovaCommandTimeout                = 805 
	 WARNING_CODE_InovaCommandForbidden              = 806 
	 WARNING_CODE_InovaArmingBlocked                 = 807 
	 WARNING_CODE_InovaDownloadDataZoneFailed        = 808 
	 WARNING_CODE_InovaUploadDataZoneFailed          = 809 
	
	// Homeegrams
	 WARNING_CODE_HomeegramCouldNotDownloadTTS       = 900 
	 WARNING_CODE_HomeegramTestResults               = 901	
)

