package bot

// const spreadsheet = "https://docs.google.com/spreadsheets/d/1MmlNT2h_fqvz8Oj2SEhwJvPIR54-nWo__M8LmqVthTo/edit?usp=sharing"

const privateString = "t.me/rfbroccoli_bot \n \n/help - π command list ααΌαα« \n/register - πΎ register αα―ααΊαααΊ\n/when - β³ αα¬αααΊα·α‘αααΊαΈα αααΊαα±α¬α·αα²?\n /attendance - π attendance ααΌααΊα·αααΊ spreadsheet"
const groupString = "t.me/rfbroccoli_bot \n \n/help - π command list ααΌαα« \n/when - β³ αα¬αααΊα·α‘αααΊαΈα αααΊαα±α¬α·αα²?\n/attendance - π attendance ααΌααΊα·αααΊ spreadsheet"
const attendanceString = "π spreadsheet link:\nhttps://docs.google.com/spreadsheets/d/1MmlNT2h_fqvz8Oj2SEhwJvPIR54-nWo__M8LmqVthTo/edit?usp=sharing\n"
const noRegisterString = "β privacy α‘α group αα²ααΎα¬ register αα±αΈααα―ααΊαα«α\nregister αα―ααΊαααΊ private chat αα­α―αα¬αα« t.me/rfbroccoli_bot"

type Class struct {
	ClassName   string `json:"class_name,omitempty"`
	BatchName   string `json:"batch_name,omitempty"`
	BatchNum    int    `json:"batch_num,omitempty"`
	Description string `json:"description,omitempty"`
	PlaylistId  string `json:"playlist_id,omitempty"`
	CoverImage  string `json:"cover_image,omitempty"`
}

var batch17 = Class{
	ClassName: "Programming With Go",
	BatchName: "Batch 17",
	BatchNum:  17,
}
