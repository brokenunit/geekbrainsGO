package cfgread

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type conf struct {
	Port        int
	Dburl       string
	Jaegerurl   string
	Sentryurl   string
	Kafkabroker string
	Someappid   string
	Someappkey  string
}

// Json read func
func Jsonread() {
	if _, err := os.Stat("../lesson_7/conf.json"); err != nil {
		log.Fatalf("Не могу проверить существование файла: %v", err)
	}
	// Создаем файловый дескриптор для файла, в котором хрнаится json-конфигурация
	file, err := os.Open("../lesson_7/conf.json") // URL: ../lesson_7/conf.json
	if err != nil {
		log.Fatalf("Не могу открыть файл: %v", err)
	}

	// Не забываем закрыть файл при выходе из функции
	defer func() {
		err := file.Close()
		if err != nil {
			log.Printf("Не могу закрыть файл: %v", err)
		}
	}()

	// Задаем переменную, в которую будем считывать конфиг
	app := conf{}

	// Задаем декодер из файла и сразу же вызываем его
	err = json.NewDecoder(file).Decode(&app)
	if err != nil {
		log.Printf("Не могу декодировать json-файл в структуру: %v", err)
	}

	// Теперь в app у нас доступны параметры из файла
	fmt.Printf("Json"+"%+v\n", app)

}

func (c *conf) Parse(data []byte) error {
	return yaml.Unmarshal(data, c)
}

// Yaml read func
func Yamlread() {
	if _, err := os.Stat("../lesson_7/conf.json"); err != nil {
		log.Fatalf("Не могу проверить существование файла: %v", err)
	}
	data, err := ioutil.ReadFile("../lesson_7/conf.yaml") // URL: ../lesson_7/conf.yaml
	if err != nil {
		log.Fatalf("Не могу открыть файл: %v", err)
	}
	// Задаем переменную, в которую будем считывать конфиг
	var config conf
	if err := config.Parse(data); err != nil {
		log.Printf("Не могу декодировать yaml-файл в структуру: %v", err)
	}
	fmt.Printf("Yaml"+"%+v\n", config)
}

func Cfgread() {
	Jsonread()
	Yamlread()
}
