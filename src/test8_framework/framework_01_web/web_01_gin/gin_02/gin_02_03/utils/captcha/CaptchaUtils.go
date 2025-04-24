package captcha

import (
	"github.com/mojocn/base64Captcha"
)

const (
	source      = "ABCDEFGHJKLMNPQRSTUVWXYZ235689"
	height      = 80
	width       = 240
	noise       = 0
	length      = 6
	dotCount    = 6
	audioLength = 3
)

// `go get github.com/mojocn/base64Captcha`

// GenerateCharCaptcha 仅字母验证码
func GenerateLetterCaptcha(length int) (captchaId, answer, base64Img string) {

	driver := base64Captcha.NewDriverString(
		height, // height
		width,  // width
		noise,  // noise
		base64Captcha.OptionShowHollowLine|base64Captcha.OptionShowSlimeLine,
		length,        // length
		source,        // source
		nil, nil, nil, // custom fonts/colors
	)
	c := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)

	captchaId, base64Img, answer, err := c.Generate()
	if err != nil {
		panic("生成仅字母验证码出错：" + err.Error())
	}

	return captchaId, answer, base64Img
}
func GenerateLetterCaptcha0() (captchaId, answer, base64Img string) {
	return GenerateLetterCaptcha(length)
}

// GenerateCharCaptcha 仅数字验证码
func GenerateDigitCaptcha() (captchaId, answer, base64Img string) {
	driver := base64Captcha.NewDriverDigit(
		height, // height
		width,  // width
		noise,  // noise
		base64Captcha.OptionShowHollowLine|base64Captcha.OptionShowSlimeLine,
		dotCount,
	)
	c := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)

	captchaId, base64Img, answer, err := c.Generate()
	if err != nil {
		panic("生成仅数字验证码出错：" + err.Error())
	}

	// 检查 answer 是否为空
	if answer == "" {
		panic("验证码答案为空")
	}

	return captchaId, answer, base64Img
}

// GenerateCharCaptcha 字符验证码
func GenerateCharCaptcha() (captchaId, answer, base64Img string) {
	//OptionShowHollowLine	中空线条	默认常用
	//OptionShowSlimeLine	粘稠线条	稍混乱
	//OptionShowSineLine	正弦曲线	像水波一样的线
	//OptionShowNoiseDot	加噪点	加随机点干扰
	//OptionUseComplexNoise	复杂噪声	更高干扰度，机器更难识别
	driver := base64Captcha.NewDriverString(
		height, // height
		width,  // width
		noise,  // noise
		base64Captcha.OptionShowHollowLine|base64Captcha.OptionShowSlimeLine,
		length,        // length
		source,        // source
		nil, nil, nil, // custom fonts/colors
	)
	c := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)

	captchaId, base64Img, answer, err := c.Generate()
	if err != nil {
		panic("生成字符验证码出错：" + err.Error())
	}

	return captchaId, answer, base64Img
}

// GenerateMathCaptcha 数字四则运算验证码
func GenerateMathCaptcha() (captchaId, answer, base64Img string) {
	// 创建四则运算 Driver（默认带干扰）
	driver := base64Captcha.NewDriverMath(
		height, // height
		width,  // width
		noise,  // noise
		base64Captcha.OptionShowSlimeLine|base64Captcha.OptionShowSineLine, // 干扰类型
		nil, nil, nil, // 字体/背景色/字体色
	)

	// 使用默认内存 store 生成验证码
	c := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)

	// 生成验证码
	captchaId, base64Img, answer, err := c.Generate()
	if err != nil {
		panic("生成数字四则运算验证码出错：" + err.Error())
	}

	return captchaId, answer, base64Img
}

// GenerateMathCaptcha 数语音验证码
func GenerateAudioCaptcha() (captchaId, answer, base64Audio string) {
	// "en", "ja", "ru", "zh"
	driver := base64Captcha.NewDriverAudio(audioLength, "zh")

	// 使用默认内存 store 生成验证码
	c := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)

	// data:audio/wav;base64,
	// 生成验证码
	captchaId, base64Audio, answer, err := c.Generate()
	if err != nil {
		panic("生成数语音验证码出错：" + err.Error())
	}

	return captchaId, answer, base64Audio
}
