package cmd

import (
	"log"
	"strings"

	internal "github.com/Garfield247/go_tour/internal/word"
	"github.com/spf13/cobra"
)

const (
	ModeUpper                     = iota + 1 // 转大写
	ModeLower                                // 转小写
	ModeUnderscoreToUpperCameCase            // 转大驼峰
	ModeUnderscoreToLowerCameCase            // 转小驼峰
	ModeCameCaseToUnderscore                 // 转下划线
)

var str string
var mode int8
var desc = strings.Join([]string{
	"该子命令支持多种单词格式转换;转换模式如下:",
	"1: 全部转大写",
	"2: 全部转小写",
	"3: 下划线转大驼峰",
	"4: 下划线转小驼峰",
	"5: 驼峰转下划线",
}, "\n")

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = internal.ToUpper(str)
		case ModeLower:
			content = internal.ToLower(str)
		case ModeUnderscoreToUpperCameCase:
			content = internal.UndersocreToUpperCameCase(str)
		case ModeUnderscoreToLowerCameCase:
			content = internal.UndersocreToLowerCameCase(str)
		case ModeCameCaseToUnderscore:
			content = internal.CameCaseToUndersocre(str)
		default:
			log.Fatalln("参不支持的格式转换,请执行help word 查看帮助文档")
		}

		log.Printf("转换结果: %s", content)
	},
}

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入转换模式")
}
