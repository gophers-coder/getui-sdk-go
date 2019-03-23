package getui

const (
	// Auth
	AUTHSIGN  = "https://restapi.getui.com/v1/%s/auth_sign"
	AUTHCLOSE = "https://restapi.getui.com/v1/%s/auth_close"

	// Alias
	BINDALIAS      = "https://restapi.getui.com/v1/%s/bind_alias"
	QUERYCID       = "https://restapi.getui.com/v1/%s/query_cid/%s"
	QUERYALIAS     = "https://restapi.getui.com/v1/%s/query_alias/%s"
	UNBINDALIAS    = "https://restapi.getui.com/v1/%s/unbind_alias"
	UNBINDALLALIAS = "https://restapi.getui.com/v1/%s/unbind_alias_all"

	// Tag
	SETTAGS = "https://restapi.getui.com/v1/%s/set_tags"
	GETTAGS = "https://restapi.getui.com/v1/%s/get_tags/%s"

	// UserBlackList
	USERBLACKLIST = "https://restapi.getui.com/v1/%s/user_blk_list"

	// UserStatus
	USETSTATUS = "https://restapi.getui.com/v1/%s/user_status/%s"

	// PushResult
	PUSHRESULT       = "https://restapi.getui.com/v1/%s/push_result"
	QUERYAPPUSER     = "https://restapi.getui.com/v1/%s/query_app_user/%s"
	QUERYAPPPUSH     = "https://restapi.getui.com/v1/%s/query_app_push/%s"
	PUSHRESULTBYTASK = "https://restapi.getui.com/v1/%s/push_result"

	// PushSingle
	PUSHSINGLE = "https://restapi.getui.com/v1/%s/push_single"

	// User Online
	ONLINEUSER = "https://restapi.getui.com/v1/%s/get_last_24hours_online_User_statistics"

	//
	SAVELISTBODY    = "https://restapi.getui.com/v1/%s/save_list_body"
	PUSHLIST        = "https://restapi.getui.com/v1/%s/push_list"
	PUSHAPP         = "https://restapi.getui.com/v1/%s/push_app"
	STOPTASK        = "https://restapi.getui.com/v1/%s/stop_task/%s"
	PUSHSINGLEBATCH = "https://restapi.getui.com/v1/%s/push_single_batch"
	GETSCHEDULETASK = "https://restapi.getui.com/v1/%s/get_schedule_task"
	DELSCHEDULETASK = "https://restapi.getui.com/v1/%s/del_schedule_task"
)
