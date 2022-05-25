package service

import (
	"log"
	"strings"

	"github.com/EriconYu/BaXingNengLiang/dto"
)

/**
 * @Description 使用八星能量分析手机号
	手机后五位代表父母 兄弟 事业 夫妻 子女
 * @Params  [phoneNumber]
 * @Returns  [err, result]
 * @Creator  Ericon Freeman
 * @Date  2022/5/24 3:40 下午
*/
func ParsePhoneNumberByBaXing(phoneNumber string) (baxingArray string, result string) {

	baxingArray, result = commonParse(phoneNumber)
	return baxingArray, result
}

func commonParse(phoneNumber string) (baxingArray string, result string) {
	var baxing, jianyi string
	result += parse1Byte_7_11(phoneNumber)
	if result != "" {
		result += "，"
	}
	//拆解4~11位
	slices := splitPhoneNumber4_11(phoneNumber)
	log.Println("手机号码分拆后：", slices)
	//  2位  3位 解析：
	for i, slice := range slices {
		if len(slice) == 2 {
			baxing, jianyi = parse2Bytes_4_11(slice)
		}
		if len(slice) == 3 {
			baxing, jianyi = parse3Bytes_4_11(slice)
		}
		if baxing != "" && jianyi != "" {
			if !strings.Contains(baxingArray, baxing) {
				result += jianyi
				if result != "" && i != len(slice)-1 {
					result += "，"
				}
			}
			if baxingArray != "" {
				baxingArray += "，"
			}
			baxingArray += baxing
			log.Println(slice, "八星："+baxing+" ;", "含义："+jianyi)
		} else {
			log.Println("error: " + slice + "没有找到对应的含义")
		}
	}
	return baxingArray, result
}

func splitPhoneNumber4_11(phoneNumber string) []string {
	var result []string
	for i := 3; i < 10; i++ {
		if phoneNumber[i] == '5' || phoneNumber[i] == '0' {
			continue
		}
		if i != 9 {
			if phoneNumber[i+1] == '5' || phoneNumber[i+1] == '0' || phoneNumber[i+2] == '5' || phoneNumber[i+2] == '0' {
				result = append(result, phoneNumber[i:i+3])
				continue
			}
		}
		result = append(result, phoneNumber[i:i+2])
	}
	return result
}

func parse1Byte_7_11(phoneNumber string) string {

	var result string
	if phoneNumber[10] == '4' {
		result +="【尽量尾号别带4】"
	}
	//手机后五位代表父母 兄弟 事业 夫妻 子女
	if phoneNumber[6] == '0' {
		result += "与父母亲缘分薄"
	}
	if phoneNumber[7] == '0' {
		result += "与兄弟姐妹缘分薄，人际关系方面交心的朋友很少，并且更有甚者兄弟姐妹有夭折或者是母亲有流产"
	}
	if phoneNumber[8] == '0' {
		result += "事业平平，甚至容易失业或者无业，所做的工作与之前学习的专业/技能不相符，工作中不易做出成绩，有才华却无处发挥"
	}
	if phoneNumber[9] == '0' {
		result += "感情少，缘分薄，不稳定，不易成婚或长久"
	}
	if phoneNumber[10] == '0' {
		result += "与子女缘分薄，更有甚者会出现子女夭折甚至流产的情况"
	}
	return result
}
func parse2Bytes_4_11(phoneNumber2Byte string) (baxing, jianyi string) {

	for _, s := range dto.TIANYI_PHONENUMBER_2BYTES {
		if s == phoneNumber2Byte {
			return dto.BAXING_ENUM[dto.TIANYI_INDEX], dto.BAXING_HANYI_2BYTES_ENUM[dto.TIANYI_INDEX]
		}
	}
	for _, s := range dto.YANNIAN_PHONENUMBER_2BYTES {
		if s == phoneNumber2Byte {
			return dto.BAXING_ENUM[dto.YANNIAN_INDEX], dto.BAXING_HANYI_2BYTES_ENUM[dto.YANNIAN_INDEX]
		}
	}
	for _, s := range dto.SHENGQI_PHONENUMBER_2BYTES {
		if s == phoneNumber2Byte {
			return dto.BAXING_ENUM[dto.SHENGQI_INDEX], dto.BAXING_HANYI_2BYTES_ENUM[dto.SHENGQI_INDEX]
		}
	}
	for _, s := range dto.FUWEI_PHONENUMBER_2BYTES {
		if s == phoneNumber2Byte {
			return dto.BAXING_ENUM[dto.FUWEI_INDEX], dto.BAXING_HANYI_2BYTES_ENUM[dto.FUWEI_INDEX]
		}
	}
	for _, s := range dto.JUEMING_PHONENUMBER_2BYTES {
		if s == phoneNumber2Byte {
			return dto.BAXING_ENUM[dto.JUEMING_INDEX], dto.BAXING_HANYI_2BYTES_ENUM[dto.JUEMING_INDEX]
		}
	}
	for _, s := range dto.HUOHAI_PHONENUMBER_2BYTES {
		if s == phoneNumber2Byte {
			return dto.BAXING_ENUM[dto.HUOHAI_INDEX], dto.BAXING_HANYI_2BYTES_ENUM[dto.HUOHAI_INDEX]
		}
	}
	for _, s := range dto.WUGUI_PHONENUMBER_2BYTES {
		if s == phoneNumber2Byte {
			return dto.BAXING_ENUM[dto.WUGUI_INDEX], dto.BAXING_HANYI_2BYTES_ENUM[dto.WUGUI_INDEX]
		}
	}
	for _, s := range dto.LIUSHA_PHONENUMBER_2BYTES {
		if s == phoneNumber2Byte {
			return dto.BAXING_ENUM[dto.LIUSHA_INDEX], dto.BAXING_HANYI_2BYTES_ENUM[dto.LIUSHA_INDEX]
		}
	}
	return "", ""
}
func parse3Bytes_4_11(phoneNumber3Byte string) (baxing, jianyi string) {

	for numberSlice, jianyi := range dto.SHENGQI_PHONENUMBER_3BYTES_MAP {
		if numberSlice == phoneNumber3Byte {
			return dto.BAXING_ENUM[dto.SHENGQI_INDEX], jianyi
		}
	}
	for numberSlice, jianyi := range dto.YANNIAN_PHONENUMBER_3BYTES_MAP {
		if numberSlice == phoneNumber3Byte {
			return dto.BAXING_ENUM[dto.YANNIAN_INDEX], jianyi
		}
	}
	for numberSlice, jianyi := range dto.LIUSHA_PHONENUMBER_3BYTES_MAP {
		if numberSlice == phoneNumber3Byte {
			return dto.BAXING_ENUM[dto.LIUSHA_INDEX], jianyi
		}
	}
	for numberSlice, jianyi := range dto.WUGUI_PHONENUMBER_3BYTES_MAP {
		if numberSlice == phoneNumber3Byte {
			return dto.BAXING_ENUM[dto.WUGUI_INDEX], jianyi
		}
	}
	for numberSlice, jianyi := range dto.HUOHAI_PHONENUMBER_3BYTES_MAP {
		if numberSlice == phoneNumber3Byte {
			return dto.BAXING_ENUM[dto.HUOHAI_INDEX], jianyi
		}
	}
	for numberSlice, jianyi := range dto.JUEMING_PHONENUMBER_3BYTES_MAP {
		if numberSlice == phoneNumber3Byte {
			return dto.BAXING_ENUM[dto.JUEMING_INDEX], jianyi
		}
	}

	// 走到这里说明没有找到对应的三个数字的，需要手动把0 或者 5拆掉
	if phoneNumber3Byte[1] == '0' {
		phoneNumber2Byte := string(append([]byte{}, phoneNumber3Byte[0], phoneNumber3Byte[2]))
		baxing, jianyi = parse2Bytes_4_11(phoneNumber2Byte)
		//jianyi += ",做事容易虎头蛇尾,三分热血"
	} else if phoneNumber3Byte[1] == '5' {
		phoneNumber2Byte := string(append([]byte{}, phoneNumber3Byte[0], phoneNumber3Byte[2]))
		baxing, jianyi = parse2Bytes_4_11(phoneNumber2Byte)
		//jianyi += ",且有主动之意"
	} else if phoneNumber3Byte[2] == '0' {
		phoneNumber2Byte := string(append([]byte{}, phoneNumber3Byte[0], phoneNumber3Byte[1]))
		baxing, jianyi = parse2Bytes_4_11(phoneNumber2Byte)
		//jianyi += ",做事容易虎头蛇尾,三分热血"
	} else if phoneNumber3Byte[2] == '5' {
		phoneNumber2Byte := string(append([]byte{}, phoneNumber3Byte[0], phoneNumber3Byte[1]))
		baxing, jianyi = parse2Bytes_4_11(phoneNumber2Byte)
		//jianyi += ",且有主动之意"
	}

	return baxing, jianyi
}
