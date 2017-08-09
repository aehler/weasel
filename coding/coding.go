package main

import (
	"fmt"
)

var text string     = "человекизвестныйсакраментоименемшерифагенриблэкаразыскиваетсявластямиштатаобвинениюхищениизолотогозапасапринадлежащегоказневзятьживымвознаграждениегарантировуетсямаршалстингрейскатсэм"
var keywords string = "вислоухойптицыкутыконическийзаяцгенриблэкуктантекупотактевонючихтопейбритвытомогавкаколотерабартоломеомагдаленисакрментоньюйоркситибарибаситиквакальщиковантаркидыаляскистингрейскатсэм"


var text1 string = "калифорния"
var text2 string = "джокджоук"

var alphabet string = "абвгдежсзиклмнопрстуфхцчшщъыьэюя"

var key = map[rune]rune {
	'а' :'ᑫ',
	'б' :'ᒲ',
	'в' :'ᑭ',
	'г' :'ᑮ',
	'д' :'ᑯ',
	'е' :'ᒷ',
	'ж' :'ᒶ',
	'з' :'ᑲ',
	'и' :'ᒺ',
	'й' :'ᑴ',
	'к' :'ᑵ',
	'л' :'ᑶ',
	'м' :'ᑷ',
	'н' :'ᑸ',
	'о' :'ᑹ',
	'п' :'ᑺ',
	'р' :'ᑻ',
	'с' :'ᑼ',
	'т' :'ᑽ',
	'у' :'ᒐ',
	'ф' :'ᒠ',
	'х' :'ᒜ',
	'ц' :'ᒓ',
	'ч' :'ᒙ',
	'ш' :'ᒥ',
	'щ' :'ᒖ',
	'ъ' :'ᒧ',
	'ы' :'ᒪ',
	'ь' :'ᒬ',
	'э' :'ᒭ',
	'ю' :'ᒮ',
	'я' :'ᒯ',
	' ' : 'ᕒ',
}

var key_california = map[rune]rune {
	'а' :'а',
	'б' :'ᒲ',
	'в' :'ᑭ',
	'г' :'ᑮ',
	'д' :'ᑯ',
	'е' :'ᒷ',
	'ж' :'ᒶ',
	'з' :'ᑲ',
	'и' :'и',
	'й' :'ᑴ',
	'к' :'к',
	'л' :'л',
	'м' :'ᑷ',
	'н' :'н',
	'о' :'о',
	'п' :'ᑺ',
	'р' :'р',
	'с' :'ᑼ',
	'т' :'ᑽ',
	'у' :'ᒐ',
	'ф' :'ф',
	'х' :'ᒜ',
	'ц' :'ᒓ',
	'ч' :'ᒙ',
	'ш' :'ᒥ',
	'щ' :'ᒖ',
	'ъ' :'ᒧ',
	'ы' :'ᒪ',
	'ь' :'ᒬ',
	'э' :'ᒭ',
	'ю' :'ᒮ',
	'я' :'я',
	' ' : 'ᕒ',
}

var key2 = map[rune]rune {

	'ы' :'ᑫ',
	'ь' :'ᒲ',
	'э' :'ᑭ',
	'ю' :'ᑮ',
	'я' :'ᑯ',
	'а' :'ᒷ',
	'б' :'ᒶ',
	'в' :'ᑲ',
	'г' :'ᒺ',
	'д' :'ᑴ',
	'е' :'ᑵ',
	'ж' :'ᑶ',
	'з' :'ᑷ',
	'и' :'ᑸ',
	'й' :'ᑹ',
	'к' :'ᑺ',
	'л' :'ᑻ',
	'м' :'ᑼ',
	'н' :'ᑽ',
	'о' :'ᒐ',
	'п' :'ᒠ',
	'р' :'ᒜ',
	'с' :'ᒓ',
	'т' :'ᒙ',
	'у' :'ᒥ',
	'ф' :'ᒖ',
	'х' :'ᒧ',
	'ц' :'ᒪ',
	'ч' :'ᒬ',
	'ш' :'ᒭ',
	'щ' :'ᒮ',
	'ъ' :'ᒯ',
	' ' : 'ᕒ',

}

func main() {

	fmt.Println(len([]rune(text)), len([]rune(keywords)))

	s := []rune{}

	for _, c := range []rune(text) {

		s = append(s, []rune{key_california[c], ' '}...)

	}

	fmt.Println(string(s))

	for pos, offset := range []rune(keywords) {

		fmt.Println([]rune(text)[pos], offset)

	}

//	for _, c := range s {
//
//		for k, v := range key {
//
//			if v == c {
//
//				fmt.Printf("%s", string(k))
//
//			}
//
//		}
//
//	}

}
