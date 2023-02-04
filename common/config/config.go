package config

import "github.com/liuhongdi/digv04/pkg/setting"

func ReadConfig(configName string, val interface{}) error {
	s, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = s.ReadSection(configName, val)
	if err != nil {
		return err
	}
	return nil
}
