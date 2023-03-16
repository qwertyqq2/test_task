package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestFirst(t *testing.T) {
	const html = `
<div class="container">
    <div class="row">
      <div class="col-lg-8">
        <p align="justify"><b>Name</b>Priyaka</p>
        <p align="justify"><b>Surname</b>Patil</p>
        <p align="justify"><b>Adress</b><br>India,Kolhapur</p>
        <p align="justify"><b>Hobbies&nbsp;</b><br>Playing</p>
        <p align="justify"><b>Eduction</b><br>12th</p>
        <p align="justify"><b>School</b><br>New Highschool</p>
       </div>
    </div>
</div>
`

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		panic(err)
	}

	doc.Find(".container").Find("[align=\"justify\"]").Each(func(_ int, s *goquery.Selection) {
		prefix := s.Find("b").Text()
		result := strings.TrimPrefix(s.Text(), prefix)
		println(result)
		fmt.Println(s.Text())
	})
}

func TestSecond(t *testing.T) {
	const html = `
	<tr id="rowId-744811" class="odd " >
		<td  class="row-type "><span data-toggle="tooltip" title="Электронный аукцион"><span class="label label-info">ЭА</span></span></td>
	<td  class="row-procedure_name "><a href="https://etp-ets.ru/44/ea21/procedure/view/249083?&backurl=LzQ0L2NhdGFsb2cvcHJvY2VkdXJlP3BhZ2U9Mg==">Поставка картриджей для лазерных устройств</a> (0372200157923000002)</td>
	<td  class="row-contract_start_price sortable sortable">104 587.50 RUB</td>
	<td  class="row-placer_name sortable sortable"><a href="/organization/view/524613?&backurl=LzQ0L2NhdGFsb2cvcHJvY2VkdXJlP3BhZ2U9Mg==" target="_blank">ГОСУДАРСТВЕННОЕ БЮДЖЕТНОЕ ДОШКОЛЬНОЕ ОБРАЗОВАТЕЛЬНОЕ УЧРЕЖДЕНИЕ ДЕТСКИЙ САД № 58 ОБЩЕРАЗВИВАЮЩЕГО ВИДА С ПРИОРИТЕТНЫМ ОСУЩЕСТВЛЕНИЕМ ДЕЯТЕЛЬНОСТИ ПО ФИЗИЧЕСКОМУ РАЗВИТИЮ ДЕТЕЙ КАЛИНИНСКОГО РАЙОНА САНКТ-ПЕТЕРБУРГА</a> </td>
	<td  class="row-customer_name "><a href="/organization/view/524613?&backurl=LzQ0L2NhdGFsb2cvcHJvY2VkdXJlP3BhZ2U9Mg==" target="_blank">ГОСУДАРСТВЕННОЕ БЮДЖЕТНОЕ ДОШКОЛЬНОЕ ОБРАЗОВАТЕЛЬНОЕ УЧРЕЖДЕНИЕ ДЕТСКИЙ САД № 58 ОБЩЕРАЗВИВАЮЩЕГО ВИДА С ПРИОРИТЕТНЫМ ОСУЩЕСТВЛЕНИЕМ ДЕЯТЕЛЬНОСТИ ПО ФИЗИЧЕСКОМУ РАЗВИТИЮ ДЕТЕЙ КАЛИНИНСКОГО РАЙОНА САНКТ-ПЕТЕРБУРГА</a> </td>
	<td  class="row-publication_datetime sortable sortable">16.03.2023 14:38:57<span class="tz"><br/>(MSK+00:00)</span></td>
	<td  class="row-request_end_give_datetime sortable sortable">24.03.2023 09:00:00<span class="tz"><br/>(MSK+00:00)</span></td>
	<td  class="row-request_review_end_datetime sortable sortable"></td>
	<td  class="row-trade_start_datetime sortable sortable">24.03.2023 11:00:00<span class="tz"><br/>(MSK+00:00)</span></td>
	<td  class="row-status ">Прием заявок</td>
	<td  class="row-events "><a href="/44/ea21/procedure/view/249083/?target=events&eventId=1343251&backurl=LzQ0L2NhdGFsb2cvcHJvY2VkdXJlP3BhZ2U9Mg==">Публикация извещения от 16.03.2023 14:38:57 (MSK+00:00)</a></td>
		<td class="row-actions">
												<a target="_blank"class=" mte-grid-action action-notice.eis-6413033f5ce5d" href="https://zakupki.gov.ru/epz/order/notice/ea20/view/common-info.html?regNumber=0372200157923000002" >Извещение в ЕИС</a>
			</td>
	</tr>
	`

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		panic(err)
	}

	doc.Find("td").Find("[class=\"row-contract_start_price sortable sortable\"]").Each(func(_ int, s *goquery.Selection) {
		prefix := s.Find("b").Text()
		result := strings.TrimPrefix(s.Text(), prefix)
		println(result)
		fmt.Println(s.Text())
	})

	doc.Find("[row-contract_start_price sortable sortable\"]").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Text())
	})

}
