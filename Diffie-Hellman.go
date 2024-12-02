package main

import (
	"fmt"
)

type DH_Endpoint struct {
	public_key1 int
	public_key2 int
	private_key int
	full_key    *int
}

// Функция создания частичного ключа шифрования
func (d *DH_Endpoint) generate_partial_key() int {
	partial_key := (d.public_key1 * d.private_key) % d.public_key2
	return partial_key
}

// Функция генерирования полного ключа
func (d *DH_Endpoint) generate_full_key(partial_key_r int) int {
	full_key := (partial_key_r * d.private_key) % d.public_key2
	d.full_key = &full_key
	return full_key
}

// Шифрование сообщения
func (d *DH_Endpoint) encrypt_message(message string) string {
	key := *d.full_key
	encrypted_message := ""
	for _, c := range message {
		encrypted_message += string(c + rune(key))
	}
	return encrypted_message
}

// Расшифровка сообщения
func (d *DH_Endpoint) decrypt_message(encrypted_message string) string {
	key := *d.full_key
	decrypted_message := ""
	for _, c := range encrypted_message {
		decrypted_message += string(c - rune(key))
	}
	return decrypted_message
}

func main() {
	c1_public := 3323
	c1_private := 2161
	c2_public := 3571
	c2_private := 2711

	companion_1 := DH_Endpoint{public_key1: c1_public, public_key2: c2_public, private_key: c1_private}
	companion_2 := DH_Endpoint{public_key1: c1_public, public_key2: c2_public, private_key: c2_private}

	// Ввод имен собеседников
	var name1, name2 string
	fmt.Print("Собеседник 1, введите ваше имя: ")
	fmt.Scanln(&name1)
	fmt.Print("Собеседник 2, введите ваше имя: ")
	fmt.Scanln(&name2)

	// Генерирование частичных ключей
	c1_partial := companion_1.generate_partial_key()
	c2_partial := companion_2.generate_partial_key()

	// Генерирование полных ключей
	companion_1.generate_full_key(c2_partial)
	companion_2.generate_full_key(c1_partial)

	// Цикл чата
	for {
		// Собеседник 1 отправляет сообщение
		var message string
		fmt.Printf("%s, введите сообщение: ", name1)
		fmt.Scanln(&message)

		if message == "exit" {
			fmt.Println("Диалог завершен")
			break
		}

		// Шифрование и отправка сообщения
		c1_encrypted := companion_1.encrypt_message(message)
		fmt.Printf("%s отправил зашифрованное сообщение: %s\n", name1, c1_encrypted)

		// Собеседник 2 получает и расшифровывает сообщение
		c2_decrypted := companion_2.decrypt_message(c1_encrypted)
		fmt.Printf("%s получил расшифрованное сообщение: %s\n", name2, c2_decrypted)

		// Собеседник 2 отвечает на сообщение
		fmt.Printf("%s, введите сообщение: ", name2)
		fmt.Scanln(&message)

		if message == "exit" {
			fmt.Println("Диалог завершен")
			break
		}

		// Шифрование и отправка ответа
		c2_encrypted := companion_2.encrypt_message(message)
		fmt.Printf("%s отправил зашифрованное сообщение: %s\n", name2, c2_encrypted)

		// Собеседник 1 получает и расшифровывает ответ
		c1_decrypted := companion_1.decrypt_message(c2_encrypted)
		fmt.Printf("%s получил расшифрованное сообщение: %s\n", name1, c1_decrypted)
	}
}
