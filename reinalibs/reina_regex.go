package reinalibs

import "regexp"

func IsReinaCalling(content string) bool {
	// Normal ueshama call
	// be careful for 〜 strings in Windows OS!!!!!
	// https://ch.nicovideo.jp/suwatoh/blomaga/ar634436
	if content == "うえしゃま\u301C" || content == "うえしゃま\uFF5E" {
		return true
	}
	if match, err := regexp.MatchString("^うえしゃま([ぁあ])+$", content); err == nil && match {
		return true
	}
	return false
}
