package parser

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(
	`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var salaryRe = regexp.MustCompile(
	`<td><span class="label">月收入：</span>([\d\-]+)元</td>`)
var genderRe = regexp.MustCompile(
	`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var heightRe = regexp.MustCompile(
	`<td><span class="label">身高：</span>([\d]+)CM</td>`)
var weightRe = regexp.MustCompile(
	`<td><span class="label">体重：</span><span field="">([\d]+)KG</span></td>`)
var marriageRe = regexp.MustCompile(
	`<td><span class="label">婚况：</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(
	`<td><span class="label">学历：</span>([^<]+)</td>`)
var occupationRe = regexp.MustCompile(
	`<td><span class="label">职业： </span>([^<]+)</td>`)
var workCityRe = regexp.MustCompile(
	`<td><span class="label">工作地：</span>([^<]+)</td>`)
var locationRe = regexp.MustCompile(
	`<td><span class="label">籍贯：</span>([^<]+)</td>`) //籍贯
var houseRe = regexp.MustCompile(
	`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carRe = regexp.MustCompile(
	`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)

func ProfileParser(context []byte, name string) engine.ParserResult {
	profile := model.Profile{}
	profile.NickName = name
	profile.Age = execInt(ageRe, context)
	profile.Salary = execString(salaryRe, context)
	profile.Gender = execString(genderRe, context)
	profile.Height = execInt(heightRe, context)
	profile.Weight = execInt(weightRe, context)
	profile.Marriage = execString(marriageRe, context)
	profile.Education = execString(educationRe, context)
	profile.Occupation = execString(occupationRe, context)
	profile.WorkCity = execString(workCityRe, context)
	profile.Location = execString(locationRe, context)
	profile.House = execString(houseRe, context)
	profile.Car = execString(carRe, context)

	result := engine.ParserResult{}
	result.Items = append(result.Items, profile)
	return result
}

func execString(re *regexp.Regexp, context []byte) string {
	match := re.FindSubmatch(context)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}

func execInt(re *regexp.Regexp, context []byte) int {
	match := re.FindSubmatch(context)
	if len(match) >= 2 {
		c, err := strconv.Atoi(string(match[1]))
		if err != nil {
			panic(err)
		}
		return c
	} else {
		return -1
	}
}
