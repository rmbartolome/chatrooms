package config

import "gitlab.falabella.com/fif/integracion/forthehorde/commons/naga"

// configEntries retorna configuraciones iniciales (env y flags)
func configEntries() []naga.ConfigEntry {
	//TODO: Agregar entradas de configuracion adicionales requeridas para el servicio
	return []naga.ConfigEntry{
		{
			VariableName: "port",
			Description:  "Puerto a utilizar",
			Shortcut:     "p",
			DefaultValue: ":8080",
		},
		{
			VariableName: "logging_level",
			Description:  "Level de detalle de logs",
			Shortcut:     "l",
			DefaultValue: "info",
		},
		{
			VariableName: "tracing_enabled",
			Description:  "Especifica si se debe configurar tracing",
			Shortcut:     "t",
			DefaultValue: false,
		},
		{
			VariableName: "metrics_enabled",
			Description:  "Especifica si se debe configurar metrics",
			Shortcut:     "m",
			DefaultValue: false,
		},
		{
			VariableName: "timeout",
			Description:  "timeout de configuración",
			DefaultValue: 2,
		},
		{
			VariableName: "url",
			Description:  "api a consumir",
			DefaultValue: "http://rickandmortyapi.com/api/",
		},
	}
}

// ReadConfiguration resuelve los valores de los flags o envs utilizando Naga
func ReadConfiguration() map[string]interface{} {
	typeResolver := naga.NewVariableTypeResolver()
	flagConfigurator := naga.NewFlagConfigurator(typeResolver)
	entries := configEntries()

	configurator := naga.NewConfigurator(flagConfigurator, typeResolver)

	config, err := configurator.Configure("", entries)

	if err != nil {
		panic("Error de configuracion")
	}

	return config
}
