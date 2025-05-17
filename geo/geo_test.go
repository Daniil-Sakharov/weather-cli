package geo_test

import (
	"go/BasGo/geo"
	"testing"
)

func TestGetMyLocation(t *testing.T) {
	//Arrange - подготовка, expected результат, данные для функции
	city := "Paris"
	expected := geo.DataGeo{
		City: "Paris",
	}
	//Act - выполняем функцию
	got, err := geo.GetMyLocation(city)
	//Assert - проверка результата с expected
	if err != nil {
		t.Error(err)
	}
	if got.City != expected.City {
		t.Errorf("Ожидалось %v, получено %v", expected, got)
	}
}
