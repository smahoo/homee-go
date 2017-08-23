package model


type Node struct {
	Id 		int		`json:"id"`
	Name 		string		`json:"name"`
	Profile 	int		`json:"profile"`
	Image 		string		`json:"image"`
	Favorite 	int		`json:"favorite"`
	Order 		int		`json:"order"`
	Protocol 	int		`json:"protocol"`
	Routing 	int		`json:"routing"`
	State 		int		`json:"state"`
	StateChanged 	int		`json:"state_changed"`
	Added 		int		`json:"added"`
	History 	int 		`json:"history"`
	CubeType 	int		`json:"cube_type"`
	Note 		string		`json:"note"`
	Services 	int 		`json:"services"`
	PhoneticName 	string 		`json:"phonetic_name"`
	Owner 		int		`json:"owner"`
	//DeniedUserIds
	Attributes 	[]Attribute 	`json:"attributes"`
}

const  (
	PROFILE_UNKNOWN                               =  -1
	PROFILE_NONE                                  =  0
	PROFILE_HOMEE                                 =  1

	// GENERIC
	PROFILE_ON_OFF_PLUG                            =  10
	PROFILE_DIMMABLE_METER_PLUG                    =  11
	PROFILE_METERING_SWITCH                        =  12
	PROFILE_METERING_PLUG                          =  13
	PROFILE_DIMMABLE_PLUG                          =  14
	PROFILE_DIMMABLE_SWITCH                        =  15
	PROFILE_ON_OFF_SWITCH                          =  16 
	PROFILE_DOUBLE_ON_OFF_SWITCH                   =  18 
	PROFILE_DIMMABLE_METERING_PLUG                 =  19 
	PROFILE_ONE_BUTTON_REMOTE                      =  20 
	PROFILE_BINARY_INPUT                           = 21 
	PROFILE_DIMMABLE_COLOR_METERING_SWITCH         = 22 
	PROFILE_DOUBLE_BINARY_INPUT                    = 23 
	PROFILE_TWO_BUTTON_BINARY_REMOTE               = 24 
	PROFILE_THREE_BUTTON_REMOTE                    = 25 
	PROFILE_FOUR_BUTTON_REMOTE                     = 26 
	PROFILE_ALARM_SENSOR                           = 27 
	PROFILE_DOUBLE_ON_OFF_PLUG                     = 28 
	PROFILE_ON_OFF_SWITCH_WITH_BINARY_INPUT        = 29 

	// Lightning
	PROFILE_SENSOR_BRIGHTNESS                    = 1000 
	PROFILE_DIMMABLE_COLOR_LIGHT                 = 1001 
	PROFILE_DIMMABLE_EXTENDED_COLOR_LIGHT        = 1002 
	PROFILE_DIMMABLE_COLOR_TEMPERATURE_LIGHT     = 1003 
	PROFILE_DIMMABLE_LIGHT                       = 1004 
	PROFILE_DIMMABLE_LIGHT_WITH_BRIGHTNESS_SENSOR   = 1005 
	PROFILE_DIMMABLE_LIGHT_WITH_BRIGHTNESS_AND_PRESENCE_SENSOR  = 1006 
	PROFILE_DIMMABLE_LIGHT_WITH_PRESENCE_SENSOR  = 1007 
	PROFILE_DIMMABLE_RGBW_LIGHT                  = 1008 
	
	// Closures
	PROFILE_OPEN_CLOSE_SENSOR                    = 2000 
	PROFILE_WINDOW_HANDLE                        = 2001 
	PROFILE_SHUTTER_POSITION_SWITCH              = 2002 
	PROFILE_OPEN_CLOSE_AND_TEMPERATURE_SENSOR    = 2003 
	PROFILE_ELECTRIC_MOTOR_METERING_SWITCH       = 2004 
	PROFILE_OPEN_CLOSE_WITH_TEMPERATURE_AND_BRIGHTNESS_SENSOR  = 2005 
	PROFILE_ELECTRIC_MOTOR_METERING_SWITCH_WITHOUT_SLAT_POSITION  = 2006 
	
	// HVAC  = 3000-3999)
	PROFILE_TEMPERATURE_AND_HUMIDITY_SENSOR              = 3001 
	PROFILE_CO2_SENSOR3002                               = 3002 
	PROFILE_ROOM_THERMOSTAT                              = 3003 
	PROFILE_ROOM_THERMOSTAT_WITH_HUMIDITY_SENSOR         = 3004 
	PROFILE_BINARY_INPUT_WITH_TEMPERATURE_SENSOR         = 3005 
	PROFILE_RADIATOR_THERMOSTAT                          = 3006 
	PROFILE_TEMPERATURE_SENSOR                           = 3009 
	PROFILE_HUMIDITY_SENSOR                              = 3010 
	PROFILE_WATER_VALVE                                  = 3011 
	PROFILE_WATER_METER                                  = 3012 
	PROFILE_WEATHER_STATION                              = 3013 
	PROFILE_NETATMO_MAIN_MODULE                          = 3014 
	PROFILE_NETATMO_OUTDOOR_MODULE                       = 3015 
	PROFILE_NETATMO_INDOOR_MODULE                        = 3016 
	PROFILE_NETATMO_RAINMODULE                           = 3017 
	PROFILE_TWO_CHANNEL_COSI_THERM                       = 3018 
	PROFILE_SIX_CHANNEL_COSI_THERM                       = 3019 
	PROFILE_NEST_THERMOSTAT_WITH_COOLING                 = 3020 
	PROFILE_NEST_THERMOSTAT_WITH_HEATING                 = 3021 
	PROFILE_NET_THERMOSTAT_WITH_HEATING_AND_COOLING      = 3022 
	PROFILE_NETATMO_WIND_MODULE                          = 3023 
	PROFILE_ELECTRICAL_HEATING                           = 3024 
	PROFILE_VALVE_DRIVE                                  = 3025 
	PROFILE_NETATMO_CAMERA                               = 3026 
	PROFILE_NETATMO_THERMOSTAT                           = 3027 
	PROFILE_NETATMO_TAGS                                 = 3028 
	
	// Alarm  = 4000-4999)
	PROFILE_MOTION_DETECTOR_WITH_TEMPERATURE_BRIGHTNESS_AND_HUMIDITY_SENSOR  = 4010 
	PROFILE_MOTION_DETECTOR                                                  = 4011 
	PROFILE_SMOKE_DETECTOR                                                   = 4012 
	PROFILE_FLOOD_DETECTOR                                                   = 4013 
	PROFILE_PRESENCE_DETECTOR                                                = 4014 
	PROFILE_MOTION_DETECTOR_WITCH_TEMPERATURE_AND_BRIGHTNESS_SENSOR          = 4015 
	PROFILE_SMOKE_DETECTOR_WITH_TEMPERATURE_SENSOR                           = 4016 
	PROFILE_FLOOD_DETECTOR_WITH_TEMPERATURE_SENSOR                           = 4017 
	PROFILE_WATCH_DOG_DEVICE                                                 = 4018 
	PROFILE_LAG                                                              = 4019 
	PROFILE_OWU                                                              = 4020 
	PROFILE_EUROVAC                                                          = 4021 
	PROFILE_OWWG3                                                            = 4022 
	PROFILE_EUROPRESS                                                        = 4023 
	PROFILE_MINIMUM_DETECTOR                                                 = 4024 
	PROFILE_MAXIMUM_DETECTOR                                                 = 4025 
	PROFILE_SMOKE_DETECTOR_AND_CO_DETECTOR                                   = 4026 
	PROFILE_SIREN                                                            = 4027 
	PROFILE_MOTION_DETECTOR_WITH_OPEN_CLOSE_TEMPERATURE_AND_BRIGHTNESS_SENSOR  = 4028 
	PROFILE_MOTION_DETECTOR_WITH_BRIGHTNESS                                  = 4029 
	PROFILE_DOOR_BELL                                                        = 4030 
	PROFILE_SMOKE_DETECTOR_AND_SIREN                                         = 4031 
	PROFILE_PROFILE_FLOOD_DETECTOR_WITH_TEMPERATURE_AND_HUMIDITY_SENSOR      = 4032 
	
	// Hager  = 5000-5999)
	PROFILE_INOVA_ALARM_SYSTEM                           = 5000 
	PROFILE_INOVA_DETECTOR                               = 5001 
	PROFILE_INOVA_SIREN                                  = 5002 
	PROFILE_INOVA_COMMAND                                = 5003 
	PROFILE_INOVA_TRANSMITTER                            = 5004 
	PROFILE_INOVA_RECEIVER                               = 5005 
	PROFILE_INOVA_KOALE                                  = 5006 
	PROFILE_INOVA_INTERNAL_TRANSMITTER                   = 5007 
	PROFILE_INOVA_CONTROL_PANEL                          = 5008 
	PROFILE_INOVA_INPUT_OUTPUT_EXTENSION                 = 5009 
	PROFILE_INOVA_MOTION_DETECTION_WITH_VOD              = 5010 
	PROFILE_INOVA_MOTION_DETECTOR                        = 5011 
	
	// Plants  = 6000-6999)
	PROFILE_KOUBACHI_PLANT_SENSOR                        = 6000



	// NODE STATES

	NODE_STATE_None                         = 0
	NODE_STATE_AVAILABLE                    = 1
	NODE_STATE_UNAVAILABLE                  = 2
	NODE_STATE_UPDATE_IN_PROGRESS           = 3
	NODE_STATE_WAITING_FOR_ATTRIBUTES       = 4
	NODE_STATE_INITIALIZING                 = 5
	NODE_STATE_USER_INTERACTION_REQUIRED	= 6
	NODE_STATE_PASSWORD_REQUIRED            = 7
	NODE_STATE_HOST_UNAVAILABLE             = 8
	NODE_STATE_DELETE_IN_PROGRESS           = 9
	NODE_STATE_COSI_CONNECTED               = 10
	NODE_STATE_BLOCKED                      = 11
)


func GetNodeProfileName(nodeProfile int) string {
	switch nodeProfile {
	case -1: return "PROFILE_UNKNOWN"
	case 0: return "PROFILE_NONE"
	case 1: return "PROFILE_HOMEE"

	// GENERIC
	case 10: return "PROFILE_ON_OFF_PLUG"
	case 11: return "PROFILE_DIMMABLE_METER_PLUG"
	case 12: return "PROFILE_METERING_SWITCH"
	case 13: return "PROFILE_METERING_PLUG"
	case 14: return "PROFILE_DIMMABLE_PLUG"
	case 15: return "PROFILE_DIMMABLE_SWITCH"
	case 16: return "PROFILE_ON_OFF_SWITCH"
	case 18: return "PROFILE_DOUBLE_ON_OFF_SWITCH"
	case 19: return "PROFILE_DIMMABLE_METERING_PLUG"
	case 20: return "PROFILE_ONE_BUTTON_REMOTE"
	case 21: return "PROFILE_BINARY_INPUT"
	case 22: return "PROFILE_DIMMABLE_COLOR_METERING_SWITCH"
	case 23: return "PROFILE_DOUBLE_BINARY_INPUT"
	case 24: return "PROFILE_TWO_BUTTON_BINARY_REMOTE"
	case 25: return "PROFILE_THREE_BUTTON_REMOTE"
	case 26: return "PROFILE_FOUR_BUTTON_REMOTE"
	case 27: return "PROFILE_ALARM_SENSOR"
	case 28: return "PROFILE_DOUBLE_ON_OFF_PLUG"
	case 29: return "PROFILE_ON_OFF_SWITCH_WITH_BINARY_INPUT"

	// Lightning
	case 1000: return "PROFILE_SENSOR_BRIGHTNESS"
	case 1001: return "PROFILE_DIMMABLE_COLOR_LIGHT"
	case 1002: return "PROFILE_DIMMABLE_EXTENDED_COLOR_LIGHT"
	case 1003: return "PROFILE_DIMMABLE_COLOR_TEMPERATURE_LIGHT"
	case 1004: return "PROFILE_DIMMABLE_LIGHT"
	case 1005: return "PROFILE_DIMMABLE_LIGHT_WITH_BRIGHTNESS_SENSOR"
	case 1006: return "PROFILE_DIMMABLE_LIGHT_WITH_BRIGHTNESS_AND_PRESENCE_SENSOR"
	case 1007: return "PROFILE_DIMMABLE_LIGHT_WITH_PRESENCE_SENSOR"
	case 1008: return "PROFILE_DIMMABLE_RGBW_LIGHT"

	// Closures
	case 2000: return "PROFILE_OPEN_CLOSE_SENSOR"
	case 2001: return "PROFILE_WINDOW_HANDLE"
	case 2002: return "PROFILE_SHUTTER_POSITION_SWITCH"
	case 2003: return "PROFILE_OPEN_CLOSE_AND_TEMPERATURE_SENSOR"
	case 2004: return "PROFILE_ELECTRIC_MOTOR_METERING_SWITCH"
	case 2005: return "PROFILE_OPEN_CLOSE_WITH_TEMPERATURE_AND_BRIGHTNESS_SENSOR"
	case 2006: return "PROFILE_ELECTRIC_MOTOR_METERING_SWITCH_WITHOUT_SLAT_POSITION"

	// HVAC  case 3000-3999)
	case 3001: return "PROFILE_TEMPERATURE_AND_HUMIDITY_SENSOR"
	case 3002: return "PROFILE_CO2_SENSOR3002"
	case 3003: return "PROFILE_ROOM_THERMOSTAT"
	case 3004: return "PROFILE_ROOM_THERMOSTAT_WITH_HUMIDITY_SENSOR"
	case 3005: return "PROFILE_BINARY_INPUT_WITH_TEMPERATURE_SENSOR"
	case 3006: return "PROFILE_RADIATOR_THERMOSTAT"
	case 3009: return "PROFILE_TEMPERATURE_SENSOR"
	case 3010: return "PROFILE_HUMIDITY_SENSOR"
	case 3011: return "PROFILE_WATER_VALVE"
	case 3012: return "PROFILE_WATER_METER"
	case 3013: return "PROFILE_WEATHER_STATION"
	case 3014: return "PROFILE_NETATMO_MAIN_MODULE"
	case 3015: return "PROFILE_NETATMO_OUTDOOR_MODULE"
	case 3016: return "PROFILE_NETATMO_INDOOR_MODULE"
	case 3017: return "PROFILE_NETATMO_RAINMODULE"
	case 3018: return "PROFILE_TWO_CHANNEL_COSI_THERM"
	case 3019: return "PROFILE_SIX_CHANNEL_COSI_THERM"
	case 3020: return "PROFILE_NEST_THERMOSTAT_WITH_COOLING"
	case 3021: return "PROFILE_NEST_THERMOSTAT_WITH_HEATING"
	case 3022: return "PROFILE_NET_THERMOSTAT_WITH_HEATING_AND_COOLING"
	case 3023: return "PROFILE_NETATMO_WIND_MODULE"
	case 3024: return "PROFILE_ELECTRICAL_HEATING"
	case 3025: return "PROFILE_VALVE_DRIVE"
	case 3026: return "PROFILE_NETATMO_CAMERA"
	case 3027: return "PROFILE_NETATMO_THERMOSTAT"
	case 3028: return "PROFILE_NETATMO_TAGS"

	// Alarm  case 4000-4999)
	case 4010: return "PROFILE_MOTION_DETECTOR_WITH_TEMPERATURE_BRIGHTNESS_AND_HUMIDITY_SENSOR"
	case 4011: return "PROFILE_MOTION_DETECTOR"
	case 4012: return "PROFILE_SMOKE_DETECTOR"
	case 4013: return "PROFILE_FLOOD_DETECTOR"
	case 4014: return "PROFILE_PRESENCE_DETECTOR"
	case 4015: return "PROFILE_MOTION_DETECTOR_WITCH_TEMPERATURE_AND_BRIGHTNESS_SENSOR"
	case 4016: return "PROFILE_SMOKE_DETECTOR_WITH_TEMPERATURE_SENSOR"
	case 4017: return "PROFILE_FLOOD_DETECTOR_WITH_TEMPERATURE_SENSOR"
	case 4018: return "PROFILE_WATCH_DOG_DEVICE"
	case 4019: return "PROFILE_LAG"
	case 4020: return "PROFILE_OWU"
	case 4021: return "PROFILE_EUROVAC"
	case 4022: return "PROFILE_OWWG3"
	case 4023: return "PROFILE_EUROPRESS"
	case 4024: return "PROFILE_MINIMUM_DETECTOR"
	case 4025: return "PROFILE_MAXIMUM_DETECTOR"
	case 4026: return "PROFILE_SMOKE_DETECTOR_AND_CO_DETECTOR"
	case 4027: return "PROFILE_SIREN"
	case 4028: return "PROFILE_MOTION_DETECTOR_WITH_OPEN_CLOSE_TEMPERATURE_AND_BRIGHTNESS_SENSOR"
	case 4029: return "PROFILE_MOTION_DETECTOR_WITH_BRIGHTNESS"
	case 4030: return "PROFILE_DOOR_BELL"
	case 4031: return "PROFILE_SMOKE_DETECTOR_AND_SIREN"
	case 4032: return "PROFILE_FLOOD_DETECTOR_WITH_TEMPERATURE_AND_HUMIDITY_SENSOR"

	// Hager  case 5000-5999)
	case 5000: return "PROFILE_INOVA_ALARM_SYSTEM"
	case 5001: return "PROFILE_INOVA_DETECTOR"
	case 5002: return "PROFILE_INOVA_SIREN"
	case 5003: return "PROFILE_INOVA_COMMAND"
	case 5004: return "PROFILE_INOVA_TRANSMITTER"
	case 5005: return "PROFILE_INOVA_RECEIVER"
	case 5006: return "PROFILE_INOVA_KOALE"
	case 5007: return "PROFILE_INOVA_INTERNAL_TRANSMITTER"
	case 5008: return "PROFILE_INOVA_CONTROL_PANEL"
	case 5009: return "PROFILE_INOVA_INPUT_OUTPUT_EXTENSION"
	case 5010: return "PROFILE_INOVA_MOTION_DETECTION_WITH_VOD"
	case 5011: return "PROFILE_INOVA_MOTION_DETECTOR"

	// Plants  case 6000-6999)
	case 6000: return "PROFILE_KOUBACHI_PLANT_SENSOR"

	}
	return "UNDEFINED"
}