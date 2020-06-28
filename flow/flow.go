package flow

import (
	"fmt"
	"github.com/kisulken/go-telegram-flow/chain"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"strings"
	"testBot/repository"
	"testBot/resources"
)

func GetBotFlow(bot *tb.Bot) *chain.Chain {
	flow, err := chain.NewChainFlow("flow1", bot)
	if err != nil {
		panic(err)
	}
	flow.SetDefaultHandler(defaultResponse).GetRoot().
		Then("searchProduct", stageFirst, tb.OnText).
		Then("selectProduct", stageSelect, tb.OnText).
		Then("orderProduct", stageOrder, tb.OnText)
	return flow
}

func defaultResponse(e *chain.Node, m *tb.Message) *chain.Node {
	e.GetFlow().GetBot().Send(m.Sender, "دستور ارسالی شما پشتیبانی نمی شود")
	return e // stays on the same stage
}

func stageFirst(e *chain.Node, m *tb.Message) *chain.Node {
	if len(m.Text) < 2 {
		e.GetFlow().GetBot().Send(m.Sender, "به نظر خطایی در اسم محصول شما وجود دارد")
		return e // stays on the same stage
	}
	log.Println(m.Sender.Recipient(), "goes through", e.GetId())
	products := resources.GetProducts(m.Text)
	if products != nil {
		e.GetFlow().GetBot().Send(m.Sender, "محصول مورد نظر خود را از لیست انتخاب کنید", products)
		return e.Next()
	}else {
		e.GetFlow().GetBot().Send(m.Sender, "محصول مورد نظر شما یافت نشد لطفا دوباره تلاش کنید")
		return e
	}
}

func stageSelect(e *chain.Node, m *tb.Message) *chain.Node {
	log.Println(m.Sender.Recipient(), "goes through", e.GetId())
	product := repository.GetProductByTitle(m.Text)
	productDetails := fmt.Sprintf("productId : %d \n productName : %s \n  productPrice : %d", product.Id, product.Title, product.Price)
	e.GetFlow().GetBot().Send(m.Sender, productDetails, resources.GetOrderButtons())
	e.GetFlow().GetBot().Send(m.Sender, "اگر مایل به خرید این محصول هستید بله و در غیر این صورت خیر را ارسال کنید ")
	return e.Next() // continue
}

func stageOrder(e *chain.Node, m *tb.Message) *chain.Node {
	text := strings.ToLower(m.Text)
	if text == "بله" {
		e.GetFlow().GetBot().Send(m.Sender, "محصول مورد نظر شما در سیستم درج شد و این محصول هم اکنون در حال پردازش می باشد")
		return nil
	} else if text == "خیر" {
		e.GetFlow().GetBot().Send(m.Sender, "بسیار عالی کار شما با ربات به پایان رسید بعدا بهمون سر بزن !")
		return nil
	} else {
		e.GetFlow().GetBot().Send(m.Sender, "لطفا کلمه بله / خیر را فقط وارد کنید ")
		return e
	}
}
