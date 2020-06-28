package resources

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"testBot/repository"
)

func GetPhoneNumberMarkup() *tb.ReplyMarkup {
	btnSharePhone := tb.ReplyButton{
		Contact: true,
		Text:    "ارسال شماره تماس",
	}
	replySharePhone := [][]tb.ReplyButton{
		{
			btnSharePhone,
		},
	}
	markup := &tb.ReplyMarkup{
		ReplyKeyboard: replySharePhone,
		ForceReply:    true,
	}
	return markup
}

func GetProducts(name string) *tb.ReplyMarkup {
	products := repository.GetProducts(name)
	allButtons := [][]tb.ReplyButton{}
	if len(products) == 0 {
		return nil
	}
	for _, v := range products {
		btnSharePhone := []tb.ReplyButton{
			{
				Text: v.Title,
			},
		}
		allButtons = append(allButtons, btnSharePhone)
	}

	markup := &tb.ReplyMarkup{
		ReplyKeyboard: allButtons,
		ForceReply:    true,
	}
	return markup
}

func GetOrderButtons() *tb.ReplyMarkup {
	yesButton := []tb.ReplyButton{
		{
			Text: " بله",
		},
	}
	noButton := []tb.ReplyButton{
		{
			Text: " خیر",
		},
	}
	orderButtons := [][]tb.ReplyButton{
		yesButton,
		noButton,
	}
	markup := &tb.ReplyMarkup{
		ReplyKeyboard: orderButtons,
		ForceReply:    true,
	}
	return markup
}
