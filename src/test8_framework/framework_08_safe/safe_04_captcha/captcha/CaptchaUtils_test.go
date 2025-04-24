package captcha

import (
	"fmt"
	"testing"
)

func TestGenerateCharCaptcha(t *testing.T) {

	//captchaId, answer, b64s := GenerateLetterCaptcha(5)
	//fmt.Println(captchaId)
	//fmt.Println(b64s)
	//fmt.Println(answer)

	//captchaId, answer, b64s := GenerateDigitCaptcha()
	//fmt.Println(captchaId)
	//fmt.Println(answer)
	//fmt.Println(b64s)

	//captchaId, answer, b64s := GenerateCharCaptcha()
	//fmt.Println(captchaId)
	//fmt.Println(answer)
	//fmt.Println(b64s)

	//captchaId, answer, b64s := GenerateMathCaptcha()
	//fmt.Println(captchaId)
	//fmt.Println(answer)
	//fmt.Println(b64s)

	captchaId, answer, base64Audio := GenerateAudioCaptcha()
	fmt.Println(captchaId)
	fmt.Println(answer)
	fmt.Println(base64Audio)

}
