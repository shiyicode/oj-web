package managers

import (
	"github.com/open-fightcoder/oj-web/models"
	"github.com/pkg/errors"
	"fmt"
)



func TeamCreat(name, avator, description string, uid int64) (int64, error) {

	//判断是否存在
	group, err := models.TeamGetbyName(name)
	if err != nil {
		return 0, fmt.Errorf("get Team failure : %s ", err.Error())
	}
	if group != nil {
		return 0, errors.New("Team is exist")
	}
	//创建
	groupinfo := &models.Team{Name:name, Uid:uid, Avator:avator, Description:description }
	createId, err := models.TeamCreate(groupinfo)
	if err != nil {
		return 0, fmt.Errorf("create group failure : %s ", err.Error())
	}

	groupinfo, _ = models.TeamGetbyName(name)

	member := &models.TeamMember{Uid:uid, Gid:groupinfo.Id}
	//加入uid
	_, err = models.MemberAdd(member)
	if err != nil {
		return 0, fmt.Errorf("add uid failure : %s ", err.Error())
	}

	return createId, err
}

func MemberAdd(gid, uid, user int64) (string, error) {

	//判断组是否存在
	group, err := models.TeamGetbyId(gid)


	if err != nil {
		return "", fmt.Errorf("get Team failure : %s ", err.Error())
	}
	if group == nil {
		return "", errors.New("Team not exist")
	}

	if group.Uid != user {
		return "", errors.New("Permition deny")
	}

	//是否需要判断用户合法？

	//判断用户所属
	id, err:= MemberCheckByGid(uid, gid)
	if err == nil {
		return "", fmt.Errorf("get member failure : %s ", err.Error())
	}
	if id != 0 {
		return "", errors.New( "Member is already in the Team")
	}

	//添加
	groupmember := new(models.TeamMember)
	groupmember.Gid = gid
	groupmember.Uid = uid
	_, err = models.MemberAdd(groupmember)
	if err != nil {
		return "" ,fmt.Errorf("add memver failure: %s", err.Error())
	}
	return "done", nil
}


func TeamRemove(id, owner int64) error{

	//判断组是否存在
	group, err := models.TeamGetbyId(id)
	if err != nil {
		return fmt.Errorf("get Team failure : %s ", err.Error())
	}
	if group == nil {
		return errors.New("Team not exist")
	}
	if group.Uid != owner {
		return errors.New("Permition deny")
	}

	//得到所有成员
	members, err := models.MembersQueryByGid(id)
	if err != nil {
		return fmt.Errorf("get members failure : %s ", err.Error())
	}

	//删除成员

	for _, member:= range *members {
		err = models.MemberRemove(member.Id)
		if err != nil{
			return fmt.Errorf("delete member failure : %s ", err.Error())
		}
	}



	//删除组
	err = models.TeamRemove(id)
	if err != nil {
		return fmt.Errorf("delete group failure : %s ", err.Error())
	}

	return nil
}

func MemberRemove(uid, gid, user int64) error {

	//判断组是否存在
	group, err := models.TeamGetbyId(gid)
	if err != nil {
		return fmt.Errorf("get Team failure : %s ", err.Error())
	}
	if group == nil {
		return errors.New("Team not exist")
	}


	//判断用户是否所属组
	id, err := MemberCheckByUid(uid, gid)
	if err != nil {
		return fmt.Errorf("get member failure : %s ", err.Error())
	}
	if id == 0 {
		return errors.New("the member doesn't belong to the group")
	}

	//判断是否有权限
	if user != uid && user != group.Uid {
		return errors.New("Permition deny")
	}


	//删除
	err = models.MemberRemove(id)
	if err != nil {
		return fmt.Errorf("delete member failure : %s ", err.Error())
	}

	return nil
}


func GetTeam(id int64) (*map[string]interface{}, error) {
	//判断组是否存在
	groupinfo, err := models.TeamGetbyId(id)
	if err != nil {
		return nil, fmt.Errorf("get Team failure : %s ", err.Error())
	}
	if groupinfo == nil {
		return nil, errors.New("Team not exist")
	}


	//获得成员
	members, err:= models.MembersQueryByGid(id)
	fmt.Println(members)
	if err != nil {
		return nil, fmt.Errorf("get members failure : %s ", err.Error())
	}

	var member_id []int64
	for _, node := range *members {
		member_id = append(member_id, node.Uid)
	}

	team := map[string]interface{}{
		"id": groupinfo.Id,
		"uid":groupinfo.Uid,
		"name":groupinfo.Name,
		"description":groupinfo.Description,
		"avator":groupinfo.Avator,
		"Member_id":member_id,
	}

	return &team, err

}

//判断uid成员是否加入gid组
func MemberCheckByUid(uid, gid int64) (int64, error) {
	groupmember, err := models.MembersQueryByUid(uid)
	if err != nil {
		return 0,  fmt.Errorf("get Team failure : %s ", err.Error())
	}

	for _, member := range *groupmember {
		if member.Gid == gid {
			return member.Id, nil
		}
	}

	return 0,errors.New("no Team")

}

//判断gid组内有没有uid成员
func MemberCheckByGid(uid, gid int64) (int64, error) {
	groupmember, err := models.MembersQueryByGid(gid)
	if err != nil {
		return 0,  fmt.Errorf("get Team failure : %s ", err.Error())
	}

	for _, member := range *groupmember {
		if member.Uid == uid {
			return member.Id, nil
		}
	}


	return 0,errors.New("no member")

}

func TeamUpdate(id, owner int64, name, avator, description string) (string, error) {

	//判断组是否存在
	group, err := models.TeamGetbyId(id)
	if err != nil {
		return "", fmt.Errorf("get Team failure : %s ", err.Error())
	}
	if group == nil {
		return "", errors.New("Team not exist")
	}
	if group.Uid != owner {
		return "", errors.New("Permition deny")
	}

	if name != "" {
		group.Name = name
	}
	if avator != "" {
		group.Avator = avator
	}
	if description != "" {
		group.Description = description
	}
	err = models.TeamUpdate(group)
	if err != nil {
		return "", fmt.Errorf("get Team failure : %s ", err.Error())
	}

	return "done", err
}