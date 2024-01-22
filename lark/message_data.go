package lark

// MessageData 飞书消息的封装结构
// 消息类型统一为 消息卡片，接收者ID类型统一为 openID
type MessageData struct {
	Content   MessageContent
	ReceiveID string
}

type MessageContent struct {
	Content string
	ImgList []string
	Title   string
	URL     string
}

/*
{
  "config": {
    "wide_screen_mode": true
  },
  "elements": [
    {
      "tag": "markdown",
      "content": "这里是卡片文本，支持使用markdown标签设置文本格式。例如：\n*斜体* 、**粗体**、~~删除线~~、[文字链接](https://www.feishu.cn)、<at id=all></at>、<font color='red'> 彩色文本 </font>"
    },
    {
      "alt": {
        "content": "",
        "tag": "plain_text"
      },
      "img_key": "img_v2_041b28e3-5680-48c2-9af2-497ace79333g",
      "tag": "img"
    }
  ],
  "header": {
    "template": "blue",
    "title": {
      "content": "这里是卡片标题",
      "tag": "plain_text"
    }
  }
}
*/
