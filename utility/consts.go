package utility

var Regions []string
var Departments []string
var Gender map[string]string

func init() {
	Regions = []string{"未选择", "朝晖", "屏峰", "莫干山"}
	Departments = []string{
		"未选择",
		"办公室",
		"活动部",
		"秘书处",
		"Touch产品部",
		"小弘工作室",
		"编辑工作室",
		"视觉影像部",
		"技术部",
		"易班文化工作站",
	}
	Gender = map[string]string{
		"0": "男",
		"1": "女",
	}
}
