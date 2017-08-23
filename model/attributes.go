package model




type Attribute struct {
	Id 		int		`json:"id"`
	NodeId		int		`json:"node_id"`
	Instance	int		`json:"instance"`
	Minimum		int		`json:"minimum"`
	Maximum		int		`json:"maximum"`
	CurrentValue	float32		`json:"current_value"`
	TargetValue	float32		`json:"target_value"`
	LastValue	float32		`json:"last_value"`
	Unit 		string		`json:"unit"`
	StepValue	float32		`json:"step_value"`
	Editable	int		`json:"editable"`
	Type 		int		`json:"type"`
	State 		int		`json:"state"`
	LastChanged	int		`json:"last_changed"`
	ChangedBy	int		`json:"changed_by"`
	BasedOn		int		`json:"based_on"`
	Data 		string		`json:"data"`
	//options []
}

const (

	ATTRIBUTE_TYPE_UNKNOWN                                   = -1 
	ATTRIBUTE_TYPE_NONE                                       = 0 
	ATTRIBUTE_TYPE_ON_OFF                                     = 1 
	ATTRIBUTE_TYPE_DIMMING_LEVEL                              = 2 
	ATTRIBUTE_TYPE_CURRENT_ENERGY_USE                         = 3 
	ATTRIBUTE_TYPE_ACCUMULATED_ENERGY_USE                     = 4 
	ATTRIBUTE_TYPE_TEMPERATURE                                = 5 
	ATTRIBUTE_TYPE_TARGET_TEMPERATURE                         = 6 
	ATTRIBUTE_TYPE_RELATIVE_HUMIDITY                          = 7 
	ATTRIBUTE_TYPE_BATTERY_LEVEL                              = 8 
	ATTRIBUTE_TYPE_STATUS_LED                                 = 9 
	ATTRIBUTE_TYPE_WINDOW_POSITION                           = 10 
	ATTRIBUTE_TYPE_BRIGHTNESS                                = 11 
	ATTRIBUTE_TYPE_FLOOD_ALARM                               = 12 
	ATTRIBUTE_TYPE_SIREN                                     = 13 
	ATTRIBUTE_TYPE_OPENCLOSE                                 = 14 
	ATTRIBUTE_TYPE_POSITION                                  = 15 
	ATTRIBUTE_TYPE_SMOKE_ALARM                               = 16 
	ATTRIBUTE_TYPE_BLACKOUT_ALARM                            = 17 
	ATTRIBUTE_TYPE_CURRENT_VALVE_POSITION                    = 18 
	ATTRIBUTE_TYPE_BINARY_INPUT                              = 19 
	ATTRIBUTE_TYPE_CO2_LEVEL                                 = 20 
	ATTRIBUTE_TYPE_UPDATE                                    = 21 
	ATTRIBUTE_TYPE_BATTERY_STATE                             = 22 
	ATTRIBUTE_TYPE_COLOR                                     = 23 
	ATTRIBUTE_TYPE_SATURATION                                = 24 
	ATTRIBUTE_TYPE_MOTION_ALARM                              = 25 
	ATTRIBUTE_TYPE_MOTION_SENSITIVITY                        = 26 
	ATTRIBUTE_TYPE_MOTION_INSENSITIVITY                      = 27 
	ATTRIBUTE_TYPE_MOTION_CANCELATION_DELAY                  = 28 
	ATTRIBUTE_TYPE_WAKE_UP_INTERVAL                          = 29 
	ATTRIBUTE_TYPE_TAMPER_ALARM                              = 30 
	ATTRIBUTE_TYPE_ACOUSTIC_SIGNAL                           = 31 
	ATTRIBUTE_TYPE_BINARY_OUTPUT                             = 32 
	ATTRIBUTE_TYPE_LINK_QUALITY                              = 33 
	ATTRIBUTE_TYPE_INNOVA_ALARM_SYSTEM_STATE                 = 34 
	ATTRIBUTE_TYPE_INNOVA_ALARM_GROUP_STATE                  = 35 
	ATTRIBUTE_TYPE_INNOVA_ALARM_INTRUSION_STATE              = 36 
	ATTRIBUTE_TYPE_INNOVA_ALARM_ERROR_STATE                  = 37 
	ATTRIBUTE_TYPE_INNOVA_ALARM_DOOR_STATE                   = 38 
	ATTRIBUTE_TYPE_INNOVA_ALARM_EXTERNAL_SENSOR              = 39 
	ATTRIBUTE_TYPE_BUTTON_STATE                              = 40 
	ATTRIBUTE_TYPE_HUE                                       = 41 
	ATTRIBUTE_TYPE_COLOR_TEMPERATURE                         = 42 
	ATTRIBUTE_TYPE_HARDWARE_REVISION                         = 43 
	ATTRIBUTE_TYPE_FIRMWARE_REVISION                         = 44 
	ATTRIBUTE_TYPE_SOFTWARE_REVISION                         = 45 
	ATTRIBUTE_TYPE_LED_STATE                                 = 46 
	ATTRIBUTE_TYPE_LED_STATE_WHEN_ON                         = 47 
	ATTRIBUTE_TYPE_LED_STATE_WHEN_OFF                        = 48 
	ATTRIBUTE_TYPE_OPEN_CASING_ALARM                         = 49 
	ATTRIBUTE_TYPE_ACCOUSTIC_AND_VISUAL_SIGNALS              = 50 
	ATTRIBUTE_TYPE_TEMPERATURE_MEASUREMENT_INTERVAL          = 51 
	ATTRIBUTE_TYPE_HIGH_TEMPERATURE_ALARM                    = 52 
	ATTRIBUTE_TYPE_HIGH_TEMPERATURE_ALARM_THRESHOLD          = 53 
	ATTRIBUTE_TYPE_LOW_TEMPERATURE_ALARM                     = 54 
	ATTRIBUTE_TYPE_LOW_TEMPERATURE_ALARM_THRESHOLD           = 55 
	ATTRIBUTE_TYPE_TAMPER_SENSITIVITY                        = 56 
	ATTRIBUTE_TYPE_TAMPER_ALARM_CANCELATION_DELAY            = 57 
	ATTRIBUTE_TYPE_BRIGHTNESS_REPORT_INTERVAL                = 58 
	ATTRIBUTE_TYPE_TEMPERATURE_REPORT_INTERVAL               = 59

	// TODO: following attribute types should be in capital letters

	ATTRIBUTE_TYPE_MotionAlarmIndicationMode        = 60
	ATTRIBUTE_TYPE_LEDBrightness                    = 61
	ATTRIBUTE_TYPE_TamperAlarmIndicationMode        = 62
 	ATTRIBUTE_TYPE_SwitchType                       = 63
	ATTRIBUTE_TYPE_TemperatureOffset                = 64
	ATTRIBUTE_TYPE_AccumulatedWaterUse              = 65
	ATTRIBUTE_TYPE_AccumulatedWaterUseLastMonth     = 66
	ATTRIBUTE_TYPE_CurrentDate                      = 67
	ATTRIBUTE_TYPE_LeakAlarm                        = 68
	ATTRIBUTE_TYPE_BatteryLowAlarm                  = 69
	ATTRIBUTE_TYPE_MalfunctionAlarm                 = 70
	ATTRIBUTE_TYPE_LinkQualityAlarm                 = 71
	ATTRIBUTE_TYPE_Mode                             = 72
	ATTRIBUTE_TYPE_CurrentState                     = 73
	ATTRIBUTE_TYPE_TargetState                      = 74
	ATTRIBUTE_TYPE_Calibration                      = 75
	ATTRIBUTE_TYPE_PresenceAlarm                    = 76
	ATTRIBUTE_TYPE_MinimumAlarm                     = 77
	ATTRIBUTE_TYPE_MaximumAlarm                     = 78
	ATTRIBUTE_TYPE_OilAlarm                         = 79
	ATTRIBUTE_TYPE_WaterAlarm                       = 80
	ATTRIBUTE_TYPE_InovaAlarmInhibition             = 81
	ATTRIBUTE_TYPE_InovaAlarmEjection               = 82
	ATTRIBUTE_TYPE_InovaAlarmCommercialRef          = 83
	ATTRIBUTE_TYPE_InovaAlarmSerialNumber           = 84
	ATTRIBUTE_TYPE_RadiatorThermostatSummerMode     = 85
	ATTRIBUTE_TYPE_InovaAlarmOperationMode          = 86
	ATTRIBUTE_TYPE_AutomaticMode                    = 87
	ATTRIBUTE_TYPE_PollingInterval                  = 88
	ATTRIBUTE_TYPE_FeedTemperature                  = 89
	ATTRIBUTE_TYPE_DisplayOrientation               = 90
	ATTRIBUTE_TYPE_ManualOperation                  = 91
	ATTRIBUTE_TYPE_DeviceTemperature                = 92
	ATTRIBUTE_TYPE_Sonometer                        = 93
	ATTRIBUTE_TYPE_AirPressure                      = 94
	ATTRIBUTE_TYPE_OutdoorRelativeHumidity          = 95
	ATTRIBUTE_TYPE_IndoorRelativeHumidity           = 96
	ATTRIBUTE_TYPE_OutdoorTemperature               = 97
	ATTRIBUTE_TYPE_IndoorTemperature                = 98
	ATTRIBUTE_TYPE_InovaAlarmAntimask               = 99
	ATTRIBUTE_TYPE_InovaAlarmBackupSupply           = 100
	ATTRIBUTE_TYPE_RainFall                         = 101
	ATTRIBUTE_TYPE_LastUpdateOnServer               = 102
	ATTRIBUTE_TYPE_InovaAlarmGeneralHomeCommand     = 103
	ATTRIBUTE_TYPE_InovaAlarmAlert                  = 104
	ATTRIBUTE_TYPE_InovaAlarmSilentAlert            = 105
	ATTRIBUTE_TYPE_InovaAlarmPreAlarm               = 106
	ATTRIBUTE_TYPE_InovaAlarmDeterrenceAlarm        = 107
	ATTRIBUTE_TYPE_InovaAlarmWarning                = 108
	ATTRIBUTE_TYPE_InovaAlarmFireAlarm              = 109
	ATTRIBUTE_TYPE_UpTime                           = 110
	ATTRIBUTE_TYPE_DownTime                         = 111
	ATTRIBUTE_TYPE_ShutterBlindMode                 = 112
	ATTRIBUTE_TYPE_ShutterSlatPosition              = 113
	ATTRIBUTE_TYPE_ShutterSlatTime                  = 114
	ATTRIBUTE_TYPE_RestartDevice                    = 115
	ATTRIBUTE_TYPE_SoilMoisture                     = 116
	ATTRIBUTE_TYPE_WaterPlantAlarm                  = 117
	ATTRIBUTE_TYPE_MistPlantAlarm                   = 118
	ATTRIBUTE_TYPE_FertilizePlantAlarm              = 119
	ATTRIBUTE_TYPE_CoolPlantAlarm                   = 120
	ATTRIBUTE_TYPE_HeatPlantAlarm                   = 121
	ATTRIBUTE_TYPE_PutPlantIntoLightAlarm           = 122
	ATTRIBUTE_TYPE_PutPlantIntoShadeAlarm           = 123
	ATTRIBUTE_TYPE_ColorMode                        = 124
	ATTRIBUTE_TYPE_TargetTemperatureLow             = 125
	ATTRIBUTE_TYPE_TargetTemperatureHigh            = 126
	ATTRIBUTE_TYPE_HVACMode                         = 127
	ATTRIBUTE_TYPE_Away                             = 128
	ATTRIBUTE_TYPE_HVACState                        = 129
	ATTRIBUTE_TYPE_HasLeaf                          = 130
	ATTRIBUTE_TYPE_SetEnergyConsumption             = 131
	ATTRIBUTE_TYPE_COAlarm                          = 132
	ATTRIBUTE_TYPE_RestoreLastKnownState            = 133
	ATTRIBUTE_TYPE_LastImageReceived                = 134
	ATTRIBUTE_TYPE_UpDown                           = 135
	ATTRIBUTE_TYPE_RequestVOD                       = 136
	ATTRIBUTE_TYPE_InovaDetectorHistory             = 137
	ATTRIBUTE_TYPE_SurgeAlarm                       = 138
	ATTRIBUTE_TYPE_LoadAlarm                        = 139
	ATTRIBUTE_TYPE_OverloadAlarm                    = 140
	ATTRIBUTE_TYPE_VoltageDropAlarm                 = 141
	ATTRIBUTE_TYPE_ShutterOrientation               = 142
	ATTRIBUTE_TYPE_OverCurrentAlarm                 = 143
	ATTRIBUTE_TYPE_SirenMode                        = 144
	ATTRIBUTE_TYPE_AlarmAutoStopTime                = 145
	ATTRIBUTE_TYPE_WindSpeed                        = 146
	ATTRIBUTE_TYPE_WindDirection                    = 147
	ATTRIBUTE_TYPE_ComfortTemperature               = 148
	ATTRIBUTE_TYPE_EcoTemperature                   = 149
	ATTRIBUTE_TYPE_ReduceTemperature                = 150
	ATTRIBUTE_TYPE_ProtectTemperature               = 151
	ATTRIBUTE_TYPE_InovaSystemTime                  = 152
	ATTRIBUTE_TYPE_InovaCorrespondentProtocol       = 153
	ATTRIBUTE_TYPE_InovaCorrespondentID             = 154
	ATTRIBUTE_TYPE_InovaCorrespondentListen         = 155
	ATTRIBUTE_TYPE_InovaCorrespondentNumber         = 156
	ATTRIBUTE_TYPE_InovaCallCycleFireProtection     = 157
	ATTRIBUTE_TYPE_InovaCallCycleIntrusion          = 158
	ATTRIBUTE_TYPE_InovaCallCycleTechnicalProtect   = 159
	ATTRIBUTE_TYPE_InovaCallCycleFaults             = 160
	ATTRIBUTE_TYPE_InovaCallCycleDeterrence         = 161
	ATTRIBUTE_TYPE_InovaCallCyclePrealarm           = 162
	ATTRIBUTE_TYPE_InovaPSTNRings                   = 163
	ATTRIBUTE_TYPE_InovaDoubleCallRings             = 164
	ATTRIBUTE_TYPE_InovaPIN                         = 165
	ATTRIBUTE_TYPE_InovaPUK                         = 166
	ATTRIBUTE_TYPE_InovaMainMediaSelection          = 167
	ATTRIBUTE_TYPE_RainFallLastHour                 = 168
	ATTRIBUTE_TYPE_RainFallToday                    = 169
	ATTRIBUTE_TYPE_IdentificationMode               = 170
	ATTRIBUTE_TYPE_ButtonDoubleClick                = 171
	ATTRIBUTE_TYPE_SirenTriggerMode                 = 172
	ATTRIBUTE_TYPE_UV                               = 173
	ATTRIBUTE_TYPE_SlatSteps                        = 174
	ATTRIBUTE_TYPE_EcoModeConfig                    = 175
	ATTRIBUTE_TYPE_ButtonLongRelease                = 176
	ATTRIBUTE_TYPE_VisualGong                       = 177
	ATTRIBUTE_TYPE_AcousticGong                     = 178
	ATTRIBUTE_TYPE_SurveillanceOnOff                = 179
	ATTRIBUTE_TYPE_NetatmoCameraURL                 = 180
	ATTRIBUTE_TYPE_SdState                          = 181
	ATTRIBUTE_TYPE_AdapterState                     = 182
	ATTRIBUTE_TYPE_NetatmoHome                      = 183
	ATTRIBUTE_TYPE_NetatmoPerson                    = 184
	ATTRIBUTE_TYPE_NetatmoLastEventPersonId         = 185
	ATTRIBUTE_TYPE_NetatmoLastEventTime             = 186
	ATTRIBUTE_TYPE_NetatmoLastEventType             = 187
	ATTRIBUTE_TYPE_NetatmoLastEventIsKnownPerson    = 188
	ATTRIBUTE_TYPE_NetatmoLastEventIsArrival        = 189
	ATTRIBUTE_TYPE_PresenceTimeout                  = 190
	ATTRIBUTE_TYPE_KnownPersonPresence              = 191
	ATTRIBUTE_TYPE_UnknownPersonPresence            = 192
	ATTRIBUTE_TYPE_Current                          = 193
	ATTRIBUTE_TYPE_Frequency                        = 194
	ATTRIBUTE_TYPE_Voltage                          = 195
	ATTRIBUTE_TYPE_PresenceAlarmCancelationDelay    = 196
	ATTRIBUTE_TYPE_PresenceAlarmDetectionDelay      = 197
	ATTRIBUTE_TYPE_PresenceAlarmThreshold           = 198
	ATTRIBUTE_TYPE_NetatmoThermostatMode            = 199
	ATTRIBUTE_TYPE_NetatmoRelayBoilerConnected      = 200
	ATTRIBUTE_TYPE_NetatmoRelayMac                  = 201
	ATTRIBUTE_TYPE_NetatmoThermostatModeTimeout     = 202
	ATTRIBUTE_TYPE_NetatmoThermostatNextChange      = 203
	ATTRIBUTE_TYPE_NetatmoThermostatPrograms        = 204
	ATTRIBUTE_TYPE_HomeeMode                        = 205
	ATTRIBUTE_TYPE_ColorWhite                       = 206
	ATTRIBUTE_TYPE_MovementAlarm                    = 207
	ATTRIBUTE_TYPE_MovementSensitivity              = 208
	ATTRIBUTE_TYPE_VibrationAlarm                   = 209
	ATTRIBUTE_TYPE_VibrationSensitivity             = 210
	ATTRIBUTE_TYPE_AverageEnergyUse                 = 211
	ATTRIBUTE_TYPE_BinaryInputMode                  = 212

)


