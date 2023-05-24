package config

/*
TransferRestart：布尔值，表示是否在传输数据时重启。
Root：字符串，表示应用程序的根目录路径。
Server：字符串，表示应用程序使用的服务器类型。
SApi：字符串，表示服务器API的地址。
SPlug：字符串，表示服务器插件的地址。
SInitialize：字符串，表示服务器初始化的地址。
SModel：字符串，表示服务器模型的地址。
SRequest：字符串，表示服务器请求的地址。
SRouter：字符串，表示服务器路由的地址。
SService：字符串，表示服务器服务的地址。
Web：字符串，表示应用程序使用的Web框架类型。
WApi：字符串，表示Web框架API的地址。
WForm：字符串，表示Web框架表单的地址。
WTable：字符串，表示Web框架表格的地址。
*/

type Autocode struct {
	TransferRestart bool   `mapstructure:"transfer-restart" json:"transfer-restart" yaml:"transfer-restart"`
	Root            string `mapstructure:"root" json:"root" yaml:"root"`
	Server          string `mapstructure:"server" json:"server" yaml:"server"`
	SApi            string `mapstructure:"server-api" json:"server-api" yaml:"server-api"`
	SPlug           string `mapstructure:"server-plug" json:"server-plug" yaml:"server-plug"`
	SInitialize     string `mapstructure:"server-initialize" json:"server-initialize" yaml:"server-initialize"`
	SModel          string `mapstructure:"server-model" json:"server-model" yaml:"server-model"`
	SRequest        string `mapstructure:"server-request" json:"server-request"  yaml:"server-request"`
	SRouter         string `mapstructure:"server-router" json:"server-router" yaml:"server-router"`
	SService        string `mapstructure:"server-service" json:"server-service" yaml:"server-service"`
	Web             string `mapstructure:"web" json:"web" yaml:"web"`
	WApi            string `mapstructure:"web-api" json:"web-api" yaml:"web-api"`
	WForm           string `mapstructure:"web-form" json:"web-form" yaml:"web-form"`
	WTable          string `mapstructure:"web-table" json:"web-table" yaml:"web-table"`
}
