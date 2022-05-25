package dto

const (
	TIANYI_INDEX  = 0
	YANNIAN_INDEX = 1
	SHENGQI_INDEX = 2
	FUWEI_INDEX   = 3
	JUEMING_INDEX = 4
	HUOHAI_INDEX  = 5
	WUGUI_INDEX   = 6
	LIUSHA_INDEX  = 7
)

var BAXING_ENUM = []string{"天医", "延年", "生气", "伏位", "绝命", "祸害", "五鬼", "六煞"}
var BAXING_HANYI_2BYTES_ENUM = []string{
	"财富方面比较不错，而且桃花旺，亦或是婚姻幸福，但是已婚者注意防桃花",
	"身体健康，易长寿",
	"富有活力，精力充足",
	"学习能力强，喜欢学习，自身运势稳定",
	"情绪易波动起伏",
	"疾病方面要多多留心，建议每年做一次正规的医疗体检",
	"心智方面需要加强锻炼",
	"感情和婚姻上容易出问题，感情不顺利"}

var TIANYI_PHONENUMBER_2BYTES = []string{"13", "31", "68", "86", "49", "94", "27", "72"}
var YANNIAN_PHONENUMBER_2BYTES = []string{"19", "91", "78", "87", "34", "43", "26", "62"}
var SHENGQI_PHONENUMBER_2BYTES = []string{"14", "41", "67", "76", "39", "93", "28", "82"}
var FUWEI_PHONENUMBER_2BYTES = []string{"11", "22", "88", "99", "66", "77", "33", "44"}
var JUEMING_PHONENUMBER_2BYTES = []string{"12", "21", "69", "96", "48", "84", "37", "73"}
var HUOHAI_PHONENUMBER_2BYTES = []string{"17", "71", "89", "98", "46", "64", "23", "32"}
var WUGUI_PHONENUMBER_2BYTES = []string{"18", "81", "79", "97", "36", "63", "24", "42"}
var LIUSHA_PHONENUMBER_2BYTES = []string{"16", "61", "47", "74", "38", "83", "29", "92"}

var SHENGQI_PHONENUMBER_3BYTES_MAP = map[string]string{
	"154": "人缘好在明处，好朋友确认到人，人际关系上边界感比较强",
	"104": "人际关系上容易反目，总是容易遇到小人",
	"145": "人际关系比较好，好朋友越来越多",
	"140": "朋友比较少，交心的更少"}
var YANNIAN_PHONENUMBER_3BYTES_MAP = map[string]string{
	"159": "大家公认的能力强人",
	"109": "能力被埋没，怀才不遇",
	"195": "自身能力越来越强大",
	"190": "能力减弱，江郎才尽，自身工作的单位也越来越衰弱，自己权力越来越小"}
var LIUSHA_PHONENUMBER_3BYTES_MAP = map[string]string{
	"106": "容易郁闷，生闷气，注意调节情绪",
	"160": "心情容易糟糕，生气",
	"156": "自己容易主动陷入到负面情绪中",
	"165": "内心负能量大，容易反社会"}
var WUGUI_PHONENUMBER_3BYTES_MAP = map[string]string{
	"709": "想改变现状",
	"790": "改变现状后不如以前",
	"759": "变动当下的处境比较吃力",
	"795": "做事东张西望，三心二意，善变",
	"452": "心智方面需要主动加强训练，自身易主动陷入到消极负面的情绪中"}
var HUOHAI_PHONENUMBER_3BYTES_MAP = map[string]string{
	"157": "注意身体健康，易出现确诊的大病",
	"107": "易无法确诊的病情",
	"175": "易因自己不注重保养导致患大病",
	"170": "易患有慢性病症"}
var JUEMING_PHONENUMBER_3BYTES_MAP = map[string]string{
	"102": "留不住钱财，或者钱财被占用导致失去流动性",
	"152": "财难受",
	"120": "谨慎投资，否则钱财容易石沉大海",
	"125": "自身在持续投资，该手机号适合投行从业者使用"}
