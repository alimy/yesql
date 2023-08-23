package yesql

import (
	"os"

	"gopkg.in/yaml.v3"
)

type yesqlConf struct {
	Version   string `yaml:"version"`
	Generator struct {
		Engine            string `yaml:"engine"`
		SqlxPkgName       string `yaml:"sqlx_package"`
		DefaultStructName string `yaml:"default_struct_name"`
		GoFileName        string `yaml:"go_file_name"`
	}
	Sql []struct {
		Queries string `yaml:"queries"`
		Gen     struct {
			Package           string `yaml:"package"`
			Out               string `yaml:"out"`
			Engine            string `yaml:"engine"`
			SqlxPkgName       string `yaml:"sqlx_package"`
			DefaultStructName string `yaml:"default_struct_name"`
			GoFileName        string `yaml:"go_file_name"`
		}
	}
}

func (s *yesqlConf) prepare() {
	for i := range s.Sql {
		if s.Sql[i].Gen.Engine == "" {
			s.Sql[i].Gen.Engine = s.Generator.Engine
		}
		if s.Sql[i].Gen.SqlxPkgName == "" {
			s.Sql[i].Gen.SqlxPkgName = s.Generator.SqlxPkgName
		}
		if s.Sql[i].Gen.DefaultStructName == "" {
			s.Sql[i].Gen.DefaultStructName = s.Generator.DefaultStructName
		}
		if s.Sql[i].Gen.GoFileName == "" {
			s.Sql[i].Gen.GoFileName = s.Generator.GoFileName
		}
	}
}

func yesqlConfFrom(path string) (*yesqlConf, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	res := &yesqlConf{}
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(res)
	res.prepare()
	return res, err
}
