package config

/**
 * @Description
 * @Author enzuo
 * @Date 2023/2/5
 **/
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
