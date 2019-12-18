package handlers

import (
	"testing"
)

//test pagination
func TestGetAdvertListHandler1(t *testing.T) {
	/*req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAdvertListHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"alive": true}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}*/
}

//test data sorting
func TestGetAdvertListHandler2(t *testing.T) {

}

//test price sorting
func TestGetAdvertListHandler3(t *testing.T) {

}

//test getting adv by id
func TestGetAdvertHandler1(t *testing.T) {

}

//test getting adv with fieds
func TestGetAdvertHandler2(t *testing.T) {

}

func TestCreateAdvertHandler1(t *testing.T) {

}
