package models

import "log"

func SearchUsersInDB(str string) (usrs []string, err error) {
	stmt, err := dbAcc.Prepare("SELECT Username FROM UserInfo WHERE Username LIKE CONCAT('%', ? ,'%') ORDER BY Username ASC LIMIT 20")
	if err != nil {
		return nil, err
	}
	usernames, err := stmt.Query(str)
	if err != nil {
		return nil, err
	}
	for usernames.Next() {
		var u string
		err = usernames.Scan(&u)
		if err != nil {
			log.Println(err.Error())
			return
		}
		usrs = append(usrs, u)
	}
	return
}

func db_AddGroupToUsersGroups(uuid int, groupid string) error {
	log.Println("adding group to db")
	stmt, err := dbAcc.Prepare("INSERT INTO GroupMapping (UserID, SubbedGroup) VALUES (?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(uuid, groupid)
	return err
}

func db_RemoveGroupFromUsersGroups(uuid int, groupid string) error {
	stmt, err := dbAcc.Prepare("DELETE FROM GroupMapping WHERE (UserID = ?) AND (SubbedGroup = ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(uuid, groupid)
	return err
}

func db_GetUsersGroups(uuid int) (sl []string, err error) {
	stmt, err := dbAcc.Prepare("SELECT g.SubbedGroup FROM GroupMapping g INNER JOIN UserInfo u ON u.UserID = g.UserID WHERE g.UserID = ?")
	if err != nil {
		return nil, err
	}
	groupNames, err := stmt.Query(uuid)
	if err != nil {
		return nil, err
	}
	for groupNames.Next() {
		var g string
		err = groupNames.Scan(&g)
		if err != nil {
			log.Println(err.Error())
			return
		}
		sl = append(sl, g)
	}
	return
}

func db_GetUsersFriends(uuid int) (sl []string, err error) {
	stmt, err := dbAcc.Prepare("SELECT Username FROM UserInfo WHERE UserID IN (SELECT f.followedID FROM FriendTest f INNER JOIN UserInfo u ON u.UserID = f.followerID WHERE f.followerID = ?)")
	if err != nil {
		return nil, err
	}
	friends, err := stmt.Query(uuid)
	if err != nil {
		return nil, err
	}
	for friends.Next() {
		var f string
		err = friends.Scan(&f)
		if err != nil {
			log.Println(err.Error())
			return
		}
		sl = append(sl, f)
	}
	return
}
