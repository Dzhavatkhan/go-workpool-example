package postman

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func postman(
	ctx context.Context,
	wg *sync.WaitGroup,
	transferPoint chan<- string,
	n int,
	mail string,
) {
	defer wg.Done();
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Почтальон ", n, " закончил, слава России!");
			return;
		default:
			fmt.Println("Я почтальон № ", n, ", взял письмо")
			time.Sleep(1 * time.Second)
			fmt.Println("Я почтальон № ", n, ", доставил письмо: ", mail)

			transferPoint <- mail
			fmt.Println("Доставил ", mail)
		}

	}

}

func PostmanPool(
	ctx context.Context,
	postmanCount int,
) <-chan string {
	wg 				  := &sync.WaitGroup{}
	mailTransferPoint := make(chan string);

	for i := 1; i <= postmanCount; i++ {
		wg.Add(1);
		go postman(ctx,wg, mailTransferPoint, i, postmanToMail(i))
	}
	go func(){
	  wg.Wait();
	  close(mailTransferPoint);
	}()	
	return mailTransferPoint;
}

func postmanToMail(postmanNumber int) string {
	ptm 	:= map[int]string{
		1:  "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam auctor, nisl eget ultricies",
		2:  "Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam",
		3:  "Quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat",
		4:  "Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat",
		5:  "Привет sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit",
		6:  "Ну как ты там sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit",
		7:  "Вообще sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit",
		8:  "Много изменилось sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit",
		9:  "Вещей, но sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit",
		10: "Я по-прежнему ничей sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit",
		11: "Привет! Хотим узнать, всё ли у вас хорошо с работой нашего сервиса. Ответьте на два вопроса в опросе",
		12: "Здравствуйте! Мы заметили подозрительный вход в ваш аккаунт. Если это были вы — игнорируйте письмо",
		13: "Добрый день! Отправляем вам чек за август во вложении. С уважением, команда поддержки",
		14: "Привет! У нас для вас сюрприз — бесплатный месяц использования премиум-функций. Активируйте сегодня!",
	}

	mail, ok := ptm[postmanNumber]
	if !ok {
		return "ШТРАФ 5000!"
	}
	return mail
}
