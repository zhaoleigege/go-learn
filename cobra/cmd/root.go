package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// 执行根命令 go run main.go
// 指定flag的值 go run main.go --license=test
// 或者 go run main.go --license test
var (
	cfgFile     string
	UserLicense string
	rootCmd     = &cobra.Command{
		Use:   "cobra",
		Short: "学习使用cobra",
		Long:  `学习使用cobra进行开发`,
		//Args: cobra.MinimumNArgs(1),
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("root 执行PersistentPreRunE")
			return nil
		},
		PreRun: func(cmd *cobra.Command, args []string) {
			fmt.Println("root 执行preRun")
		},
		// Run方法存在才会执行该命令下的应用
		Run: func(cmd *cobra.Command, args []string) {
			//cmd.Help()
			fmt.Println("root 执行run")
			fmt.Println("args: ", args)
		},
		PostRun: func(cmd *cobra.Command, args []string) {
			fmt.Println("root 执行postRun")
		},
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("root 执行PersistentPostRunE")
			return nil
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "配置文件的地址")
	rootCmd.PersistentFlags().StringVar(&UserLicense, "license", "注册文件", "注册文件的信息")
}

type Config struct {
	Name    string `mapstruct:"name"`
	Version string `mapstruct:"version"`
	Count   int64  `mapstruct:"count"`
}

var conf = new(Config)

// viper的使用 https://www.liwenzhou.com/posts/Go/viper_tutorial/
func initConfig() {
	fmt.Printf("license: %s\n", UserLicense)

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigFile("conf.yaml")
	}

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// 保存配置信息的值到结构体中
	if err := viper.Unmarshal(conf); err != nil {
		panic(err)
	}

	fmt.Println(conf)
}
