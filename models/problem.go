package models

import (
	"math"
	"strings"

	. "github.com/open-fightcoder/oj-web/common/store"
)

type Problem struct {
	Id                 int64  `form:"id" json:"id"`
	Flag               int    `form:"flag" json:"flag"`                             //1-普通题目 2-用户题目
	Status             int    `form:"status" json:"status"`                         //申请状态
	UserId             int64  `form:"user_id" json:"user_id"`                       //题目提供者
	Difficulty         string `form:"difficulty" json:"difficulty"`                 //题目难度
	CaseData           string `form:"caseData" json:"caseData"`                     //测试数据
	Title              string `form:"title" json:"title"`                           //题目标题
	Description        string `form:"description" json:"description"`               //题目描述
	InputDes           string `form:"inputDes" json:"inputDes"`                     //输入描述
	OutputDes          string `form:"outputDes" json:"outputDes"`                   //输出描述
	InputCase          string `form:"inputCase" json:"inputCase"`                   //测试输入
	OutputCase         string `form:"outputCase" json:"outputCase"`                 //测试输出
	Hint               string `form:"hint" json:"hint"`                             //题目提示(可以为对样例输入输出的解释)
	TimeLimit          int    `form:"timeLimit" json:"timeLimit"`                   //时间限制
	MemoryLimit        int    `form:"memoryLimit" json:"memoryLimit"`               //内存限制
	Tag                int64  `form:"tag" json:"tag"`                               //题目标签
	IsSpecialJudge     bool   `form:"isSpecialJudge" json:"isSpecialJudge"`         //是否特判
	SpecialJudgeSource string `form:"specialJudgeSource" json:"specialJudgeSource"` //特判程序源代码
	Code               string `form:"code" json:"code"`                             //标准程序
	LanguageLimit      string `form:"languageLimit" json:"languageLimit"`           //语言限制
	Remark             string `form:"remark" json:"remark"`                         //备注
}

func ProblemCreate(problem *Problem) (int64, error) {
	return OrmWeb.Insert(problem)
}

func ProblemRemove(id int64) error {
	_, err := OrmWeb.Id(id).Delete(&Problem{})
	return err
}

func ProblemUpdate(problem *Problem) error {
	_, err := OrmWeb.AllCols().ID(problem.Id).Update(problem)
	return err
}

func ProblemGetById(id int64) (*Problem, error) {
	problem := new(Problem)
	has, err := OrmWeb.Id(id).Get(problem)

	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return problem, nil
}

func ProblemGetByUserId(userId int64, currentPage int, perPage int) ([]*Problem, error) {
	problemList := make([]*Problem, 0)
	err := OrmWeb.Where("user_id=?", userId).Limit(perPage, (currentPage-1)*perPage).Find(&problemList)
	if err != nil {
		return nil, err
	}
	return problemList, nil
}

func ProblemCountByUserId(userId int64) (int64, error) {
	problem := &Problem{}
	count, err := OrmWeb.Where("user_id=?", userId).Count(problem)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func getNum(tag string) int {
	tagArr := map[string]int{
		"零": 0,
		"一": 1,
		"二": 2,
		"三": 3,
		"四": 4,
		"五": 5,
		"六": 6,
		"七": 7,
		"八": 8,
		"九": 9,
	}
	num := 0
	if tag != "" {
		strs := strings.Split(tag, ",")
		for i := 0; i < len(strs); i++ {
			num += int(math.Pow(2, float64(tagArr[strs[i]])))
		}
	}
	return num
}

func ProblemGetIdsByConds(origins []int64, tag string) ([]*Problem, error) {
	session := OrmWeb.NewSession()
	if len(origins) != 0 {
		session.In("user_id", origins)
	}
	tagNum := getNum(tag)
	if tagNum != 0 {
		session.Where("tag & ? > 0", tagNum)
	}
	problemList := make([]*Problem, 0)

	err := session.Cols("id").Find(&problemList)
	if err != nil {
		return nil, err
	}
	return problemList, nil
}

func ProblemGetProblem(origins []int64, tag string, sortKey string, isAscKey string, currentPage int, perPage int) ([]*Problem, error) {
	session := OrmWeb.NewSession()
	if len(origins) != 0 {
		session.In("user_id", origins)
	}
	tagNum := getNum(tag)
	if tagNum != 0 {
		session.Where("tag & ? > 0", tagNum)
	}
	if isAscKey == "asc" {
		session.Asc(sortKey).Limit(perPage, (currentPage-1)*perPage)
	} else {
		session.Desc(sortKey).Limit(perPage, (currentPage-1)*perPage)
	}
	problemList := make([]*Problem, 0)

	err := session.Find(&problemList)
	if err != nil {
		return nil, err
	}
	return problemList, nil
}

func ProblemCountProblem(origins []int64, tag string) (int64, error) {
	session := OrmWeb.NewSession()
	if len(origins) != 0 {
		session.In("user_id", origins)
	}
	tagNum := getNum(tag)
	if tagNum != 0 {
		session.Where("tag & ? > 0", tagNum)
	}
	problem := &Problem{}
	count, err := session.Count(problem)
	if err != nil {
		return 0, err
	}
	return count, nil
}
