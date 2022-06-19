package bot

// const spreadsheet = "https://docs.google.com/spreadsheets/d/1MmlNT2h_fqvz8Oj2SEhwJvPIR54-nWo__M8LmqVthTo/edit?usp=sharing"

const groupString = "t.me/rf_bot_joi \n \n/help - 📃 command list ပြပါ \n/register - 💾 register လုပ်ရန်\n/when - ⏳ လာမယ့်အတန်းက ဘယ်တော့လဲ?\n /attendance - 🗒 attendance ဖြည့်ရန် spreadsheet"
const privateString = "t.me/rf_bot_joi \n \n/help - 📃 command list ပြပါ \n/when - ⏳ လာမယ့်အတန်းက ဘယ်တော့လဲ?\n/attendance - 🗒 attendance ဖြည့်ရန် spreadsheet"
const attendanceString = "spreadsheet link:\nhttps://docs.google.com/spreadsheets/d/1MmlNT2h_fqvz8Oj2SEhwJvPIR54-nWo__M8LmqVthTo/edit?usp=sharing\n"

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
